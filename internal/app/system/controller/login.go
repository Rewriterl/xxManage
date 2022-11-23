package controller

import (
	"context"
	"github.com/Rewriterl/xxManage/v1/api/v1/system"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/model/dto"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/service"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/library/libUtils"
)

var (
	Login = loginController{}
)

type loginController struct{}

func (l loginController) Login(ctx context.Context, req *system.UserLoginReq) (userLoginRes *system.UserLoginRes, err error) {
	var (
		user        *dto.LoginUserRes
		token       string
		permissions []string
		menuList    []*dto.UserMenus
	)

	ip := libUtils.GetClientIp(ctx)
	//userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.User().GetUserByUserNameAndPassWord(ctx, req)
	if err != nil {
		// 登录失败日志
		return
	}
	err = service.User().UpdateLoginInfo(ctx, user.Id, ip)
	if err != nil {
		return
	}
	// 默认限制多点登录
	key := gconv.String(user.Id) + "_" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncrypt(gtime.Date())
	user.UserPassword = ""
	user.UserSalt = ""
	token, err = service.Token().GenerateToken(ctx, key, user)
	if err != nil {
		return
	}
	menuList, permissions, err = service.User().GetAdminRules(ctx, user.Id)
	if err != nil {
		return
	}
	userLoginRes = &system.UserLoginRes{
		UserInfo:    user,
		Token:       token,
		MenuList:    menuList,
		Permissions: permissions,
	}
	return
}
