package crawler

import (
	"github.com/PuerkitoBio/goquery"
	goose "github.com/advancedlogic/GoOse"
)

func (article *Article) GetArticle() (err error) {
	doc, err := goquery.NewDocument(article.URL)
	if err != nil {
		return err
	}
	article.RootDoc = doc

	rawHTML, err := doc.Html()
	if err != nil {
		return err
	}
	article.RawHTML = rawHTML
	return nil
}

func (article *Article) ParseArticle() (err error) {
	g := goose.New()
	if extracted, e := g.ExtractFromRawHTML(article.RawHTML, article.URL); err == nil {
		article.Article = extracted
		article.ParseContent()
		article.ParseDateTime()
		err = e
	}
	return
}
