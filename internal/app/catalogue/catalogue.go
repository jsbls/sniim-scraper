package catalogue

import (
	"strings"

	"github.com/everitosan/snimm-scrapper/internal/app/market"
	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
	"github.com/sirupsen/logrus"
)

type optionsKey struct {
	key     string
	options scrapper.SelectOptionsAsMap
}

type requesterParams struct {
	okChan    chan optionsKey
	errorChan chan error
	scrapper  scrapper.Scrapper
	markets   []market.SubMarket
	key       string
}

const keySeparator = "-"

/*
This function have the responsability to get the products with the help of the scrapper and save them as json files
*/
func SaveCataloguesFromMarkets(baseUrl string, markets []market.Market, dstDirectory string) error {
	productScrapper := scrapper.NewProductScrapper(baseUrl)
	writter, err := NewCatalogueWritter(dstDirectory)
	if err != nil {
		return err
	}

	routines := 0
	routinesCount := 0

	okChan := make(chan optionsKey)
	errorChan := make(chan error)

	for _, market := range markets {
		for _, category := range market.Categories {
			for _, subcategory := range category.SubCategories {

				keys := []string{market.Name, category.Name, subcategory.Name}
				key := strings.Join(keys, keySeparator)

				params := requesterParams{
					okChan:    okChan,
					errorChan: errorChan,
					key:       key,
					markets:   subcategory.SubMarkets,
					scrapper:  productScrapper,
				}
				routines = routines + len(subcategory.SubMarkets)
				go request(params)
			}
		}
	}

	logrus.Printf("Waiting for %d responses", routines)

	for routinesCount < routines {
		select {
		case err := <-errorChan:
			logrus.Warn(err)
		case p := <-okChan:
			optionsLen := len(p.options)
			if optionsLen > 0 {
				errW := writter.SaveMapToJsonFile(p.key, p.options,  "product")
				if errW != nil {
					logrus.Error(errW)
				}
			}
			// logrus.Printf("%s has %d", p.key, optionsLen)
		}
		routinesCount += 1
	}

	return nil
}

func request(params requesterParams) {
	for _, mrkt := range params.markets {
		options, err := params.scrapper.RequestFromSource(mrkt.Url)
		key := mrkt.Name + keySeparator + params.key
		key = strings.ReplaceAll(key, " ", "")
		key = strings.ReplaceAll(key, "/", "")
		key = strings.ReplaceAll(key, "\\", "")

		if err != nil {
			params.errorChan <- err
		}

		res := optionsKey{
			key:     key,
			options: options,
		}

		params.okChan <- res
	}
}
