package utils

import (
	"github.com/gocolly/colly"
)

const KeyCatalogueSeparator = "-"

type FilterScrapper interface {
	Extract(*colly.HTMLElement, string) // should search by selector
}
