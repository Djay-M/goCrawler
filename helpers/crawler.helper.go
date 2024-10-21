package helpers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var ProductFinders = [...]string{"/p", "/products", "/product"}

type ResponseObj struct {
	Website      string   `json:"website"`
	ProductLinks []string `json:"productLinks"`
}

func CrawlerWebHelper(website string, responseDataChan chan ResponseObj) {
	var linkData []string
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

	// called before an HTTP request is triggered
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// triggered when the scraper encounters an error
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
		responseDataChan <- ResponseObj{Website: r.Request.URL.String(), ProductLinks: linkData}
	})

	// fired when the server responds
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("a", func(h *colly.HTMLElement) {
		hrefLink := h.Attr("href")

		if len(hrefLink) > 0 {
			for _, productVal := range ProductFinders {
				if strings.Contains(hrefLink, productVal) {
					if strings.HasPrefix(hrefLink, "http") {
						linkData = append(linkData, hrefLink)
					} else {
						urlArr := strings.Split(h.Request.URL.String(), "/")
						parentUrl := urlArr[0] + "//" + urlArr[2] + "/"
						linkData = append(linkData, parentUrl+hrefLink)
					}
					break
				}
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		responseDataChan <- ResponseObj{Website: r.Request.URL.String(), ProductLinks: linkData}
	})

	c.Visit(website)
	return
}
