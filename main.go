package main

import (
	"fmt"
	"github.com/kellydunn/golang-geo"
)

func main() {
	geo.SetGoogleAPIKey("")
	p := geo.NewPoint(42.25, 120.2)
	res, err := geo.GoogleGeocoder.ReverseGeocode(p)
	if err != nil {
		fmt.Errorf("Error creating query string: %v", err)
	}
}
