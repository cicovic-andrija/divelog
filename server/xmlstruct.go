package server

import (
	"encoding/xml"
)

type SubDecoder struct {
	XMLDecoder *xml.Decoder
}

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
