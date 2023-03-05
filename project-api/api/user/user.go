package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	common "test.com/project-common"
	"test.com/project-common/logs"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/model"
	"test.com/project-user/pkg/repo"
	"time"
)

type HandlerLogin struct {
	cache repo.Cache
}

func New() *HandlerLogin {
	return &HandlerLogin{
		cache: dao.Rc,
	}
}

// GetCaptcha 获取手机验证码
func (h *HandlerLogin) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	mobile := ctx.PostForm("mobile")
	//2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, result.Fail(model.NoLegalMobile, "不合法"))
		return
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
		err := h.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		defer cancel()
		if err != nil {
			log.Printf("验证码存入redis出错,cause by:%v", err)
		}
	}()
	ctx.JSON(http.StatusOK, result.Success("123456"))
}
