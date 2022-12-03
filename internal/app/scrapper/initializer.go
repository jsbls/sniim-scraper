package scrapper

import (
	"strings"

	"github.com/everitosan/sniim-scrapper/internal/app/form"
	"github.com/everitosan/sniim-scrapper/internal/app/market"
	"github.com/everitosan/sniim-scrapper/internal/app/utils"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
	"github.com/sirupsen/logrus"
)

func GetCatlogues(baseUrl string, repo repository.Repository) error {

	markets, err := repo.Market.GetAll()

	if err != nil || len(markets) == 0 {
		mS := market.NewMarketScrapper(baseUrl)
		markets, err = mS.RequestFromSource()
		if err != nil {
			return err
		}
		repo.Market.Save(markets)
	} else {
		logrus.Printf("%d Markets detected in storage, request is avoid", len(markets))
	}

	okChan := make(chan *form.FormScrapper)
	errorChan := make(chan error)

	routines := 0
	routinesCount := 0

	// Every category will be requested in a go routine, 8 should be created at the moment of writting this code
	for _, mrkt := range markets {
		for _, invtory := range mrkt.Inventories {
			for _, cat := range invtory.Categories {
				keys := []string{mrkt.Name, invtory.Name, cat.Name}
				routines = routines + len(cat.SubCategories)
				go request(baseUrl, cat, okChan, errorChan, keys)
			}
		}
	}

	logrus.Printf("Waiting for %d responses", routines)

	inputs := form.NewInputContainer()
	params := make([]form.FormParams, 0)

	for routinesCount < routines {
		select {
		case err := <-errorChan:
			logrus.Warn(err)
		case formScrapper := <-okChan:
			for selectType, options := range formScrapper.Inputs.GetInputs() {
				inputs.AddOptions(selectType, options)
			}
			if formScrapper.Params.Params != nil {
				params = append(params, formScrapper.Params)
			}
			// fmt.Printf("\r %d of %d", routinesCount+1, routines)
		}
		routinesCount += 1
	}

	// Save to db
	repo.Params.Save(params)

	for selectType, options := range inputs.GetInputs() {
		switch selectType {
		case form.ProductType:
			repo.Product.Save(options)
		case form.DestinyType:
			repo.ProductDestiny.Save(options)
		case form.OriginType:
			repo.ProductSource.Save(options)
		case form.PerPriceType:
			repo.PricePresentation.Save(options)
		case form.WeekType:
			repo.Week.Save(options)
		case form.MonthType:
			repo.Month.Save(options)
		case form.YearType:
			repo.Year.Save(options)
		}
	}

	return nil
}

func request(
	baseUrl string,
	cat market.Catergory,
	okChan chan *form.FormScrapper,
	errorChan chan error,
	keys []string,
) {
	req := utils.NewRequester(baseUrl)
	for _, subCat := range cat.SubCategories {
		formScrapper := form.NewFormScrapper()
		key := append(keys, subCat.Name)
		html, err := req.SyncR(subCat.Url)

		if err != nil {
			errorChan <- err
		} else {
			formScrapper.GetFormInputs(html, strings.Join(key, utils.KeyCatalogueSeparator))
			okChan <- formScrapper
		}

	}
}
