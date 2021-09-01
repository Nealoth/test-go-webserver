package repository

import (
	"restapi/internal/database"
	"restapi/internal/models"
	"restapi/internal/util"
)

func CreateUser(u *models.User) error {
	db, dbConnectionErr := database.GetConnection()
	defer util.CloseDb(db)

	if dbConnectionErr != nil {
		util.Logger().Error(dbConnectionErr)
		return dbConnectionErr
	}

	_, dbQueryErr := db.Exec("INSERT INTO users (id, email, password) VALUES (nextval('users_seq'), $1, $2)", u.Email, u.Password)

	if dbQueryErr != nil {
		util.Logger().Error(dbQueryErr)
		return dbQueryErr
	}

	return nil
}

func GetUsers() (users []models.User, err error) {
	db, dbConnectionErr := database.GetConnection()
	defer util.CloseDb(db)

	if dbConnectionErr != nil {
		util.Logger().Error(dbConnectionErr)
		return nil, dbConnectionErr
	}

	res, queryErr := db.Query("SELECT * FROM users")

	if queryErr != nil {
		util.Logger().Error(queryErr)
		return nil, queryErr
	}

	for res.Next() {
		u := models.User{}

		scanErr := res.Scan(&u.ID, &u.Email, &u.Password)

		if scanErr != nil {
			util.Logger().Error(scanErr)
			return nil, scanErr
		}

		users = append(users, u)
	}

	return users, nil
}
