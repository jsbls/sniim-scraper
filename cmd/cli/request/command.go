package request

import (
	"fmt"

	_ "embed"

	"github.com/everitosan/snimm-scrapper/internal/app/consult"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const allFlag = "all"
const listFlag = "list"
const indexFlag = "index"

func Command(sniimAddr string, consultRepo consult.ConsultRepository) *cobra.Command {
	requestCommand := &cobra.Command{
		Use:   "request",
		Short: "Request information",
		Long:  "Request information from a consult register",
		Run: func(cmd *cobra.Command, args []string) {

			list, _ := cmd.Flags().GetBool(listFlag)
			index, _ := cmd.Flags().GetInt32(indexFlag)
			// all, _ := cmd.Flags().GetBool(allFlag)
			consults, err := consultRepo.GetAll()

			if index != -1 {
				if int(index) < len(consults) {
					selectedConsult := consults[index]
					consult.Scrap(sniimAddr, selectedConsult)
				} else {
					logrus.Warn("Indice inválido")
				}
				return
			}

			if list {
				if err != nil {
					logrus.Fatal(err)
				}
				for index, consult := range consults {
					fmt.Printf("(%d) - %s\n", index, consult.String())
				}
				return
			}

		},
	}

	requestCommand.Flags().BoolP(listFlag, "l", false, "List all requests")
	requestCommand.Flags().Int32P(indexFlag, "i", -1, "Request consult by index")
	requestCommand.Flags().BoolP(allFlag, "a", false, "Request all consults registered")

	return requestCommand
}
