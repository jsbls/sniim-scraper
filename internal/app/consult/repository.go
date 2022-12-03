package consult

type ConsultRepository interface {
	SaveOne(Consult) error
	DeleteOne(index int) error
	GetAll() ([]Consult, error)
}

type ConsultResponseRepository interface {
	Save([][]RegisterConcept) error
}
