package entity


type Mensagem struct {
	IdMensagem uint   `json:"id_mensagem" gorm:"primaryKey;autoIncrement"`
	Nome       string `json:"nome" gorm:"not null"`
	Email      string `json:"email" gorm:"not null"`
	Mensagem   string `json:"mensagem" gorm:"not null"`
	Monitor    string `json:"monitor" gorm:"not null"`
}


func (Mensagem) TableName() string {
	return "mensagens"
}

