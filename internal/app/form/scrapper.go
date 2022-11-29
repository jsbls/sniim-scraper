package form

import (
	"strings"

	"github.com/everitosan/snimm-scrapper/internal/app/utils"
	"github.com/gocolly/colly"
)

// Enums
type formType int64

const (
	Form0 formType = iota
)


/*
* Form scraps an html and create a form struct with it's selects by category
* when finishes, every type represents a single input in the form
 */

type formScrapper struct {
	Inputs inputContainer // help to store all available options of a select
}

func NewFormScrapper() *formScrapper {
	return &formScrapper{
		Inputs: *NewInputContainer(),
	}
}

func (f *formScrapper) GetFormInputs(html *colly.HTMLElement, keyJoined string) {
	html.ForEach("table", func(_ int, table *colly.HTMLElement) {
		tableId := table.Attr("id")
		keys := strings.Split(keyJoined, utils.KeyCatalogueSeparator)

		switch tableId {
		case "tblDatos":
			From0Srapper(table, keys, f)
		case "tblFiltro":
			From0Srapper(table, keys, f)
		}
	})
}
