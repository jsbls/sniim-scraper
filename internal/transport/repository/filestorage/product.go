package filestorage

type productFileRepository struct {
	optionSelectFileRepository
}

func NewProductFileRepository(dst, fileName string) (*productFileRepository, error) {
	var pR productFileRepository
	optionRepo, err := NewOptionSelectFileRepository(dst, fileName)
	if err != nil {
		return &pR, err
	}
	return &productFileRepository{
		optionSelectFileRepository: *optionRepo,
	}, nil
}
