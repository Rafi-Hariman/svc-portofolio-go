package repository

import (
	"svc-portofolio-golang/utils/database"
)

func (db *oracleAuthRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
