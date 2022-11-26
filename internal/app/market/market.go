package market

import "fmt"

/*
*	Full Market
 */

type Market struct {
	Name       string     // Mercados nacionales
	Categories []Category // Precios de mercado, Anuarios estadísticos
}

func NewMarket(name string) Market {
	Categories := make([]Category, 0)

	return Market{
		Name:       name,
		Categories: Categories,
	}
}

func (m *Market) IsNotEmpty() bool {
	return m.Name != ""
}

func (m *Market) Print() {
	fmt.Println(m.Name)

	for _, cat := range m.Categories {
		fmt.Printf("\t%s\n", cat.Name)

		for _, sub := range cat.SubCategories {
			fmt.Printf("\t\t%s\n", sub.Name)

			for _, submarket := range sub.SubMarkets {
				fmt.Printf("\t\t\t%s -> %s\n", submarket.Name, submarket.Url)
			}
		}
	}
}
