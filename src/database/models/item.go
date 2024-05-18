package models

type Item struct {
	Id           int    `json:"-" db:"id"`
	OrderUid     string `json:"-" db:"order_uid"`
	Chrt_id      int    `json:"chrt_id" db:"chrt_id"`
	Track_number string `json:"track_number" db:"track_number"`
	Price        int    `json:"price" db:"price"`
	Rid          string `json:"rid" db:"rid"`
	Name         string `json:"name" db:"name"`
	Sale         int    `json:"sale" db:"sale"`
	Size         string `json:"size" db:"size"`
	Total_price  int    `json:"total_price" db:"total_price"`
	Nm_id        int    `json:"nm_id" db:"nm_id"`
	Brand        string `json:"brand" db:"brand"`
	Status       int    `json:"status" db:"status"`
}
