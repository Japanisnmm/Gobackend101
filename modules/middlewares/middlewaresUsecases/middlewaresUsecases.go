package middlewaresUsecases

import (
	"github.com/Japanisnmm/GoBackend101/modules/middlewares"
	"github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
    FindAccessToken(userId, accessToken string) bool
	FindRole()([]*middlewares.Role,error)

}


type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepositories
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMiddlewaresRepositories) IMiddlewaresUsecase {
	return &middlewaresUsecase{
 
		middlewaresRepository: middlewaresRepository,
	}
}

func (u *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}
func(u *middlewaresUsecase) FindRole()([]*middlewares.Role,error){
    roles , err :=  u.middlewaresRepository.FindRole()
	if err != nil {
          return nil, err
	}
	return roles , nil
}