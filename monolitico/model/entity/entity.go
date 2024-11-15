package entity

import (
	"time"

)

type Solicitacao struct {
	IdSolicitacao  uint      `json:"id_solicitacao" gorm:"primaryKey;autoIncrement"`
	Nome           string    `json:"nome" gorm:"not null"`
	Email          string    `json:"email" gorm:"not null"`
	Matricula      string    `json:"matricula" gorm:"not null"`
	Vinculo        string    `json:"vinculo" gorm:"not null"`
	Observacoes    string    `json:"observacoes" gorm:"not null"`
	Data           time.Time `json:"data" gorm:"not null"`
	HorarioInicio  time.Time `json:"horario_inicio" gorm:"not null"`
	HorarioTermino time.Time `json:"horario_termino" gorm:"not null"`
}

type Mensagem struct {
	IdMensagem uint   `json:"id_mensagem" gorm:"primaryKey;autoIncrement"`
	Nome       string `json:"nome" gorm:"not null"`
	Email      string `json:"email" gorm:"not null"`
	Mensagem   string `json:"mensagem" gorm:"not null"`
	Monitor    string `json:"monitor" gorm:"not null"`
}

type Reserva struct {
	IdReserva     uint `json:"id_reserva" gorm:"primaryKey;autoIncrement"`
	FkSolicitacao uint `json:"fk_solicitacao" gorm:"foreignKey:IdSolicitacao"`
	Data           time.Time `json:"data" gorm:"not null"`
	HorarioInicio  time.Time `json:"horario_inicio" gorm:"not null"`
	HorarioTermino time.Time `json:"horario_termino" gorm:"not null"`
	Ocupado       bool `json:"is_ocupado" gorm:"not null"`
}

type Usuario struct {
	IdUsuario uint `json:"id_usuario" gorm:"primaryKey;autoIncrement"`
	Usuario string `json:"usuario" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Senha string `json:"senha" gorm:"not null"`
	Tipo string `json:"tipo" gorm:"not null"`
}

func CriarReserva(solicitacao *Solicitacao) (*Reserva) {
	reserva := Reserva{
		FkSolicitacao: solicitacao.IdSolicitacao,
		Data: solicitacao.Data,
		HorarioInicio: solicitacao.HorarioInicio,
		HorarioTermino: solicitacao.HorarioTermino,
		Ocupado: true,
	}

	return &reserva
}

/*

func CriarSolicitacaoDto(solicitacao *Solicitacao) (*dto.SolicitacaoDto) {
	solicitacaoDto := dto.SolicitacaoDto{
		Nome: solicitacao.Nome,
		Email: solicitacao.Email,
		Matricula: solicitacao.Matricula,
		Vinculo: solicitacao.Vinculo,
		Observacoes: solicitacao.Observacoes,
		Data: solicitacao.Data,
		HorarioInicio: solicitacao.HorarioInicio,
		HorarioTermino: solicitacao.HorarioTermino,
	}

	return &solicitacaoDto
}

func CriarMensagemDto(mensagem *Mensagem) (*dto.MensagemDto) {
	mensagemDto := dto.MensagemDto{
		Nome: mensagem.Nome,
		Email: mensagem.Email,
		Mensagem: mensagem.Mensagem,
		Monitor: mensagem.Monitor,
	}

	return &mensagemDto
}

func CriarReservaDto(reserva *Reserva) (*dto.ReservaDto) {
	reservaDto := dto.ReservaDto{
		FkSolicitacao: reserva.FkSolicitacao,
		Ocupado: reserva.Ocupado,
	}

	return &reservaDto
}
*/

func (Solicitacao) TableName() string {
	return "solicitacoes"
}

func (Mensagem) TableName() string {
	return "mensagens"
}

func (Reserva) TableName() string {
	return "reservas"
}

func (Usuario) TableName() string {
	return "usuarios"
}
