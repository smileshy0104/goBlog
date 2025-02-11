package main

import (
	"embed"
	"goBlog/lib/config"
)

//go:embed config/*
var f embed.FS

func main() {
	//config.ConfigInit(f)
	config.ConfigInit2(f)
	println(config.AppConf.GetString("viewer.Title"))
}
