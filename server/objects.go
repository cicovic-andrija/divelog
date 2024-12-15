package server

import (
	"fmt"
)

type All struct {
	DiveSites []*DiveSite `json:"dive_sites"`
	DiveTrips []*DiveTrip `json:"dive_trips"`
	Dives     []*Dive     `json:"dives"`
}

type SiteHead struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SiteFull struct {
	*DiveSite
	LinkedDives []*DiveHead `json:"linked_dives"`
}

type DiveHead struct {
	ID               int    `json:"id"`
	ShortLabel       string `json:"short_label"`
	DateTimeInPretty string `json:"date_time_in_pretty"`
}

type DiveFull struct {
	*Dive
	DiveSiteName     string `json:"dive_site_name"`
	DateTimeInPretty string `json:"date_time_in_pretty"`
}

type Trip struct {
	ID          int         `json:"id"`
	Label       string      `json:"label"`
	LinkedDives []*DiveHead `json:"linked_dives"`
}

func NewDiveHead(dive *Dive, diveSite *DiveSite) *DiveHead {
	return &DiveHead{
		ID:               dive.ID,
		ShortLabel:       fmt.Sprintf("No. %d | %s", dive.Number, diveSite.ShortName()),
		DateTimeInPretty: dive.datetime.Format("January 2 2006, 15:04"),
	}
}

func NewDiveFull(dive *Dive, diveSite *DiveSite) *DiveFull {
	return &DiveFull{
		Dive:             dive,
		DiveSiteName:     diveSite.Name,
		DateTimeInPretty: dive.datetime.Format("January 2 2006, 15:04"),
	}
}

func NewSiteFull(site *DiveSite, allDives []*Dive) *SiteFull {
	s := &SiteFull{DiveSite: site}
	for _, dive := range allDives {
		if dive.DiveSiteID == site.ID {
			s.LinkedDives = append(s.LinkedDives, NewDiveHead(dive, site))
		}
	}
	return s
}