package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Pedro-Bezerra/reservalab-mono/database"
	"github.com/Pedro-Bezerra/reservalab-mono/model/entity"
	"github.com/Pedro-Bezerra/reservalab-mono/model/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RealizarAuth(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		log.Fatalf("Erro na obtenção do cookie: %v", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		log.Fatalf("erro no processamento do token: %v", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var usuario entity.Usuario
		database.DB.First(&usuario, claims["sub"])

		if usuario.IdUsuario == 0 || usuario.Tipo != utils.ADMIN {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("usuario", usuario)
		c.Next()
		
	} else {
		log.Fatalf("erro na validação do token: %v", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}