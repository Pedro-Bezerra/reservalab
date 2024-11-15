package dto

import (
	"time"

	"github.com/Pedro-Bezerra/reservalab-mono/model/entity"
)

type SolicitacaoDto struct {
	Nome           string    `json:"nome" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	Matricula      string    `json:"matricula" binding:"required"`
	Vinculo        string    `json:"vinculo" binding:"required"`
	Observacoes    string    `json:"observacoes" binding:"required"`
	Data           time.Time `json:"data" binding:"required"`
	HorarioInicio  time.Time `json:"horario_inicio" binding:"required"`
	HorarioTermino time.Time `json:"horario_termino" binding:"required"`
}

type MensagemDto struct {
	Nome     string `json:"nome" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Mensagem string `json:"mensagem" binding:"required"`
	Monitor  string `json:"monitor" binding:"required"`
}

type ReservaDto struct {
	FkSolicitacao uint `json:"fk_solicitacao" binding:"required"`
	Ocupado       bool `json:"is_ocupado" binding:"required"`
}

type UsuarioDto struct {
	Usuario string `json:"usuario" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

type Login struct {
	Email string `json:"email" binding:"required"`
	Senha string `json:"senha" binding:"required"`
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

func CriarSolicitacao(dto *SolicitacaoDto) *entity.Solicitacao {
	solicitacao := entity.Solicitacao{
		Nome:           dto.Nome,
		Email:          dto.Email,
		Matricula:      dto.Matricula,
		Vinculo:        dto.Vinculo,
		Observacoes:    dto.Observacoes,
		Data:           dto.Data,
		HorarioInicio:  dto.HorarioInicio,
		HorarioTermino: dto.HorarioTermino,
	}

	return &solicitacao
}

func CriarReserva(dto *ReservaDto) *entity.Reserva {
	reserva := entity.Reserva{
		FkSolicitacao: dto.FkSolicitacao,
		Ocupado:       dto.Ocupado,
	}

	return &reserva
}
