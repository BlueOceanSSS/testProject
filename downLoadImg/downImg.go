package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	wangyiUrl   = "https://www.163.com/"
	bilibiliUrl = "https://www.bilibili.com/"
)

func main() {
	imgUrls := getImgUrl(wangyiUrl)
	for _, imgUrl := range imgUrls {
		downLoadImg(imgUrl)
	}
}

func downLoadImg(imgUrl string) {
	resp, _ := http.Get(imgUrl)
	imgBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	ioutil.WriteFile("/Users/bytedance/go/src/testProject/Img/"+strconv.Itoa(int(time.Now().UnixNano()))+".jpg", imgBytes, 0644)
}
