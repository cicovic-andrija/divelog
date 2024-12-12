package main

import (
	"fmt"
	"os"
	"path/filepath"

	"src.acicovic.me/divelog/subsurface"
)

// Subsurface Decoder Validator

func main() {
	if len(os.Args) != 2 {
		fmt.Println("provide file name as the first program argument")
		os.Exit(0x1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("failed to open file: %v", err)
		os.Exit(0x2)
	}
	defer file.Close()

	if err := subsurface.DecodeSubsurfaceDatabase(file, Handler{fname: os.Args[1]}); err != nil {
		fmt.Printf("decoding error: %v", err)
		os.Exit(0x3)
	}
}

type Handler struct {
	fname string
}

func (h Handler) HandleBegin() {
	fmt.Printf("SUBSURFACE_DATABASE %q\n", filepath.Base(h.fname))
}

func (h Handler) HandleEnd() {
	fmt.Printf("END.\n")
}

func (h Handler) HandleSkip(element string) {
	fmt.Printf(">>>>>\nSKIP ELEMENT %q\n<<<<<\n", element)
}

func (h Handler) HandleHeader(program string, version string) {
	fmt.Printf("\tHEADER\n")
	fmt.Printf("\t\tPROGRAM = %q\n", program)
	fmt.Printf("\t\tVERSION = %q\n", version)

}

func (h Handler) HandleDiveSite(uuid string, name string, coords string) int {
	fmt.Printf("\tDIVE_SITE\n")
	fmt.Printf("\t\tUUID = %q\n\t\tNAME = %q\n\t\tCOORDS = %q\n", uuid, name, coords)
	return 0
}

func (h Handler) HandleGeoData(id int, cat int, label string) {
	fmt.Printf("\t\tGEO_DATA\n")
	fmt.Printf("\t\t\tCATEGORY = %d\n\t\t\tLABEL = %q\n", cat, label)
}

func (h Handler) HandleDiveTrip(label string) int {
	fmt.Printf("\tDIVE_TRIP %q\n", label)
	return 0
}

func (h Handler) HandleDive(tripID int, diveNum string) int {
	fmt.Printf("\t\tDIVE\n")
	fmt.Printf("\t\t\tNUMBER = %s\n", diveNum)
	return 0
}
