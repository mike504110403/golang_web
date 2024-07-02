package cookiedemo

import (
	"fmt"
	"net/http"

	"github.com/mike504110403/goutils/log"
)

func TestCookieHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("testCookie")
	if err != nil {
		log.Error(fmt.Sprintf("Get cookie failed: %s", err.Error()))
		return
	}
	log.Info(fmt.Sprintf("Get cookie: %s", c.Value))

	cookie := &http.Cookie{
		Name:   "testCookie",
		Value:  "testCookieValue",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/", // cookie有效路径
	}
	http.SetCookie(w, cookie)

	// 應在資料返回前設定cookie
	w.Write([]byte("Set cookie success"))
}
