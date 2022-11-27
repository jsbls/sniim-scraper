package scrapper

import (
	"strings"

	"github.com/everitosan/snimm-scrapper/internal/app/market"
	"github.com/everitosan/snimm-scrapper/internal/app/utils"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/sirupsen/logrus"
)

func InitCatlogues(baseUrl string, repo repository.Repository) error {

	markets, err := repo.Market.GetAll()

	if err != nil || len(markets) == 0 {
		mS := market.NewMarketScrapper(baseUrl)
		markets, err = mS.RequestFromSource()
		if err != nil {
			return err
		}
		repo.Market.Save(markets)
	} else {
		logrus.Println("Markets detected in storage, request is avoid")
	}

	// Scrappers for products
	pS1 := utils.NewOptionSelectScrapper("select[id=ddlProducto]")
	pS2 := utils.NewOptionSelectScrapper("select[name=prod]")
	// Scrappers for product sources
	pSS1 := utils.NewOptionSelectScrapper("select[id=ddlOrigen]")
	// Scrappers for product destinity
	pSD1 := utils.NewOptionSelectScrapper("select[id=ddlDestino]")
	// Scrappers for price presentation
	pP1 := utils.NewOptionSelectScrapper("select[id=ddlPrecios]")

	cS := NewCatalogueScrapper(baseUrl)
	cS.AddScrapper(ProductType, pS1)
	cS.AddScrapper(ProductType, pS2)
	cS.AddScrapper(SourceType, pSS1)
	cS.AddScrapper(DestinyType, pSD1)
	cS.AddScrapper(PricePresentationType, pP1)

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
			// fmt.Printf("\r %d of %d", routinesCount+1, routines)
		}
		routinesCount += 1
	}

	// Finaly we store them
	products := append(pS1.GetOptions(), pS2.GetOptions()...)
	repo.Product.Save(products)

	repo.ProductSource.Save(pSS1.GetOptions())
	repo.ProductDestiny.Save(pSD1.GetOptions())
	repo.PricePresentation.Save(pP1.GetOptions())
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
