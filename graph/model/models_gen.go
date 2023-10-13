// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthSchema struct {
	AccessToken string `json:"accessToken"`
	ID          string `json:"id"`
}

type Event struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	Location  string `json:"location"`
	EndDate   string `json:"endDate"`
	CreatedBy *User  `json:"createdBy"`
}

type EventInput struct {
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Location  string `json:"location"`
}

type EventResponse struct {
	ID *int `json:"id,omitempty"`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type UserInput struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type UserResponse struct {
	ID *int `json:"id,omitempty"`
}
