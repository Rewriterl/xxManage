package service

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"time"
)

type sysInit struct{}

func SysInit() ISysInit {
	return &sysInit{}
}

type ISysInit interface {
	LoadConfigFile() (err error)
}

func (s *sysInit) LoadConfigFile() (err error) {
	var (
		ctx = context.TODO()
	)
	fileName, _ := libUtils.ParseFilePath("/manifest/config/config.yaml")
	c1, err := g.Cfg(fileName).Get(ctx, "database.default")
	if err != nil {
		return
	}
	dbConfig := c1.MapStrVar()
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Link:             dbConfig["link"].String(),
				Debug:            dbConfig["debug"].Bool(),
				Charset:          dbConfig["charset"].String(),
				DryRun:           dbConfig["dryRun"].Bool(),
				MaxIdleConnCount: dbConfig["maxIdle"].Int(),
				MaxOpenConnCount: dbConfig["maxOpen"].Int(),
				MaxConnLifeTime:  dbConfig["maxLifetime"].Duration() * time.Second,
			},
		},
	})
	c2, err := g.Cfg(fileName).Get(ctx, "redis.default")
	if err != nil {
		return
	}
	redisConfig := c2.Map()
	err = gredis.SetConfigByMap(redisConfig)
	if err != nil {
		return
	}
	return
}
