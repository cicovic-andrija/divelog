package server

import (
	"fmt"
	"strings"
	"time"
)

// TODO: There are two yellowstone reefs, both with 1 dive
// TODO: There is one Tugboat rozi with no dives

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

	Coordinates string   `json:"coordinates,omitempty"`
	GeoLabels   []string `json:"geo_labels,omitempty"`

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

	Rating5         int      `json:"rating5,omitempty"`
	Visibility5     int      `json:"visibility5,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	Salinity        string   `json:"salinity,omitempty"`
	DateTimeIn      string   `json:"date_time_in,omitempty"`
	OperatorDM      string   `json:"operator_dm,omitempty"`
	Buddy           string   `json:"buddy,omitempty"`
	Notes           string   `json:"notes,omitempty"`
	Suit            string   `json:"suit,omitempty"`
	CylSize         string   `json:"cyl_size,omitempty"`
	CylType         string   `json:"cyl_type,omitempty"`
	StartPressure   string   `json:"start_pressure,omitempty"`
	EndPressure     string   `json:"end_pressure,omitempty"`
	Gas             string   `json:"gas,omitempty"`
	Weights         string   `json:"weights,omitempty"`
	WeightsType     string   `json:"weights_type,omitempty"`
	DCModel         string   `json:"dc_model,omitempty"`
	DepthMax        string   `json:"depth_max,omitempty"`
	DepthMean       string   `json:"depth_mean,omitempty"`
	TempWaterMin    string   `json:"temp_water_min,omitempty"`
	TempAir         string   `json:"temp_air,omitempty"`
	SurfacePressure string   `json:"surface_pressure,omitempty"`

	datetime time.Time
}

func (s *DiveSite) String() string {
	return fmt.Sprintf("S%d:[%s]", s.ID, s.Name)
}

func (s *DiveSite) ShortName() string {
	return strings.TrimSpace(strings.Split(s.Name, ",")[0])
}

func (t *DiveTrip) String() string {
	return fmt.Sprintf("T%d:[%s]", t.ID, t.Label)
}

func (d *Dive) String() string {
	return fmt.Sprintf("D%d:[%s]", d.ID, d.datetime.Format(time.DateOnly))
}

func (d *Dive) Normalize() {
	if strings.HasPrefix(d.Salinity, "1000") {
		d.Salinity = "fresh water"
	} else if strings.HasPrefix(d.Salinity, "1030") {
		d.Salinity = "salt water"
	} else {
		d.Salinity = ""
	}

	if d.Gas == "" {
		d.Gas = "air"
	} else { // e.g. "32%"
		d.Gas = "nitrox " + d.Gas
	}

	d.CylType = cylTypeMappings[d.CylType]
}

func (d *Dive) IsTaggedWith(tag string) bool {
	if tag == "" {
		return true
	}
	for _, t := range d.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

var cylTypeMappings = map[string]string{
	"AL100": "aluminium",
	"HP100": "steel",
	"HP130": "steel",
}
