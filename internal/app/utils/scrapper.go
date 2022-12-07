package utils

import (
	"github.com/gocolly/colly"
)

const KeyCatalogueSeparator = "-"

type FilterScraper interface {
	Extract(string, *colly.HTMLElement, string) // should search by selector
}
