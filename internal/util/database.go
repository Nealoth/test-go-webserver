package util

import "database/sql"

func CloseDb(db *sql.DB) {
	err := db.Close()
	if err != nil {
		Logger().Error(err)
	}
}