package server

import (
	"fmt"
	"time"
)

type DiveLog struct {
	Metadata         DiveLogMetadata
	DiveSites        []*DiveSite
	DiveTrips        []*DiveTrip
	Dives            []*Dive
	sourceToSystemID map[string]int
}

type DiveLogMetadata struct {
	Program        string `json:"program"`
	ProgramVersion string `json:"program_version"`
	Source         string `json:"source"`
	Units          string `json:"units"`
}

type DiveSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Coordinates string   `json:"coordinates"`
	GeoLabels   []string `json:"geo_labels"`

	sourceID string
}

type DiveTrip struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

type Dive struct {
	ID         int `json:"id"`
	Number     int `json:"number"`
	DiveSiteID int `json:"dive_site_id"`
	DiveTripID int `json:"dive_trip_id"`

	DateTimeIn string `json:"date_time_in"`

	datetime time.Time
}

func (s *DiveSite) String() string {
	return fmt.Sprintf("S%d:[%s]", s.ID, s.Name)
}

func (t *DiveTrip) String() string {
	return fmt.Sprintf("T%d:[%s]", t.ID, t.Label)
}

func (d *Dive) String() string {
	return fmt.Sprintf("D%d:[%s]", d.ID, d.datetime.Format(time.DateOnly))
}
