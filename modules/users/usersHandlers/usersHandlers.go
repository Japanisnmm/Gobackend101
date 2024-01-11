package usersHandlers

import (


	"github.com/Japanisnmm/GoBackend101/config"
	"github.com/Japanisnmm/GoBackend101/modules/entities"
	"github.com/Japanisnmm/GoBackend101/modules/users"
	"github.com/Japanisnmm/GoBackend101/modules/users/usersUsecases"
	"github.com/gofiber/fiber/v2"
)
type userHandlersErrCode string
const (
	signUpCustomeErr  userHandlersErrCode = "users-001"
    signInErr  userHandlersErrCode = "users-002"

)

type IUsersHandler interface {
	SignUpCustomer(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
}

type usersHandler struct {
	cfg config.IConfig
	usersUsecase usersUsecases.IUsersUsecase
}

func UsersHandler(cfg config.IConfig, usersUsecase usersUsecases.IUsersUsecase) IUsersHandler {
     return &usersHandler{
		cfg: cfg,
		usersUsecase: usersUsecase,
	 }
}

func (h *usersHandler) SignUpCustomer(c *fiber.Ctx) error {
	//req body parser
	req := new(users.UserRegisterReq)
	if err := c.BodyParser(req); err != nil{
		return entities.NewResponse(c).Error(
          fiber.ErrBadRequest.Code,
		  string(signUpCustomeErr),
          err.Error(),

		).Res()
	}

	//Email valid
	if  !req.IsEmail() {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signUpCustomeErr),
		    "email pattern is invalid",
  
		  ).Res()
	}
	//Insert 
	result , err := h.usersUsecase.IsertCustomer(req)
	if err != nil {
		switch err.Error(){
		case "username has been used ":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpCustomeErr),
				err.Error(),
	  
			  ).Res()
		case "email has been used ":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpCustomeErr),
				err.Error(),
	  
			  ).Res()

		default:
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(signUpCustomeErr),
				err.Error(),
	  
			  ).Res()

		}
	}
	
	return entities.NewResponse(c).Success(fiber.StatusCreated, result).Res()
}

func (h *usersHandler) SignIn(c *fiber.Ctx) error  {
	req := new(users.UserCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signInErr),
			err.Error(),
		).Res()
	}
	
	
	passport, err := h.usersUsecase.GetPassport(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signInErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}