package models

type Ticket struct {
	TicketID int64 `json:"ticketId"`
	UserID   int64 `json:"userId"`
	EventId  int64 `json:"eventId"`
}
