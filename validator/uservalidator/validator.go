package uservalidator

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
	IsEmailUnique(email string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
