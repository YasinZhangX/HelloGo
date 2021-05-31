package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()

		if id, exist := c.Get("requestId"); exist {
			logger.Info("incoming request",
				zap.String("requestId", strconv.Itoa(id.(int))),
				zap.String("path", c.Request.URL.Path),
				zap.Int("status", c.Writer.Status()),
				zap.String("elapsed", time.Since(s).String()))
		} else {
			logger.Info("incoming request",
				zap.String("path", c.Request.URL.Path),
				zap.Int("status", c.Writer.Status()),
				zap.String("elapsed", time.Since(s).String()))
		}

	}, func(c *gin.Context) {
		c.Set("requestId", rand.Int())

		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

		time.Sleep(time.Nanosecond)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
