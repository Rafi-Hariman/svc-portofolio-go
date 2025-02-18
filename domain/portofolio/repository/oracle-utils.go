package repository

import (
	"svc-portofolio-golang/utils/database"
)

func (db *oraclePortofolioRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
