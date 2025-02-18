package repository

import (
	"database/sql"

	"svc-portofolio-golang/domain/portofolio"
)

/*
*
default const that you should believe and please don't change the value
*/
const MYSQL = "mysql"
const ORACLE = "oracle"

/*
*
default const that you should believe and please don't change the value
*/
const SELECT_DISTINCT = "select-distinct"
const SELECT = "select"
const INSERT = "insert"
const UPDATE = "update"
const DELETE = "delete"

/*
*
two lines below is the const of table name, please change the value of this const
*/
const MYSQL_TABLE = "portofolio"
const ORACLE_TABLE = "portofolio"

/*
*
the struct of mysql
*/
type mysqlPortofolioRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for mysql
*/
func NewMysqlPortofolioRepository(databaseConnection *sql.DB) portofolio.MysqlRepository {
	return &mysqlPortofolioRepository{databaseConnection}
}

/*
*
the struct of oracle
*/
type oraclePortofolioRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for oracle
*/
func NewOraclePortofolioRepository(databaseConnection *sql.DB) portofolio.OracleRepository {
	return &oraclePortofolioRepository{databaseConnection}
}
