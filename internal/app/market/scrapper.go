package market

import (
	"fmt"
	"strings"

	"github.com/everitosan/snimm-scrapper/internal"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type marketScrapper struct {
	sourceAddress string
}

func NewMarketScrapper(baseUrl string) *marketScrapper {
	return &marketScrapper{
		sourceAddress: baseUrl + "/nuevo/mapa.asp",
	}
}

func (mrktScrapper *marketScrapper) request(okChan chan []Market, errChan chan error) {
	c := colly.NewCollector()

	markets := make([]Market, 0)
	tmpMarket := NewMarket("")
	tmpInventory := NewInventory("")
	tmpCategory := NewCategory("")

	c.OnHTML("table", func(table *colly.HTMLElement) {
		table.ForEach("td", func(_ int, td *colly.HTMLElement) {
			bgcolor := td.Attr("bgcolor")
			tdClass := td.Attr("class")
			content := td.Text
			// Clean spaces
			content = strings.TrimSpace(content)
			// Clean ": "
			content = strings.ReplaceAll(content, ":", "")

			switch {
			case bgcolor == "228833": // Market
				if strings.Contains(strings.ToLower(content), "mercado") {
					if tmpMarket.IsNotEmpty() {
						if tmpCategory.IsNotEmpty() {
							tmpInventory.Categories = append(tmpInventory.Categories, tmpCategory)
							tmpCategory = NewCategory("")
						}

						if tmpInventory.IsNotEmpty() {
							tmpMarket.Inventories = append(tmpMarket.Inventories, tmpInventory)
							tmpInventory = NewInventory("")
						}
						markets = append(markets, tmpMarket)
						// Reset vars
						tmpMarket = NewMarket("")
					}
					tmpMarket.Name = content
				} else {

					if tmpCategory.IsNotEmpty() {
						tmpInventory.Categories = append(tmpInventory.Categories, tmpCategory)
						tmpCategory = NewCategory("")
					}

					if tmpInventory.IsNotEmpty() {
						tmpMarket.Inventories = append(tmpMarket.Inventories, tmpInventory)
						tmpInventory = NewInventory("")
					}

					if tmpMarket.IsNotEmpty() {
						markets = append(markets, tmpMarket)
						tmpMarket = NewMarket("")
					}
				}
			case bgcolor == "88aaff": //Category
				if tmpMarket.IsNotEmpty() && content != "" {
					if tmpInventory.IsNotEmpty() {
						if tmpCategory.IsNotEmpty() {
							tmpInventory.Categories = append(tmpInventory.Categories, tmpCategory)
							tmpCategory = NewCategory("")
						}

						tmpMarket.Inventories = append(tmpMarket.Inventories, tmpInventory)
						tmpInventory = NewInventory("")
					}

					tmpInventory.Name = content
					tmpInventory.SetType(content)
				}
			case bgcolor != "" && bgcolor != "ffccaa":
				if tmpInventory.IsNotEmpty() && content != "" {
					if tmpCategory.IsNotEmpty() {
						tmpInventory.Categories = append(tmpInventory.Categories, tmpCategory)
						tmpCategory = NewCategory("")
					}

					tmpCategory.Name = content
				}
			}

			if tdClass != "" && tmpCategory.IsNotEmpty() {
				td.ForEach("a", func(_ int, a *colly.HTMLElement) {
					tmpSubMarket := SubCategory{
						Name: strings.TrimSpace(a.Text),
						Url:  a.Attr("href"),
					}
					tmpCategory.SubCategories = append(tmpCategory.SubCategories, tmpSubMarket)
				})
			}

		})

		if tmpMarket.IsNotEmpty() {
			markets = append(markets, tmpMarket)
			tmpMarket = NewMarket("")
		}

		okChan <- markets
	})

	c.OnError(func(r *colly.Response, err error) {
		msg := fmt.Errorf("%w: Request URL: %s Error: %v", internal.ErrRequest, r.Request.URL, err)
		errChan <- msg
	})

	logrus.Printf("🍱 Scrapping Markets %s", mrktScrapper.sourceAddress)
	c.Visit(mrktScrapper.sourceAddress)

}

func (mrktScrapper *marketScrapper) RequestFromSource() ([]Market, error) {
	markets := make([]Market, 0)

	responseOk := make(chan []Market)
	responseErr := make(chan error)

	go mrktScrapper.request(responseOk, responseErr)

	select {
	case err := <-responseErr:
		return markets, err
	case markets = <-responseOk:
		logrus.Debugf("🍱 %d markets retrieved", len(markets))
		return markets, nil
	}
}
