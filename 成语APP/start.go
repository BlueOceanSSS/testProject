package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	/*
		appid:602142
		appname:just_test_app
		secret:a5c3cb1b239748a0bc618d29e8c6b99a
	*/
	showapi_appid = "602142"
	showapi_sign  = "a5c3cb1b239748a0bc618d29e8c6b99a"
	keyword       = "火"
	page          = "1"
	rows          = "10"
	idiomApi      = "https://route.showapi.com/1196-1?" +
		"showapi_appid=" + showapi_appid +
		"&showapi_sign=" + showapi_sign +
		"&keyword=" + keyword +
		"&page=" + page +
		"&rows=" + rows
)

func main() {
	resp, err := http.Get(idiomApi)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("访问失败")
		return
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("数据读取失败")
	}
	respStr := string(bytes)
	fmt.Println(respStr)
}
