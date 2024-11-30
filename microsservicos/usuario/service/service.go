package service

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/Pedro-Bezerra/usuario/database"
	"github.com/Pedro-Bezerra/usuario/dto"
	"github.com/Pedro-Bezerra/usuario/entity"
	"github.com/Pedro-Bezerra/usuario/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CadastrarMonitor(usuarioDto *dto.UsuarioDto) (bool, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(usuarioDto.Senha), 10)

	if err != nil {
		log.Fatalf("Erro na encriptação da senha: %v", err)
		return false, errors.New("ERRO NA ENCRIPTAÇÃO DA SENHA")
	}

	var monitor entity.Usuario = entity.Usuario{Nome: usuarioDto.Nome, Email: usuarioDto.Email, Senha: string(hash), Tipo: utils.MONITOR}

	if err := database.DB.Create(&monitor).Error; err != nil {
		log.Fatalf("Erro no cadastro do monitor: %v", err)
		return false, errors.New("ERRO NO CADASTRO DO MONITOR")
	}

	return true, nil
}

func Login(login *dto.Login) (string, error) {
	var usuario entity.Usuario

	if err := database.DB.Where("email = ?", login.Email).Find(&usuario).Error; err != nil {
		log.Fatalf("Erro no login do monitor: %v", err)
		return "", errors.New("ERRO NO LOGIN DO MONITOR")
	}

	err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(login.Senha))

	if err != nil {
		return "", errors.New("SENHA ERRADA")
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": usuario.IdUsuario,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		log.Fatalf("Erro na geração do token JWT: %v", err)
		return "", errors.New("ERRO NA GERAÇÃO DO TOKEN JWT")
	}

	return tokenString, nil

}

