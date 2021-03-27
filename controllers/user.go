package user

import (
	"net/http"
	"strconv"

	"github.com/crowdeco/demo-user-service/services"
	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 10, 64)

	res, err := services.GetUserProfile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
