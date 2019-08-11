package internal

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ParserGoquery parse html using goquery
type ParserGoquery struct {
}

// GetLink ...
func (g *ParserGoquery) GetLink(u string) string {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var siteName string
	title := doc.Find("head title").Text()
	doc.Find("head meta").Each(func(i int, s *goquery.Selection) {
		op, _ := s.Attr("property")
		con, _ := s.Attr("content")
		if op == "og:site_name" {
			siteName = con
		}
	})
	return ToLink(strings.TrimSpace(title), siteName, u)
}
