package request

import (
	_ "embed"

	"github.com/everitosan/snimm-scrapper/internal/app/consult"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const allFlag = "all"
const indexFlag = "index"
const saveFlag = "save"

func Command(sniimAddr string, consultRepo consult.ConsultRepository, responseRepo consult.ConsultResponseRepository) *cobra.Command {
	requestCommand := &cobra.Command{
		Use:   "request",
		Short: "Request information",
		Long:  "Request information from a consult register",
		Run: func(cmd *cobra.Command, args []string) {

			index, _ := cmd.Flags().GetInt32(indexFlag)
			shouldSave, _ := cmd.Flags().GetBool(saveFlag)
			consults, err := consultRepo.GetAll()

			if err != nil {
				logrus.Fatal(err)
			}

			switch {
			case index != -1:
				/*
				* Case for makinng a single request
				 */
				if int(index) >= len(consults) {
					logrus.Warnf("No existe consulta número %d", index)
					return
				}

				selectedConsult := consults[index]
				results, err := consult.Scrap(sniimAddr, selectedConsult)

				if err != nil {
					logrus.Fatal(err)
				}

				if len(results) == 0 {
					logrus.Warn("No hay resultados de la búsqueda.")
					return
				}

				if shouldSave {
					err = responseRepo.Save(results)
				} else {

					PrintResultTable(results)
				}

				if err != nil {
					logrus.Fatal(err)
				}
				return

			}

		},
	}

	requestCommand.Flags().Int32P(indexFlag, "i", -1, "Request consult by index")
	requestCommand.Flags().BoolP(allFlag, "a", false, "Request all consults registered")
	requestCommand.Flags().BoolP(saveFlag, "s", false, "Save the results of a request")

	return requestCommand
}
