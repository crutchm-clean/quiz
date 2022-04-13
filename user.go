package Kahoot

type User struct {
	Id       string `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
