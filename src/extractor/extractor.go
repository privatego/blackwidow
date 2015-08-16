package extractor

import (
	"fmt"
	"github.com/opesun/goquery"
	"model"
)

func Extract(webEntity *model.WebEntity) {
	if webEntity == nil || webEntity.Body == nil {
		return
	}
	html := *webEntity.Body
	nodes, err := goquery.ParseString(html)

	if err != nil {
		fmt.Println("goquery ParseString error : ", err)
		panic(err)
	} else {
		title := nodes.Find("title").Text()
		fmt.Println("title : ", title)
	}
}

func ExtractHtml(crawlUrl *model.CrawlURL) {
	fmt.Println(">>> ExtractHtml ... ")
	var doc *Document
	var e error
	if doc, e = NewDocument(crawlUrl.Url); e != nil {
		panic(e.Error())
	}

	doc.Find("tr").Each(func(i int, ss *goquery.Selection) {
		s := *ss.Find("td")
		if s.Length == 6 {
			port, _ := strconv.Atoi(s.Eq(1).Text())
			if port > 1 {
				//var pInfo ProxyInfo
				fmt.Println(">>> port : ", port)
			}
		}
	})

}
