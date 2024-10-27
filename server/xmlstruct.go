package server

import (
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

var (
	ErrNilReader     = errors.New("io.Reader is nil")
	ErrInvalidFormat = errors.New("XML database is not in the valid format")
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

func DecodeSubsurfaceDatabase(r io.Reader) (*DiveLog, error) {
	if r == nil {
		return nil, ErrNilReader
	}

	var (
		divelog  = &DiveLog{}
		decoder  = NewSubDecoder(r)
		startTag *xml.StartElement
		err      error
	)

	// <divelog ...>
	if startTag, err = decoder.ExpectStart("divelog"); err != nil {
		return nil, err
	}
	if program, ok := FindAttribute(startTag, "program"); ok {
		divelog.Program = program
	}
	if version, ok := FindAttribute(startTag, "version"); ok {
		divelog.Version = version
	}

	// <settings>
	if err = decoder.SkipElement("settings"); err != nil {
		return nil, err
	}
	// </settings>

	// <divesites>
	if _, err = decoder.ExpectStart("divesites"); err != nil {
		return nil, err
	}

	for {
		if startTag, err = decoder.NextOrEnd("site", "divesites"); err != nil {
			return nil, err
		}
		if startTag != nil {
			// <site ...> 1..N
			if siteXML, err := DecodeSiteXML(decoder, startTag); err != nil {
				return nil, err
			} else {
				divelog.DiveSites = append(divelog.DiveSites, &DiveSite{
					ProgramID:   siteXML.UUID,
					Name:        siteXML.Name,
					Coordinates: siteXML.GPS,
				})
			}
			// </site>
		} else {
			// </divesites>
			break
		}
	}

	// </divelog>
	return divelog, nil
}

func DecodeSiteXML(decoder *SubDecoder, tok *xml.StartElement) (*SiteXML, error) {
	siteXML := &SiteXML{}
	err := decoder.XMLDecoder.DecodeElement(siteXML, tok)
	return siteXML, err
}

func NewSubDecoder(r io.Reader) *SubDecoder {
	return &SubDecoder{
		XMLDecoder: xml.NewDecoder(r),
	}
}

func (d *SubDecoder) Token() (tok xml.Token, err error) {
	for {
		tok, err = d.XMLDecoder.Token()
		if err != nil {
			return
		}
		switch t := tok.(type) {
		case xml.CharData:
			if len(strings.TrimSpace(string(t))) > 0 {
				return
			}
		default:
			return
		}
	}
}

func (d *SubDecoder) ExpectStart(tag string) (*xml.StartElement, error) {
	tok, err := d.Token()
	if err != nil {
		return nil, ErrInvalidFormat
	}
	switch tok := tok.(type) {
	case xml.StartElement:
		if tok.Name.Local == tag {
			return &tok, nil
		}
	}
	return nil, ErrInvalidFormat
}

func (d *SubDecoder) ExpectAnyStart() (*xml.StartElement, error) {
	tok, err := d.Token()
	if err != nil {
		return nil, ErrInvalidFormat
	}
	switch tok := tok.(type) {
	case xml.StartElement:
		return &tok, nil
	}
	return nil, ErrInvalidFormat
}

func (d *SubDecoder) ExpectEnd(tag string) (*xml.EndElement, error) {
	tok, err := d.Token()
	if err != nil {
		return nil, ErrInvalidFormat
	}
	switch tok := tok.(type) {
	case xml.EndElement:
		if tok.Name.Local == tag {
			return &tok, nil
		}
	}
	return nil, ErrInvalidFormat
}

func (d *SubDecoder) NextOrEnd(next string, end string) (*xml.StartElement, error) {
	tok, err := d.Token()
	if err != nil {
		return nil, ErrInvalidFormat
	}
	switch tok := tok.(type) {
	case xml.StartElement:
		if tok.Name.Local == next {
			return &tok, nil
		}
	case xml.EndElement:
		if tok.Name.Local == end {
			return nil, nil
		}
	}
	return nil, ErrInvalidFormat
}

func (d *SubDecoder) SkipElement(tag string) error {
	if _, err := d.ExpectStart(tag); err != nil {
		return err
	}
	return d.XMLDecoder.Skip()
}

func FindAttribute(tok *xml.StartElement, name string) (val string, ok bool) {
	for _, attr := range tok.Attr {
		if attr.Name.Local == name {
			ok = true
			val = attr.Value
			return
		}
	}
	return
}
