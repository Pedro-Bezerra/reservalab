package controller

import (
	"net/http"
	"github.com/Pedro-Bezerra/mensagem/dto"
	"github.com/Pedro-Bezerra/mensagem/service"
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
		return 
	}

	c.JSON(http.StatusCreated, mensagemEnviada)
}

