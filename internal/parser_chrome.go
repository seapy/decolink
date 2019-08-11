package internal

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

// ParserChrome parse html using chrome headless
type ParserChrome struct {
}

// GetLink ...
func (g *ParserChrome) GetLink(u string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var title string
	var siteName string
	var ok bool
	err := chromedp.Run(ctx,
		chromedp.Navigate(u),
		chromedp.Text("head > title", &title, chromedp.ByQuery),
		chromedp.AttributeValue("meta[property='og:site_name']", "content", &siteName, &ok, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	return ToLink(strings.TrimSpace(title), siteName, u)
}
