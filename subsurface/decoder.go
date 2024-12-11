package subsurface

import (
	"encoding/xml"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	ErrNilReader     = errors.New("io.Reader is nil")
	ErrNilHandler    = errors.New("subsurface.Handler is nil")
	ErrInvalidFormat = errors.New("XML database is not in the valid format")
)

type Handler interface {
	HandleBegin()
	HandleEnd()
	HandleHeader(program string, version string)
	HandleSkip(element string)
	HandleDiveSite(uuid string, name string, coords string) int
	HandleGeoData(siteID int, cat int, label string)
}

type Decoder struct {
	XMLDecoder *xml.Decoder
}

func DecodeSubsurfaceDatabase(r io.Reader, h Handler) error {
	if r == nil {
		return ErrNilReader
	}
	if h == nil {
		return ErrNilHandler
	}

	var (
		decoder = &Decoder{
			XMLDecoder: xml.NewDecoder(r),
		}
		startTag *xml.StartElement
		err      error
	)

	// <divelog ...>
	if startTag, err = decoder.ExpectStart("divelog"); err != nil {
		return err
	}
	h.HandleBegin()

	program, _ := FindAttribute(startTag, "program")
	version, _ := FindAttribute(startTag, "version")
	h.HandleHeader(program, version)

	// <settings>
	if err = decoder.SkipElement("settings"); err != nil {
		return err
	}
	h.HandleSkip("settings")
	// </settings>

	// <divesites>
	if _, err = decoder.ExpectStart("divesites"); err != nil {
		return err
	}

	for {
		if startTag, err = decoder.NextOrEnd("site", "divesites"); err != nil {
			return err
		}
		if startTag != nil {
			// <site ...> 1..N
			if siteXML, err := DecodeSiteXML(decoder, startTag); err != nil {
				return err
			} else {
				siteID := h.HandleDiveSite(siteXML.UUID, siteXML.Name, siteXML.GPS)
				for _, geoData := range siteXML.Geos {
					if cat, err := strconv.Atoi(geoData.Cat); err == nil {
						h.HandleGeoData(siteID, cat, geoData.Value)
					} else {
						return ErrInvalidFormat
					}

				}
			}
			// </site>
		} else {
			// </divesites>
			break
		}
	}

	h.HandleEnd()
	// </divelog>

	return nil
}

func DecodeSiteXML(decoder *Decoder, tok *xml.StartElement) (*SiteXML, error) {
	siteXML := &SiteXML{}
	err := decoder.XMLDecoder.DecodeElement(siteXML, tok)
	return siteXML, err
}

func (d *Decoder) Token() (tok xml.Token, err error) {
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

func (d *Decoder) ExpectStart(tag string) (*xml.StartElement, error) {
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

func (d *Decoder) ExpectAnyStart() (*xml.StartElement, error) {
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

func (d *Decoder) ExpectEnd(tag string) (*xml.EndElement, error) {
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

func (d *Decoder) NextOrEnd(next string, end string) (*xml.StartElement, error) {
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

func (d *Decoder) SkipElement(tag string) error {
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
