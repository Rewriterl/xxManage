package service

import (
	"github.com/Rewriterl/xxManage/v1/internal/app/common/model"
	commonService "github.com/Rewriterl/xxManage/v1/internal/app/common/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"sync"
)

type token struct {
	options *model.TokenOptions
	token   commonService.IToken
	lock    *sync.Mutex
}

var tokenService = &token{
	options: nil,
	token:   nil,
	lock:    &sync.Mutex{},
}

func Token() commonService.IToken {
	if tokenService.token == nil {
		tokenService.lock.Lock()
		defer tokenService.lock.Unlock()
		if tokenService.token == nil {
			ctx := gctx.New()
			err := g.Cfg().MustGet(ctx, "token").Struct(&tokenService.options)
			liberr.ErrIsNil(ctx, err)
			tokenService.token = commonService.GfToken(tokenService.options)
		}
	}
	return tokenService.token
}
