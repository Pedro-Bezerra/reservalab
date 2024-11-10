package service

import (
	"log"

	"github.com/Pedro-Bezerra/reservalab-mono/database"
	"github.com/Pedro-Bezerra/reservalab-mono/model/dto"
	"github.com/Pedro-Bezerra/reservalab-mono/model/entity"
)

func EnviarMensagem(mensagemDto *dto.MensagemDto) (*entity.Mensagem, error) {

	var mensagem entity.Mensagem = *dto.CriarMensagem(mensagemDto)

	if err := database.DB.Create(&mensagem).Error; err != nil {
		log.Fatalf("Erro no armazenamento da mensagem: %v", err)
		return nil, err
	}

	return &mensagem, nil
}

//c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})