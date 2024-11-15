package routes

import (
	"github.com/Pedro-Bezerra/reservalab-mono/model/controller"
	"github.com/gin-gonic/gin"
)

func RotasMensagem(router *gin.Engine) {
	mensagemRotas := router.Group("/mensagem")
	{
		mensagemRotas.POST("", controller.EnviarMensagem)
	}
}

func RotasSolicitacao(router *gin.Engine) {
	solicitacaoRotas := router.Group("/solicitacao")
	{
		solicitacaoRotas.POST("", controller.EnviarSolicitacao)
		solicitacaoRotas.PUT("/reagendamento", controller.ReagendarReserva)
	}
}

func RotasUsuario(router *gin.Engine) {
	rotasUsuario := router.Group("/usuario")
	{
		rotasUsuario.POST("/cadastro", controller.CadastrarMonitor)
		rotasUsuario.POST("/login", controller.Login)
	}
}
