package controller

import (
	"context"
	"github.com/Rewriterl/xxManage/v1/api/v1/system"
)

var (
	User = userController{}
)

type userController struct{}

func (c userController) login(ctx context.Context, req *system.UserLoginReq) (userLoginResp system.UserLoginResp, err error) {
	//var (
	//	user        *dto.LoginUserRes
	//	token       string
	//	permissions []string
	//	menuList    []*dto.UserMenus
	//)
	//
	//ip := libUtils.GetClientIp(ctx)
	//userAgent := libUtils.GetUserAgent(ctx)
	//user, err = service.User().GetUserByUserNameAndPassWord(ctx, req)
	//if err != nil {
	//	return
	//}
	//err = service.User().UpdateLoginInfo(ctx, user.Id, ip)
	//if err != nil {
	//	return
	//}
	//// 默认限制多点登录
	//key := gconv.String(user.Id) + "_" + gmd5.MustEncryptString(user.UserName)
	//user.UserPassword = ""
	//user.UserSalt = ""
	//token, err = service.Token().GenerateToken(ctx, key, user)
	//if err != nil {
	//	return
	//}
	//service.User()
	return
}
