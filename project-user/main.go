package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, "webcenter", ":80")
}
