package request

import (
	"os"

	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
)

func PrintResultTable(results [][]consult.RegisterConcept) {
	logrus.Debugf("Showing %d results", len(results))

	headerRow := table.Row{}
	contentRows := make([]table.Row, 0)

	for _, register := range results[0] {
		headerRow = append(headerRow, register.Name)
	}

	for _, registers := range results {
		tmpRow := table.Row{}

		for _, register := range registers {
			tmpRow = append(tmpRow, register.Value)
		}

		contentRows = append(contentRows, tmpRow)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(headerRow)
	t.AppendRows(contentRows)

	t.Render()
}
