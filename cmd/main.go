package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/netbsder/uranus/module"
	"github.com/netbsder/uranus/utils"

	"github.com/gin-gonic/gin"
	"github.com/netbsder/uranus"
)

func main() {
	config := utils.GetAppConfig()

	var engine *gin.Engine

	if config.App.RunMode == "release" {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(gin.Recovery())
	} else {
		engine = gin.Default()
	}

	// 注册服务模块
	uranus.Register("actuator", module.Actuator{Engine: engine})
	utils.GetLogger().Info("Register modules success!", zap.Any("modules", uranus.Modules()))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.App.Port),
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.GetLogger().Fatal("listen: ", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.GetLogger().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.GetLogger().Fatal("Server Shutdown:", zap.Error(err))
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		utils.GetLogger().Info("timeout of 5 seconds.")
	}
	utils.GetLogger().Info("Server exiting")
}
