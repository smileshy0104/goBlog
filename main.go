package main

import (
	"embed"
	"goBlog/lib/config"
	"goBlog/lib/global"
	"goBlog/lib/initialize"
	"log"
)

//go:embed config/*
var f embed.FS

func main() {
	// 1. 配置管理初始化
	global.GVA_VP = config.ConfigInit() // 加载配置文件（如 config.yaml）
	// 作用：通过 Viper 库读取配置文件，存储到全局变量 GVA_VP

	config.ConfigInitByEmbed(f)
	println(config.AppConf.GetString("viewer.Title"))

	// 2. 日志管理初始化
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 3. 数据库连接
	global.GVA_DB = initialize.Gorm() // 使用 GORM 建立数据库连接
	// 典型实现：根据配置创建 *gorm.DB 实例
}
