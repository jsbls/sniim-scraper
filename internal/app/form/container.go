package form

type SelectCategory int64

const (
	ProductType SelectCategory = iota
	OriginType
	DestinyType
	PerPriceType
	WeekType
	MonthType
	YearType
)

type inputContainer struct {
	inputs map[SelectCategory][]OptionSelect // help to store all available options of a select
}

func NewInputContainer() *inputContainer {
	inputs := make(map[SelectCategory][]OptionSelect)
	return &inputContainer{
		inputs: inputs,
	}
}

func (f *inputContainer) AddOption(selectCat SelectCategory, inputOption OptionSelect) {
	prev, exists := f.inputs[selectCat]
	if !exists {
		f.inputs[selectCat] = make([]OptionSelect, 0)
		prev = f.inputs[selectCat]
	}

	f.inputs[selectCat] = append(prev, inputOption)
}

func (f *inputContainer) AddOptions(selectCat SelectCategory, inputOptions []OptionSelect) {
	prev, exists := f.inputs[selectCat]
	if !exists {
		f.inputs[selectCat] = make([]OptionSelect, 0)
		prev = f.inputs[selectCat]
	}

	f.inputs[selectCat] = append(prev, inputOptions...)
}

func (f *inputContainer) GetInputs() map[SelectCategory][]OptionSelect {
	return f.inputs
}
