package main

import (
	//client "golang_web/httpclientdemo"
	template "golang_web/templatedemo"
	"net/http"

	mlog "github.com/mike504110403/goutils/log"
)

func Init() {
	mlog.Init(mlog.Config{
		EnvMode: mlog.EnvMode("debug"),
		LogType: mlog.LogType("console"),
	})
}

func main() {
	Init()
	//client.PutDemo()
	http.HandleFunc("/", template.HelloHandlerFunc)
	http.ListenAndServe(":8086", nil)
}
