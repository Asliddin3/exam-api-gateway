package main

import (
	"fmt"

	"github.com/Asliddin3/exam-api-gateway/api"
	"github.com/Asliddin3/exam-api-gateway/config"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	"github.com/Asliddin3/exam-api-gateway/services"
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v2"

	r "github.com/Asliddin3/exam-api-gateway/storage/redis"
	"github.com/gomodule/redigo/redis"
)

func main() {
	var (
		casbinEnforcer *casbin.Enforcer
	)

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	)
	fmt.Println(psqlString)
	enf, err := gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		log.Error("gorm adapter error", logger.Error(err))
		return
	}

	casbinEnforcer, err = casbin.NewEnforcer(cfg.AuthConfigPath, enf)
	if err != nil {
		log.Error("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", logger.Error(err))
		return
	}

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		Redis:          r.NewRedisRepo(pool),
		CasbinEnforcer: casbinEnforcer,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
