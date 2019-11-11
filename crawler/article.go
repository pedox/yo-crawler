package crawler

import (
	"github.com/PuerkitoBio/goquery"
	goose "github.com/advancedlogic/GoOse"
)

//Article extended article
type Article struct {
	URL          string `json:"URL"`
	CleanContent string `json:"clean_content"`
	RawContent   string `json:"raw_content"`
	RawHTML      string `json:"raw_html"`
	RootDoc      *goquery.Document
	*goose.Article
}
