package main

import (
	"log"

	"github.com/everitosan/snimm-scrapper/internal/app/catalogue"
	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
	"github.com/everitosan/snimm-scrapper/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.LoadConfig()

	if config.DEBUG {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug Log level")
	}

	mS := scrapper.NewMarketScrapper(config.SNIIM_ADDR)
	// pS := scrapper.NewProductScrapper(config.SNIIM_ADDR)

	markets, err := mS.RequestFromSource()

	if err != nil {
		log.Fatal(err)
	}

	catalogue.SaveCataloguesFromMarkets(config.SNIIM_ADDR, markets, config.CATALOGUE_SRC)

}
