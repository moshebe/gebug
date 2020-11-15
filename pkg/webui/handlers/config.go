package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moshebe/gebug/pkg/config"
	"github.com/moshebe/gebug/pkg/osutil"
	"io/ioutil"
	"net/http"
	"os"
)

func errorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"message": msg,
	})
}

func HandleGetConfig(c *gin.Context) {
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
	// TODO: handle POST request holding the location and the configuration content
	// should create the directory if needed (.gebug) and the config.yaml file with the relevant content
	// TODO: add here validations for the configuration values
}
