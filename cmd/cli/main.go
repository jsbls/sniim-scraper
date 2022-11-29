package main

import (
	"github.com/everitosan/snimm-scrapper/internal/app/scrapper"
	"github.com/everitosan/snimm-scrapper/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	config := config.LoadConfig()

	if config.DEBUG {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug Log level")
	}

	rootCmd := &cobra.Command{Use: "snimm-cli"}

	rootCmd.AddCommand(getInitCalaogueCommand(config.SNIIM_ADDR, config.CATALOGUE_SRC))
	rootCmd.Execute()
}

func getInitCalaogueCommand(sniimAddr, catalogueSrc string) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Retrieve information from source and create catalogues",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			scrapper.InitCatalogues(sniimAddr, catalogueSrc)
		},
	}
}
