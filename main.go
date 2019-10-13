package main

import (
	"log"
	"paris/config"
	"paris/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	gin.SetMode(config.C.GetString("server.mode"))
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.L.WriterLevel(logrus.DebugLevel)}), gin.RecoveryWithWriter(logger.L.WriterLevel(logrus.ErrorLevel)))
	if err := engine.Run(config.C.GetString("server.port")); err != nil {
		log.Fatalf("Fatal error gin: %v\n", err)
	}
}
