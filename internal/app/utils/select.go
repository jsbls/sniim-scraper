package utils

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type OptionSelect struct {
	Id          string `json:"id"`          // Identifier of the product
	Name        string `json:"name"`        // Name of the product
	Market      string `json:"market"`      // Market where the procust belongs
	Inventory   string `json:"inventory"`   // Inventory where the product belongs
	Category    string `json:"category"`    // Category where the product belongs
	SubCategory string `json:"subcategory"` // Subcategory where the procust belongs
}

type selectScrapper struct {
	selector string
	options  []OptionSelect
}

func NewOptionSelectScrapper(selector string) *selectScrapper {
	return &selectScrapper{
		options:  make([]OptionSelect, 0),
		selector: selector,
	}
}

func (pS *selectScrapper) GetOptions() []OptionSelect {
	return pS.options
}

func (pS *selectScrapper) AddOption(p OptionSelect) {
	pS.options = append(pS.options, p)
}

func (pS *selectScrapper) Extract(url string, html *colly.HTMLElement, keyJoined string) {
	html.ForEach(pS.selector, func(_ int, sel *colly.HTMLElement) {
		sel.ForEach("option", func(_ int, option *colly.HTMLElement) {
			keys := strings.Split(keyJoined, KeyCatalogueSeparator)
			p := OptionSelect{
				Name:        option.Text,
				Id:          option.Attr("value"),
				Market:      keys[0],
				Inventory:   keys[1],
				Category:    keys[2],
				SubCategory: keys[3],
			}
			pS.AddOption(p)
		})
	})

	if len(pS.options) == 0 {
		logrus.Warnf("%s in %s has %d", pS.selector, url, len(pS.options))
	}
}

type OptionSelectRepository interface {
	GetGroupName() string
	GetAll() ([]OptionSelect, error)
	Save([]OptionSelect) error
}
