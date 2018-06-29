package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shokunin/contrib/ginrus"
	websim "github.com/shokunin/websim/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true, "websim"))

	// Start routes
	router.GET("/health", websim.HealthCheck)
	router.GET("/status/:status", websim.SetStatus)
	router.GET("/timer/:bucketstart/:bucketend", websim.Timer)
	router.GET("/timersize/:bucketstart/:bucketend/:mb", websim.TimerSize)
	router.GET("/timersizebytes/:bucketstart/:bucketend/:bytes", websim.TimerSizeBytes)
	// RUN rabit run
	router.Run() // listen and serve on 0.0.0.0:8080
}
