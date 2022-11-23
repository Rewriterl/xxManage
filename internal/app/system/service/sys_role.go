package service

import (
	"context"
	commonService "github.com/Rewriterl/xxManage/v1/internal/app/common/service"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/consts"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/dao"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

type IRole interface {
	GetRoleList(ctx context.Context) (list []*entity.SysRole, err error)
}

type roleImpl struct{}

var roleService = roleImpl{}

// 从数据库获取所有角色
func (r *roleImpl) getRoleListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var roles []*entity.SysRole
		//从数据库获取
		err = dao.SysRole.Ctx(ctx).
			Order(dao.SysRole.Columns().ListOrder + " asc," + dao.SysRole.Columns().Id + " asc").
			Scan(&roles)
		liberr.ErrIsNil(ctx, err, "获取角色失败")
		value = roles
	})
	return
}

func (r *roleImpl) GetRoleList(ctx context.Context) (list []*entity.SysRole, err error) {
	//TODO implement me
	cache := commonService.Cache()
	//从缓存获取
	iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysRole, r.getRoleListFromDb, 0, consts.CacheSysAuthTag)
	if iList != nil {
		err = gconv.Struct(iList, &list)
	}
	return
}

func Role() IRole {
	return &roleService
}
