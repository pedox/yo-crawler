package main

import (
	"net/http"
	"github.com/pedox/yo-crawler/crawler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type PostInput struct {
	Url string `json:"url"`
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/api/fetcher", func(c echo.Context) (err error) {
		f := new(PostInput)
		if err := c.Bind(f); err != nil {
			return err
		}

		// newsURL := "https://news.detik.com/berita/d-4779088/rio-capella-heran-pembukaan-kongres-undang-anies-bukan-jokowi-ini-kata-nasdem?tag_from=wp_hl_judul&_ga=2.125698890.1964938698.1573258684-1888490124.1573258684"
		article := crawler.Article{
			URL: f.Url,
		}

		if article.GetArticle() != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "pengambilan berita gagal")
			// .Error("Get news gagal")
		}

		if article.ParseArticle() != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "parsing data gagal")
			// t.Error("Parsing data gagal")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"title":   article.Title,
			"content": article.CleanContent,
			"raw":     article.RawContent,
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8885"))
}
