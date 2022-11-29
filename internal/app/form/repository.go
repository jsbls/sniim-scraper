package form

type OptionSelectRepository interface {
	GetGroupName() string
	GetAll() ([]OptionSelect, error)
	Save([]OptionSelect) error
}
