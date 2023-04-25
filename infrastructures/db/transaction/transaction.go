package transaction

import (
	"database/sql"
)

func Transaction(tx *sql.Tx) {
	if err := recover(); err != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	} else {
		if err := tx.Commit(); err != nil {
			panic(err)
		}
	}
}
