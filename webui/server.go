package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/moshebe/gebug/pkg/config"
	"github.com/moshebe/gebug/pkg/osutil"
	"go.uber.org/zap"
)

type server struct {
	port     int
	location string
	logger   *zap.Logger
}

const (
	frontendDir       = "./frontend/dist"
	locationEnvName   = "VUE_APP_GEBUG_PROJECT_LOCATION"
	defaultServerPort = 3030
)

func errorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"message": msg,
	})
}

func (s server) handleGetConfig(c *gin.Context) {
	response := struct {
		Location   string         `json:"location"`
		HasProject bool           `json:"has_project"`
		Config     *config.Config `json:"config"`
	}{}

	response.HasProject = osutil.FileExists(s.location)
	response.Location = s.location

	s.logger.Debug("Received Get config request", zap.String("location", response.Location), zap.Bool("hasProject", response.HasProject))
	if !response.HasProject {
		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
		return
	}

	configData, err := ioutil.ReadFile(s.location)
	if err != nil {
		s.logger.Error("Failed to read configuration file", zap.String("location", s.location), zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "unable to read config file")
		return
	}
	cnfg, err := config.Load(configData)
	if err != nil {
		s.logger.Error("Failed to load configuration from file content", zap.String("data", string(configData)), zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "unable to load config from file")
		return
	}

	response.Config = cnfg

	s.logger.Debug("Response data", zap.Bool("hasProject", response.HasProject), zap.Any("config", *response.Config))
	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (s server) handleCreateConfig(c *gin.Context) {
	data := struct {
		Config config.Config `json:"config"`
	}{}

	s.logger.Debug("Got save config request")
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		s.logger.Error("Failed to decode request body", zap.Error(err))
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	s.logger.Debug("Received config from client", zap.Any("config", data.Config))
	if !osutil.FileExists(s.location) {
		configDir := filepath.Dir(s.location)
		if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
			s.logger.Error("Failed to create directory and its parents", zap.String("path", s.location), zap.Error(err))
			errorResponse(c, http.StatusInternalServerError, "create config directory")
			return
		}
	}

	configFile, err := os.OpenFile(s.location, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		s.logger.Error("Failed to open config file", zap.String("path", s.location), zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "open config file")
		return
	}
	defer func() {
		_ = configFile.Close()
	}()

	err = data.Config.Write(configFile)
	if err != nil {
		s.logger.Error("Failed to write config file", zap.String("path", s.location), zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, "write config file")
		return
	}

	s.logger.Info("Configuration saved successfully", zap.Any("config", data.Config))
	c.JSON(http.StatusOK, gin.H{})
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *server) Start() error {
	s.location = config.ResolvePath(s.location)
	s.logger.Debug("Resolved project location", zap.String("path", s.location))

	fs := http.FileServer(http.Dir(frontendDir))
	http.Handle("/", fs)

	r := gin.Default()
	r.Use(corsMiddleware())
	r.Use(static.Serve("/", static.LocalFile(frontendDir, true)))

	r.GET("/config", s.handleGetConfig)
	r.POST("/config", s.handleCreateConfig)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	r.NoRoute(func(c *gin.Context) {
		c.File(frontendDir + "/index.html")
	})

	return r.Run()
}

func main() {
	logConfig := zap.NewDevelopmentConfig()
	if os.Getenv("VERBOSE") == "" {
		gin.SetMode(gin.ReleaseMode)
		logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("Initializing...")

	location := os.Getenv(locationEnvName)
	if location == "" {
		logger.Sugar().Fatalf("Could not find project location, make sure %q was set correctly", locationEnvName)
	}
	s := &server{
		port:     defaultServerPort,
		location: location,
		logger:   logger,
	}
	logger.Info("Starting server...", zap.Int("port", s.port), zap.String("location", s.location))
	err = s.Start()
	if err != nil {
		logger.Fatal("Unexpected error while running server", zap.Error(err))
	}
}
