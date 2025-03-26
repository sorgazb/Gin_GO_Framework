# Servidor Web Basico en Go
````go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//	c.String(200, "Hola mundo")
		c.JSON(200, gin.H{
			"message": "Hola mundo",
		})
	})

	r.Run(":8080")
}

````