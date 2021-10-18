package models

type PayloadStudentCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
