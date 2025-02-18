package repository

import (
	"svc-portofolio-golang/utils/database"
)

func (db *mysqlAuthRepository) GenerateID() (id uint64, err error) {
	return database.GenerateID(db.sqlDB)
}

func (db *mysqlAuthRepository) GenerateUUID() (uuid string, err error) {
	return database.GenerateUUID(db.sqlDB)
}

func (db *mysqlAuthRepository) Exec(queryConfig ...database.QueryConfig) (err error) {
	err = database.ExecTransaction(db.sqlDB, queryConfig...)
	return
}
