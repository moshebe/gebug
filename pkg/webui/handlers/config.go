package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/moshebe/gebug/pkg/config"
	"github.com/moshebe/gebug/pkg/osutil"
)

func errorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"message": msg,
	})
}

func HandleGetConfig(c *gin.Context) {
	// TODO: try to get from env if not exists
	workDir := c.Query("path")
	if workDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("failed to get current working directory, error: ", err)
			errorResponse(c, http.StatusBadRequest, "missing parameter 'path'")
			return
		}
		workDir = cwd
	}
	response := struct {
		Location   string         `json:"location"`
		HasProject bool           `json:"has_project"`
		Config     *config.Config `json:"config"`
	}{}

	location := config.ResolvePath(workDir)

	hasProject := osutil.FileExists(location)
	if !hasProject {
		response.Location = workDir
		c.JSON(http.StatusOK, gin.H{
			"data": response,
		})
		return
	}

	response.Location = location

	configData, err := ioutil.ReadFile(location)
	if err != nil {
		fmt.Println("failed to get current working directory, error: ", err)
		errorResponse(c, http.StatusInternalServerError, "unable to read config file from: "+workDir)
		return
	}
	cnfg, err := config.Load(configData)
	if err != nil {
		fmt.Println("failed to load configuration, error: ", err)
		errorResponse(c, http.StatusInternalServerError, "unable to load config file from: "+workDir)
		return
	}

	response.Config = cnfg

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func HandleCreateConfig(c *gin.Context) {
	// TODO: add here validations for the configuration values
	data := struct {
		Location string        `json:"location"`
		Config   config.Config `json:"config"`
	}{}

	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		errorResponse(c, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	configPath := config.ResolvePath(data.Location)
	log.Println("Resolved config path: ", configPath)
	log.Println("exists: ", osutil.FileExists(configPath))

	if !osutil.FileExists(configPath) {
		configDir := filepath.Dir(configPath)
		if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
			errorResponse(c, http.StatusInternalServerError, "create config directory: "+err.Error())
			return
		}
	}

	configFile, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "open config file: "+err.Error())
		return
	}
	defer func() {
		_ = configFile.Close()
	}()

	err = data.Config.Write(configFile)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "write config file: "+err.Error())
		return
	}

	log.Println("Finished successfully. ", data.Config)
	c.JSON(http.StatusOK, gin.H{})
}
