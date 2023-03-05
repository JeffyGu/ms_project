package login_service_v1

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	common "test.com/project-common"
	"test.com/project-common/logs"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/repo"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	//1. 获取参数
	mobile := msg.Mobile
	//2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		return nil, errors.New("手机号不合法")
	}
	//3.生成验证码
	code := "123456"
	//4. 调用短信平台发送验证码（放入go协程）
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("调用短信平台发送短信:info")
		logs.LG.Debug("调用短信平台发送短信:Debug")
		zap.L().Error("调用短信平台发送短信:Error")
		//发送成功 存入redis
		fmt.Println(mobile, code)
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err := ls.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		defer cancel()
		if err != nil {
			log.Printf("验证码存入redis出错,cause by:%v", err)
		}
	}()
	return &CaptchaResponse{}, nil
}
