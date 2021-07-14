package user

import (
	"database/sql"
	"fmt"

	"github.com/crowdeco/demo-user-service/database"
	"github.com/crowdeco/demo-user-service/database/schema"
	"github.com/crowdeco/demo-user-service/query"
	logger_ "github.com/crowdeco/logger"
	"github.com/jinzhu/copier"
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

	res.Id = id

	return &res, nil
}

func CreateUser(v *UserProfile) (*UserProfile, error) {
	stmt, err := database.Client.Prepare(query.QueryCreateUser)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to prepare create user statement: %s", err))
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(v.Name, v.Email)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to execute create user: %s", err))
		return nil, err
	}

	v.Id, err = res.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get last insert id: %s", err))
		return nil, err
	}

	return v, nil
}

func CreateUser2(r *UserProfile) (*UserProfile, error) {
	db := database.Database

	user := schema.User{}
	copier.Copy(&user, &r)

	err := db.Create(&user).Error
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	r.Id = user.Id

	return r, nil
}
