package data

// Customer contains details for single customer
type Customer struct {
	ID       int64  `json:"customerID"`
	Fullname string `json:"fullname"`
}
