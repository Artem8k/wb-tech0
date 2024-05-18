package models

type Delivery struct {
	Id       int    `json:"-" db:"id"`
	OrderUid string `json:"-" db:"order_uid"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Zip      string `json:"zip" db:"zip"`
	City     string `json:"city" db:"city"`
	Adress   string `json:"adress" db:"adress"`
	Region   string `json:"region" db:"region"`
	Email    string `json:"email" db:"email"`
}
