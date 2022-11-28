package v1

import (
	"fmt"

	jwthandler "github.com/Asliddin3/exam-api-gateway/api/token"
	"github.com/Asliddin3/exam-api-gateway/config"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	"github.com/Asliddin3/exam-api-gateway/services"
	"github.com/Asliddin3/exam-api-gateway/storage/repo"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func GetClaims(h handlerV1, c *gin.Context) (*jwthandler.CustomClaims, error) {

	var (
		claims = jwthandler.CustomClaims{}
	)

	strToken := c.GetHeader("Authorization")
	fmt.Println(h.cfg.SignKey)

	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) { return []byte(h.cfg.SignKey), nil })

	if err != nil {
		fmt.Println(err)
		h.log.Error("invalid access token")
		return nil, err
	}
	rawClaims := token.Claims.(jwt.MapClaims)

	claims.Sub = rawClaims["sub"].(string)
	claims.Exp = rawClaims["exp"].(float64)
	claims.Role = rawClaims["role"].(string)
	// fmt.Printf("%T type of value in map %v",rawClaims["exp"],rawClaims["exp"])
	// fmt.Printf("%T type of value in map %v",rawClaims["iat"],rawClaims["iat"])

	claims.Token = token
	return &claims, nil

}
