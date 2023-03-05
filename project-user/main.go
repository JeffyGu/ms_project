package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/project-common"
	"test.com/project-user/config"
	"test.com/project-user/router"
)

func main() {
	r := gin.Default()

	////从配置中读取日志配置，初始化日志
	//lc := &logs.LogConfig{
	//	DebugFileName: "E:\\go-code\\ms_project_hub\\logs\\debug\\project-debug.log",
	//	InfoFileName:  "E:\\go-code\\ms_project_hub\\logs\\info\\project-info.log",
	//	WarnFileName:  "E:\\go-code\\ms_project_hub\\logs\\error\\project-error.log",
	//	MaxSize:       500,
	//	MaxAge:        28,
	//	MaxBackups:    3,
	//}
	//err := logs.InitLogger(lc)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	router.InitRouter(r)
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr)
}
