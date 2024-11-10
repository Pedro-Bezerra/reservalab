package controller

import (
	"net/http"

	"github.com/Pedro-Bezerra/reservalab-mono/model/dto"
	"github.com/Pedro-Bezerra/reservalab-mono/model/service"
	"github.com/gin-gonic/gin"
)

func EnviarMensagem(c *gin.Context) {
	var mensagemDto dto.MensagemDto
	
	if err := c.ShouldBindJSON(&mensagemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return 
	}

	mensagemEnviada, err := service.EnviarMensagem(&mensagemDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
	}

	c.JSON(http.StatusCreated, mensagemEnviada)
}

func EnviarSolicitacao(c *gin.Context) {
	var solicitacaoDto dto.SolicitacaoDto

	if err := c.ShouldBindJSON(&solicitacaoDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	solicitacaoEnviada, err := service.EnviarSolicitacao(&solicitacaoDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
	}

	c.JSON(http.StatusCreated, solicitacaoEnviada)
}