package servers

import (
	"github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresHandlers"
	"github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresRepositories"
	"github.com/Japanisnmm/GoBackend101/modules/middlewares/middlewaresUsecases"
	"github.com/Japanisnmm/GoBackend101/modules/monitor/monitorHandlers"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server  *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, server *server,mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
    return &moduleFactory {
		router:r,
		server: server,
		mid : mid,
	}

}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
    usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
	

}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)
    m.router.Get("/", handler.HealthCheck)
}