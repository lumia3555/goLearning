package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImgUrls chan string
	chanTask    chan string
	waitGroup   sync.WaitGroup
)

func SpiderPrettyImg(url string) (urls []string) {
	pageStr := GetPageStr(url)
	//fmt.Println(pageStr)

	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果:\n", len(results))
	for _, result := range results {
		url := result[1]
		fmt.Println(url)
		urls = append(urls, url)
	}
	return
}

func GetFilenameFromUrl(url string, dirPath string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	filename = dirPath + filename
	//fmt.Println(fileName)
	return
}

func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	if err != nil {
		HandleError(err, "http.Get(url)")
		return
	}
	defer resp.Body.Close()

	//ioutil.ReadAll(resp.Body)read tcp 192.168.20.50:57178->175.6.244.4:80: wsarecv:
	// An existing connection was forcibly closed by the remote host.
	fBytes, e := ioutil.ReadAll(resp.Body)
	HandleError(e, "ioutil.ReadAll(resp.Body)")
	err = ioutil.WriteFile(filename, fBytes, 0644)
	HandleError(err, "http.Get(url)")
	if err != nil {
		return false
	} else {
		return true
	}
}

func SpiderImgUrls(url string) {
	//获取一个页面下的所有图片链接
	urls := SpiderPrettyImg(url)
	//将所有图片超链接丢入数据管道
	for _, url := range urls {
		chanImgUrls <- url
	}

	//通知当前协程任务完成
	chanTask <- url
	waitGroup.Done()
}

func DownloadImg() {
	for url := range chanImgUrls {
		filename := GetFilenameFromUrl(url, "D:/BJBlockChain1801/demos/W4/day4/img/")
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s下载成功！\n", filename)
		} else {
			fmt.Printf("%s下载失败!!!!!!!!!!!!\n", filename)
		}
	}
	waitGroup.Done()
}

func CheckIfAllSpidersOk() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s完成爬取任务\n", url)
		count++
		if count == 147 {
			close(chanImgUrls)
			break
		}
	}
	waitGroup.Done()
}

func main() {
	//初始化数据管道
	chanImgUrls = make(chan string, 1000000)
	chanTask = make(chan string, 147)

	//爬虫协程：源源不断地往管道中添加图片链接
	for i := 1; i < 148; i++ {
		waitGroup.Add(1)
		go SpiderImgUrls("http://www.umei.cc/tags/meinv_" + strconv.Itoa(i) + ".htm")
	}

	//开辟任务统计协程，如果147个任务全部完成，则关闭数据管道
	waitGroup.Add(1)
	go CheckIfAllSpidersOk()

	//下载协程：源源不断地从管道中读取地址并下载
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}
