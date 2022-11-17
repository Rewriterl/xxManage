package router

import (
	"github.com/Rewriterl/xxManage/v1/internal/app/common/controller"
	"github.com/Rewriterl/xxManage/v1/internal/app/common/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().MiddlewareCORS)
		group.Group("/pub", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.Captcha,
			)
		})
	})
}
