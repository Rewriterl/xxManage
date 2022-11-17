package controller

import (
	"context"
	"github.com/Rewriterl/xxManage/v1/api/v1/common"
	"github.com/Rewriterl/xxManage/v1/internal/app/common/service"
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
