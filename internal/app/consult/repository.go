package consult

type ConsultRepository interface {
	SaveOne(Consult) error
	GetAll() ([]Consult, error)
}
