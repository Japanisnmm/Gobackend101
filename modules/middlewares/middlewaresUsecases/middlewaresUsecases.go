package middlewaresUsecases

import "github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepositories
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMiddlewaresRepositories) IMiddlewaresUsecase {
	return &middlewaresUsecase{
 
		middlewaresRepository: middlewaresRepository,
	}
}