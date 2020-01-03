package wattyway

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"huckleberry.app/server/models"
)

func Scrape(c *colly.Collector, bookmark *models.Bookmark) {
	firstSlideVisited := false
	// Get price
	c.OnHTML(".listing-price", func(e *colly.HTMLElement) {
		price := strings.ReplaceAll(e.Text, e.ChildText("small"), "")
		price = strings.ReplaceAll(price, "$", "")
		price = strings.ReplaceAll(price, " ", "")
		price = strings.ReplaceAll(price, ",", "")
		price = strings.ReplaceAll(price, ".00", "")
		var err error
		bookmark.Price, err = strconv.ParseUint(price, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
	})
	// Get address
	c.OnHTML(".listing-address", func(e *colly.HTMLElement) {
		bookmark.Address = e.Text
	})
	// Get preview image url
	c.OnHTML(".slides", func(e *colly.HTMLElement) {
		if firstSlideVisited {
			return
		}
		bookmark.ImgURL = e.ChildAttr("img", "src")
		firstSlideVisited = true
	})
}
