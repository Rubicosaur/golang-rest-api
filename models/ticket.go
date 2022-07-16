package models

type Ticket struct {
	Id                string `json:"id"`
	PassengerName     string `json:"passengername"`
	Departurelocation string `json:"departurelocation"`
	ArrivalLocation   string `json:"arrivallocation"`
}

func (t *Ticket) GenerateId() {

	t.Id = "foo"
}
