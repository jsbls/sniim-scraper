package market

/*
	SubCategory
*/

type SubCatergory struct {
	Name       string      // Agrícolas
	SubMarkets []SubMarket // Frutas y hortalizas, Flores, Granos, Azúcar
}

func NewSubCategory(name string) SubCatergory {
	SubMarkets := make([]SubMarket, 0)

	return SubCatergory{
		Name:       name,
		SubMarkets: SubMarkets,
	}
}

func (m *SubCatergory) IsNotEmpty() bool {
	return m.Name != ""
}
