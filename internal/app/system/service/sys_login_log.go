package service

import (
	"context"
	"github.com/Rewriterl/xxManage/v1/internal/app/system/model/dto"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/tiger1103/gfast/v3/api/v1/system"
)

type ISysLoginLog interface {
	Invoke(ctx context.Context, data *dto.LoginLogParams)
	List(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error)
	DeleteLoginLogByIds(ctx context.Context, ids []int) (err error)
	ClearLoginLog(ctx context.Context) (err error)
}

type sysLoginLogImpl struct {
	Pool *grpool.Pool
}

func (s sysLoginLogImpl) Invoke(ctx context.Context, data *dto.LoginLogParams) {
	//TODO implement me
	s.Pool.Add(ctx, func(ctx context.Context) {
		//写入日志数据
		User().LoginLog(ctx, data)
	},
	)
	return
}

func (s sysLoginLogImpl) List(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (s sysLoginLogImpl) DeleteLoginLogByIds(ctx context.Context, ids []int) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s sysLoginLogImpl) ClearLoginLog(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}

var (
	sysLoginLogService = sysLoginLogImpl{
		Pool: grpool.New(100),
	}
)

func SysLoginLog() ISysLoginLog {
	return &sysLoginLogService
}
