package service

import (
	"github.com/Rewriterl/xxManage/v1/internal/app/system/model/dto"
)

type IUser interface {
	GetUserByUserName(username string) (user *dto.LoginUserRes, err error)
}
