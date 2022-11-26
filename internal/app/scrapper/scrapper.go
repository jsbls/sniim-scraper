package scrapper

import (
	"fmt"
	"strings"

	"github.com/everitosan/snimm-scrapper/internal/app/market"
	"github.com/everitosan/snimm-scrapper/internal/app/product"
	"github.com/everitosan/snimm-scrapper/internal/app/utils"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/sirupsen/logrus"
)

func InitCatlogues(baseUrl string, repo repository.Repository) error {

	mS := market.NewMarketScrapper(baseUrl)
	markets, err := mS.RequestFromSource()

	if err != nil {
		return err
	}

	repo.Market.Save(markets)

	// Scrappers for products
	pS1 := product.NewProductScrapper("select[id=ddlProducto]")
	pS2 := product.NewProductScrapper("select[name=prod]")

	cS := NewCatalogueScrapper(baseUrl)
	cS.AddScrapper(ProductType, pS1)
	cS.AddScrapper(ProductType, pS2)

	okChan := make(chan bool)
	errorChan := make(chan error)

	routines := 0
	routinesCount := 0

	for _, mrkt := range markets {
		for _, invtory := range mrkt.Inventories {
			for _, cat := range invtory.Categories {
				keys := []string{mrkt.Name, invtory.Name, cat.Name}
				routines = routines + len(cat.SubCategories)
				go request(cS, cat, okChan, errorChan, keys)
			}
		}
	}

	logrus.Printf("Waiting for %d responses", routines)

	for routinesCount < routines {
		select {
		case err := <-errorChan:
			logrus.Warn(err)
		case p := <-okChan:
			logrus.Debug(p)
			fmt.Printf("\r %d of %d", routinesCount+1, routines)
		}
		routinesCount += 1
	}

	products := append(pS1.GetProducts(), pS2.GetProducts()...)
	repo.Product.Save(products)
	return nil
}

func request(cS *catalogueScrapper, cat market.Catergory, okChan chan bool, errorChan chan error, keys []string) {
	for _, subCat := range cat.SubCategories {
		key := append(keys, subCat.Name)
		err := cS.RequestFromSource(subCat.Url, strings.Join(key, utils.KeyCatalogueSeparator))

		if err != nil {
			errorChan <- err
		}

		okChan <- true
	}
}
