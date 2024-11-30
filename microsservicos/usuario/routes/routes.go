package routes

import (
	"github.com/Pedro-Bezerra/usuario/middleware"
	"github.com/Pedro-Bezerra/usuario/controller"
	"github.com/gin-gonic/gin"
)


func RotasMonitor(router *gin.Engine) {
	rotasUsuario := router.Group("/usuario")
	{
		rotasUsuario.POST("/cadastro", middleware.RealizarAuth, controller.CadastrarMonitor)
		rotasUsuario.POST("/login", controller.Login)
	}
}
