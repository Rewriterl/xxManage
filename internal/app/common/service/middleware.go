package service

import "github.com/gogf/gf/v2/net/ghttp"

type IMiddleware interface {
	MiddlewareCORS(r *ghttp.Request)
}

type middlewareImpl struct{}

func (m middlewareImpl) MiddlewareCORS(r *ghttp.Request) {
	//TODO implement me
	options := r.Response.DefaultCORSOptions()
	r.Response.CORS(options)
	r.Middleware.Next()
}

var middleService = middlewareImpl{}

func Middleware() IMiddleware {
	return &middleService
}
