package service

import (
	"context"
	"fmt"
	"github.com/Rewriterl/xxManage/v1/api/v1/system"
	commonService "github.com/Rewriterl/xxManage/v1/internal/app/common/service"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/dao"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/model/dto"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mssola/user_agent"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

type IUser interface {
	GetUserByUserNameAndPassWord(ctx context.Context, req *system.UserLoginReq) (user *dto.LoginUserRes, err error)
	UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
	GetAdminRules(ctx context.Context, userId uint64) (menuList []*dto.UserMenus, permissions []string, err error)
	NotCheckAuthAdminIds(ctx context.Context) *gset.Set
	LoginLog(ctx context.Context, params *dto.LoginLogParams)
}

type userImpl struct {
	CasbinUserPrefix string
}

func (u *userImpl) LoginLog(ctx context.Context, params *dto.LoginLogParams) {
	userAgent := user_agent.New(params.UserAgent)
	browser, _ := userAgent.Browser()
	sysLoginLog := entity.SysLoginLog{
		LoginName:     params.Username,
		Ipaddr:        params.Ip,
		LoginLocation: libUtils.GetCityByIp(params.Ip),
		Browser:       browser,
		Os:            userAgent.OS(),
		Status:        params.Status,
		Msg:           params.Msg,
		LoginTime:     gtime.Now(),
		Module:        params.Module,
	}
	_, err := dao.SysLoginLog.Ctx(ctx).Insert(sysLoginLog)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (u *userImpl) NotCheckAuthAdminIds(ctx context.Context) *gset.Set {
	notCheckAuthIds := []int{1}
	return gset.NewFrom(notCheckAuthIds)
}

func (u *userImpl) GetAllMenus(ctx context.Context) (menus []*dto.UserMenus, err error) {
	var ruleInfoRes []*dto.SysAuthRuleInfoRes
	ruleInfoRes, err = Rule().GetIsMenuList(ctx)
	if err != nil {
		return
	}
	menus = make([]*dto.UserMenus, len(ruleInfoRes))
	for k, v := range ruleInfoRes {
		var menu *dto.UserMenu
		userMenu := u.setMenuData(menu, v)
		menus[k] = &dto.UserMenus{UserMenu: userMenu}
	}
	menus = u.GetMenusTree(menus, 0)
	return
}

func (u *userImpl) setMenuData(menu *dto.UserMenu, entity *dto.SysAuthRuleInfoRes) *dto.UserMenu {
	menu = &dto.UserMenu{
		Id:        entity.Id,
		Pid:       entity.Pid,
		Name:      gstr.CaseCamelLower(gstr.Replace(entity.Name, "/", "_")),
		Component: entity.Component,
		Path:      entity.Path,
		MenuMeta: &dto.MenuMeta{
			Icon:        entity.Icon,
			Title:       entity.Title,
			IsLink:      "",
			IsHide:      entity.IsHide == 1,
			IsKeepAlive: entity.IsCached == 1,
			IsAffix:     entity.IsAffix == 1,
			IsIframe:    entity.IsIframe == 1,
		},
	}
	if menu.MenuMeta.IsIframe || entity.IsLink == 1 {
		menu.MenuMeta.IsLink = entity.LinkUrl
	}
	return menu
}

func (u *userImpl) GetMenusTree(menus []*dto.UserMenus, pid uint) []*dto.UserMenus {
	returnList := make([]*dto.UserMenus, 0, len(menus))
	for _, menu := range menus {
		if menu.Pid == pid {
			menu.Children = u.GetMenusTree(menus, menu.Id)
			returnList = append(returnList, menu)
		}
	}
	return returnList
}

func (u *userImpl) GetAdminRoleIds(ctx context.Context, userId uint64) (roleIds []uint, err error) {
	enforcer, e := commonService.CasbinEnforcer(ctx)
	if e != nil {
		err = e
		return
	}
	//查询关联角色规则
	groupPolicy := enforcer.GetFilteredGroupingPolicy(0, fmt.Sprintf("%u%d", u.CasbinUserPrefix, userId))
	if len(groupPolicy) > 0 {
		roleIds = make([]uint, len(groupPolicy))
		//得到角色id的切片
		for k, v := range groupPolicy {
			roleIds[k] = gconv.Uint(v[1])
		}
	}
	return
}

// GetAdminRole 获取用户角色
func (u *userImpl) GetAdminRole(ctx context.Context, userId uint64, allRoleList []*entity.SysRole) (roles []*entity.SysRole, err error) {
	var roleIds []uint
	roleIds, err = u.GetAdminRoleIds(ctx, userId)
	if err != nil {
		return
	}
	roles = make([]*entity.SysRole, 0, len(allRoleList))
	for _, v := range allRoleList {
		for _, id := range roleIds {
			if id == v.Id {
				roles = append(roles, v)
			}
		}
		if len(roles) == len(roleIds) {
			break
		}
	}
	return
}

func (u *userImpl) GetPermissions(ctx context.Context, roleIds []uint) (userButtons []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 获取角色对应的菜单id
		enforcer, err := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, err)
		menuIds := map[int64]int64{}
		for _, roleId := range roleIds {
			// 查询当前权限
			policy := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, p := range policy {
				i := gconv.Int64(p[1])
				menuIds[i] = i
			}
		}
		Rule()
	})
	return
}

