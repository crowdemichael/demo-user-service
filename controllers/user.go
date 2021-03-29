package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crowdeco/demo-user-service/domain/user"
	common "github.com/crowdeco/demo-user-service/grpc/common"
	grpc "github.com/crowdeco/demo-user-service/grpc/user"
	"github.com/crowdeco/demo-user-service/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)

	res, err := services.GetUserProfile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {
	body := user.UserProfile{}
	c.Bind(&body)

	res, err := services.CreateUser(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

type UserHandler struct {
	// grpc.UsersServer
}

// func NewUserHandler() grpc.UsersServer {
// 	return &UserHandler{}
// }

func (UserHandler) GetUserProfile(c context.Context, r *grpc.IdInput) (*common.Response, error) {
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

func (UserHandler) CreateUser(c context.Context, r *grpc.User) (*common.Response, error) {
	dto := user.UserProfile{}

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
