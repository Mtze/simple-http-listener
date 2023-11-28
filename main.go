package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	port = 80
)

func logResponse(c *gin.Context) {
	// Log the request
	log.WithFields(log.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
		"origin": c.Request.RemoteAddr,
	}).Info("Request received")

	// Log body content
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.WithError(err).Info("Failed to read request body")
	} else {
		var out bytes.Buffer
		json.Indent(&out, body, "", "  ")
		fmt.Print(out.String())
	}

	// Send the response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	if is_debug := os.Getenv("DEBUG"); is_debug == "true" {
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Start the server
	log.Info("Starting server...")

	r := gin.Default()
	r.POST("/", logResponse)
	log.Info("Server started")

	log.Panic(r.Run(fmt.Sprintf(":%d", port)))
}
