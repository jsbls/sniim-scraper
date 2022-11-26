package market

/*
	Category
*/

type Catergory struct {
	Name          string        `json:"name"`          // Agrícolas
	SubCategories []SubCategory `json:"subcategories"` // Frutas y hortalizas, Flores, Granos, Azúcar
}

func NewCategory(name string) Catergory {
	SubCategories := make([]SubCategory, 0)

	return Catergory{
		Name:          name,
		SubCategories: SubCategories,
	}
}

func (m *Catergory) IsNotEmpty() bool {
	return m.Name != ""
}
