package httpclientdemo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	mlog "github.com/mike504110403/goutils/log"
)

func GetDemo() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		mlog.Error(fmt.Sprintf("Get request failed: %s", err.Error()))
		return
	}
	if resp == nil {
		mlog.Error("Response is nil")
		return
	}
	defer resp.Body.Close()
	mlog.Info(fmt.Sprintf("Response status: %s", resp.Status))
	// Read response
	closer := resp.Body
	bytes, err := io.ReadAll(closer)
	if err != nil {
		mlog.Error(fmt.Sprintf("Read response failed: %s", err.Error()))
		return
	} else {
		if len(bytes) > 0 {
			mlog.Info(fmt.Sprintf("Response: %s", string(bytes)))
		} else {
			mlog.Info("Response is empty")
		}
	}
}

func PostDemo() {
	url := "https://www.shirdon.com/comment/add"
	body := "{\"userId\":1,\"articleId\":1,\"comment\":\"This is a comment\"}"
	res, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))
	if err != nil {
		mlog.Error(fmt.Sprintf("Post request failed: %s", err.Error()))
		return
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		mlog.Error(fmt.Sprintf("Read response failed: %s", err.Error()))
		return
	} else {
		if len(b) > 0 {
			mlog.Info(fmt.Sprintf("Response: %s", string(b)))
		} else {
			mlog.Info("Response is empty")
		}
	}
}

func PutDemo() {
	client := &http.Client{}
	url := "https://www.shirdon.com/comment/update"
	payload := strings.NewReader("{\"userId\":1,\"articleId\":1,\"comment\":\"This is a updated comment\"}")
	req, err := http.NewRequest(http.MethodPut, url, payload)
	if err != nil {
		mlog.Error(fmt.Sprintf("Create request failed: %s", err.Error()))
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		mlog.Error(fmt.Sprintf("Put request failed: %s", err.Error()))
		return
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		mlog.Error(fmt.Sprintf("Read response failed: %s", err.Error()))
		return
	} else {
		if len(b) > 0 {
			mlog.Info(fmt.Sprintf("Response: %s", string(b)))
		} else {
			mlog.Info("Response is empty")
		}
	}
}

func DeleteDemo() {
	client := &http.Client{}
	url := "https://www.shirdon.com/comment/delete"
	payload := strings.NewReader("{\"userId\":1,\"articleId\":1,\"comment\":\"This is a updated comment\"}")
	req, err := http.NewRequest(http.MethodDelete, url, payload)
	if err != nil {
		mlog.Error(fmt.Sprintf("Create request failed: %s", err.Error()))
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		mlog.Error(fmt.Sprintf("Delete request failed: %s", err.Error()))
		return
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		mlog.Error(fmt.Sprintf("Read response failed: %s", err.Error()))
		return
	} else {
		if len(b) > 0 {
			mlog.Info(fmt.Sprintf("Response: %s", string(b)))
		} else {
			mlog.Info("Response is empty")
		}
	}
}
