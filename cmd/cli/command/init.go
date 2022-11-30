package command

import (
	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/spf13/cobra"
)

func InitCommand(sniimAddr string, rContainer repository.Repository) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Retrieve information from source and create catalogues",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			scrapper.InitCatalogues(sniimAddr, rContainer)
		},
	}
}
