package connections

import (
	"log"
	"net"

	"github.com/crowdeco/demo-user-service/configs"
	"github.com/crowdeco/demo-user-service/controllers"
	"github.com/crowdeco/demo-user-service/grpc/user"

	"google.golang.org/grpc"
)

type UserGrpc struct{}

func (UserGrpc) Run() {
	l, err := net.Listen("tcp", configs.Env.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	user.RegisterUsersServer(srv, controllers.UserHandler{})
	
	log.Fatal(srv.Serve(l))
}
