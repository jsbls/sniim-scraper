package consult

import (
	"fmt"
	"log"

	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/everitosan/sniim-scrapper/internal/app/form"
	"github.com/everitosan/sniim-scrapper/internal/transport/repository"
)

func askInputs(rContainer repository.Repository, consult *consult.Consult) bool {

	subCategoryChoosen := consult.SubCategory

	// Ask for required inputs
	formParams, err := rContainer.Params.GetBySubCategory(subCategoryChoosen)
	if err != nil {
		log.Fatal(err)
	}

	var selectOptions []form.OptionSelect
	msg := ""
	dateDetected := false

	for _, param := range formParams.Params {

		switch param.Filter {
		case form.ProductType:
			selectOptions, _ = rContainer.Product.GetBySubCategory(subCategoryChoosen)
			msg = "Seleccona un producto"
		case form.OriginType:
			selectOptions, _ = rContainer.ProductSource.GetBySubCategory(subCategoryChoosen)
			msg = "Seleccona un origen"
		case form.DestinyType:
			selectOptions, _ = rContainer.ProductDestiny.GetBySubCategory(subCategoryChoosen)
			msg = "Seleccona un destino"
		case form.PerPriceType:
			selectOptions, _ = rContainer.PricePresentation.GetBySubCategory(subCategoryChoosen)
			msg = "Elije la presentación del precio"
		case form.WeekType:
			msg = "Elije la semana de consulta"
			selectOptions, _ = rContainer.Week.GetBySubCategory(subCategoryChoosen)
			dateDetected = true
		case form.MonthType:
			selectOptions, _ = rContainer.Month.GetBySubCategory(subCategoryChoosen)
			msg = "Elije el mes de consulta"
			dateDetected = true
		case form.YearType:
			selectOptions, _ = rContainer.Year.GetBySubCategory(subCategoryChoosen)
			msg = "Elije el año de consulta"
			dateDetected = true
		}

		options := make([]string, 0, len(selectOptions))
		for _, product := range selectOptions {
			options = append(options, product.Name)
		}
		index, _, err := getOptionsPrompt(msg, options)
		if err != nil {
			log.Fatal(err)
		}
		consult.AddParameter(param.UrlParam, selectOptions[index].Id)
		consult.AddTextParameter(fmt.Sprint(param.Filter), selectOptions[index].Name)
	}

	return dateDetected
}
