package database

import (
	"database/sql"
	"fmt"
	"restapi/internal/configuration"
	"restapi/internal/util"
)
import _ "github.com/lib/pq"

var sqlConnectionString string

func GetConnection() (*sql.DB, error) {
	conf := configuration.ConfigurationManager().GetCached()
	open, err := sql.Open("postgres", getConnectionString(conf.DbConfiguration))

	if err != nil {
		util.Logger().Fatal(err)
		return nil, err
	}

	if err := open.Ping(); err != nil {
		util.Logger().Fatal(err)
		return nil, err
	}

	return open, nil
}

func getConnectionString(dbConf *configuration.DatabaseConfiguration) string {
	if sqlConnectionString == "" {
		sqlConnectionString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConf.User, dbConf.Password, dbConf.DbName)
	}

	return sqlConnectionString
}
