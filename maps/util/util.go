package util

import "fmt"

func GetMapsShowMarkerURL(lat, lng float64) string {
	return fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%g,%g", lat, lng)
}
