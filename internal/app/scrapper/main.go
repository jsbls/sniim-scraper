package scrapper

import (
	"log"

	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository/filestorage"
)

func InitCatalogues(sniimAddr, catalogueSrc string) {
	// Repositories
	marketRepo, _ := filestorage.NewMarketFileRepository(catalogueSrc)
	productRepo, _ := filestorage.NewOptionSelectFileRepository(catalogueSrc, "product")
	productSourceRepo, _ := filestorage.NewOptionSelectFileRepository(catalogueSrc, "productSource")
	productDestinyRepo, _ := filestorage.NewOptionSelectFileRepository(catalogueSrc, "productDestiny")
	pricePresentationRepo, _ := filestorage.NewOptionSelectFileRepository(catalogueSrc, "pricePresentation")
	weekRepo, _ := filestorage.NewOptionSelectFileRepository(catalogueSrc, "week")
	monthRepo, _ := filestorage.NewOptionSelectFileRepository(catalogueSrc, "month")
	yearRepo, _ := filestorage.NewOptionSelectFileRepository(catalogueSrc, "year")

	rContainer := repository.Repository{
		Market:            marketRepo,
		Product:           productRepo,
		ProductSource:     productSourceRepo,
		ProductDestiny:    productDestinyRepo,
		PricePresentation: pricePresentationRepo,
		Week:              weekRepo,
		Month:             monthRepo,
		Year:              yearRepo,
	}
	// Retrieve and save with repositories
	err := GetCatlogues(sniimAddr, rContainer)

	if err != nil {
		log.Fatal(err)
	}
}
