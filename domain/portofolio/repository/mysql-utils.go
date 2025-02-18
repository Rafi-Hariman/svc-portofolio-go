package repository

import (
	"svc-portofolio-golang/utils/database"
)

func (db *mysqlPortofolioRepository) GenerateID() (id uint64, err error) {
	return database.GenerateID(db.sqlDB)
}

func (db *mysqlPortofolioRepository) GenerateUUID() (uuid string, err error) {
	return database.GenerateUUID(db.sqlDB)
}

func (db *mysqlPortofolioRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
