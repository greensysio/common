package direction

import (
	"bitbucket.org/greensys-tech/common/log"
	maps2 "bitbucket.org/greensys-tech/common/maps"
	"context"
	"errors"
	"os"

	cmContext "bitbucket.org/greensys-tech/common/context"
	"googlemaps.github.io/maps"
)

type (
	// RouteData holds the basic direction info from Direction api
	RouteData struct {
		Summary       string
		Warnings      []string
		OverviewPoint string
		Points        []*maps2.LatLng
	}
)

// RequestDirection request Google Direction api
func RequestDirection(ctx *cmContext.CustomContext, origin, destination maps2.LatLng, waypoints []maps2.LatLng, pointDensity int) (*RouteData, error) {
	var client *maps.Client
	var err error
	var apiKey = os.Getenv("GOOGLE_API_KEY")
	if len(apiKey) > 0 {
		client, err = maps.NewClient(maps.WithAPIKey(apiKey), maps.WithRateLimit(2))
	} else {
		log.Fatal("Please specify an API Key: GOOGLE_API_KEY in ENV")
		return nil, errors.New("please specify an API Key: GOOGLE_API_KEY in ENV")
	}
	if err != nil {
		return nil, err
	}

	r := &maps.DirectionsRequest{
		Origin:       origin.String(),
		Destination:  destination.String(),
		Alternatives: false,
		Language:     ctx.GetLocale(),
		Region:       "",
		Waypoints:    latLngsToString(waypoints),
	}

	routes, _, err := client.Directions(context.Background(), r)
	if err == nil && len(routes) > 0 {
		route := routes[0]
		result := RouteData{
			Summary:       route.Summary,
			Warnings:      route.Warnings,
			OverviewPoint: route.OverviewPolyline.Points,
			Points:        getLegsPoints(route, pointDensity),
		}
		return &result, nil
	}
	return nil, err
}

/*
* getLegsPoints return total points in LatLng from route.Legs
* density default = 1, (return all points)
 */
func getLegsPoints(route maps.Route, density int) []*maps2.LatLng {
	if density <= 0 {
		density = 1
	}
	var resultPoints []*maps2.LatLng
	lenLegs := len(route.Legs)
	for _, leg := range route.Legs {
		var totalLegPoints []*maps2.LatLng
		for _, step := range leg.Steps {
			points, _ := maps2.DecodePolyline(step.Points)
			// select points
			lenP := len(points)
			if lenP > 0 && density > 1 { // must check density > 1 for getting last point working exactly
				selectedPLen := int(lenP / density)
				if selectedPLen == 0 {
					selectedPLen = 1
				}
				selectedPoints := make([]*maps2.LatLng, selectedPLen+1) // +1 for last point
				lastIndex := -1
				for i, p := range points {
					if i%density == 0 {
						lastIndex++
						selectedPoints[lastIndex] = p
						if lastIndex == selectedPLen-1 {
							selectedPoints[lastIndex+1] = points[lenP-1] // get last point
							break
						}
					}
				}
				totalLegPoints = append(totalLegPoints, selectedPoints...)
			} else {
				totalLegPoints = append(totalLegPoints, points...)
			}
		}
		if (lenLegs) == 1 { // 1 leg (no next legs)
			resultPoints = totalLegPoints
		} else {
			resultPoints = append(resultPoints, totalLegPoints...)
		}
	}
	return resultPoints
}
