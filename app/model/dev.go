package model

type Dev struct {
	ID			int64 	`json:"id,omitempty"`
	Name 		string 	`json:"name,omitempty"`
	Email 		string 	`json:"email,omitempty"`
	Username	string 	`json:"username,omitempty"`
	Password 	string 	`json:"password,omitempty"`
}