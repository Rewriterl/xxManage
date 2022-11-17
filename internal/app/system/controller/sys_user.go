package controller

import (
	"context"
	"xxManage/api/v1/system"
)

var (
	User = user{}
)

type user struct{}

func (c user) login(ctx context.Context, req system.UserLoginReq) (userLoginResp system.UserLoginResp, err error) {
	//var (
	//	user        *dto.LoginUserRes
	//	token       string
	//	permissions []string
	//	menuList    []*dto.UserMenus
	//)
	//
	//ip:=libUtils.GetClientIp(ctx)
	//userAgent:=libUtils.GetUserAgent(ctx)
	return
}
