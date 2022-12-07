package initial

import (
	"github.com/everitosan/sniim-scrapper/internal/app/scraper"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
	"github.com/spf13/cobra"
)

func Command(sniimAddr string, rContainer repository.Repository) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create catalogues",
		Long:  "Retrieve information from source and create catalogues",
		Run: func(cmd *cobra.Command, args []string) {
			scraper.InitCatalogues(sniimAddr, rContainer)
		},
	}
}
