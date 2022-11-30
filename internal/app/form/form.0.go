package form

import (
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

/*
* Form0Scrapper is based in md docs
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

	return &form0Inputs{
		inputs: []FormInput{product, origin, destiny, perprice, week, month, year},
	}

}

func From0Srapper(container *colly.HTMLElement, keys []string, f *FormScrapper) {

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
			p := OptionSelect{
				Name:        option.Text,
				Id:          option.Attr("value"),
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
