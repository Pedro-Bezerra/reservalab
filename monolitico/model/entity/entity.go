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
	Ocupado       bool `json:"is_ocupado" gorm:"not null"`
}

func (Solicitacao) TableName() string {
	return "solicitacoes"
}

func (Mensagem) TableName() string {
	return "mensagens"
}

func (Reserva) TableName() string {
	return "reservas"
}
