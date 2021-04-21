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
	page          = "1"
	rows          = "10"
)

type Idiom struct {
	Title      string
	Spell      string
	Content    string
	Sample     string
	Derivation string
}

var (
	IdiomsMap map[string]Idiom
)

func main() {
	IdiomsMap = make(map[string]Idiom)
	jsonStr, _ := getJson(false, "四")
	ParseJson2Idioms(jsonStr)
	storage()
	for k, v := range IdiomsMap {
		jsonStr, _ := getJson(true, k)
		idiom := v
		ParseJson2Idiom(jsonStr, idiom)
	}
	storage()
}

func storage() {
	//数据持久化
	dstFile, _ := os.OpenFile(macFileName+"成语大全.txt", os.O_CREATE|os.O_WRONLY, 0666)
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(IdiomsMap)
	if err != nil {
		fmt.Println("写出文件失败，err = ", err)
	}
}

func ParseJson2Idiom(jsonStr string, idiom Idiom) {
	tempMap := make(map[string]interface{})
	var title string
	json.Unmarshal([]byte(jsonStr), &tempMap)
	dataSlice := tempMap["showapi_res_body"].(map[string]interface{})["data"].(map[string]interface{})
	//fmt.Println(dataSlice)
	for key, value := range dataSlice {
		switch key {
		case "spell":
			idiom.Spell = value.(string)
		case "content":
			idiom.Content = value.(string)
		case "derivation":
			idiom.Derivation = value.(string)
		case "samples":
			idiom.Sample = value.(string)
		case "title":
			title = value.(string)
		}
	}
	IdiomsMap[title] = idiom
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

func getJson(isAccurate bool, keyword string) (jsonStr string, err error) {
	idiomApi := "http://route.showapi.com/1196-1?" +
		"showapi_appid=" + showapi_appid +
		"&showapi_sign=" + showapi_sign +
		"&keyword=" + keyword +
		"&page=" + page +
		"&rows=" + rows
	url := idiomApi
	if isAccurate {
		idiomApiAccurate := "http://route.showapi.com/1196-2?showapi_appid=" + showapi_appid + "&showapi_sign=" + showapi_sign + "&keyword=" + keyword
		url = idiomApiAccurate
	}
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
