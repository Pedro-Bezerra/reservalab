package routes

import (
	"github.com/Pedro-Bezerra/reserva/controller"
	"github.com/gin-gonic/gin"
)


func RotasSolicitacao(router *gin.Engine) {
	solicitacaoRotas := router.Group("/solicitacao")
	{
		solicitacaoRotas.POST("", controller.EnviarSolicitacao)
		solicitacaoRotas.PUT("/reagendamento", controller.ReagendarReserva)
	}
}

