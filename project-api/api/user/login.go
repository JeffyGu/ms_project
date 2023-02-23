package login

import "github.com/gin-gonic/gin"

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	g := r.Group("/project/login")
	h := HandlerLogin{}
	g.POST("/getCaptcha", h.GetCaptcha)
}
