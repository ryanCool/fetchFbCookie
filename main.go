package main


import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
	"github.com/tebeka/selenium/log"
)

const (
	//设置常量 分别设置chromedriver.exe的地址和本地调用端口
	seleniumPath = `/Users/ryan/Downloads/chromedriver`
	port         = 9515
)

func main() {
	//1.开启selenium服务
	//设置selium服务的选项,设置为空。根据需要设置。
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//延迟关闭服务
	defer service.Stop()



	//2.调用浏览器
	//设置浏览器兼容性，我们设置浏览器名称为chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",

	}
	caps.SetLogLevel(log.Performance,log.All)
	//调用浏览器urlPrefix: 测试参考：DefaultURLPrefix = "http://127.0.0.1:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	//延迟退出chrome
	defer wd.Quit()



	//3.对页面元素进行操作
	//获取百度页面
	if err := wd.Get("https://www.facebook.com/ads/library/?active_status=all&ad_type=all&country=ALL&view_all_page_id=341166659618884&sort_data[direction]=desc&sort_data[mode]=relevancy_monthly_grouped"); err != nil {
		panic(err)
	}
	cs,err:=wd.GetCookies()
	if err != nil{
		fmt.Println("err=",err)
	}

	fmt.Println(cs)
	logs,err:=wd.Log(log.Performance)
	if err != nil{
		fmt.Println("err=",err)
	}

	for _,log :=range logs{
		fmt.Println("log=",log.Message)
	}
	//睡眠20秒后退出
	time.Sleep(60 * time.Second)
}


type NetworkLog struct {
	Message struct {
		Method string `json:"method"`
		Params struct {
			AssociatedCookies []interface{} `json:"associatedCookies"`
			Headers           struct {
				Authority      string `json:":authority"`
				Method         string `json:":method"`
				Path           string `json:":path"`
				Scheme         string `json:":scheme"`
				Accept         string `json:"accept"`
				AcceptEncoding string `json:"accept-encoding"`
				AcceptLanguage string `json:"accept-language"`
				ContentLength  string `json:"content-length"`
				ContentType    string `json:"content-type"`
				Origin         string `json:"origin"`
				Referer        string `json:"referer"`
				SecFetchDest   string `json:"sec-fetch-dest"`
				SecFetchMode   string `json:"sec-fetch-mode"`
				SecFetchSite   string `json:"sec-fetch-site"`
				UserAgent      string `json:"user-agent"`
			} `json:"headers"`
			RequestID string `json:"requestId"`
		} `json:"params"`
	} `json:"message"`
	Webview string `json:"webview"`
}