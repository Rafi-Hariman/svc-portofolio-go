package auth

import (
	"svc-portofolio-golang/utils/database"
	"svc-portofolio-golang/valueobject"
)

/*
*
why there's only one usecase interface while there can more than one repository interface?...
... because, at DDD (Domain Design Driven), there's only one set of usecase and...
... the function name inside the usecase should be unique and represent the business process...
... tl;dr: function name is telling what exactly are they doing.
*/
type Usecase interface {
	GetAll(param map[string]interface{}) ([]valueobject.Auth, error)
	GetOne(param map[string]interface{}) (valueobject.Auth, error)
	Count(param map[string]interface{}) (valueobject.Auth, error)

	Store(payload valueobject.AuthPayloadInsert) (valueobject.AuthPayloadInsert, error)
	Update(payload valueobject.AuthPayloadUpdate) error
	Delete(payload valueobject.AuthPayloadDelete) error

	ProcessStore(payload valueobject.AuthPayloadInsert) ([]database.QueryConfig, error)
	ProcessUpdate(payload valueobject.AuthPayloadUpdate) ([]database.QueryConfig, error)
	ProcessDelete(payload valueobject.AuthPayloadDelete) ([]database.QueryConfig, error)

	/// new usecase for login

	DeleteUserLogin(payload valueobject.AuthLoginPayloadDelete) error
	GetOneUserLogin(param map[string]interface{}) (valueobject.AuthLogin, error)
	GetAllUserLogin(param map[string]interface{}) ([]valueobject.AuthLogin, error)
	StoreRegister(payload valueobject.AuthLoginPayloadInsert) (valueobject.AuthLoginPayloadInsert, error)

	ProcessStoreRegister(payload valueobject.AuthLoginPayloadInsert) ([]database.QueryConfig, error)
}
