// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Translation -.
type Login struct {
	Name  string `json:"Name"       example:"amir"`
	Email string `json:"Email"  example:"example@example.cpm"`
}
