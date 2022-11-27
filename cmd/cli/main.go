package main

import (
	"log"

	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
	"github.com/everitosan/snimm-scrapper/internal/config"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository/filestorage"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.LoadConfig()

	if config.DEBUG {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug Log level")
	}

	// Repositories
	marketRepo, _ := filestorage.NewMarketFileRepository(config.CATALOGUE_SRC)
	productRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "product")
	productSourceRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "productSource")
	productDestinyRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "productDestiny")
	pricePresentationRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "pricePresentation")

	rContainer := repository.Repository{
		Market:            marketRepo,
		Product:           productRepo,
		ProductSource:     productSourceRepo,
		ProductDestiny:    productDestinyRepo,
		PricePresentation: pricePresentationRepo,
	}
	// Retrieve and save with repositories
	err := scrapper.InitCatlogues(config.SNIIM_ADDR, rContainer)

	if err != nil {
		log.Fatal(err)
	}

}
