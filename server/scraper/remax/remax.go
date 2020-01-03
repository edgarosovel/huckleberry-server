package remax

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"huckleberry.app/server/models"
)

func Scrape(c *colly.Collector, bookmark *models.Bookmark) {
	firstPriceVisited := false
	firstAddressVisited := false
	firstImageVisited := false

	// Get price
	c.OnHTML(".price", func(e *colly.HTMLElement) {
		if firstPriceVisited {
			return
		}
		price := strings.ReplaceAll(e.Text, "$", "")
		price = strings.ReplaceAll(price, " ", "")
		price = strings.ReplaceAll(price, ",", "")

		var err error
		bookmark.Price, err = strconv.ParseUint(price, 10, 64)
		if err != nil {
			fmt.Println(err)
			// not able to parse
		}
		firstPriceVisited = true
	})
	// Get address
	c.OnHTML(".address", func(e *colly.HTMLElement) {
		if firstAddressVisited {
			return
		}
		bookmark.Address = e.Text
		firstAddressVisited = true
	})
	c.OnHTML(".address-supplementary", func(e *colly.HTMLElement) {
		// there's only 1 .address-supplementary
		bookmark.Address += e.Text
	})
	// Get preview image url
	c.OnHTML("img[data-srcset]", func(e *colly.HTMLElement) {
		if firstImageVisited {
			return
		}
		if e.Attr("data-srcset") != "" {
			images := strings.Split(e.Attr("data-srcset"), ", ") // data-srcset has an array of urls with different sizes
			bookmark.ImgURL = images[1]                          // second url is 320 px big
			firstImageVisited = true
		}
	})
}
