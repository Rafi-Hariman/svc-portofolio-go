package portofolio

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
	GetAll(param map[string]interface{}) ([]valueobject.Portofolio, error)
	GetOne(param map[string]interface{}) (valueobject.Portofolio, error)
	Count(param map[string]interface{}) (valueobject.Portofolio, error)

	Store(payload valueobject.PortofolioPayloadInsert) (valueobject.PortofolioPayloadInsert, error)
	Update(payload valueobject.PortofolioPayloadUpdate) error
	Delete(payload valueobject.PortofolioPayloadDelete) error

	ProcessStore(payload valueobject.PortofolioPayloadInsert) ([]database.QueryConfig, error)
	ProcessUpdate(payload valueobject.PortofolioPayloadUpdate) ([]database.QueryConfig, error)
	ProcessDelete(payload valueobject.PortofolioPayloadDelete) ([]database.QueryConfig, error)
}
