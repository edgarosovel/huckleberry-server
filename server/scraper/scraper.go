package scraper

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"

	"huckleberry.app/server/models"
	"huckleberry.app/server/scraper/realtor"
	"huckleberry.app/server/scraper/remax"
	"huckleberry.app/server/scraper/wattyway"
)

func GetBookmarkInfo(bookmark *models.Bookmark) error {
	fmt.Println(bookmark)
	URL, err := url.Parse(bookmark.URL)
	if err != nil { // bad URL
		return err
	}

	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	// c := colly.NewCollector()

	prepare(c, &err)
	getTitle(c, bookmark)
	getDescription(c, bookmark)

	switch URL.Hostname() {
	case "www.remax.ca":
		remax.Scrape(c, bookmark)
	case "www.wattyway.ca":
		wattyway.Scrape(c, bookmark)
	case "www.realtor.ca":
		realtor.Scrape(c, bookmark)

	}

	c.Visit(bookmark.URL)

	return err
}

func prepare(c *colly.Collector, err *error) {
	// Error handler
	c.OnError(func(r *colly.Response, request_err error) {
		err = &request_err
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Before making a request
	c.OnRequest(func(r *colly.Request) {
		// Cookie for realtor.ca
		if os.Getenv("REALTOR_COOKIE") != "" {
			r.Headers.Set("Cookie", os.Getenv("REALTOR_COOKIE"))
		}
		fmt.Println("Scraping", r.URL.String())
	})
}

func getTitle(c *colly.Collector, bookmark *models.Bookmark) {
	c.OnHTML("title", func(e *colly.HTMLElement) {
		bookmark.Title = e.Text
	})
	c.OnHTML("meta[property='og:title']", func(e *colly.HTMLElement) {
		if e.Attr("content") != "" {
			fmt.Println("meta[property='og:title']", e.Attr("content"))
			bookmark.Title = e.Attr("content")
		}
	})
}

func getDescription(c *colly.Collector, bookmark *models.Bookmark) {
	c.OnHTML("meta[property='og:description']", func(e *colly.HTMLElement) {
		if e.Attr("content") != "" {
			bookmark.Description = e.Attr("content")
		}
	})
	c.OnHTML("meta[name='Description']", func(e *colly.HTMLElement) {
		if e.Attr("content") != "" {
			bookmark.Description = e.Attr("content")
		}
	})
	c.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
		if e.Attr("content") != "" {
			bookmark.Description = e.Attr("content")
		}
	})
}
