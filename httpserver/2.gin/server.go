package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello GIN!")

		//c.JSON(http.StatusOK, gin.H{
		//	"message": "hello GIN!",
		//})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8081

}
