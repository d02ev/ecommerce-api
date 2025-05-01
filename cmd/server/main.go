package main

import (
	"fmt"

	"github.com/d02ev/ecommerce-api/pkg/config"
	"github.com/d02ev/ecommerce-api/pkg/db"
	"github.com/d02ev/ecommerce-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Load(); err != nil {
		panic("Failed to load config: " + err.Error());
	}

	// create Gin
	r := gin.New();

	// logger initialization
	logger.Init(viper.GetString("LOG_LEVEL"));
	// database initialization
	db.Init();

	r.Use(gin.LoggerWithWriter(logger.Log.Out));
	r.Use(gin.RecoveryWithWriter(logger.Log.Out));

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "pong" });
	});

	port := viper.GetString("PORT");

	logger.Log.Info(fmt.Sprintf("Server at http://localhost:%s", port));
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		logger.Log.Fatal("Failed to run server: " + err.Error());
	}
}