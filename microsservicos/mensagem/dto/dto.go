package dto

import (
	"github.com/Pedro-Bezerra/mensagem/entity"
)



type MensagemDto struct {
	Nome     string `json:"nome" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Mensagem string `json:"mensagem" binding:"required"`
	Monitor  string `json:"monitor" binding:"required"`
}


func CriarMensagem(dto *MensagemDto) *entity.Mensagem {
	mensagem := entity.Mensagem{
		Nome:     dto.Nome,
		Email:    dto.Email,
		Mensagem: dto.Mensagem,
		Monitor:  dto.Monitor,
	}

	return &mensagem
}
