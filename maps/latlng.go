package maps

import (
	"bitbucket.org/greensys-tech/common/maps/mathutil"
	"math"
	"strconv"
	"strings"
)

// LatLng represents a location on the Earth.
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// ParseLatLng will parse a string representation of a Lat,Lng pair.
func ParseLatLng(location string) (LatLng, error) {
	l := strings.Split(location, ",")
	lat, err := strconv.ParseFloat(l[0], 64)
	if err != nil {
		return LatLng{}, err
	}
	lng, err := strconv.ParseFloat(l[1], 64)
	if err != nil {
		return LatLng{}, err
	}
	return LatLng{Lat: lat, Lng: lng}, nil
}

func (l *LatLng) String() string {
	return strconv.FormatFloat(l.Lat, 'f', -1, 64) +
		"," +
		strconv.FormatFloat(l.Lng, 'f', -1, 64)
}

// CalculateDistance2Points function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func CalculateDistance2Points(lat1, lng1, lat2, lng2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2 float64
	la1 = lat1 * math.Pi / 180
	lo1 = lng1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lng2 * math.Pi / 180

	// calculate
	h := mathutil.Hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*mathutil.Hsin(lo2-lo1)

	return math.Abs(2 * mathutil.EarthRadius * math.Asin(math.Sqrt(h)))
}
