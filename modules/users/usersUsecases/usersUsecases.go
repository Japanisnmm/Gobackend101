package usersUsecases

import (
	"fmt"

	"github.com/Japanisnmm/GoBackend101/config"
	"github.com/Japanisnmm/GoBackend101/modules/users"
	"github.com/Japanisnmm/GoBackend101/modules/users/usersRepositories"
	"golang.org/x/crypto/bcrypt"
)

type IUsersUsecase interface {
  IsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error)
  GetPassport(req *users.UserCredential) (*users.UserPassport,error)

}

type usersUsecase struct {
	cfg  config.IConfig
	usersRepository usersRepositories.IUsersRepository
}
func UsersUsecase (cfg  config.IConfig,usersRepository usersRepositories.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		cfg: cfg,
		usersRepository: usersRepository,
	}
}

func (u *usersUsecase) IsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error) {
   //hasing pass
	
	if err := req.BcrypHashing(); err != nil {
		return nil, err
	}

//insert user
   result, err := u.usersRepository.InsertUser(req, false)
   if err != nil {
	return nil, err
   }

	return result ,nil
}




func (u *usersUsecase) GetPassport(req *users.UserCredential) (*users.UserPassport,error){
	user, err := u.usersRepository.FindOneUserByEmail(req.Email)
	if err != nil {
		return nil , err
	}
	//compare pass
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password)); err != nil {
		return nil , fmt.Errorf("password is invalid")
	}
	//set passport
	passport := &users.UserPassport{
		User : &users.User{
			Id: user.Id,
			Email: user.Email,
			Username: user.Username,
			RoleId: user.RoleId,
		},
		Token :nil,
	}
	return passport,nil
}