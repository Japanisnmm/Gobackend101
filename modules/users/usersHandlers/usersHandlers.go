package usersHandlers

import (
	"strings"

	"github.com/Japanisnmm/GoBackend101/config"
	"github.com/Japanisnmm/GoBackend101/modules/entities"
	"github.com/Japanisnmm/GoBackend101/modules/users"
	"github.com/Japanisnmm/GoBackend101/modules/users/usersUsecases"
	gobackendauth "github.com/Japanisnmm/GoBackend101/pkg/GoBackendauth"
	"github.com/gofiber/fiber/v2"
)
type userHandlersErrCode string
const (
	signUpCustomeErr  userHandlersErrCode = "users-001"
    signInErr  userHandlersErrCode = "users-002"
	refreshPassportErr  userHandlersErrCode = "users-003"
	signOutErr  userHandlersErrCode = "users-004"
	signUpAdminErr  userHandlersErrCode = "users-005"
    generateAdminTokenErr  userHandlersErrCode = "users-006"
	getUserProfileErr  userHandlersErrCode = "users-007"
)

type IUsersHandler interface {
	SignUpCustomer(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	RefreshPassport(c *fiber.Ctx) error
	SignOut(c *fiber.Ctx) error
	SignUpAdmin(c *fiber.Ctx) error
	GenerateAdminToken(c *fiber.Ctx) error
	GetUserProfile(c *fiber.Ctx) error
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
func (h  *usersHandler)  SignUpAdmin(c *fiber.Ctx) error {
    
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

func (h *usersHandler) GenerateAdminToken(c *fiber.Ctx) error {
    adminToken ,err :=  gobackendauth.NewGobackendAuth(
		gobackendauth.Admin,
		 h.cfg.Jwt(),
		nil,
		)
		if err != nil  {
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string (generateAdminTokenErr),
				err.Error(),
			).Res()
		}
	
	
	
	
	
	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		&struct {
			Token string `json:"token"`
		}{
		    Token:adminToken.SignToken(),	
		},
	).Res()
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

func (h *usersHandler) RefreshPassport(c *fiber.Ctx) error  {
	req := new(users.UserRefreshCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(refreshPassportErr),
			err.Error(),
		).Res()
	}
	
	
	passport, err := h.usersUsecase.RefreshPassport(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signInErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (h *usersHandler) SignOut(c *fiber.Ctx) error {
	req := new(users.UserRemoveCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signOutErr),
			err.Error(),
		).Res()
	}

	if err := h.usersUsecase.DeleteOauth(req.OauthId); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signOutErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, nil).Res()
}


func (h *usersHandler) GetUserProfile(c *fiber.Ctx) error {
   // set params
   userId := strings.Trim(c.Params("user_id")," ")

	// Get profile
result, err := h.usersUsecase.GetUserProfile(userId)
if err != nil {
	switch err.Error(){
	case "get user failed: sql: no rows in result set":
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(getUserProfileErr),
			err.Error(),
		 ).Res()
	default:
		return entities.NewResponse(c).Error(
           fiber.ErrInternalServerError.Code,
		   string(getUserProfileErr),
		   err.Error(),
		).Res()
	}
}
	
	
	return entities.NewResponse(c).Success(fiber.StatusOK,result).Res()
}

