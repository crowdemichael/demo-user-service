package app

import (
	"github.com/crowdeco/demo-user-service/controllers"
)

func mapUrls() {
	v1 := router.Group("api/v1/")
	{
		userGroup := v1.Group("users")
		{
			userGroup.GET("/:id", controllers.GetUserProfile)
			userGroup.POST("/", controllers.CreateUser)
		}
	}
}
