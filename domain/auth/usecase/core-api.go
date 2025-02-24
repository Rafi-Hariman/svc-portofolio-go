package usecase

import (
	"fmt"
	"regexp"
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

/// new usecase for login

func (auth authUsecase) StoreRegister(payload valueobject.AuthLoginPayloadInsert) (valueobject.AuthLoginPayloadInsert, error) {

	existingEmail, _ := auth.GetAllUserLogin(map[string]interface{}{
		"AND": map[string]interface{}{
			"email": payload.Data[0].Email,
		},
	})

	existingName, _ := auth.GetAllUserLogin(map[string]interface{}{
		"AND": map[string]interface{}{
			"name": payload.Data[0].Name,
		},
	})

	if len(existingName) > 0 {
		return payload, fmt.Errorf("user with the name %s already exists", payload.Data[0].Name)
	}

	if len(existingEmail) > 0 {
		return payload, fmt.Errorf("user with the email %s already exists", payload.Data[0].Email)
	}

	for i := range payload.Data {
		isValidName := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(payload.Data[i].Name)
		if !isValidName {
			return payload, fmt.Errorf("name %s contains invalid characters, only letters are allowed", payload.Data[i].Name)
		}

		payload.Data[i].ID, _ = auth.mysqlRepository.GenerateID()
		payload.Data[i].UUID, _ = auth.mysqlRepository.GenerateUUID()
		payload.Data[i].UserInput = payload.User

		payload.Data[i].Name = payload.Data[i].Name
		payload.Data[i].Email = payload.Data[i].Email
		payload.Data[i].Password = payload.Data[i].Password
	}

	queryConfig, err := auth.ProcessStoreRegister(payload)

	if err != nil {
		return payload, err
	}

	return payload, auth.mysqlRepository.Exec(queryConfig...)
}

func (auth authUsecase) GetAllUserLogin(param map[string]interface{}) (response []valueobject.AuthLogin, err error) {
	response, err = auth.mysqlRepository.GetAllUserLogin(param)
	return
}

func (auth authUsecase) GetOneUserLogin(param map[string]interface{}) (response valueobject.AuthLogin, err error) {
	response, err = auth.mysqlRepository.GetOneUserLogin(param)
	return
}

func (auth authUsecase) DeleteUserLogin(payload valueobject.AuthLoginPayloadDelete) (err error) {
	queryConfig, err := auth.ProcessDeleteUserLogin(payload)

	if err != nil {
		return
	}

	return auth.mysqlRepository.Exec(queryConfig...)
}
