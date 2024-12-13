package server

import (
	"fmt"
	"os"

	"src.acicovic.me/divelog/subsurface"
)

var inmemDatabase DiveLog

func buildDatabase(sourceFullPath string) error {
	file, err := os.Open(sourceFullPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", sourceFullPath, err)
	}
	defer file.Close()

	if err = subsurface.DecodeSubsurfaceDatabase(file, &SubsurfaceCallbackHandler{}); err != nil {
		return fmt.Errorf("failed to decode database in %s: %v", sourceFullPath, err)
	}

	inmemDatabase.Metadata.Source = sourceFullPath
	return nil
}

type SubsurfaceCallbackHandler struct {
	lastSiteID int
	lastTripID int
	lastDiveID int
}

func (p *SubsurfaceCallbackHandler) HandleBegin() {
	inmemDatabase.DiveSites = make([]*DiveSite, 1, 100)
	inmemDatabase.DiveTrips = make([]*DiveTrip, 1, 100)
	inmemDatabase.Dives = make([]*Dive, 1, 100)
	inmemDatabase.sourceToSystemID = make(map[string]int)
}

func (p *SubsurfaceCallbackHandler) HandleDive(ddh subsurface.DiveDataHolder) int {
	dive := &Dive{
		ID:         p.lastDiveID + 1,
		Number:     ddh.DiveNumber,
		DiveSiteID: inmemDatabase.sourceToSystemID[ddh.DiveSiteUUID],
		DiveTripID: ddh.DiveTripID,

		datetime: ddh.DateTime,
	}
	// TODO: assert dive number is not IntNull
	fmt.Printf("build: %v\n", dive)

	siteID, ok := inmemDatabase.sourceToSystemID[ddh.DiveSiteUUID]
	if ok {
		// TODO: assert site is not null
		dive.DiveSiteID = siteID
		fmt.Printf("link: %v -> %v\n", dive, inmemDatabase.DiveSites[siteID])
	} else {
		// TODO
	}

	// TODO: assert dive trip is not null
	dive.DiveTripID = ddh.DiveTripID
	fmt.Printf("link: %v -> %v\n", dive, inmemDatabase.DiveTrips[ddh.DiveTripID])

	inmemDatabase.Dives = append(inmemDatabase.Dives, dive)
	p.lastDiveID++
	return dive.ID
}

func (p *SubsurfaceCallbackHandler) HandleDiveSite(uuid string, name string, coords string) int {
	site := &DiveSite{
		ID:          p.lastSiteID + 1,
		Name:        name,
		Coordinates: coords,

		sourceID: uuid,
	}
	fmt.Printf("build: %v\n", site)

	inmemDatabase.sourceToSystemID[site.sourceID] = site.ID
	fmt.Printf("map: sourceToSystemID %q -> %d\n", site.sourceID, site.ID)

	inmemDatabase.DiveSites = append(inmemDatabase.DiveSites, site)
	p.lastSiteID++
	return site.ID
}

func (p *SubsurfaceCallbackHandler) HandleDiveTrip(label string) int {
	trip := &DiveTrip{
		ID:    p.lastTripID + 1,
		Label: label,
	}
	fmt.Printf("build: %v\n", trip)

	inmemDatabase.DiveTrips = append(inmemDatabase.DiveTrips, trip)
	p.lastTripID++
	return trip.ID
}

func (p *SubsurfaceCallbackHandler) HandleEnd() {
	// TODO: assert len(slice) == lastID
}

func (p *SubsurfaceCallbackHandler) HandleGeoData(siteID int, cat int, label string) {
	// TODO: assert site exists
	site := inmemDatabase.DiveSites[siteID]
	for _, lbl := range site.GeoLabels {
		if lbl == label {
			return
		}
	}
	site.GeoLabels = append(site.GeoLabels, label)
}

func (p *SubsurfaceCallbackHandler) HandleHeader(program string, version string) {
	inmemDatabase.Metadata.Program = program
	inmemDatabase.Metadata.ProgramVersion = version
	inmemDatabase.Metadata.Units = "metric" // DEVNOTE: make configurable?
}

func (p *SubsurfaceCallbackHandler) HandleSkip(element string) {
	// do nothing
}
