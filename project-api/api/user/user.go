package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"test.com/project-api/pkg/model"
	common "test.com/project-common"
	login_service_v1 "test.com/project-user/pkg/service/login.service.v1"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	code, err := LoginServiceClient.GetCaptcha(c, &login_service_v1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(model.NoLegalMobile, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(code))
}
