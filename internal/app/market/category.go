package market

import (
	"strings"
)

/*
*	Market category
 */

type CategoryType int32

const (
	Price CategoryType = iota
	HistoricYear
)

type Category struct {
	Name          string         // Precios de mercado || Anuarios estadísticos
	CatType       CategoryType   // Price | Historic
	SubCategories []SubCatergory // Agrícolas, Pecuarios, Pesqueros
}

func NewCategory(name string) Category {
	SubCategories := make([]SubCatergory, 0)

	return Category{
		Name:          name,
		SubCategories: SubCategories,
	}
}

func (m *Category) SetType(strType string) {
	if strings.Contains(strings.ToLower(strType), "anuario") {
		m.CatType = HistoricYear
	} else if strings.Contains(strings.ToLower(strType), "precios") {
		m.CatType = Price
	}
}

func (m *Category) IsNotEmpty() bool {
	return m.Name != ""
}
