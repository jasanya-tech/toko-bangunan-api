package helpers

import "database/sql"

func Transaction(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			panic(errRollback)
		}
		panic(err)
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			panic(errCommit)
		}
	}
}
