package product

import (
	"strings"

	"github.com/everitosan/snimm-scrapper/internal/app/utils"
	"github.com/gocolly/colly"
)

type productScrapper struct {
	selector string
	products []Product
}

func NewProductScrapper(selector string) *productScrapper {
	return &productScrapper{
		products: make([]Product, 0),
		selector: selector,
	}
}

func (pS *productScrapper) GetProducts() []Product {
	return pS.products
}

func (pS *productScrapper) AddProduct(p Product) {
	pS.products = append(pS.products, p)
}

func (pS *productScrapper) Extract(html *colly.HTMLElement, keyJoined string) {
	html.ForEach(pS.selector, func(_ int, sel *colly.HTMLElement) {
		sel.ForEach("option", func(_ int, option *colly.HTMLElement) {
			keys := strings.Split(keyJoined, utils.KeyCatalogueSeparator)
			p := Product{
				Name:        option.Text,
				Id:          option.Attr("value"),
				Market:      keys[0],
				Inventory:   keys[1],
				Category:    keys[2],
				SubCategory: keys[3],
			}
			pS.AddProduct(p)
		})

	})
}
