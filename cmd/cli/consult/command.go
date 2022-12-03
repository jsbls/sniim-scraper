package consult

import (
	"fmt"

	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const createCosultFlag = "create"
const listCosultFlag = "list"

func Command(rContainer repository.Repository) *cobra.Command {
	consultCommand := &cobra.Command{
		Use:   "consult",
		Short: "Manage consults",
		Long:  "Can create and list consults",
		Run: func(cmd *cobra.Command, args []string) {

			list, _ := cmd.Flags().GetBool(listCosultFlag)
			create, _ := cmd.Flags().GetBool(createCosultFlag)

			switch {
			case list:
				/*
				* Case for listing consults
				 */
				consults, err := rContainer.Consult.GetAll()
				if err != nil {
					logrus.Fatal(err)
				}

				for index, consult := range consults {
					fmt.Printf("(%d) - %s\n", index, consult.String())
				}
				return
			case create:
				/*
				* Case create a consult
				 */
				// Ask category and subcategory
				consult := askBreadCrumb(rContainer.Params)

				// Ask for required inputs
				dateDetected := askInputs(rContainer, consult)

				// Ask dates
				if !dateDetected {
					askDates(consult)
				}

				rContainer.Consult.SaveOne(*consult)
				return
			}

		},
	}

	consultCommand.Flags().BoolP(createCosultFlag, "c", false, "Create a consult register")
	consultCommand.Flags().BoolP(listCosultFlag, "l", false, "List all consults")

	return consultCommand
}
