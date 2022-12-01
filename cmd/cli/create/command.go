package create

import (
	"fmt"

	"github.com/everitosan/snimm-scrapper/internal/transport/repository"
	"github.com/spf13/cobra"
)

func Command(rContainer repository.Repository) *cobra.Command {
	return &cobra.Command{
		Use:   "create-request",
		Short: "Define a request",
		Long:  "Define a request specification by interactive cli",
		Run: func(cmd *cobra.Command, args []string) {

			// Ask category and subcategory
			consult := askBreadCrumb(rContainer.Params)

			// Ask for required inputs
			dateDetected := askInputs(rContainer, consult)

			// Ask dates
			if !dateDetected {
				askDates(consult)
			}

			fmt.Printf("Query %q\n", consult.ToUrl())
			rContainer.Consult.SaveOne(*consult)

		},
	}
}
