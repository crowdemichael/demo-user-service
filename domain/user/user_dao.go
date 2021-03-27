package user

import (
	"database/sql"
	"fmt"

	"github.com/crowdeco/demo-user-service/database"
	"github.com/crowdeco/demo-user-service/query"
	logger_ "github.com/crowdeco/logger"
)

var logger = logger_.NewLogger()

// sql

func GetUserProfile(id int64) (*UserProfile, error) {
	stmt, err := database.Client.Prepare(query.QueryGetUserProfile)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to prepare get user profile statement: %s", err))
		return nil, err
	}
	defer stmt.Close()

	var res UserProfile

	row := stmt.QueryRow(id)

	err = row.Scan(&res.Name, &res.Email)
	switch err {
	case nil:
		// it's fine
	case sql.ErrNoRows:
		logger.Info(fmt.Sprintf("%s", err))
		return nil, err
	default:
		logger.Error(fmt.Sprintf("error when trying to execute get user profile: %s", err))
		return nil, err
	}

	return &res, nil
}

// gorm
