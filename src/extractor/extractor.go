package extractor

type extractor interface {
	Fetch(url string) (body string, urls []string, err error)
}

type HtmlExtractor struct {
	body string
	urls []string
}

func (h *HtmlExtractor) Extract(url string) (string, []string, error) {

	if res, ok := (*h)[url]; ok {
		return res.b
	}
}
