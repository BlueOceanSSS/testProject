package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	/*
		appid:602142
		appname:just_test_app
		secret:a5c3cb1b239748a0bc618d29e8c6b99a
	*/
	winFileName   = "C:\\Users\\DSH\\GolandProjects\\testProject\\成语APP\\"
	macFileName   = "/Users/bytedance/go/src/testProject/成语APP/"
	showapi_appid = "602142"
	showapi_sign  = "a5c3cb1b239748a0bc618d29e8c6b99a"
	keyword       = "耳"
	page          = "1"
	rows          = "10"
	idiomApi      = "http://route.showapi.com/1196-1?" +
		"showapi_appid=" + showapi_appid +
		"&showapi_sign=" + showapi_sign +
		"&keyword=" + keyword +
		"&page=" + page +
		"&rows=" + rows
)

type Idiom struct {
	Title      string
	Spell      string
	Content    string
	Sample     string
	Derivation string
}

var IdiomsMap map[string]Idiom

func main() {
	IdiomsMap = make(map[string]Idiom)
	jsonStr, _ := getJson(idiomApi)
	ParseJson2Idioms(jsonStr)
	//数据持久化
	dstFile, _ := os.OpenFile(macFileName+"成语大全.json", os.O_CREATE|os.O_WRONLY, 0666)
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(IdiomsMap)
	if err != nil {
		fmt.Println("写出文件失败，err = ", err)
		return
	}
	fmt.Println("写出文件成功")
}

func ParseJson2Idioms(jsonStr string) {
	tempMap := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &tempMap)
	dataSlice := tempMap["showapi_res_body"].(map[string]interface{})["data"].([]interface{})
	for _, value := range dataSlice {
		title := value.(map[string]interface{})["title"].(string)
		idiom := Idiom{Title: title}
		IdiomsMap[title] = idiom
	}
}

func getJson(url string) (jsonStr string, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("访问失败")
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("数据读取失败")
		return
	}
	jsonStr = string(bytes)
	return
}
