package server

type All struct {
	DiveSites []*DiveSite `json:"dive_sites"`
	DiveTrips []*DiveTrip `json:"dive_trips"`
	Dives     []*Dive     `json:"dives"`
}
