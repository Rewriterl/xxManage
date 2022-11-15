package main

import (
	"github.com/gogf/gf/v2/frame/g"
	_ "xxManage/internal/packed"
	"xxManage/internal/service"

	"github.com/gogf/gf/v2/os/gctx"

	"xxManage/internal/cmd"
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
