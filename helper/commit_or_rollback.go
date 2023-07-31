package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	error := recover()

	if error != nil {
		err := tx.Rollback()
		PanicIfError(err)
		panic(error)
	} else {
		err := tx.Commit()
		PanicIfError(err)
	}
}
