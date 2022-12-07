package form

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

/*
* Form0Scraper is based in md docs
 */

type form0Inputs struct {
	inputs []FormInput
}

func newForm0Inputs() *form0Inputs {

	product := FormInput{Filter: ProductType, Selector: "ddlProducto", UrlParam: "ProductoId"}
	origin := FormInput{Filter: OriginType, Selector: "ddlOrigen", UrlParam: "OrigenId"}
	destiny := FormInput{Filter: DestinyType, Selector: "ddlDestino", UrlParam: "DestinoId"}
	perprice := FormInput{Filter: PerPriceType, Selector: "ddlPrecios", UrlParam: "PreciosPorId"}
	week := FormInput{Filter: WeekType, Selector: "ddlSemanaSemanal", UrlParam: "Semana"}
	month := FormInput{Filter: MonthType, Selector: "ddlMesSemanal", UrlParam: "Mes"}
	year := FormInput{Filter: YearType, Selector: "ddlAnioSemana", UrlParam: "Anio"}
	yearq := FormInput{Filter: YearType, Selector: "ddlAnioQuincena", UrlParam: "Anio"}

	return &form0Inputs{
		inputs: []FormInput{product, origin, destiny, perprice, week, month, year, yearq},
	}

}

// Const map for months
var months = map[string]string{
	"enero":      "1",
	"febrero":    "2",
	"marzo":      "3",
	"abril":      "4",
	"mayo":       "5",
	"junio":      "6",
	"julio":      "7",
	"agosto":     "8",
	"septiembre": "9",
	"octubre":    "10",
	"noviembre":  "11",
	"diciembre":  "12",
}

func From0Scraper(container *colly.HTMLElement, keys []string, f *FormScraper) {

	form0 := newForm0Inputs()

	container.ForEach("select", func(_ int, sel *colly.HTMLElement) {
		selectId := sel.Attr("id")
		var selectCat SelectCategory
		found := false

		for _, formInput := range form0.inputs {
			if selectId == formInput.Selector {
				selectCat = formInput.Filter
				f.Params.AddParam(formInput)
				found = true
				break
			}
		}

		if !found {
			logrus.Warnf("No criteria for input %s %s", selectId, keys)
			return
		}

		sel.ForEach("option", func(_ int, option *colly.HTMLElement) {
			value := option.Attr("value")
			if selectCat == MonthType {
				month, exists := months[strings.ToLower(value)]
				if exists {
					value = month
				}
			}
			p := OptionSelect{
				Name:        option.Text,
				Id:          value,
				Market:      keys[0],
				Inventory:   keys[1],
				Category:    keys[2],
				SubCategory: keys[3],
				FormType:    Form0Type,
			}
			f.Inputs.AddOption(selectCat, p)
		})
	})

}
