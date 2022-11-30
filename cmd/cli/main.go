package main

import (
	"github.com/everitosan/snimm-scrapper/cmd/cli/command"
	"github.com/everitosan/snimm-scrapper/cmd/cli/create"
	"github.com/everitosan/snimm-scrapper/internal/config"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository/filestorage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	config := config.LoadConfig()

	if config.DEBUG {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug Log level")
	}

	// Repositories
	marketRepo, _ := filestorage.NewMarketFileRepository(config.CATALOGUE_SRC)
	paramsRepo, _ := filestorage.NewParamsFileRepository(config.CATALOGUE_SRC, "params")
	consultRepo, _ := filestorage.NewConsultFileRepository(config.CATALOGUE_SRC, "consults")

	productRepo, _ := filestorage.NewProductFileRepository(config.CATALOGUE_SRC, "product")
	productSourceRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "productSource")
	productDestinyRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "productDestiny")
	pricePresentationRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "pricePresentation")
	weekRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "week")
	monthRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "month")
	yearRepo, _ := filestorage.NewOptionSelectFileRepository(config.CATALOGUE_SRC, "year")

	rContainer := repository.Repository{
		Market:            marketRepo,
		Params:            paramsRepo,
		Consult:           consultRepo,
		Product:           productRepo,
		ProductSource:     productSourceRepo,
		ProductDestiny:    productDestinyRepo,
		PricePresentation: pricePresentationRepo,
		Week:              weekRepo,
		Month:             monthRepo,
		Year:              yearRepo,
	}

	rootCmd := &cobra.Command{Use: "snimm-cli"}

	rootCmd.AddCommand(command.InitCommand(config.SNIIM_ADDR, rContainer))
	rootCmd.AddCommand(create.Command(rContainer))
	rootCmd.Execute()
}
