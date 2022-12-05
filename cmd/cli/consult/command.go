package consult

import (
	"fmt"
	"log"

	"github.com/everitosan/sniim-scrapper/cmd/cli/request"
	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const createCosultFlag = "create"
const listCosultFlag = "list"
const deleteCosultFlag = "delete"
const saveConsultFlag = "save"

func Command(sniiimAddr string, rContainer repository.Repository) *cobra.Command {
	consultCommand := &cobra.Command{
		Use:   "consult",
		Short: "Manage consults",
		Long:  "Can create and list consults",
		Run: func(cmd *cobra.Command, args []string) {

			list, _ := cmd.Flags().GetBool(listCosultFlag)
			create, _ := cmd.Flags().GetBool(createCosultFlag)
			delete, _ := cmd.Flags().GetInt16(deleteCosultFlag)
			save, _ := cmd.Flags().GetBool(saveConsultFlag)

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
				newConsult := askBreadCrumb(rContainer.Params)

				// Ask for required inputs
				dateDetected := askInputs(rContainer, newConsult)

				// Ask dates
				if !dateDetected {
					askDates(newConsult)
				}

				results, err := consult.Scrap(sniiimAddr, *newConsult)

				if err != nil {
					logrus.Fatal(err)
				}

				request.PrintResultTable(results)

				if save {
					rContainer.Consult.SaveOne(*newConsult)
				} else {
					res, err := confirmPropmpt("¿Desea guardar la consulta?")
					if err != nil {
						logrus.Fatal(err)
					}
					if res == "y" {
						rContainer.Consult.SaveOne(*newConsult)
					}
				}

				return
			case delete != -1:
				/*
				* Delete case
				 */
				err := rContainer.Consult.DeleteOne(int(delete))
				if err != nil {
					log.Fatal(err)
				}
				return
			}

		},
	}

	consultCommand.Flags().BoolP(createCosultFlag, "c", false, "Create a consult")
	consultCommand.Flags().BoolP(saveConsultFlag, "s", false, "Save a consult register")
	consultCommand.Flags().BoolP(listCosultFlag, "l", false, "List all consults")
	consultCommand.Flags().Int16P(deleteCosultFlag, "d", -1, "Delete a consult register")

	return consultCommand
}
