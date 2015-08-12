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
