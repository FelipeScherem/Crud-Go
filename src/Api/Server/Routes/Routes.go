package rotas

import (
	"github.com/gin-gonic/gin"
	controllerProdutos "projeto404/src/Api/Controllers/ControllerProdutos"
	controllerUsuarios "projeto404/src/Api/Controllers/ControllerUsuarios"
)

// ConfiguraRotas Define os endpoints
func ConfiguraRotas(router *gin.Engine) *gin.Engine {
	// API routes
	api := router.Group("/api/v1")
	{
		// Usuarios
		api.GET("/usuario", controllerUsuarios.ListarUsuarios)
		api.GET("/usuario/:id", controllerUsuarios.ListarUsuarios)
		api.POST("/usuario", controllerUsuarios.CriarUsuarios)
		api.PUT("/usuario/:id", controllerUsuarios.AtualizarUsuarios)
		api.DELETE("/usuario/:id", controllerUsuarios.DeletarUsuarios)

		// Produtos
		api.GET("/produtos", controllerProdutos.ListarProduto)
		api.GET("/produtos/:id", controllerProdutos.ListarProduto)
		api.POST("/produtos/", controllerProdutos.CriarProduto)
		api.PUT("/produtos/:id", controllerProdutos.AtualizarProduto)
		api.DELETE("/produtos/:id", controllerProdutos.DeletarProduto)
	}

	return router
}
