package models

type Event struct {
	EventID     int64  `json:"eventid"`
	Title       string `json:"title"`
	ImageUrl    string `json:"imageurl"`
	Date        string `json:"date"`
	Venue       string `json:"venue"`
	Description string `json:"description"`
	Time        string `json:"time"`
	Amount      string `json:"amount"`
	Capacity    string `json:"capacity"`
}

type Location struct {
	LocationId int64   `json:"locationid"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
