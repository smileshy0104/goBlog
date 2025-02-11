package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// 响应json数据
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "现在是入门教程"
	jsonStr, _ := json.Marshal(indexData)
	// 返回json数据
	w.Write(jsonStr)
}

// 渲染html页面响应
func indexHtml(w http.ResponseWriter, r *http.Request) {
	// 创建模板
	t := template.New("index.html")
	// 获取文件当前路径
	viewPath, _ := os.Getwd()
	// 解析模板
	t, _ = t.ParseFiles(viewPath + "/template/index.html")
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "现在是入门教程"
	// 渲染模板
	err := t.Execute(w, indexData)
	fmt.Println(err)
}

func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
