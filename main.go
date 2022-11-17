package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"xxManage/internal/system/cmd"
	"xxManage/internal/system/service"

	"github.com/gogf/gf/v2/os/gctx"
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
