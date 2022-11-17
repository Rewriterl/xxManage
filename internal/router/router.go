package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonRouter "xxManage/internal/app/common/router"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		commonRouter.BindController(group)
	})
}
