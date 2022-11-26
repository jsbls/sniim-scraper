// Retrieves all market entry points from map site

package scrapper

import (
	"fmt"
	"strings"

	"github.com/everitosan/snimm-scrapper/internal"
	"github.com/everitosan/snimm-scrapper/internal/app/market"
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

func (mrktScrapper *marketScrapper) request(okChan chan []market.Market, errChan chan error) {
	c := colly.NewCollector()

	markets := make([]market.Market, 0)
	tmpMarket := market.NewMarket("")
	tmpCategory := market.NewCategory("")
	tmpSubCategory := market.NewSubCategory("")

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
						if tmpSubCategory.IsNotEmpty() {
							tmpCategory.SubCategories = append(tmpCategory.SubCategories, tmpSubCategory)
							tmpSubCategory = market.NewSubCategory("")
						}

						if tmpCategory.IsNotEmpty() {
							tmpMarket.Categories = append(tmpMarket.Categories, tmpCategory)
							tmpCategory = market.NewCategory("")
						}
						markets = append(markets, tmpMarket)
						// Reset vars
						tmpMarket = market.NewMarket("")
					}
					tmpMarket.Name = content
				} else {

					if tmpSubCategory.IsNotEmpty() {
						tmpCategory.SubCategories = append(tmpCategory.SubCategories, tmpSubCategory)
						tmpSubCategory = market.NewSubCategory("")
					}

					if tmpCategory.IsNotEmpty() {
						tmpMarket.Categories = append(tmpMarket.Categories, tmpCategory)
						tmpCategory = market.NewCategory("")
					}

					if tmpMarket.IsNotEmpty() {
						markets = append(markets, tmpMarket)
						tmpMarket = market.NewMarket("")
					}
				}
			case bgcolor == "88aaff": //Category
				if tmpMarket.IsNotEmpty() && content != "" {
					if tmpCategory.IsNotEmpty() {
						if tmpSubCategory.IsNotEmpty() {
							tmpCategory.SubCategories = append(tmpCategory.SubCategories, tmpSubCategory)
							tmpSubCategory = market.NewSubCategory("")
						}

						tmpMarket.Categories = append(tmpMarket.Categories, tmpCategory)
						tmpCategory = market.NewCategory("")
					}

					tmpCategory.Name = content
					tmpCategory.SetType(content)
				}
			case bgcolor != "" && bgcolor != "ffccaa":
				if tmpCategory.IsNotEmpty() && content != "" {
					if tmpSubCategory.IsNotEmpty() {
						tmpCategory.SubCategories = append(tmpCategory.SubCategories, tmpSubCategory)
						tmpSubCategory = market.NewSubCategory("")
					}

					tmpSubCategory.Name = content
				}
			}

			if tdClass != "" && tmpSubCategory.IsNotEmpty() {
				td.ForEach("a", func(_ int, a *colly.HTMLElement) {
					tmpSubMarket := market.SubMarket{
						Name: a.Text,
						Url:  a.Attr("href"),
					}
					tmpSubCategory.SubMarkets = append(tmpSubCategory.SubMarkets, tmpSubMarket)
				})
			}

		})

		if tmpMarket.IsNotEmpty() {
			markets = append(markets, tmpMarket)
			tmpMarket = market.NewMarket("")
		}

		okChan <- markets
	})

	c.OnError(func(r *colly.Response, err error) {
		msg := fmt.Errorf("%w: Request URL: %s Error: %v", internal.ErrRequest, r.Request.URL, err)
		errChan <- msg
	})

	logrus.Printf("🕸️  Getting %s", mrktScrapper.sourceAddress)
	c.Visit(mrktScrapper.sourceAddress)

}

func (mrktScrapper *marketScrapper) RequestFromSource() ([]market.Market, error) {
	markets := make([]market.Market, 0)

	responseOk := make(chan []market.Market)
	responseErr := make(chan error)

	go mrktScrapper.request(responseOk, responseErr)

	select {
	case err := <-responseErr:
		return markets, err
	case markets = <-responseOk:
		logrus.Printf("%d markets retrieved", len(markets))
		return markets, nil
	}
}
