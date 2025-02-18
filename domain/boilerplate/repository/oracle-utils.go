package repository

import (
	"svc-portofolio-golang/utils/database"
)

func (db *oracleBoilerplateRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
