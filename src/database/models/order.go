package models

type Order struct {
	OrderUid           string   `json:"order_uid" db:"uid"`
	Track_number       string   `json:"track_number" db:"track_number"`
	Entry              string   `json:"entry" db:"entry"`
	Delivery           Delivery `json:"delivery" db:"delivery"`
	Payment            Payment  `json:"payment" db:"payment"`
	Items              []Item   `json:"items" db:"items"`
	Locale             string   `json:"locale" db:"locale"`
	Internal_signature string   `json:"internal_signature" db:"internal_signature"`
	Customer_id        string   `json:"customer_id" db:"customer_id"`
	Delivery_service   string   `json:"delivery_service" db:"delivery_service"`
	Shardkey           string   `json:"shardkey" db:"shardkey"`
	Sm_id              int      `json:"sm_id" db:"sm_id"`
	Date_created       string   `json:"date_created" db:"date_created"`
	Oof_shard          string   `json:"oof_shard" db:"oof_shard"`
}
