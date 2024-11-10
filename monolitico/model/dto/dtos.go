package dto

type MensagemDto struct {
	Nome string `json:"nome" binding:"required"`
	Email string `json:"email" binding:"required, email"`
	Mensagem string `json:"mensagem" binding:"required"`
	Monitor string `json:"monitor" binding:"required"`
}