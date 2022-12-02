package consult

import (
	"github.com/gocolly/colly"
)

func RegisterToMap(registers []RegisterConcept) map[string]string {
	result := make(map[string]string)
	for _, register := range registers {
		result[register.Name] = register.Value
	}

	return result
}

func Table0Scrapper(table *colly.HTMLElement, consult Consult) [][]RegisterConcept {
	headers := make([]string, 0)
	rows := make([][]RegisterConcept, 0)

	// Extract the headers
	table.ForEach("td[class=titDATtab2]", func(_ int, td *colly.HTMLElement) {
		headers = append(headers, td.Text)
	})

	// Extract the content
	table.ForEach("tr", func(index int, tr *colly.HTMLElement) {
		row := make([]RegisterConcept, 0)
		tr.ForEach("td[class=Datos2]", func(index int, td *colly.HTMLElement) {
			row = append(row, RegisterConcept{Name: headers[index], Value: td.Text})
		})
		if len(row) > 0 {
			// When finishes all td[class=Datos2] from a td
			row = append(consult.GetParamsAsConcepts(), row...)
			rows = append(rows, row)
		}
	})

	return rows
}
