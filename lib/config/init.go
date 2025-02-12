package config

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goBlog/lib/global"
	"os"
)

// ConfigInit 初始化配置文件
func ConfigInit(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = global.ConfigDefaultFile
				case gin.ReleaseMode:
					config = global.ConfigReleaseFile
				case gin.TestMode:
					config = global.ConfigTestFile
				}
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config)
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", global.ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	// viper 初始化
	v := viper.New()

	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
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

// ConfigInitByEmbed 从embed中读取配置文件
func ConfigInitByEmbed(f embed.FS) {
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
