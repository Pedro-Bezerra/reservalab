package dto

type UsuarioDto struct {
	Nome string `json:"Nome" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

type Login struct {
	Email string `json:"email" binding:"required"`
	Senha string `json:"senha" binding:"required"`
}
