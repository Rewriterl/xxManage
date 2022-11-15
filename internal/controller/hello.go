package controller

import (
	"context"
	"xxManage/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	//all, err := dao.SysUser.Ctx(ctx).All()
	//res = new(v1.HelloRes)
	return
}
