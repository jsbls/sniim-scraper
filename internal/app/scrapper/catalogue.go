package scrapper

import (
	"fmt"
	"strings"

	"github.com/everitosan/snimm-scrapper/internal"
	"github.com/everitosan/snimm-scrapper/internal/app/utils"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type FilterType int16

const (
	ProductType FilterType = iota
	SourceType
	DestinyType
	PricePresentationType
)

type catalogueScrapper struct {
	baseAddress string
	Filters     map[FilterType][]utils.FilterScrapper
}

func NewCatalogueScrapper(base string) *catalogueScrapper {
	return &catalogueScrapper{
		baseAddress: base,
		Filters:     make(map[FilterType][]utils.FilterScrapper),
	}
}

func (cS *catalogueScrapper) AddScrapper(fType FilterType, scrapper utils.FilterScrapper) {
	_, exists := cS.Filters[fType]

	if !exists {
		cS.Filters[fType] = make([]utils.FilterScrapper, 0)
	}
	cS.Filters[fType] = append(cS.Filters[fType], scrapper)
}

func (cS *catalogueScrapper) RequestFromSource(url string, keyJoined string) error {

	var err error
	finalUrl := cS.fixUrl(url)

	successChan := make(chan *colly.HTMLElement)
	errorChan := make(chan error)

	go cS.request(finalUrl, successChan, errorChan)

	select {
	case err = <-errorChan:
		logrus.Debug(err)
	case html := <-successChan:
		for _, filters := range cS.Filters {
			for _, filter := range filters {
				filter.Extract(url, html, keyJoined)
			}
		}
	}

	return err
}

func (cS *catalogueScrapper) request(url string, successChan chan *colly.HTMLElement, errorChan chan error) {
	c := colly.NewCollector()

	c.OnHTML("html", func(html *colly.HTMLElement) {
		successChan <- html
	})

	c.OnError(func(r *colly.Response, err error) {
		msg := fmt.Errorf("%w: Request URL: %s Error: %v", internal.ErrRequest, r.Request.URL, err)
		errorChan <- msg
	})

	logrus.Debug("🕸️  Getting %s", url)

	c.Visit(url)

}

// Help fix the url when iframe detected
func (cS *catalogueScrapper) fixUrl(url string) string {

	finalUrl := url

	if url[0:1] == "/" {
		finalUrl = cS.baseAddress + url
	}

	if strings.Contains(url, "opcion=") {
		endPart := strings.Split(finalUrl, "opcion=")[1]
		basePart := strings.Split(finalUrl, "?")[0]
		end := strings.LastIndex(basePart, "/")

		finalUrl = basePart[:end+1] + endPart
	}

	return finalUrl
}
