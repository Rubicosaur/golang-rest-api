package models

type ticket struct {
	Id                string `json:"id"`
	PassengerName     string `json:"passengername"`
	Departurelocation string `json:"departurelocation"`
	ArrivalLocation   string `json:"arrivallocation"`
}

func GenerateTicket(id, name, surname, departureLoc, arrivalLoc string) ticket {

	t := ticket{}

	t.Id = id
	t.PassengerName = name

	return t
}
