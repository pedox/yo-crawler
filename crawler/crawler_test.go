package crawler

import (
	"testing"

	goose "github.com/advancedlogic/GoOse"
)

func TestHTMLParsed(t *testing.T) {

	newsURL := "https://news.detik.com/berita/d-4779088/rio-capella-heran-pembukaan-kongres-undang-anies-bukan-jokowi-ini-kata-nasdem?tag_from=wp_hl_judul&_ga=2.125698890.1964938698.1573258684-1888490124.1573258684"
	article := Article{
		URL: newsURL,
	}

	if article.GetArticle() != nil {
		t.Error("Get news gagal")
	}

	if article.ParseArticle() != nil {
		t.Error("Parsing data gagal")
	}

	t.Log("Judul:", article.Title)
	t.Log("Konten:", article.RawContent)
	t.Log("Konten:", article.CleanContent)
	t.Log("Tag:", article.Tags)
	t.Log("Tanggal Publish:", article.PublishDate)

}

func TestGoOse(t *testing.T) {
	g := goose.New()
	d, err := g.ExtractFromURL("https://www.nytimes.com/2019/11/10/us/politics/joe-biden-ukraine.html?action=click&module=Top%20Stories&pgtype=Homepage")

	if err == nil {
		t.Log("title", d.Title)
		t.Log("content", d.CleanedText)
	}

}
