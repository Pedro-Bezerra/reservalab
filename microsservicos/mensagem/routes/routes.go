package routes

import (
	"github.com/Pedro-Bezerra/mensagem/controller"
	"github.com/gin-gonic/gin"
)

func RotasMensagem(router *gin.Engine) {
	mensagemRotas := router.Group("/mensagem")
	{
		mensagemRotas.POST("", controller.EnviarMensagem)
	}
}

