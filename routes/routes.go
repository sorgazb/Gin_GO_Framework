package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		//	c.String(200, "Hola mundo")

		/*
			c.JSON(200, gin.H{
				"message": "Hola mundo",
			})
		*/

		c.String(http.StatusOK, "Hola mundo")
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola, %s", nombre)
	})
}
