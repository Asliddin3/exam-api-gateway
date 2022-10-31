package api

import (
	_ "github.com/Asliddin3/exam-api-gateway/api/docs" //swag
	v1 "github.com/Asliddin3/exam-api-gateway/api/handlers/v1"
	"github.com/Asliddin3/exam-api-gateway/config"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	"github.com/Asliddin3/exam-api-gateway/services"
	"github.com/Asliddin3/exam-api-gateway/storage/repo"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Redis          repo.RedisRepo
}

// New ...
// @title           Review api
// @version         1.0
// @description     This is reivew api
// @termsOfService  not much usefull

// @contact.name   Asliddin
// @contact.url    https://t.me/asliddindeh
// @contact.email  asliddinvstalim@gmail.com

// @host      localhost:8070
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.Redis,
	})

	api := router.Group("/v1")
	// Customer routers
	api.POST("/customer", handlerV1.CreateCustomer)
	api.GET("/customer/list", handlerV1.GetListCustomers)
	api.PATCH("/customer/update", handlerV1.UpdateCustomer)
	api.DELETE("/customer/delete/:id", handlerV1.DeleteCustomer)
	api.GET("/customer/post/:id", handlerV1.GetCustomerPostById)
	api.GET("/customer/:id", handlerV1.GetCustomerInfo)
	// Post routers
	api.GET("/post/page", handlerV1.ListPostForPage)
	api.GET("/post/:id", handlerV1.GetPost)
	api.POST("/post", handlerV1.CreatePost)
	api.PATCH("/post/update", handlerV1.UpdatePost)
	api.DELETE("/post/delete/:id", handlerV1.DeletePost)
	api.GET("/post/list", handlerV1.GetListPosts)
	// review routers
	api.GET("/review/:id", handlerV1.GetReviewById)
	api.POST("/review", handlerV1.CreateReview)
	api.DELETE("/review/delete/:id", handlerV1.DeleteReview)
	api.POST("/register", handlerV1.Register)
	api.POST("/confirm", handlerV1.GetVerification)
	api.POST("/login", handlerV1.Login)
	api.GET("/post/search/:page/:limit/:parametrs/:orderby", handlerV1.SearchPost)
	// api.GET("/search")
	// register customer

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
