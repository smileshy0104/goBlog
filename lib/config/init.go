package config

import (
	"embed"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"goBlog/lib/global"
)

func ConfigInit(f embed.FS) *viper.Viper {
	//var config string
	//config = "config.yaml"
	// viper 初始化
	v := viper.New()

	dir := "config"
	direntries, err := f.ReadDir(dir)
	if err != nil {
		panic(fmt.Errorf("读取目录失败：%s\n", err))
	}

	for _, de := range direntries {
		if de.IsDir() {
			continue
		}
		filepath := de.Name()
		v.SetConfigFile(filepath)
		v.SetConfigType("yaml")
	}

	if err = v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("无法读取配置文件：%s\n", err))
	}
	// 监控配置文件
	v.WatchConfig()
	// 配置文件变化事件监听
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	// 配置解析
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(err)
	}

	return v
}
