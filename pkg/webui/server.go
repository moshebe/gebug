package webui

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/moshebe/gebug/pkg/webui/handlers"
)

type Server struct {
	Port   int
	cancel context.CancelFunc
}

const frontendDir = "./frontend/dist"

//func (s *Server) handleGetConfig(writer http.ResponseWriter, request *http.Request) {
//	ips, ok := request.URL.Query()["ip"]
//
//	if !ok || len(ips[0]) < 1 {
//		respondError(writer, errors.New("missing ip param"), http.StatusBadRequest)
//		return
//	}
//
//	ip := ips[0]
//
//	if net.ParseIP(ip) == nil {
//		respondError(writer, errors.New("invalid ip address"), http.StatusBadRequest)
//		return
//	}
//
//	info, err := s.dataStore.Get(ip)
//	if err != nil {
//		respondError(writer, errors.WithMessage(err, "get ip information from data store"), http.StatusInternalServerError)
//		return
//	}
//
//	if info == nil {
//		respondError(writer, errors.New("ip was not found in data store"), http.StatusNotFound)
//		return
//	}
//
//	res, err := json.Marshal(info)
//	if err != nil {
//		respondError(writer, errors.WithMessage(err, "marshal ip info"), http.StatusInternalServerError)
//		return
//	}
//
//	_, err = writer.Write(res)
//	if err != nil {
//		log.Println("Failed to write response, error: ", err)
//	}
//}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // TODO: do we really want to enable all?
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

func (s *Server) Start() error {
	fs := http.FileServer(http.Dir(frontendDir))
	http.Handle("/", fs)

	os.Setenv("PORT", "3030")

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(static.Serve("/", static.LocalFile(frontendDir, true)))
	//r.Use(static.Serve("/js", static.LocalFile("js", false)))
	//r.Use(static.Serve("/fonts", static.LocalFile("fonts", false)))

	r.GET("/config", handlers.HandleGetConfig)
	r.POST("/config", handlers.HandleCreateConfig)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.File(frontendDir + "/index.html")
	})

	//ctx, cancel := context.WithCancel(context.Background())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//s := &Server{
	//	cancel:    cancel,
	//}
	//router := mux.NewRouter()
	//var v1 = router.PathPrefix("/v1").Subrouter()
	//v1.HandleFunc("/config", s.handleFindCountry).Methods("GET")

	//s.server = &http.Server{Addr: ":" + strconv.Itoa(s.config.ListenPort), Handler: router}

	// http.HandleFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {

	// })
	fmt.Println("Server listening on port ", s.Port)

	return http.ListenAndServe(":"+strconv.Itoa(s.Port), nil)
}
