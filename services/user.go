package services

import (
	"github.com/crowdeco/demo-user-service/domain/user"
)

func GetUserProfile(id int64) (*user.UserProfile, error) {
	return user.GetUserProfile(id)
}
