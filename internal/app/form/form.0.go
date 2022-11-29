package form

import (
	"github.com/gocolly/colly"
)

/*
* Form0Scrapper is based in md docs
 */

func From0Srapper(container *colly.HTMLElement, keys []string, f *formScrapper) {
	container.ForEach("select", func(_ int, sel *colly.HTMLElement) {
		selectId := sel.Attr("id")
		var selectCat SelectCategory

		switch selectId {
		case "ddlProducto":
			selectCat = ProductType
		case "ddlOrigen":
			selectCat = OriginType
		case "ddlDestino":
			selectCat = DestinyType
		case "ddlPrecios":
			selectCat = PerPriceType
		case "ddlSemanaSemanal":
			selectCat = WeekType
		case "ddlMesSemanal":
			selectCat = MonthType
		case "ddlAnioSemana":
			selectCat = YearType
		}

		sel.ForEach("option", func(_ int, option *colly.HTMLElement) {
			p := OptionSelect{
				Name:        option.Text,
				Id:          option.Attr("value"),
				Market:      keys[0],
				Inventory:   keys[1],
				Category:    keys[2],
				SubCategory: keys[3],
				FormType:    Form0,
			}
			f.Inputs.AddOption(selectCat, p)
		})
	})

}
