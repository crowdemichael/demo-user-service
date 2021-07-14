package controllers

import (
	"net/http"
	"strconv"

	"github.com/crowdeco/demo-user-service/domain/user"
	"github.com/crowdeco/demo-user-service/services"
	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   res,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"data":   res,
	})
}
