package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var (
	rePhone = `(1[3456789]\d)(\d{4})(\d{4})`
	reEmail = `[1-9]\d{4,}@qq.com`
)

const (
	phoneUrl = "https://www.jihaoba.com/shouji/all/"
	emailUrl = "https://club.autohome.com.cn/bbs/thread/ec014de6c2b25cca/80053303-21.html"

)


func HanderErr(err error,when string){
	if err != nil{
		fmt.Println(when,err)
		os.Exit(1)
	}
}


func main()  {
	content := getResp(emailUrl)
	//fmt.Println(content)
	re := regexp.MustCompile(reEmail)
	//数字是几就是匹配的次数
	strings := re.FindAllStringSubmatch(content, -1)
	//fmt.Println(strings)
	for _, phone := range strings {
		fmt.Println(phone)
	}
	fmt.Println(len(strings))
}

func getResp(url string) string {
	resp, err := http.Get(url)
	HanderErr(err, "http.Get(\""+url+"\")")
	bytes, _ := ioutil.ReadAll(resp.Body)
	content := string(bytes)
	return content
}