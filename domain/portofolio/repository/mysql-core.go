package repository

import (
	"svc-portofolio-golang/utils/database"
	"svc-portofolio-golang/valueobject"
)

func (db *mysqlPortofolioRepository) GetAll(param map[string]interface{}) (response []valueobject.Portofolio, err error) {
	var result valueobject.Portofolio

	builder := database.New(MYSQL, MYSQL_TABLE, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{"id", "uuid"},
		Where:  param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query, err := db.sqlDB.Query(builder.Result.Query, builder.Result.Value...)

	if err != nil {
		return
	}

	defer query.Close()

	for query.Next() {
		err = query.Scan(
			&result.ID,
			&result.UUID,
		)

		if err != nil {
			return
		}

		response = append(response, result)
	}

	return
}

func (db *mysqlPortofolioRepository) GetOne(param map[string]interface{}) (response valueobject.Portofolio, err error) {
	builder := database.New(MYSQL, MYSQL_TABLE, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{"id", "uuid"},
		Where:  param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query := db.sqlDB.QueryRow(builder.Result.Query, builder.Result.Value...)

	err = query.Scan(
		&response.ID,
		&response.UUID,
	)

	return
}

func (db *mysqlPortofolioRepository) Count(param map[string]interface{}) (response valueobject.Portofolio, err error) {
	builder := database.New(MYSQL, MYSQL_TABLE, SELECT)

	delete(param, "LIMIT")

	builder.OnSelect = database.OnSelect{
		Column: []string{"count(*)"},
		Where:  param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query := db.sqlDB.QueryRow(builder.Result.Query, builder.Result.Value...)

	err = query.Scan(
		&response.Count,
	)

	return
}

func (db *mysqlPortofolioRepository) Store(column []string, data []interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(MYSQL, MYSQL_TABLE, INSERT)

	builder.OnInsert = database.OnInsert{
		Column: column,
		Data:   data,
	}

	err = builder.QueryBuilder()
	return
}

func (db *mysqlPortofolioRepository) Update(param map[string]interface{}, data map[string]interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(MYSQL, MYSQL_TABLE, UPDATE)

	builder.OnUpdate = database.OnUpdate{
		Where: param,
		Data:  data,
	}

	err = builder.QueryBuilder()
	return
}

func (db *mysqlPortofolioRepository) Delete(param map[string]interface{}) (builder database.QueryConfig, err error) {
	builder = database.New(MYSQL, MYSQL_TABLE, DELETE)

	builder.OnDelete = database.OnDelete{
		Where: param,
	}

	err = builder.QueryBuilder()
	return
}
