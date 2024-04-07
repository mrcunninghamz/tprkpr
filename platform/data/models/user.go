package models

type User struct {
	ID         string
	Name       string
	Paydays    []Payday
	Worksheets []Worksheet
}
