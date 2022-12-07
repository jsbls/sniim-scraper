package consult

import (
	_ "embed"

	"github.com/everitosan/sniim-scrapper/internal/app/utils"
	"github.com/gocolly/colly"
	"gopkg.in/yaml.v3"
)

//go:embed endpoints.yaml
var endpointsStr string

func getUrl(consult Consult) (url string, err error) {
	url = ""
	var data []CategoryEndpoint
	err = yaml.Unmarshal([]byte(endpointsStr), &data)

	if err != nil {
		return url, err
	}

	for _, endpoint := range data {
		if endpoint.Category == consult.Category {
			for _, subcat := range endpoint.Subcategories {
				if subcat.Name == consult.SubCategory {
					url = endpoint.UrlPrefix + subcat.Url + consult.ToUrl()
					break
				}
			}
		}
	}

	return url, nil
}

func Scrap(sniimAddr string, consult Consult) ([][]RegisterConcept, error) {
	var registers [][]RegisterConcept

	subUrl, err := getUrl(consult)
	if err != nil {
		return registers, err
	}

	url := sniimAddr + subUrl

	requester := utils.NewRequester(sniimAddr)

	html, err := requester.SyncR(url)
	if err != nil {
		return registers, err
	}

	html.ForEach("table", func(_ int, table *colly.HTMLElement) {
		tableId := table.Attr("id")

		switch tableId {
		case "tblResultados":
			registers = Table0Scrapper(table, consult)
		}
	})

	return registers, nil
}
