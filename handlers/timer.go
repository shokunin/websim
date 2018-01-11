package websim

import (
	"github.com/gin-gonic/gin"
	"math/rand"
  "strconv"
  "time"
)

func Timer(c *gin.Context) {
  s, _ := strconv.Atoi(c.Param("bucketstart"))
  e, _ := strconv.Atoi(c.Param("bucketend"))
	randsleep := s + rand.Intn(e-s)
  time.Sleep(time.Duration(randsleep) * time.Millisecond)
	c.JSON(200, gin.H{
		"message":     "OK",
		"bucketstart": c.Param("bucketstart"),
		"bucketend":   c.Param("bucketend"),
		"randsleep":   randsleep,
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
