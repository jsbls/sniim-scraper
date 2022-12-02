package main

import (
	"github.com/everitosan/snimm-scrapper/cmd/cli/consult"
	"github.com/everitosan/snimm-scrapper/cmd/cli/initial"
	"github.com/everitosan/snimm-scrapper/cmd/cli/request"
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
	consultResponseRepo, _ := filestorage.NewConsultResponseFileRepository(config.CATALOGUE_SRC, "consultsResponses")

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
		ConsultResponse:   consultResponseRepo,
		Product:           productRepo,
		ProductSource:     productSourceRepo,
		ProductDestiny:    productDestinyRepo,
		PricePresentation: pricePresentationRepo,
		Week:              weekRepo,
		Month:             monthRepo,
		Year:              yearRepo,
	}

	rootCmd := &cobra.Command{Use: "snimm-cli"}

	rootCmd.AddCommand(initial.Command(config.SNIIM_ADDR, rContainer))
	rootCmd.AddCommand(consult.Command(rContainer))
	rootCmd.AddCommand(request.Command(config.SNIIM_ADDR, rContainer.Consult, rContainer.ConsultResponse))
	rootCmd.Execute()
}
