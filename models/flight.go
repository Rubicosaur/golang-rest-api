package models

type Flight struct {
	Id                string `json:"id"`
	PlaneModel        string `json:"planemodel"`
	DepartureLocation string `json:"departurelocation"`
	ArrivalLocation   string `json:"arrivallocation"`
}
