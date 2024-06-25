package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Redis struct {
		Addr     string
		Password string
		DB       int
	}

	Auth struct { // jwt鉴权配置
		AccessSecret string // jwt密钥
		AccessExpire int64  // 有效期，单位：秒
	}

	Mysql struct {
		DataSource   string
		OutPath      string
		ModelPkgPath string
	}
}
