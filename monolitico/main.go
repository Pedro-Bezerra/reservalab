package main

import (
	"github.com/Pedro-Bezerra/reservalab-mono/config"
	"github.com/Pedro-Bezerra/reservalab-mono/database"
	"github.com/Pedro-Bezerra/reservalab-mono/model/entity"
	"github.com/Pedro-Bezerra/reservalab-mono/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()

	database.ConectarDB()

	database.DB.AutoMigrate(&entity.Solicitacao{}, &entity.Mensagem{}, &entity.Reserva{}, &entity.Usuario{})
}

func main() { 
	router := gin.Default()

	routes.RotasMensagem(router)
	routes.RotasSolicitacao(router)
	routes.RotasUsuario(router)

	router.Run(":8080")
}
