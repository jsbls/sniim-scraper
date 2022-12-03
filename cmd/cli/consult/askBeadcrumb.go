package consult

import (
	"log"

	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/everitosan/sniim-scrapper/internal/app/form"
)

func askBreadCrumb(paramsRepo form.ParamsRepository) *consult.Consult {
	// Ask for categories
	categories, _ := paramsRepo.GetCategories()
	_, categoryChoosen, err := getOptionsPrompt("Selecciona unacategoría", categories)
	if err != nil {
		log.Fatal(err)
	}

	// Ask for subcategories
	subcats, _ := paramsRepo.GetSubCategories(categoryChoosen)
	_, subCategoryChoosen, err := getOptionsPrompt("Selecciona una subcategoría", subcats)
	if err != nil {
		log.Fatal(err)
	}

	consult := consult.NewConsult(categoryChoosen, subCategoryChoosen)

	return consult
}
