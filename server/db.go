package server

import (
	"fmt"
	"os"
	"time"

	"src.acicovic.me/divelog/subsurface"
)

var _inmemDatabase DiveLog

func buildDatabase() error {
	file, err := os.Open(_inmemDatabase.Metadata.Source)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", _inmemDatabase.Metadata.Source, err)
	}
	defer file.Close()

	if err = subsurface.DecodeSubsurfaceDatabase(file, &SubsurfaceCallbackHandler{}); err != nil {
		return fmt.Errorf("failed to decode database in %s: %v", _inmemDatabase.Metadata.Source, err)
	}

	return nil
}

type SubsurfaceCallbackHandler struct {
	lastSiteID int
	lastTripID int
	lastDiveID int
}

func (p *SubsurfaceCallbackHandler) HandleBegin() {
	_inmemDatabase.DiveSites = make([]*DiveSite, 1, 100)
	_inmemDatabase.DiveTrips = make([]*DiveTrip, 1, 100)
	_inmemDatabase.Dives = make([]*Dive, 1, 100)
	_inmemDatabase.sourceToSystemID = make(map[string]int)
}

func (p *SubsurfaceCallbackHandler) HandleDive(ddh subsurface.DiveDataHolder) int {
	dive := &Dive{
		ID:     p.lastDiveID + 1,
		Number: ddh.DiveNumber,

		Rating5:         ddh.Rating,
		Visibility5:     ddh.Visibility,
		Tags:            ddh.Tags,
		Salinity:        ddh.WaterSalinity,
		DateTimeIn:      ddh.DateTime.Format(time.RFC3339),
		OperatorDM:      ddh.DiveMasterOrOperator,
		Buddy:           ddh.Buddy,
		Notes:           ddh.Notes,
		Suit:            ddh.Suit,
		CylSize:         ddh.CylinderSize,
		CylType:         ddh.CylinderDescription,
		StartPressure:   ddh.CylinderStartPressure,
		EndPressure:     ddh.CylinderEndPressure,
		Gas:             ddh.CylinderGas,
		Weights:         ddh.Weight,
		WeightsType:     ddh.WeightType,
		DCModel:         ddh.DiveComputerModel,
		DepthMax:        ddh.DepthMax,
		DepthMean:       ddh.DepthMean,
		TempWaterMin:    ddh.TemperatureWaterMin,
		TempAir:         ddh.TemperatureAir,
		SurfacePressure: ddh.SurfacePressure,

		datetime: ddh.DateTime,
	}
	fmt.Printf("build: %v\n", dive)
	assert(dive.ID == len(_inmemDatabase.Dives), "invalid Dive.ID")

	siteID, ok := _inmemDatabase.sourceToSystemID[ddh.DiveSiteUUID]
	assert(ok, "DiveDataHolder.DiveSiteUUID is not mapped to DiveSite.ID")
	dive.DiveSiteID = siteID
	assert(siteID > 0 && siteID < len(_inmemDatabase.DiveSites), "invalid dive site ID mapping")
	assert(_inmemDatabase.DiveSites[siteID] != nil, "DiveSite ptr is nil")
	fmt.Printf("link: %v -> %v\n", dive, _inmemDatabase.DiveSites[siteID])

	dive.DiveTripID = ddh.DiveTripID
	assert(ddh.DiveTripID > 0 && ddh.DiveTripID < len(_inmemDatabase.DiveTrips), "invalid dive trip ID")
	assert(_inmemDatabase.DiveTrips[ddh.DiveTripID] != nil, "DiveTrip ptr is nil")
	fmt.Printf("link: %v -> %v\n", dive, _inmemDatabase.DiveTrips[ddh.DiveTripID])

	dive.Normalize()

	_inmemDatabase.Dives = append(_inmemDatabase.Dives, dive)
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
	assert(site.ID == len(_inmemDatabase.DiveSites), "invalid DiveSite.ID")

	_inmemDatabase.sourceToSystemID[site.sourceID] = site.ID
	fmt.Printf("map: sourceToSystemID %q -> %d\n", site.sourceID, site.ID)

	_inmemDatabase.DiveSites = append(_inmemDatabase.DiveSites, site)
	p.lastSiteID++

	return site.ID
}

func (p *SubsurfaceCallbackHandler) HandleDiveTrip(label string) int {
	trip := &DiveTrip{
		ID:    p.lastTripID + 1,
		Label: label,
	}
	fmt.Printf("build: %v\n", trip)
	assert(trip.ID == len(_inmemDatabase.DiveTrips), "invalid DiveTrip.ID")

	_inmemDatabase.DiveTrips = append(_inmemDatabase.DiveTrips, trip)
	p.lastTripID++

	return trip.ID
}

func (p *SubsurfaceCallbackHandler) HandleEnd() {
	assert(len(_inmemDatabase.Dives)-1 == p.lastDiveID, "invalid Dives slice length")
	assert(len(_inmemDatabase.DiveSites)-1 == p.lastSiteID, "invalid DiveSites slice length")
	assert(len(_inmemDatabase.DiveTrips)-1 == p.lastTripID, "invalid DiveTrips slice length")
}

func (p *SubsurfaceCallbackHandler) HandleGeoData(siteID int, cat int, label string) {
	assert(_inmemDatabase.DiveSites[siteID] != nil, "DiveSite ptr is nil")
	site := _inmemDatabase.DiveSites[siteID]
	for _, lbl := range site.GeoLabels {
		if lbl == label {
			return
		}
	}
	site.GeoLabels = append(site.GeoLabels, label)
}

func (p *SubsurfaceCallbackHandler) HandleHeader(program string, version string) {
	_inmemDatabase.Metadata.Program = program
	_inmemDatabase.Metadata.ProgramVersion = version
	_inmemDatabase.Metadata.Units = "metric" // DEVNOTE: make configurable?
}

func (p *SubsurfaceCallbackHandler) HandleSkip(element string) {
	// do nothing
}
