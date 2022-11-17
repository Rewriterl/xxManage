package service

import (
	"xxManage/internal/app/system/model/dto"
)

type IUser interface {
	GetUserByUserName(username string) (user *dto.LoginUserRes, err error)
}
