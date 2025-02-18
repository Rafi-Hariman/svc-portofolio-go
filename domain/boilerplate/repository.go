package boilerplate

import (
	"svc-portofolio-golang/utils/database"
	"svc-portofolio-golang/valueobject"
)

type MysqlRepository interface {
	Exec(...database.QueryConfig) error
	GenerateID() (uint64, error)
	GenerateUUID() (string, error)

	GetAll(param map[string]interface{}) ([]valueobject.Boilerplate, error)
	GetOne(param map[string]interface{}) (valueobject.Boilerplate, error)
	Count(param map[string]interface{}) (valueobject.Boilerplate, error)

	Store(column []string, data []interface{}) (database.QueryConfig, error)
	Update(param map[string]interface{}, data map[string]interface{}) (database.QueryConfig, error)
	Delete(param map[string]interface{}) (database.QueryConfig, error)
}

type OracleRepository interface {
	Exec(...database.QueryConfig) error

	GetAll(param map[string]interface{}) (response []valueobject.Boilerplate, err error)
	GetOne(param map[string]interface{}) (response valueobject.Boilerplate, err error)

	Store(column []string, data []interface{}) (database.QueryConfig, error)
	Update(param map[string]interface{}, data map[string]interface{}) (database.QueryConfig, error)
	Delete(param map[string]interface{}) (database.QueryConfig, error)
}
