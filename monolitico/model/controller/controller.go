package controller

import (
	"log"
	"net/http"
	"strconv"
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
		return 
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
		return
	}

	c.JSON(http.StatusCreated, solicitacaoEnviada)
}

func ReagendarReserva(c *gin.Context) {
	var solicitacaoDto dto.SolicitacaoDto
	idSolicitacao := c.Query("idSolicitacao")

	id64, err := strconv.ParseUint(idSolicitacao, 10, 64)
	if err != nil {
		log.Fatalf("Erro na passagem da solicitação antiga: %v", err)
	}

	idUint := uint(id64)
	

	if err := c.ShouldBindJSON(&solicitacaoDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	solicitacaoAtualizada, err := service.AtualizarSolicitacao(idUint, &solicitacaoDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, solicitacaoAtualizada)
}

func CadastrarMonitor(c *gin.Context) {
	var usuarioDto dto.UsuarioDto

	if err := c.ShouldBindJSON(&usuarioDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	_, err := service.CadastrarMonitor(&usuarioDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func Login(c *gin.Context) {
	var login dto.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	token, err := service.Login(&login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24 * 30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}