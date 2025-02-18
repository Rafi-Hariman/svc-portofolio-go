package usecase

import (
	"svc-portofolio-golang/domain/auth"
)

/*
*
the struct of usecases
*/
type authUsecase struct {
	mysqlRepository auth.MysqlRepository
}

/*
*
the initiator function for usecase
*/
func NewAuthUsecase(mysql auth.MysqlRepository) auth.Usecase {
	return &authUsecase{
		mysqlRepository: mysql,
	}
}
