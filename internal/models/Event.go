package models

type Event struct {
	EventID     int64   `form:"eventid"`
	Title       string  `form:"title"`
	ImageUrl    string  `form:"imageurl"`
	Date        string  `form:"date"`
	Venue       string  `form:"venue"`
	Description string  `form:"description"`
	Time        string  `form:"time"`
	Amount      float32 `form:"amount"`
	Capacity    int     `form:"capacity"`
}

type Location struct {
	LocationId int64   `json:"locationid"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
