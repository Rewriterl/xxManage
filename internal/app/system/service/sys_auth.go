package service

import (
	"context"
	commonService "github.com/Rewriterl/xxManage/v1/internal/app/common/service"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/consts"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/dao"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/model/dto"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

type IRule interface {
	GetIsMenuList(ctx context.Context) ([]*dto.SysAuthRuleInfoRes, error)
	GetIsButtonList(ctx context.Context) ([]*dto.SysAuthRuleInfoRes, error)
}

type ruleImpl struct {
}

func (r *ruleImpl) GetIsButtonList(ctx context.Context) ([]*dto.SysAuthRuleInfoRes, error) {
	//TODO implement me
	list, err := r.GetMenuList(ctx)
	if err != nil {
		return nil, err
	}
	ruleInfoRes := make([]*dto.SysAuthRuleInfoRes, 0, len(list))
	for _, v := range ruleInfoRes {
		if v.MenuType == 2 {
			ruleInfoRes = append(ruleInfoRes, v)
		}
	}
	return ruleInfoRes, nil
}

func (r *ruleImpl) getMenuListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var v []*dto.SysAuthRuleInfoRes
		err = dao.SysAuthRule.Ctx(ctx).
			Fields(dto.SysAuthRuleInfoRes{}).Order("weigh desc,id asc").Scan(&v)
		liberr.ErrIsNil(ctx, err, "获取菜单数据失败")
		value = v
	})
	return
}

func (r *ruleImpl) GetMenuList(ctx context.Context) (list []*dto.SysAuthRuleInfoRes, err error) {
	cache := commonService.Cache()
	menuCache := cache.GetOrSetFuncLock(ctx, consts.CacheSysAuthMenu, r.getMenuListFromDb, 0, consts.CacheSysAuthTag)
	if menuCache != nil {
		err = gconv.Struct(menuCache, &list)
		liberr.ErrIsNil(ctx, err)
	}
	return
}

func (r *ruleImpl) GetIsMenuList(ctx context.Context) ([]*dto.SysAuthRuleInfoRes, error) {
	//TODO implement me
	list, err := r.GetMenuList(ctx)
	if err != nil {
		return nil, err
	}
	ruleInfoRes := make([]*dto.SysAuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.MenuType == 0 || v.MenuType == 1 {
			ruleInfoRes = append(ruleInfoRes, v)
		}
	}
	return ruleInfoRes, nil
}

var ruleService = ruleImpl{}

func Rule() IRule {
	return &ruleService
}
