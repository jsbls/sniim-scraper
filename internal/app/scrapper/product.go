package scrapper

import (
	"fmt"
	"strings"
	"time"

	"github.com/everitosan/snimm-scrapper/internal"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type SelectOptionsAsMap = map[string]string

type Product struct {
	Id   string
	Name string
}

/*
* This scrapper retrieves all select options from a SubMarket url
 */
type ProductScrapper struct {
	sourceAddress string
}

func NewProductScrapper(snimmBase string) *ProductScrapper {
	return &ProductScrapper{
		sourceAddress: snimmBase,
	}
}

func (ps *ProductScrapper) request(url string, productsChan chan []Product, errorChan chan error) {
	products := make([]Product, 0)
	c := colly.NewCollector()

	c.OnHTML("select[id=ddlProducto]", func(sel *colly.HTMLElement) {
		sel.ForEach("option", func(_ int, option *colly.HTMLElement) {
			p := Product{
				Name: option.Text,
				Id:   option.Attr("value"),
			}
			products = append(products, p)
		})

		productsChan <- products
	})

	c.OnHTML("select[name=prod]", func(sel *colly.HTMLElement) {
		sel.ForEach("option", func(_ int, option *colly.HTMLElement) {
			p := Product{
				Name: option.Text,
				Id:   option.Attr("value"),
			}
			products = append(products, p)
		})

		productsChan <- products
	})

	c.OnHTML("html", func(sel *colly.HTMLElement) {
		time.Sleep(500 * time.Millisecond)
		msg := fmt.Errorf("%w: Request URL: %s", internal.ErrRequestSearchTiemout, url)
		errorChan <- msg
	})

	c.OnError(func(r *colly.Response, err error) {
		msg := fmt.Errorf("%w: Request URL: %s Error: %v", internal.ErrRequest, r.Request.URL, err)
		errorChan <- msg
	})

	logrus.Debug("🕸️  Getting %s", url)

	c.Visit(url)
}

func (ps *ProductScrapper) RequestFromSource(url string) (SelectOptionsAsMap, error) {
	var products []Product
	var err error
	options := make(map[string]string)

	finalUrl := ps.fixUrl(url)

	productsChan := make(chan []Product)
	errorChan := make(chan error)

	go ps.request(finalUrl, productsChan, errorChan)

	select {
	case err = <-errorChan:
		logrus.Debug(err)
	case products = <-productsChan:
		logrus.Debugf("%d Products detected", len(products))
	}

	for _, product := range products {
		options[product.Id] = product.Name
	}
	return options, nil

}

func (ps *ProductScrapper) fixUrl(url string) string {

	finalUrl := url

	if url[0:1] == "/" {
		finalUrl = ps.sourceAddress + url
	}

	if strings.Contains(url, "opcion=") {
		endPart := strings.Split(finalUrl, "opcion=")[1]
		basePart := strings.Split(finalUrl, "?")[0]
		end := strings.LastIndex(basePart, "/")

		finalUrl = basePart[:end+1] + endPart
	}

	return finalUrl
}
