package portofolio

import (
	"svc-portofolio-golang/utils/database"
	"svc-portofolio-golang/valueobject"
)

type MysqlRepository interface {
	Exec(...database.QueryConfig) error
	GenerateID() (uint64, error)
	GenerateUUID() (string, error)

	GetAll(param map[string]interface{}) ([]valueobject.Portofolio, error)
	GetOne(param map[string]interface{}) (valueobject.Portofolio, error)
	Count(param map[string]interface{}) (valueobject.Portofolio, error)

	Store(column []string, data []interface{}) (database.QueryConfig, error)
	Update(param map[string]interface{}, data map[string]interface{}) (database.QueryConfig, error)
	Delete(param map[string]interface{}) (database.QueryConfig, error)
}

type OracleRepository interface {
	Exec(...database.QueryConfig) error

	GetAll(param map[string]interface{}) (response []valueobject.Portofolio, err error)
	GetOne(param map[string]interface{}) (response valueobject.Portofolio, err error)

	Store(column []string, data []interface{}) (database.QueryConfig, error)
	Update(param map[string]interface{}, data map[string]interface{}) (database.QueryConfig, error)
	Delete(param map[string]interface{}) (database.QueryConfig, error)
}
