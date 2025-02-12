package main

import (
	"context"
	"embed"
	"fmt"
	"goBlog/lib/config"
	"goBlog/lib/global"
	"goBlog/lib/initialize"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed config/*
var f embed.FS

func main() {
	// 1. 配置管理初始化
	global.GVA_VP = config.ConfigInit() // 加载配置文件（如 config.yaml）
	// 作用：通过 Viper 库读取配置文件，存储到全局变量 GVA_VP

	config.ConfigInitByEmbed(f)
	//println(config.AppConf.GetString("viewer.Title"))

	// 2. 日志管理初始化
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 3. 数据库连接
	global.GVA_DB = initialize.Gorm() // 使用 GORM 建立数据库连接
	// 典型实现：根据配置创建 *gorm.DB 实例

	// 4. 路由初始化
	r := initialize.InitRouter()

	////定时任务
	//serviceJob := business.NewTaskService()
	//serviceJob.StartScheduledTasks()

	// 5、监听http服务
	srv := &http.Server{
		Addr:    config.AppConf.GetString("address"),
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("web服务器 启动失败 : %v\n", err)
		}
	}()

	log.Println("web 服务器启动成功： ", config.AppConf.GetString("address"))

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var done = make(chan struct{}, 1)
	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("web server shutdown error: %v", err)
		} else {
			fmt.Println("web server shutdown ok")
		}
		done <- struct{}{}
	}()

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		fmt.Println("web server shutdown timeout")
	case <-done:
	}

	fmt.Println("program exit ok")
}
