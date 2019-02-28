package api

type Shop struct {
	Id       int    `json:"shop_id"`
	Name     string `json:"shop_name"`
	Vacation bool   `json:"is_vacation"`
}
