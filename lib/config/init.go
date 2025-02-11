package config

import (
	"bytes"
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
		if err = v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("无法读取配置文件：%s\n", err))
		}
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

var AppConf *viper.Viper

func ConfigInit2(f embed.FS) {
	AppConf = viper.New()
	//添加根路径
	dir := "config"
	// 获取目录的文件名
	dirEntries, err := f.ReadDir(dir)
	if err != nil {
		fmt.Println("错误", err.Error())
	}
	for _, de := range dirEntries {
		if !de.IsDir() {
			file, _ := f.ReadFile(dir + "/" + de.Name())
			// 如果你的配置文件没有写扩展名，那么这里需要声明你的配置文件属于什么格式
			AppConf.SetConfigType("yaml")
			_ = AppConf.MergeConfig(bytes.NewBuffer(file))
		}
	}
}
