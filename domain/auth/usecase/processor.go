package usecase

import (
	"svc-portofolio-golang/utils/database"
	"svc-portofolio-golang/valueobject"
)

func (auth authUsecase) ProcessStore(payload valueobject.AuthPayloadInsert) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		/**
		add data you wanted to insert on this interface{}...
		*/
		data := []interface{}{
			[]interface{}{
				x.ID,
				x.UUID,
				x.UserInput,
			},
		}

		/**
		column on data and this line should have same order
		*/
		column := []string{
			"id", "uuid", "user_input",
		}

		queryInsert, err := auth.mysqlRepository.Store(column, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (auth authUsecase) ProcessUpdate(payload valueobject.AuthPayloadUpdate) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"uuid": x.Param.UUID, // add the parameter to update the row
			},
		}
		var data = map[string]interface{}{
			"user_update": x.Body.UserUpdate, // add the data to update the row
		}

		queryUpdate, err := auth.mysqlRepository.Update(param, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryUpdate)
	}
	return
}

func (auth authUsecase) ProcessDelete(payload valueobject.AuthPayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"uuid": x.UUID, // add the parameter to delete the row
			},
		}

		queryDelete, err := auth.mysqlRepository.Delete(param)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryDelete)
	}
	return
}

/// New processor

func (auth authUsecase) ProcessStoreRegister(payload valueobject.AuthLoginPayloadInsert) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {

		data := []interface{}{
			[]interface{}{
				x.ID,
				x.UUID,
				x.UserInput,
				x.Name,
				x.Password,
				x.Email,
			},
		}

		column := []string{
			"id",
			"uuid",
			"user_input",
			"name",
			"password",
			"email",
		}

		queryInsert, err := auth.mysqlRepository.StoreRegister(column, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (auth authUsecase) ProcessDeleteUserLogin(payload valueobject.AuthLoginPayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"uuid": x.UUID,
			},
		}

		queryDelete, err := auth.mysqlRepository.DeleteUserLogin(param)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
