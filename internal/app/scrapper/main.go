package scrapper

import (
	"log"

	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
)

func InitCatalogues(sniimAddr string, rContainer repository.Repository) {

	// Retrieve and save with repositories
	err := GetCatlogues(sniimAddr, rContainer)

	if err != nil {
		log.Fatal(err)
	}
}
