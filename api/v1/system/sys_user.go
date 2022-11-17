package system

import (
	"github.com/gogf/gf/v2/frame/g"
	dto2 "xxManage/internal/system/model/dto"
)

// 对外提供的输入输出结构定义

type UserLoginReq struct {
	g.Meta     `path:"/login" tags:"登录" method:"post" summary:"用户登录"`
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}
type UserLoginResp struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *dto2.LoginUserRes `json:"userInfo"`
	Token       string             `json:"token"`
	MenuList    []*dto2.UserMenus  `json:"menuList"`
	Permissions []string           `json:"permissions"`
}
