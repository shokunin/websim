package websim

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func SetStatus(c *gin.Context) {
	s, _ := strconv.Atoi(c.Param("status"))
	c.JSON(s, gin.H{
		"message": "OK",
		"status":  s,
	})
}
