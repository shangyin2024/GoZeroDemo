package svc

import (
	"gozero_demo/internal/common/gormx"
	"gozero_demo/internal/config"
	"gozero_demo/internal/dal/query"
	"gozero_demo/internal/middleware"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
)

type ServiceContext struct {
	Config                        config.Config
	RequestLogMiddleware          rest.Middleware
	ParameterValidationMiddleware rest.Middleware

	DBQuery *query.Query
	DBRaw   *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := initDB(c)
	return &ServiceContext{
		Config:                        c,
		RequestLogMiddleware:          middleware.NewRequestLogMiddleware().Handle,
		ParameterValidationMiddleware: middleware.NewParameterValidationMiddleware().Handle,

		DBQuery: query.Use(dbConn),
		DBRaw:   dbConn,
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

	return dbConn
}
