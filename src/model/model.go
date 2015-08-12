package model

//URL实体
type CrawlURL struct {
	Url    string   //url链接
	Depth  int      //深度
	Childs []string //包含的URL
}

type WebEntity struct {
	Title string //标题

	Body *string  //网页内容
	Url  string   //访问链接
	Tags []string //标签

	Source string //来源
}
