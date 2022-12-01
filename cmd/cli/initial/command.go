package initial

import (
	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/spf13/cobra"
)

func Command(sniimAddr string, rContainer repository.Repository) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create catalogues",
		Long:  "Retrieve information from source and create catalogues",
		Run: func(cmd *cobra.Command, args []string) {
			scrapper.InitCatalogues(sniimAddr, rContainer)
		},
	}
}
