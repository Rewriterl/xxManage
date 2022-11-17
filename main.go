package main

import (
	"github.com/Rewriterl/xxManage/v1/internal/app/system/service"
	"github.com/Rewriterl/xxManage/v1/internal/cmd"
	"github.com/gogf/gf/v2/frame/g"
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
