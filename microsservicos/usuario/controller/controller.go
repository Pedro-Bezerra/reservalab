package controller

import (
	"net/http"
	"github.com/Pedro-Bezerra/usuario/dto"
	"github.com/Pedro-Bezerra/usuario/service"
	"github.com/gin-gonic/gin"
)

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

