package usecase

import (
	"svc-portofolio-golang/domain/portofolio"
)

/**
the struct of usecases
*/
type portofolioUsecase struct {
	mysqlRepository portofolio.MysqlRepository
}

/**
the initiator function for usecase
*/
func NewPortofolioUsecase(mysql portofolio.MysqlRepository) portofolio.Usecase {
	return &portofolioUsecase{
		mysqlRepository: mysql,
	}
}
