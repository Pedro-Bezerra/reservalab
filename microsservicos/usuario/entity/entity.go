package entity

type Usuario struct {
	IdUsuario uint `json:"id_usuario" gorm:"primaryKey;autoIncrement"`
	Nome string `json:"nome" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Senha string `json:"senha" gorm:"not null"`
	Tipo string `json:"tipo" gorm:"not null"`
}

func (Usuario) TableName() string {
	return "monitores"
}
