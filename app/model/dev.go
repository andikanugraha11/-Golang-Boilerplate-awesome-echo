package model

type Dev struct {
	ID			int 	`json:"id,omitempty" col:"id"`
	Name 		string 	`json:"name,omitempty" col:"name"`
	Email 		string 	`json:"email,omitempty" col:"email"`
	Username	string 	`json:"username,omitempty" col:"username"`
	Password 	string 	`json:"password,omitempty" col:"password"`
}