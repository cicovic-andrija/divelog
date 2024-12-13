package subsurface

import (
	"encoding/xml"
)

type SiteXML struct {
	XMLName xml.Name `xml:"site"`
	UUID    string   `xml:"uuid,attr"`
	Name    string   `xml:"name,attr"`
	GPS     string   `xml:"gps,attr"`
	Geos    []GeoXML `xml:"geo"`
}

type GeoXML struct {
	Cat   string `xml:"cat,attr"`
	Value string `xml:"value,attr"`
}

type DiveXML struct {
	Number            string               `xml:"number,attr"`
	Rating            string               `xml:"rating,attr"`
	Visibility        string               `xml:"visibility,attr"`
	SAC               string               `xml:"sac,attr"`
	Tags              string               `xml:"tags,attr"`
	DiveSiteUUID      string               `xml:"divesiteid,attr"`
	WaterSalinity     string               `xml:"watersalinity,attr"`
	Date              string               `xml:"date,attr"`
	Time              string               `xml:"time,attr"`
	Duration          string               `xml:"duration,attr"`
	DiveMaster        string               `xml:"divemaster"`
	Buddy             string               `xml:"buddy"`
	Notes             string               `xml:"notes"`
	Suit              string               `xml:"suit"`
	Cylinder          CylinderXML          `xml:"cylinder"`
	WeightSystem      WeightSystemXML      `xml:"weightsystem"`
	TemperatureManual TemperatureManualXML `xml:"divetemperature"`
	DiveComputer      DiveComputerXML      `xml:"divecomputer"`
}

type CylinderXML struct {
	Size         string `xml:"size,attr"`
	WorkPressure string `xml:"workpressure,attr"`
	Description  string `xml:"description,attr"`
	Start        string `xml:"start,attr"`
	End          string `xml:"end,attr"`
	O2           string `xml:"o2,attr"`
}

type WeightSystemXML struct {
	Weight      string `xml:"weight,attr"`
	Description string `xml:"description,attr"`
}

type TemperatureManualXML struct {
	Air   string `xml:"air,attr"`
	Water string `xml:"water,attr"`
}

type DiveComputerXML struct {
	Model           string             `xml:"model,attr"`
	DeviceID        string             `xml:"deviceid,attr"`
	DiveID          string             `xml:"diveid,attr"`
	DepthInfo       DepthInfoXML       `xml:"depth"`
	TemperatureInfo TemperatureInfoXML `xml:"temperature"`
	SurfaceInfo     SurfaceInfoXML     `xml:"surface"`
}

type DepthInfoXML struct {
	Max  string `xml:"max,attr"`
	Mean string `xml:"mean,attr"`
}

type TemperatureInfoXML struct {
	WaterMin string `xml:"water,attr"`
}

type SurfaceInfoXML struct {
	Pressure string `xml:"pressure,attr"`
}
