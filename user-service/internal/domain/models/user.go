package models

type User struct {
	Uuid    string
	Name    string
	Surname string
	Email   string
	Role    string
}

type UpdateUser struct {
	Name    *string
	Surname *string
	Email   *string
	Role    *string
}
