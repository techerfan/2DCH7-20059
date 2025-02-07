package tablevalidator

type Repository interface {
	DoesTableExist(id uint) (bool, error)
	DoesTableExistByTableNum(no uint8) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
