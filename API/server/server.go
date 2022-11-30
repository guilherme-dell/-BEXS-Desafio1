package server

import (
	"log"
	"context"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"time"
)

type fireStoreConfig struct {
	ctx	context.Context
	client *firestore.Client
}

var FireStore fireStoreConfig

func Start() {

	StartFireEnvp()
	StartFire()

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	server := gin.Default()
	server.GET("/:roleID", pegarRole)
	server.GET("/roles", todosRoles)
	server.GET("/rolesHoje", rolesHoje)
	server.POST("/", marcarRole)
	server.PUT("/:roleID", alteraRole)
	server.DELETE("/:roleID", apagarRole)
	if err := server.Run(":8070"); err != nil {
		log.Fatal(err)
	}
}