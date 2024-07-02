package templatedemo

import (
	"fmt"
	"html/template"
	"net/http"

	mlog "github.com/mike504110403/goutils/log"
)

func HelloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// 解析範本
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	mlog.Error(fmt.Sprintf("Get home directory failed: %s", err.Error()))
	// 	return
	// }
	templatePath := "./template/template_demo.html"
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		mlog.Error(fmt.Sprintf("Parse template failed: %s", err.Error()))
		return
	}
	data := struct {
		Title1 string
		Title2 string
		Title3 string
	}{
		Title1: "第一行",
		Title2: "第二行",
		Title3: "第三行",
	}
	// 執行範本
	err = tmpl.Execute(w, data)
	if err != nil {
		mlog.Error(fmt.Sprintf("Execute template failed: %s", err.Error()))
		return
	}
}
