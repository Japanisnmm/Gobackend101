package servers

import (
	"github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresHandlers"
	"github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresRepositories"
	"github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresUsecases"
	"github.com/Japanisnmm/GoBackend101/modules/monitor/monitorHandlers"
	"github.com/Japanisnmm/GoBackend101/modules/users/usersHandlers"
	"github.com/Japanisnmm/GoBackend101/modules/users/usersRepositories"
	"github.com/Japanisnmm/GoBackend101/modules/users/usersUsecases"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
}

type moduleFactory struct {
	r fiber.Router
	s  *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server,mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
    return &moduleFactory {
		r:    r,
		s:   s,
		mid : mid,
	}

}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
    usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
	

}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)
    m.r.Get("/", handler.HealthCheck)
}

func (m *moduleFactory) UsersModule(){
	repository := usersRepositories.UsersRepository(m.s.db)
    usecase := usersUsecases.UsersUsecase(m.s.cfg,repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)
    // /v1/users/sign
	router := m.r.Group("/users")

	router.Post("/signup",handler.SignUpCustomer)
	router.Post("/signin",handler.SignIn)
    router.Post("/refresh",handler.RefreshPassport)
    router.Post("/signout",handler.SignOut)
	router.Post("/signup-admin",handler.SignUpAdmin)

    router.Get("/:user_id",m .mid.JwtAuth(),m.mid.ParamsCheck() ,handler.GetUserProfile)
    router.Get("/admin/secret",m .mid.JwtAuth(),m.mid.Authorize(2) ,handler.GenerateAdminToken)
}