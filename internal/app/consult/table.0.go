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

func Table0Scrapper(table *colly.HTMLElement, consult Consult) []map[string]string {
	results := make([]map[string]string, 0)
	row := make([]RegisterConcept, 0)

	table.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
		tr.ForEach("td", func(index int, td *colly.HTMLElement) {
			class := td.Attr("class")
			switch class {
			case "titDATtab2":
				row = append(row, RegisterConcept{Name: td.Text})
			case "Datos2":
				row[index].Value = td.Text

				if index+1 == len(row) {
					row = append(row, consult.GetParamsAsConcepts()...)
					results = append(results, RegisterToMap(row))
				}
			}
		})
	})

	return results
}
