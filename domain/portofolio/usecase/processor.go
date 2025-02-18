package usecase

import (
	"svc-portofolio-golang/utils/database"
	"svc-portofolio-golang/valueobject"
)

func (portofolio portofolioUsecase) ProcessStore(payload valueobject.PortofolioPayloadInsert) (queryConfig []database.QueryConfig, err error) {
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

		queryInsert, err := portofolio.mysqlRepository.Store(column, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryInsert)
	}
	return
}

func (portofolio portofolioUsecase) ProcessUpdate(payload valueobject.PortofolioPayloadUpdate) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Data {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"uuid": x.Param.UUID, // add the parameter to update the row
			},
		}
		var data = map[string]interface{}{
			"user_update": x.Body.UserUpdate, // add the data to update the row
		}

		queryUpdate, err := portofolio.mysqlRepository.Update(param, data)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryUpdate)
	}
	return
}

func (portofolio portofolioUsecase) ProcessDelete(payload valueobject.PortofolioPayloadDelete) (queryConfig []database.QueryConfig, err error) {
	for _, x := range payload.Param {
		var param = map[string]interface{}{
			"AND": map[string]interface{}{
				"uuid": x.UUID, // add the parameter to delete the row
			},
		}

		queryDelete, err := portofolio.mysqlRepository.Delete(param)

		if err != nil {
			return queryConfig, err
		}

		queryConfig = append(queryConfig, queryDelete)
	}
	return
}
