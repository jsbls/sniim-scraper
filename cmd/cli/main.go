package main

import (
	"log"

	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
	"github.com/everitosan/snimm-scrapper/internal/config"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.LoadConfig()

	if config.DEBUG {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug Log level")
	}

	// Repositories
	marketRepo, _ := repository.NewMarketFileRepository(config.CATALOGUE_SRC)
	productRepo, _ := repository.NewProductRepository(config.CATALOGUE_SRC)
	rContainer := repository.Repository{
		Market:  marketRepo,
		Product: productRepo,
	}
	// Retrieve and save with repositories
	err := scrapper.InitCatlogues(config.SNIIM_ADDR, rContainer)

	if err != nil {
		log.Fatal(err)
	}

	// products := cS.Filters[0].GetResults()
	// println(len(products))
	// catalogue.SaveCataloguesFromMarkets(config.SNIIM_ADDR, markets, config.CATALOGUE_SRC)

}
