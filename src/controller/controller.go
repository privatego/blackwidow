package controller

import (
	"download"
	"extractor"
	//"fmt"
	"model"
	"utils"
)

func CrawlController() {

	var crawlUrl model.CrawlURL = model.CrawlURL{}
	crawlUrl.Url = utils.URL_SEED
	webEntity := download.Download(&crawlUrl)
	extractor.Extract(webEntity)
	//fmt.Println("web content : ", *webEntity.Body)

}
