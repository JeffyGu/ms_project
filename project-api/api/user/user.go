package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	common "test.com/project-common"
	"time"
)

type HandlerLogin struct {
}

// GetCaptcha 获取手机验证码
func (HandlerLogin) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	mobile := ctx.PostForm("mobile")
	//2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		ctx.JSON(200, result.Fail(2001, "不合法"))
		return
	}
	//3.生成验证码
	code := "123456"
	//4. 发送验证码
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("调用短信平台发送短信")
		//发送成功 存入redis
		fmt.Println(mobile, code)
	}()
	ctx.JSON(200, result.Success("123456"))
}
