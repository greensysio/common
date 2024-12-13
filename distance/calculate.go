package distance

import (
	"bytes"
	"fmt"
	"time"

	"github.com/greensysio/common/log"
	"googlemaps.github.io/maps"
)

// ------------------ Input struct and Input preprocessing ------------------------

// Location is the same as a point (lat, Lng)
type Location struct {
	Lat float64
	Lng float64
}

// Distance stores Origin Location and Destination Location
type Distance struct {
	Origin      Location
	Destination Location
}

// ToString format Location as a standard format Lat,Lng
func (l Location) ToString() string {
	return fmt.Sprint(l.Lat, ",", l.Lng)
}

func formatLocationParam(locations []Location) string {
	locationLen := len(locations)
	var buffer bytes.Buffer
	for i := 0; i < locationLen; i++ {
		if i > 0 {
			buffer.WriteString("|")
		}
		buffer.WriteString(locations[i].ToString())
	}
	return buffer.String()
}

func formatDistances(distances []Distance) (string, string) {
	var originLocations []Location
	var destinationLocations []Location
	for _, distance := range distances {
		originLocations = append(originLocations, distance.Origin)
		destinationLocations = append(destinationLocations, distance.Destination)
	}
	originsFormatted := formatLocationParam(originLocations)
	destinationFormatted := formatLocationParam(destinationLocations)
	return originsFormatted, destinationFormatted
}

// ----------------- Output struct and Output preprocessing

// DistanceRow including distance & duration of a distance
type DistanceRow struct {
	// Distance in meter
	Distance int
	// The length of time it takes to travel this route, expressed in seconds
	Duration time.Duration
}

// DistanceResult includes arrays of DistanceRows, TotalDistance and TotalDuration
type DistanceResult struct {
	DistanceRows []DistanceRow
	// unit: meter
	TotalDistance int
	// unit: second
	TotalDuration float64
}

// StatusCode includes status codes corresponding Google Distance Matrix response
type StatusCode string

const (
	// OK Status OK
	OK StatusCode = "OK"
	// Max destination per request
	MaxElementsExceeded int = 10
)

type errorMsg struct {
	code  string
	cause string
}

func (e errorMsg) Error() string {
	return fmt.Sprintf("code=%s, cause= %s", e.code, e.cause)
}

// CalculateDistances caculates distances
// return a DistanceResult or an error
func CalculateDistances(distances []Distance) (*DistanceResult, error) {
	resp := &DistanceResult{}
	loopNumber := int(len(distances) / (MaxElementsExceeded + 1))
	for i := 0; i <= loopNumber; i++ {
		temp := make([]Distance, MaxElementsExceeded)
		if i != loopNumber {
			temp = distances[i*MaxElementsExceeded : (i+1)*MaxElementsExceeded]
		} else {
			// Last loop
			temp = distances[i*MaxElementsExceeded:]
		}

		originsFormatted, destinationsFormatted := formatDistances(temp)
		respTemp, err := calculateDistance(originsFormatted, destinationsFormatted)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		proRsTemp, err := processResult(respTemp)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		resp.DistanceRows = append(resp.DistanceRows, proRsTemp.DistanceRows...)
		resp.TotalDistance += proRsTemp.TotalDistance
		resp.TotalDuration += proRsTemp.TotalDuration
	}
	return resp, nil
}

func processResult(resp *maps.DistanceMatrixResponse) (*DistanceResult, error) {
	var usedElements []*maps.DistanceMatrixElement
	// lookup useful elements
	for rowIndex, elements := range resp.Rows {
		element := elements.Elements[rowIndex]
		if StatusCode(element.Status) != OK {
			log.Debugf("Distance_Matrix_Responses: a response element failed, status=%s", element.Status)
			return nil, errorMsg{element.Status, "Google Distance Matrix api return error status"}
		}
		usedElements = append(usedElements, element)
	}
	// normalize the result
	var distanceRow []DistanceRow
	var totalDistance = 0
	var totalDuration float64 = 0
	for _, resultElement := range usedElements {
		distanceRow = append(distanceRow,
			DistanceRow{
				Distance: resultElement.Distance.Meters,
				Duration: resultElement.Duration,
			})
		totalDistance += resultElement.Distance.Meters
		totalDuration += resultElement.Duration.Seconds()
	}

	return &DistanceResult{DistanceRows: distanceRow, TotalDistance: totalDistance, TotalDuration: totalDuration}, nil
}
