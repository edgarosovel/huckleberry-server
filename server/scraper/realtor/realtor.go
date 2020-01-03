package realtor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"huckleberry.app/server/models"
)

func Scrape(c *colly.Collector, bookmark *models.Bookmark) {
	c.OnHTML("#btnAcceptTOSModal", func(e *colly.HTMLElement) {
		//
	})
	// Get price
	c.OnHTML("#listingPrice", func(e *colly.HTMLElement) {
		price := strings.ReplaceAll(e.Text, "$", "")
		price = strings.ReplaceAll(price, " ", "")
		price = strings.ReplaceAll(price, ",", "")
		var err error
		bookmark.Price, err = strconv.ParseUint(price, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
	})
	// Get address
	c.OnHTML("#listingAddress", func(e *colly.HTMLElement) {
		bookmark.Address = e.Text
	})
	// Get preview image url
	c.OnHTML("#propimg_1", func(e *colly.HTMLElement) {
		bookmark.ImgURL = e.Attr("src")
	})
}
