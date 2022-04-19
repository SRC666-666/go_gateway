package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/e421083458/gateway_demo/router"
	"github.com/e421083458/golang_common/lib"
)

func main() {
	//如果configPath：“"./conf/dev/"”为空，则从命令行中`-config=./conf/prod/`中读取
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
