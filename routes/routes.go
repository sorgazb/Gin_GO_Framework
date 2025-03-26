package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

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

	r.POST("/usuarios", func(c *gin.Context) {
		var nuevoUsuario Usuario
		if err := c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el JSON"})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y correo son campos requeridos"})
		}

		usuarios = append(usuarios, nuevoUsuario)

		c.JSON(http.StatusOK, gin.H{"message": "Usuario Registrado", "datos": usuarios})
	})
}
