package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	wangyiUrl   = "https://www.163.com/"
	bilibiliUrl = "http://www.jj20.com/"
	windowsFile = "C:\\Users\\DSH\\GolandProjects\\testProject\\Img\\"
)

var (
	//创建一个管道
	chSem = make(chan int, 3)
	//创建一个异步等待
	downloadWG sync.WaitGroup
	//创建一个进程锁
	randomMT sync.Mutex
)

/*程序入口*/
func main() {
	//获取wangyiUrl内的所有的图片链接
	imgUrls := getImgUrl(bilibiliUrl)
	//使用切片的方式遍历每一个图片地址，然后异步下载
	for _, imgUrl := range imgUrls {
		getImgAsync(imgUrl)
	}
	//异步下载时，等待downloadWG全部释放（等待所有进程结束）
	downloadWG.Wait()
}

/*异步下载图片*/
func getImgAsync(imgUrl string) {
	//downloadWG+1，代表开启了一个线程
	downloadWG.Add(1)
	//多线程启动
	go func() {
		//time.Sleep(1)
		//管道开始，管道内的代码会被异步处理
		chSem <- 123
		//下载图片的操作
		downLoadImg(imgUrl)
		//管道结束
		<-chSem
		//每次下载结束后，downloadWG-1
		downloadWG.Done()
	}()
}

/*下载图片函数，不多阐述了*/
func downLoadImg(imgUrl string) {
	fmt.Println("开始下载。。。")
	resp, _ := http.Get(imgUrl)
	imgBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	err := ioutil.WriteFile(windowsFile+getRandomName(), imgBytes, 0644)
	if err != nil {
		fmt.Println("下载失败")
	} else {
		fmt.Println("下载成功")
	}
}

/*获得一个文件名字，高并发下载时会导致文件名重复*/
func getRandomName() string {
	//时间戳字符串
	timeStamp := strconv.Itoa(int(time.Now().UnixNano()))
	//获得一个随机字符串
	randomInt := strconv.Itoa(getRandomInt(100, 1000))
	//return完整文件名
	return timeStamp + randomInt + "lll.jpg"
}

/*获得一个随机数，加randomMT锁，相当于同步处理*/
func getRandomInt(start, end int) (ret int) {
	randomMT.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret = start + r.Intn(end-start)
	randomMT.Unlock()
	return
}
