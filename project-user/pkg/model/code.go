package model

import (
	"test.com/project-common/errs"
)

var (
	RedisError    = errs.NewError(999, "redis错误")
	DBError       = errs.NewError(998, "DB错误")
	NoLegalMobile = errs.NewError(10102001, "手机号不合法")
	CaptchaError  = errs.NewError(10102002, "验证码不正确")
	EmailExist    = errs.NewError(10102003, "邮箱已存在")
	AccountExist  = errs.NewError(101020034, "账号已存在")
	MobileExist   = errs.NewError(101020034, "手机号已存在")
)
