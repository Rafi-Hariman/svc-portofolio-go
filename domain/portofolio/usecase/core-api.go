package usecase

import (
	"svc-portofolio-golang/valueobject"
)

func (portofolio portofolioUsecase) GetAll(param map[string]interface{}) (response []valueobject.Portofolio, err error) {
	response, err = portofolio.mysqlRepository.GetAll(param)
	return
}

func (portofolio portofolioUsecase) GetOne(param map[string]interface{}) (response valueobject.Portofolio, err error) {
	response, err = portofolio.mysqlRepository.GetOne(param)
	return
}

func (portofolio portofolioUsecase) Count(param map[string]interface{}) (response valueobject.Portofolio, err error) {
	response, err = portofolio.mysqlRepository.Count(param)
	return
}

func (portofolio portofolioUsecase) Store(payload valueobject.PortofolioPayloadInsert) (valueobject.PortofolioPayloadInsert, error) {
	for i := range payload.Data {
		payload.Data[i].ID, _ = portofolio.mysqlRepository.GenerateID()
		payload.Data[i].UUID, _ = portofolio.mysqlRepository.GenerateUUID()
		payload.Data[i].UserInput = payload.User
	}

	queryConfig, err := portofolio.ProcessStore(payload)

	if err != nil {
		return payload, err
	}

	return payload, portofolio.mysqlRepository.Exec(queryConfig...)
}

func (portofolio portofolioUsecase) Update(payload valueobject.PortofolioPayloadUpdate) (err error) {
	for i := range payload.Data {
		payload.Data[i].Body.UserUpdate = payload.User
	}

	queryConfig, err := portofolio.ProcessUpdate(payload)

	if err != nil {
		return
	}

	return portofolio.mysqlRepository.Exec(queryConfig...)
}

func (portofolio portofolioUsecase) Delete(payload valueobject.PortofolioPayloadDelete) (err error) {
	queryConfig, err := portofolio.ProcessDelete(payload)

	if err != nil {
		return
	}

	return portofolio.mysqlRepository.Exec(queryConfig...)
}
