package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"xxManage/internal/app/common/controller"
	"xxManage/internal/app/common/service"
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
