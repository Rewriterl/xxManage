package router

import (
	commonRouter "github.com/Rewriterl/xxManage/v1/internal/app/common/router"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		commonRouter.BindController(group)
	})
}
