package app

import (
	"github.com/crowdeco/demo-user-service/configs"
	"github.com/crowdeco/demo-user-service/connections"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	grpc := connections.UserGrpc{}
	go grpc.Run()

	mapUrls()
	port := configs.Env.AppPort
	router.Run(port)
}
