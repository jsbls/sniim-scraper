package form

type OptionSelectRepository interface {
	GetGroupName() string
	GetAll() ([]OptionSelect, error)
	Save([]OptionSelect) error
	GetBySubCategory(string) ([]OptionSelect, error)
}

type ProductRepository interface {
	OptionSelectRepository
}

type ParamsRepository interface {
	Save([]FormParams) error
	GetAll() ([]FormParams, error)
	GetCategories() ([]string, error)
	GetSubCategories(category string) ([]string, error)
	GetBySubCategory(subcat string) (FormParams, error)
}
