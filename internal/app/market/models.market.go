package market

import "fmt"

/*
*	Full Market
 */

type Market struct {
	Name        string      `json:"name"`        // Mercados nacionales
	Inventories []Inventory `json:"inventories"` // Precios de mercado, Anuarios estadísticos
}

func NewMarket(name string) Market {
	Inventories := make([]Inventory, 0)

	return Market{
		Name:        name,
		Inventories: Inventories,
	}
}

func (m *Market) IsNotEmpty() bool {
	return m.Name != ""
}

func (m *Market) Print() {
	fmt.Println(m.Name)

	for _, inv := range m.Inventories {
		fmt.Printf("\t%s\n", inv.Name)

		for _, sub := range inv.Categories {
			fmt.Printf("\t\t%s\n", sub.Name)

			for _, submarket := range sub.SubCategories {
				fmt.Printf("\t\t\t%s -> %s\n", submarket.Name, submarket.Url)
			}
		}
	}
}
