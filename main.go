package main

import (
	controllerLivros "projeto404/src/Api/Controllers/ControllerLivros"
	controllerUsuarios "projeto404/src/Api/Controllers/ControllerUsuarios"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Set up Gin router
	router := gin.Default()

	// API routes
	api := router.Group("/api/v1")
	{
		// Usuarios
		api.GET("/usuario", controllerUsuarios.ListarUsuarios)
		api.GET("/usuario/:id", controllerUsuarios.ListarUsuarios)
		api.POST("/usuario", controllerUsuarios.CriarUsuarios)
		api.PUT("/usuario/:id", controllerUsuarios.AtualizarUsuarios)
		api.DELETE("/usuario/:id", controllerUsuarios.DeletarUsuarios)

		// Livros
		api.GET("/livros", controllerLivros.ListarLivro)
		api.GET("/livros/:id", controllerLivros.ListarLivro)
		api.POST("/livros/", controllerLivros.CriarLivro)
		api.PUT("/livros/:id", controllerLivros.AtualizarLivro)
		api.DELETE("/livros/:id", controllerLivros.DeletarLivro)

	}

	log.Println("Server started on port 8080")
	router.Run(":8080")
}
