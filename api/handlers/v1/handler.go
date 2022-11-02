package v1

import (
	jwthandler "github.com/Asliddin3/exam-api-gateway/api/token"
	"github.com/Asliddin3/exam-api-gateway/config"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	"github.com/Asliddin3/exam-api-gateway/services"
	"github.com/Asliddin3/exam-api-gateway/storage/repo"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	storage        repo.AdminRepo
	cfg            config.Config
	redis          repo.RedisRepo
	jwthandler     jwthandler.JWTHandler
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Storage        repo.AdminRepo
	Redis          repo.RedisRepo
	JwtHandler     jwthandler.JWTHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		storage:        c.Storage,
		redis:          c.Redis,
		jwthandler:     c.JwtHandler,
	}
}
