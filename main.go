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

	_authHttpDeliver "svc-portofolio-golang/domain/auth/delivery/http"
	_authRepository "svc-portofolio-golang/domain/auth/repository"
	_authUsecase "svc-portofolio-golang/domain/auth/usecase"

	_portofolioHttpDeliver "svc-portofolio-golang/domain/portofolio/delivery/http"
	_portofolioRepository "svc-portofolio-golang/domain/portofolio/repository"
	_portofolioUsecase "svc-portofolio-golang/domain/portofolio/usecase"
)

func main() {
	routers := gin.Default()

	mysql := ConnectMySQL()
	//orcl:= ConnectOracle()

	boilerplateMysqlRepository := _boilerplateRepository.NewMysqlBoilerplateRepository(mysql)
	boilerplateUsecase := _boilerplateUsecase.NewBoilerplateUsecase(boilerplateMysqlRepository)
	_boilerplateHttpDeliver.NewBoilerplateHttpHandler(boilerplateUsecase, routers)

	authMysqlRepository := _authRepository.NewMysqlAuthRepository(mysql)
	authUsecase := _authUsecase.NewAuthUsecase(authMysqlRepository)
	_authHttpDeliver.NewAuthHttpHandler(authUsecase, routers)

	portofolioMysqlRepository := _portofolioRepository.NewMysqlPortofolioRepository(mysql)
	portofolioUsecase := _portofolioUsecase.NewPortofolioUsecase(portofolioMysqlRepository)
	_portofolioHttpDeliver.NewPortofolioHttpHandler(portofolioUsecase, routers)

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
