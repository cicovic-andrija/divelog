package server

type DiveLog struct {
	Program        string
	Version        string
	DiveSites      []*DiveSite
	Dives          []*Dive
	programIDIDMap map[string]int
}

type DiveSite struct {
	ID          int    `json:"id"`
	ProgramID   string `json:"program_id"`
	Name        string `json:"name"`
	Coordinates string `json:"coordinates"`
}

type Dive struct {
	ID         int `json:"id"`
	Number     int `json:"number"`
	DiveSiteID int `json:"dive_site_id"`
}
