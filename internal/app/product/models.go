package product

type Product struct {
	Id          string `json:"id"`          // Identifier of the product
	Name        string `json:"name"`        // Name of the product
	Market      string `json:"market"`      // Market where the procust belongs
	Inventory   string `json:"inventory"`   // Inventory where the product belongs
	Category    string `json:"category"`    // Category where the product belongs
	SubCategory string `json:"subcategory"` // Subcategory where the procust belongs
}
