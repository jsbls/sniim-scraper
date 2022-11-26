package product

type ProductRepository interface {
	Save([]Product) error
	GetDstName() string
}
