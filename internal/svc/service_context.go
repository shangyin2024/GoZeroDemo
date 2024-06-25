package svc

import (
	"gozero_demo/internal/common/gormx"
	"gozero_demo/internal/config"
	"gozero_demo/internal/dal/query"
	"gozero_demo/internal/middleware"

	"gorm.io/gorm"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
)

type ServiceContext struct {
	Config                        config.Config
	RequestLogMiddleware          rest.Middleware
	ParameterValidationMiddleware rest.Middleware

	DBQuery     *query.Query
	DBRaw       *gorm.DB
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := initDB(c)
	redisClient := initRedis(c)

	return &ServiceContext{
		Config:                        c,
		RequestLogMiddleware:          middleware.NewRequestLogMiddleware().Handle,
		ParameterValidationMiddleware: middleware.NewParameterValidationMiddleware().Handle,

		DBQuery:     query.Use(dbConn),
		DBRaw:       dbConn,
		RedisClient: redisClient,
	}
}

func initDB(c config.Config) *gorm.DB {
	dbConn, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger:      gormx.NewLog(true),
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	// 添加默认排序规则，防止未指定排序时查询数据出现随机排序
	err = dbConn.Callback().Query().Before("gorm:query").Register("default_order", func(db *gorm.DB) {
		if db.Statement.SQL.String() == "" {
			db.Order("created_at desc,id desc")
		}
	})
	if err != nil {
		panic(err)
	}

	if err := dbConn.Use(otelgorm.NewPlugin(otelgorm.WithDBName("GoZeroDemoGORM"))); err != nil {
		panic(err)
	}

	return dbConn
}

func initRedis(c config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	// 开启 tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	return rdb
}
