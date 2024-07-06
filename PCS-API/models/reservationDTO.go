package models

type ReservationDTO struct {
	Reservation
	Bill    Bill         `json:"bill"`
	Service []ServiceDTO `json:"service"`
}
