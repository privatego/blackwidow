package controller

import (
	"download"
	"extractor"
	"fmt"
	"model"
	"utils"
)

func CrawlController() {
	fmt.Println(" CrawlController ")
	var crawlUrl model.CrawlURL = model.CrawlURL{}
	crawlUrl.Url = utils.URL_SEED
	//webEntity := download.Download(&crawlUrl)
	//extractor.Extract(webEntity)
	extractor.ExtractHtml(&crawlUrl)
	//fmt.Println("web content : ", *webEntity.Body)

}
