package main

import (
	"websim/handlers"
	"github.com/gin-gonic/gin"
	"github.com/shokunin/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	router := gin.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true, "websim"))

	// Start routes
	router.GET("/health", websim.HealthCheck)
	router.GET("/timer/:bucketstart/:bucketend", websim.Timer)
	router.GET("/timersize/:bucketstart/:bucketend/:mb", websim.TimerSize)

	// RUN rabit run
	router.Run() // listen and serve on 0.0.0.0:8080
}
