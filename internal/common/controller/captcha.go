package controller

import (
	"context"
	"xxManage/api/v1/common"
	"xxManage/internal/common/service"
)

type captchaController struct{}

var Captcha = captchaController{}

func (c captchaController) Captcha(ctx context.Context, req *common.CaptchaReq) (userCaptchaRes *common.CaptchaRes, err error) {
	var (
		verifyImgString, base64stringC string
	)
	verifyImgString, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	userCaptchaRes = &common.CaptchaRes{
		Key: verifyImgString,
		Img: base64stringC,
	}
	return
}
