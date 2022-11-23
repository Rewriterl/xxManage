package router

import (
	"github.com/Rewriterl/xxManage/v1/internal/app/common/service"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().MiddlewareCORS)
		group.Bind(
			controller.Login,
		)
	})
}
