package model

import  "time"

type EntityFields struct {
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

type Person struct {
	EntityFields
	Id        uint      `json:"id"`
    Firstname string    `json:"firstname, omitempty"`
    Lastname  string    `json:"lastname, omitempty"`
    Address   Address   `json:"address, omitempty"`
    AddressId uint		`json:"-"`
}

type Address struct {
	EntityFields
	Id    uint    `json:"-"`
    City  string  `json:"city, omitempty"`
    State string  `json:"state, omitempty"`
}
