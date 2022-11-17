package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"xxManage/internal/app/system/cmd"
	"xxManage/internal/app/system/service"
)

func init() {
	ctx := gctx.New()
	err := service.SysInit().LoadConfigFile()
	if err != nil {
		g.Log().Panic(ctx, err)
	}
}

func main() {
	cmd.Main.Run(gctx.New())
}
