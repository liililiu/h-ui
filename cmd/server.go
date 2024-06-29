package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/dao"
	"h-ui/middleware"
	"h-ui/model/constant"
	"h-ui/router"
	"h-ui/service"
	"h-ui/util"
	"net/http"
	"os"
)

func runServer() {
	defer releaseResource()
	initFile()
	middleware.InitLog()
	dao.InitSqliteDB(port)
	middleware.InitCron()
	service.InitHysteria2()
	r := gin.Default()
	router.Router(r)
	for {
		webServer, err := service.NewServer()
		if err != nil {
			panic(err)
		}
		if err := webServer.StartServer(r); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func releaseResource() {
	dao.CloseSqliteDB()
	service.ReleaseHysteria2()
}

func initFile() {
	var dirs = []string{constant.LogDir, constant.SqliteDBDir, constant.BinDir, constant.ExportPathDir}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				panic(fmt.Sprintf("%s create err: %v", item, err))
			}
		}
	}
}
