package websim

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Timer(c *gin.Context) {
	var randsleep int
	s, _ := strconv.Atoi(c.Param("bucketstart"))
	e, _ := strconv.Atoi(c.Param("bucketend"))
	if e > s {
		randsleep = s + rand.Intn(e-s)
	} else {
		randsleep = s
	}
	time.Sleep(time.Duration(randsleep) * time.Millisecond)
	c.JSON(200, gin.H{
		"message":      "OK",
		"bucketstart":  c.Param("bucketstart"),
		"bucketend":    c.Param("bucketend"),
		"randsleep_ms": randsleep,
	})
}

func TimerSize(c *gin.Context) {
	var randsleep int
	s, _ := strconv.Atoi(c.Param("bucketstart"))
	e, _ := strconv.Atoi(c.Param("bucketend"))
	m, _ := strconv.Atoi(c.Param("mb"))
	var mydata []byte
	for i := 0; i < (m * 1000000); i++ {
		mydata = append(mydata, 0)
	}
	if e > s {
		randsleep = s + rand.Intn(e-s)
	} else {
		randsleep = s
	}
	time.Sleep(time.Duration(randsleep) * time.Millisecond)
	c.JSON(200, gin.H{
		"message":      "OK",
		"bucketstart":  c.Param("bucketstart"),
		"bucketend":    c.Param("bucketend"),
		"randsleep_ms": randsleep,
		"size_bytes":   len(mydata),
		"data":         mydata,
	})
}

func TimerSizeBytes(c *gin.Context) {
	s, _ := strconv.Atoi(c.Param("bucketstart"))
	e, _ := strconv.Atoi(c.Param("bucketend"))
	bytes, _ := strconv.Atoi(c.Param("bytes"))
	var mydata []byte
	for i := 0; i < (bytes); i++ {
		mydata = append(mydata, 0)
	}
	randsleep := s + rand.Intn(e-s)
	time.Sleep(time.Duration(randsleep) * time.Millisecond)
	c.JSON(200, gin.H{
		"message":      "OK",
		"bucketstart":  c.Param("bucketstart"),
		"bucketend":    c.Param("bucketend"),
		"randsleep_ms": randsleep,
		"size_bytes":   len(mydata),
		"data":         mydata,
	})
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
