package consult

import (
	"log"

	"github.com/everitosan/sniim-scrapper/internal/app/consult"
	"github.com/everitosan/sniim-scrapper/internal/app/form"
)

func askBreadCrumb(paramsRepo form.ParamsRepository) *consult.Consult {
	// Ask for categories
	categories, _ := paramsRepo.GetCategories()
	_, categoryChosen, err := getOptionsPrompt("Selecciona una categoría", categories)
	if err != nil {
		log.Fatal(err)
	}

	// Ask for subcategories
	subcats, _ := paramsRepo.GetSubCategories(categoryChosen)
	_, subCategoryChosen, err := getOptionsPrompt("Selecciona una subcategoría", subcats)
	if err != nil {
		log.Fatal(err)
	}

	consult := consult.NewConsult(categoryChosen, subCategoryChosen)

	return consult
}
