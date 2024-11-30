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

type Reserva struct {
	IdReserva     uint `json:"id_reserva" gorm:"primaryKey;autoIncrement"`
	FkSolicitacao uint `json:"fk_solicitacao" gorm:"foreignKey:IdSolicitacao"`
	Data           time.Time `json:"data" gorm:"not null"`
	HorarioInicio  time.Time `json:"horario_inicio" gorm:"not null"`
	HorarioTermino time.Time `json:"horario_termino" gorm:"not null"`
	Ocupado       bool `json:"is_ocupado" gorm:"not null"`
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


func (Solicitacao) TableName() string {
	return "solicitacoes"
}

func (Reserva) TableName() string {
	return "reservas"
}


