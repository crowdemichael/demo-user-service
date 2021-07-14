package connections

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/crowdeco/demo-user-service/configs"
	dto "github.com/crowdeco/demo-user-service/domain/user"
	"github.com/crowdeco/demo-user-service/grpc/common"
	"github.com/crowdeco/demo-user-service/grpc/user"
	"github.com/crowdeco/demo-user-service/services"
	"github.com/jinzhu/copier"

	"google.golang.org/grpc"
)

type UserGrpc struct{}

func (UserGrpc) Run() {
	l, err := net.Listen("tcp", configs.Env.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	user.RegisterUsersServer(srv, UserHandler{})

	log.Fatal(srv.Serve(l))
}

type UserHandler struct {
	// grpc.UsersServer
}

func (UserHandler) GetUserProfile(c context.Context, r *user.IdInput) (*common.Response, error) {
	println("i am coming id: ", r.Id)
	res, err := services.GetUserProfile(r.Id)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(res)

	resp := &common.Response{}
	resp.Status = http.StatusOK
	resp.Data = data
	resp.Message = "success"

	return resp, nil
}

func (UserHandler) CreateUser(c context.Context, r *user.User) (*common.Response, error) {
	dto := dto.UserProfile{}

	copier.Copy(&dto, &r)

	res, err := services.CreateUser(&dto)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(res)

	resp := &common.Response{}
	resp.Status = http.StatusOK
	resp.Data = data
	resp.Message = "success"

	return resp, nil
}
