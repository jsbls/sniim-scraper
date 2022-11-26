package market

import (
	"strings"
)

/*
*	Market Inventory
 */

type InventoryType int32

const (
	Price InventoryType = iota
	HistoricYear
)

type Inventory struct {
	Name       string        `json:"name"`       // Precios de mercado || Anuarios estadísticos
	CatType    InventoryType `json:"category"`   // Price | Historic
	Categories []Catergory   `json:"categories"` // Agrícolas, Pecuarios, Pesqueros
}

func NewInventory(name string) Inventory {
	Categories := make([]Catergory, 0)

	return Inventory{
		Name:       name,
		Categories: Categories,
	}
}

func (m *Inventory) SetType(strType string) {
	if strings.Contains(strings.ToLower(strType), "anuario") {
		m.CatType = HistoricYear
	} else if strings.Contains(strings.ToLower(strType), "precios") {
		m.CatType = Price
	}
}

func (m *Inventory) IsNotEmpty() bool {
	return m.Name != ""
}
