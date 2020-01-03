package scraper

import (
	"fmt"
	"net/url"

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
		r.Headers.Set("Cookie", "ll-visitor-id=1cfa9ca4-b91f-4198-a6b5-c5876dbaca86; ai_user=Pxvda|2019-10-25T23:06:28.294Z; __AntiXsrfToken=9f5f0f78e13d499f902b0ec1c2d2db01; Language=1; app_mode=1; Province=QuerÃ©taro; Country=Mexico; GUID=4942d9d2-c340-4d27-96af-13370a6624e7; DG_IID=60D52527-AEED-3647-A74D-278E559C4E58; DG_UID=9BC3601A-6C4E-3A45-B2EA-14D969073963; DG_ZID=75D0580D-9A32-342F-BA61-16B5CDC5BB61; DG_ZUID=915FD68B-730B-335B-9F5D-563C2DF1B3F7; DG_HID=BB8408B6-4362-35CD-8102-26CAC996A75A; DG_SID=177.248.227.16:QqXefcuwApS0e4o8oyxfsmbedE/YDFquIy6ptRvHU10; gig_bootstrap_3_mrQiIl6ov44s2X3j6NGWVZ9SDDtplqV7WgdcyEpGYnYxl7ygDWPQHqQqtpSiUfko=ver2; _ga=GA1.2.1539205010.1578011138; _gid=GA1.2.1680428996.1578011138; _fbp=fb.1.1578011139515.1907941343; cmsdraft=False; TermsOfUseAgreement=2018-06-07; ll-visitor-id=1cfa9ca4-b91f-4198-a6b5-c5876dbaca86; mapZoomLevel=ZoomLevel=3; ViewedListings=5; PreferredMeasurementUnits=2; ai_session=WJ3PE|1578011138136|1578014111847.31; _4c_=fVLNbpwwEH6VyOd4F4MxZm9VKlWREqmt0nNk7GGxwmJkzJI02lOfpq%2FRvljHW3ZJk7Yc0Pibb36%2BmXkmUwMd2bC8kAnjWSqkLC%2FJAzwNZPNMdB%2F%2F%2B%2FgbfUs2pAmhHzbr9TRNKw%2BqDc6vtFpHk8IQVIB1yjLJeCrXGa3AeOd2VLvOOMokZakQdBh3Oxuo2kM3Au297TRQP%2FbgA7kk2hnAQqxclasE3%2BErvmgmog0dtkJ6b9D%2B8O7%2By%2FX7yMyzMk3yhCWr3yoYNoCE3jsz6nAfnvqYcILqYjAP6DCwtxruJ2tCE%2BMzIRa0AbttAsKFiElM7yMFrcmihul11Iyeo8QRvVHddlRbmENv3HYL5uIa50xq1Q6A2Efv9lE4Qp9G8D%2B%2F%2FfgelHfouXJjF%2FwTOm7h0eoIfYbBGuiCVa3zV263A2%2B1al%2Bkq7ybBoi9XjXe7eCiKBF1Q8yiNJoeavD%2ByMDXYEOsvGxwxnDvC0yPcB9HHkW1DkvGKLwXZIMO1nULG7E7b1Gov4XQOIOeO6%2BMjSxsFUcZyQZqNbZxzyaOS7dqGKw2MDwE15PDJXmcb7EseYZGgQcQUKgUPIkfMrw181GSOk8TKCuOh6Y55SIXtDQqoYku0sRUlSxSTuacErMWeBlZnmGSvT3l0MpIBTKnMuM55bVQtGJVThNgRZ7VVVqLmix9MS6xryyf%2B2Ly1FbfzhnZQk7zPJM5kycyP4vo9zNbvJSMV86lfCt5ns886X%2FHluxvsfXbxv5D706DWVZ1np8sRYJykHYen3rlTwtexCVNp6InR1Gm%2FE9qRA6Hwy8%3D")
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
