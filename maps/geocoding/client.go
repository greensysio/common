// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main contains a simple command line tool for Geocoding API
// Documentation: https://developers.google.com/maps/documentation/geocoding/
package geocoding

import (
	"github.com/greensysio/common/log"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	cmContext "github.com/greensysio/common/context"
	myMaps "github.com/greensysio/common/maps"
	"googlemaps.github.io/maps"
)

type (
	// GeocodingRequest request params
	GeocodingRequest struct {
		Address  string
		Ward     *string
		District *string
		Province *string
		Country  *string
		LatLng   *myMaps.LatLng
	}
	// GeocodingResult hold Lat, Lng, ...
	GeocodingResult struct {
		Location         myMaps.LatLng
		PlaceID          string
		FormattedAddress string
	}
)

// GeoCode geocode (Lat, Lng, ...) from address
func GeoCode(c *cmContext.CustomContext, request GeocodingRequest) (*GeocodingResult, error) {
	var client *maps.Client
	var err error
	var apiKey = os.Getenv("GOOGLE_API_KEY")
	if len(apiKey) > 0 {
		client, err = maps.NewClient(maps.WithAPIKey(apiKey), maps.WithRateLimit(2))
	} else {
		log.Fatal("Please specify an API Key: GOOGLE_API_KEY in ENV")
		os.Exit(2)
		return nil, errors.New("Please specify an API Key: GOOGLE_API_KEY in ENV")
	}
	if err != nil {
		return nil, err
	}

	r := &maps.GeocodingRequest{
		Address:  request.Address,
		Language: c.GetLocale(),
		Region:   "",
	}
	r.Components = make(map[maps.Component]string)
	if request.Ward != nil {
		r.Components[maps.ComponentLocality] = *request.Ward
	}
	if request.District != nil && request.Province != nil {
		r.Components[maps.ComponentAdministrativeArea] = fmt.Sprintf("%s, %s", *request.District, *request.Province)
	} else {
		if request.District != nil {
			r.Components[maps.ComponentAdministrativeArea] = *request.District
		} else if request.Province != nil {
			r.Components[maps.ComponentAdministrativeArea] = *request.Province
		}
	}
	if request.LatLng != nil {
		r.LatLng = &maps.LatLng{
			Lat: request.LatLng.Lat,
			Lng: request.LatLng.Lng,
		}
	}

	resp, err := client.Geocode(context.Background(), r)
	if err == nil && len(resp) > 0 {
		address := resp[0].FormattedAddress
		lastIndex := strings.LastIndex(resp[0].FormattedAddress, ",") // exclude country
		if lastIndex > 0 {
			address = resp[0].FormattedAddress[0:lastIndex]
		}

		return &GeocodingResult{
			FormattedAddress: address,
			PlaceID:          resp[0].PlaceID,
			Location: myMaps.LatLng{
				Lat: resp[0].Geometry.Location.Lat,
				Lng: resp[0].Geometry.Location.Lng,
			},
		}, nil
	}
	return nil, err
}
