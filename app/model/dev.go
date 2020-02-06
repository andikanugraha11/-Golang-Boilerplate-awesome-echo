package model

type Dev struct {
	ID			int64 `json:"id"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Username	string `json:"username"`
	Password 	string `json:"password"`
}