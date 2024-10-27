package server

import (
	"strings"
	"testing"
)

const (
	HeaderEx = ` <divelog program='subsurface' version='3'>
<settings>
<fingerprint model='9b2750e4' serial='00001f47' deviceid='e92802e6' diveid='5f70bc90' data='0b001f00120008007c00'/>
</settings><divesites></divesites></divelog> `

	DiveSiteEx_1 = ` <site uuid='dc3051aa' name='Green Bay, Protaras' gps='35.001439 34.067126'>
  <geo cat='1' origin='0' value='Mediterranean Sea, Eastern Basin'/>
  <geo cat='2' origin='0' value='Cyprus'/>
  <geo cat='3' origin='0' value='Ammochostos'/>
  <geo cat='5' origin='0' value='Protaras'/>
  <geo cat='6' origin='3' value='Protaras'/>
</site> `
)

var (
	DiveSite_1 = &SiteXML{
		UUID: "dc3051aa",
		Name: "Green Bay, Protaras",
		GPS:  "35.001439 34.067126",
	}
)

func TestDecodeHeader(t *testing.T) {
	diveLog, err := DecodeSubsurfaceDatabase(strings.NewReader(HeaderEx))
	if err != nil {
		t.Fatalf("failed to decode header: %v", err)
	}
	t.Logf("decoded header: program=%q version=%q", diveLog.Program, diveLog.Version)
}

func TestDecodeSite(t *testing.T) {
	validate := func(ex string, expected *SiteXML) {
		decoder := NewSubDecoder(strings.NewReader(ex))
		start, err := decoder.ExpectStart("site")
		if err != nil {
			t.Fatalf("failed to decode first token: %v (ex: %s)", err, ex)
		}
		decoded, err := DecodeSiteXML(decoder, start)
		if err != nil {
			t.Fatalf("failed to decode the site element: %v (ex: %s)", err, ex)
		}
		if decoded.UUID != expected.UUID {
			t.Fatalf("expected UUID %q got %q", expected.UUID, decoded.UUID)
		}
		if decoded.Name != expected.Name {
			t.Fatalf("expected Name %q got %q", expected.Name, decoded.Name)
		}
		if decoded.GPS != expected.GPS {
			t.Fatalf("expected GPS %q got %q", expected.GPS, decoded.GPS)
		}
		t.Logf("%v", decoded)
		t.Logf("successfully decoded %q", ex)
	}

	validate(DiveSiteEx_1, DiveSite_1)
}
