package service

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
)

type captchaImpl struct {
	driver *base64Captcha.DriverString
	store  base64Captcha.Store
}

func (c captchaImpl) GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error) {
	//TODO implement me
	fonts := c.driver.ConvertFonts()
	newCaptcha := base64Captcha.NewCaptcha(fonts, c.store)
	idKeyC, base64stringC, err = newCaptcha.Generate()
	return
}

func (c captchaImpl) VerifyString(id, answer string) bool {
	//TODO implement me
	newCaptcha := base64Captcha.NewCaptcha(c.driver, c.store)
	answer = gstr.ToLower(answer)
	return newCaptcha.Verify(id, answer, true)
}

type ICaptcha interface {
	GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error)
	VerifyString(id, answer string) bool
}

var (
	captcha = captchaImpl{
		driver: &base64Captcha.DriverString{
			Height:          80,
			Width:           240,
			NoiseCount:      50,
			ShowLineOptions: 20,
			Length:          4,
			Source:          "abcdefghjklmnopqrstuvwxyz23456789",
			Fonts:           []string{"chromohv.ttf"},
		},
		store: base64Captcha.DefaultMemStore,
	}
)

func Captcha() ICaptcha {
	return &captcha
}
