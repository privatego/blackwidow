package download

import (
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
)

func Download(crawlUrl *model.CrawlURL) *model.WebEntity {
	fmt.Println("Crawl URL : ", crawlUrl.Url)
	if crawlUrl == nil || crawlUrl.Url == "" {
		return nil
	}

	resp, err := http.Get(crawlUrl.Url)
	if err != nil {
		fmt.Println("Http Get error : ", err.Error())
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Http Resp Body error : ", err.Error())
		return nil
	}
	//fmt.Println("success Body : ", string(body))
	content := string(body)
	webEntity := model.WebEntity{}
	webEntity.Url = crawlUrl.Url
	webEntity.Body = &content
	return &webEntity
}