func (u *userImpl) GetAdminRules(ctx context.Context, userId uint64) (menuList []*dto.UserMenus, permissions []string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		isSuperAdmin := false
		u.NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
			if gconv.Uint64(v) == userId {
				isSuperAdmin = true
				return false
			}
			return true
		})
		// 获取用户菜单
		roleList, err2 := Role().GetRoleList(ctx)
		liberr.ErrIsNil(ctx, err2, "获取角色列表失败")
		roles, err2 := u.GetAdminRole(ctx, userId, roleList)
		liberr.ErrIsNil(ctx, err2)
		name := make([]string, len(roles))
		roleIds := make([]uint, len(roles))
		for k, v := range roles {
			name[k] = v.Name
			roleIds[k] = v.Id
		}

		if isSuperAdmin {
			//超管获取所有菜单
			permissions = []string{"*/*/*"}
			menuList, err = u.GetAllMenus(ctx)
			liberr.ErrIsNil(ctx, err)
		} else {
			menuList, err = u.GetAdminMenusByRoleIds(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
			permissions, err = u.GetPermissions(ctx, roleIds)
			liberr.ErrIsNil(ctx, err)
		}
	})
	return
}

func (u *userImpl) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(id).Update(g.Map{
			dao.SysUser.Columns().LastLoginIp:   ip,
			dao.SysUser.Columns().LastLoginTime: gtime.Now(),
		})
		liberr.ErrIsNil(ctx, err, "用户登录信息更新失败")
	})
	return
}

func (u *userImpl) GetUserByUserNameAndPassWord(ctx context.Context, req *system.UserLoginReq) (user *dto.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err = u.getUserByUserName(ctx, req.Username)
		liberr.ErrIsNil(ctx, err)
		liberr.ValueIsNil(user, "用户名或密码错误")
		if libUtils.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
			liberr.ErrIsNil(ctx, gerror.New("用户名或密码错误"))
		}
		if user.UserStatus == 0 {
			liberr.ErrIsNil(ctx, gerror.New("账号已被禁用"))
		}
	})
	return
}

var (
	userService = userImpl{
		CasbinUserPrefix: "u_",
	}
)

func (u *userImpl) getUserByUserName(ctx context.Context, username string) (user *dto.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user = &dto.LoginUserRes{}
		err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().UserName, username).Scan(user)
		liberr.ErrIsNil(ctx, err, "用户名或密码为空")
	})
	return
}

func (u *userImpl) GetAdminMenusByRoleIds(ctx context.Context, roleIds []uint) (menus []*dto.UserMenus, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, err2 := commonService.CasbinEnforcer(ctx)
		liberr.ErrIsNil(ctx, err2)
		menuIds := map[int64]int64{}
		for _, roleId := range roleIds {
			// 查询当前权限
			policy := enforcer.GetFilteredPolicy(0, gconv.String(roleId))
			for _, v := range policy {
				i := gconv.Int64(v[1])
				menuIds[i] = i
			}
		}
		// 获取所有开启的菜单
		allMenus, err2 := Rule().GetIsMenuList(ctx)
		liberr.ErrIsNil(ctx, err2)
		menus = make([]*dto.UserMenus, 0, len(allMenus))
		for _, menu := range allMenus {
			if _, ok := menuIds[gconv.Int64(menu.Id)]; gstr.Equal(menu.Condition, "nocheck") || ok {
				var roleMenu *dto.UserMenu
				roleMenu = u.setMenuData(roleMenu, menu)
				menus = append(menus, &dto.UserMenus{UserMenu: roleMenu})
			}
		}
		menus = u.GetMenusTree(menus, 0)
	})
	return
}

func User() IUser {
	return &userService
}
