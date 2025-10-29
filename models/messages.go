package models

type Message struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	Author    string  `json:"author"`
	Text      string  `json:"text"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
