package service

import (
	"log"
	"github.com/Pedro-Bezerra/mensagem/database"
	"github.com/Pedro-Bezerra/mensagem/dto"
	"github.com/Pedro-Bezerra/mensagem/entity"
)

func EnviarMensagem(mensagemDto *dto.MensagemDto) (*entity.Mensagem, error) {

	var mensagem entity.Mensagem = *dto.CriarMensagem(mensagemDto)

	if err := database.DB.Create(&mensagem).Error; err != nil {
		log.Fatalf("Erro no armazenamento da mensagem: %v", err)
		return nil, err
	}

	return &mensagem, nil
}
