package repository

import (
	"database/sql"

	"svc-portofolio-golang/domain/auth"
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
const MYSQL_TABLE = "auth"
const ORACLE_TABLE = "auth"

/*
*
the struct of mysql
*/
type mysqlAuthRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for mysql
*/
func NewMysqlAuthRepository(databaseConnection *sql.DB) auth.MysqlRepository {
	return &mysqlAuthRepository{databaseConnection}
}

/*
*
the struct of oracle
*/
type oracleAuthRepository struct {
	sqlDB *sql.DB
}

/*
*
the initiator function for oracle
*/
func NewOracleAuthRepository(databaseConnection *sql.DB) auth.OracleRepository {
	return &oracleAuthRepository{databaseConnection}
}
