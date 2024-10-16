package main

import (
	//client "golang_web/httpclientdemo"
	//template "golang_web/templatedemo"
	cookiedemo "golang_web/cookieDemo"
	sqldemo "golang_web/sqlDemo"
	"net/http"

	mlog "github.com/mike504110403/goutils/log"
)

func Init() {
	// 初始化log
	mlog.Init(mlog.Config{
		EnvMode: mlog.EnvMode("debug"),
		LogType: mlog.LogType("console"),
	})
	// 初始化sql
	sqldemo.Init(nil)
}

func main() {
	Init()
	//client.PutDemo()
	//http.HandleFunc("/", template.HelloHandlerFunc)
	http.HandleFunc("/", cookiedemo.TestCookieHandler)
	http.ListenAndServe(":8086", nil)
}
