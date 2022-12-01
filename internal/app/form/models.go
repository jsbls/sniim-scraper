package form

/*
* Option Select is a struct to save an option of a select input
 */

type OptionSelect struct {
	Id          string   `json:"id"`          // Identifier of the product
	Name        string   `json:"name"`        // Name of the product
	Market      string   `json:"market"`      // Market where the procust belongs
	Inventory   string   `json:"inventory"`   // Inventory where the product belongs
	Category    string   `json:"category"`    // Category where the product belongs
	SubCategory string   `json:"subcategory"` // Subcategory where the procust belongs
	FormType    FormType `json:"formType"`    // Form Type it shloud use
}

/*
* Form Input is a struct to map concept, selector and url param of an input
 */

type FormInput struct {
	Filter   SelectCategory `json:"category"`
	Selector string         `json:"selector"`
	UrlParam string         `json:"urlParam"`
}

/*
* formParams Represent a set of parameters a form contains
 */

type FormParams struct {
	Category    string      `json:"category"`    // Category where de params belongs
	SubCategory string      `json:"subcategory"` // Subcategory where de params belongs
	FormType    FormType    `json:"formType"`    // Form Type where de params belongs
	Params      []FormInput `json:"params"`      // Params of the form
}

func NewFormParams(keys []string, formType FormType) *FormParams {
	return &FormParams{
		Category:    keys[2],
		SubCategory: keys[3],
		FormType:    formType,
		Params:      make([]FormInput, 0),
	}
}

func (fP *FormParams) AddParam(param FormInput) {
	fP.Params = append(fP.Params, param)
}
