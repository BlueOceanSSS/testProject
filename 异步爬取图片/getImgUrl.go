package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reImg = `<img[\s\S]+?src="(http[\s\S]+?)"`
)

func getImgUrl(url string) []string {
	html := getHtml(url)
	re := regexp.MustCompile(reImg)
	imgs := re.FindAllStringSubmatch(html, -1)
	imgUrls := make([]string, 0)
	for _, url := range imgs {
		imgUrls = append(imgUrls, url[1])
	}
	return imgUrls
}

func getHtml(url string) string {
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	html := string(bytes)
	return html
}
