package usecase

import (
	"svc-portofolio-golang/valueobject"
)

func (auth authUsecase) GetAll(param map[string]interface{}) (response []valueobject.Auth, err error) {
	response, err = auth.mysqlRepository.GetAll(param)
	return
}

func (auth authUsecase) GetOne(param map[string]interface{}) (response valueobject.Auth, err error) {
	response, err = auth.mysqlRepository.GetOne(param)
	return
}

func (auth authUsecase) Count(param map[string]interface{}) (response valueobject.Auth, err error) {
	response, err = auth.mysqlRepository.Count(param)
	return
}

func (auth authUsecase) Store(payload valueobject.AuthPayloadInsert) (valueobject.AuthPayloadInsert, error) {
	for i := range payload.Data {
		payload.Data[i].ID, _ = auth.mysqlRepository.GenerateID()
		payload.Data[i].UUID, _ = auth.mysqlRepository.GenerateUUID()
		payload.Data[i].UserInput = payload.User
	}

	queryConfig, err := auth.ProcessStore(payload)

	if err != nil {
		return payload, err
	}

	return payload, auth.mysqlRepository.Exec(queryConfig...)
}

func (auth authUsecase) Update(payload valueobject.AuthPayloadUpdate) (err error) {
	for i := range payload.Data {
		payload.Data[i].Body.UserUpdate = payload.User
	}

	queryConfig, err := auth.ProcessUpdate(payload)

	if err != nil {
		return
	}

	return auth.mysqlRepository.Exec(queryConfig...)
}

func (auth authUsecase) Delete(payload valueobject.AuthPayloadDelete) (err error) {
	queryConfig, err := auth.ProcessDelete(payload)

	if err != nil {
		return
	}

	return auth.mysqlRepository.Exec(queryConfig...)
}
