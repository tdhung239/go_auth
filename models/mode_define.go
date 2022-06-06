package models

type User struct {
	Email    string `orm:"pk" json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
	Phone    int    `json:"phone"`
	Adress   string `json:"adress"`
	Admin    bool   `json:"admin"`
}
