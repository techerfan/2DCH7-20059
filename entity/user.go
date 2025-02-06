package entity

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type User struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      Gender `json:"Gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
