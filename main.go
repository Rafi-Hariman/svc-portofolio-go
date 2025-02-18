package main

import (
	"github.com/gin-gonic/gin"

	"database/sql"
	"log"
	"os"

	"svc-portofolio-golang/utils/database"

	_boilerplateHttpDeliver "svc-portofolio-golang/domain/boilerplate/delivery/http"
	_boilerplateRepository "svc-portofolio-golang/domain/boilerplate/repository"
	_boilerplateUsecase "svc-portofolio-golang/domain/boilerplate/usecase"
)

func main() {
	routers := gin.Default()

	mysql := ConnectMySQL()
	//orcl:= ConnectOracle()

	boilerplateMysqlRepository := _boilerplateRepository.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUsecase.NewBoilerplateUsecase(boilerplateMysqlRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	routers.Run(":" + os.Getenv("PORT"))
}

func ConnectMySQL() (mysql *sql.DB) {
	mysql, err := database.SetupMysqlDatabaseConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}

func ConnectOracle() (oracle *sql.DB) {
	oracle, err := database.SetupOracleDatabaseConnection()

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}
