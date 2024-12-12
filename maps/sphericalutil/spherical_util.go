package sphericalutil

import (
	"bitbucket.org/greensys-tech/common/maps"
	"bitbucket.org/greensys-tech/common/maps/mathutil"
	"math"
)

/**
 * Returns the heading from one LatLng to another LatLng. Headings are
 * expressed in degrees clockwise from North within the range [-180,180).
 *
 * @return The heading in degrees clockwise from north.
 */
func ComputeHeading(from, to maps.LatLng) float64 {
	// http://williams.best.vwh.net/avform.htm#Crs
	fromLat := mathutil.ToRadians(from.Lat)
	fromLng := mathutil.ToRadians(from.Lng)
	toLat := mathutil.ToRadians(to.Lat)
	toLng := mathutil.ToRadians(to.Lng)
	dLng := toLng - fromLng
	heading := math.Atan2(math.Sin(dLng)*math.Cos(toLat), math.Cos(fromLat)*math.Sin(toLat)-math.Sin(fromLat)*math.Cos(toLat)*math.Cos(dLng))
	return mathutil.Wrap(mathutil.ToDegrees(heading), -180, 180)
}

/**
 * Returns the LatLng resulting from moving a distance from an origin
 * in the specified heading (expressed in degrees clockwise from north).
 *
 * @param from     The LatLng from which to start.
 * @param distance The distance to travel.
 * @param heading  The heading in degrees clockwise from north.
 */
func computeOffset(from *maps.LatLng, distance, heading float64) *maps.LatLng {
	distance /= mathutil.EarthRadius
	heading = mathutil.ToRadians(heading)
	// http://williams.best.vwh.net/avform.htm#LL
	fromLat := mathutil.ToRadians(from.Lat)
	fromLng := mathutil.ToRadians(from.Lng)
	cosDistance := math.Cos(distance)
	sinDistance := math.Sin(distance)
	sinFromLat := math.Sin(fromLat)
	cosFromLat := math.Cos(fromLat)
	sinLat := cosDistance*sinFromLat + sinDistance*cosFromLat*math.Cos(heading)
	dLng := math.Atan2(sinDistance*cosFromLat*math.Sin(heading), cosDistance-sinFromLat*sinLat)
	return &maps.LatLng{Lat: mathutil.ToDegrees(math.Asin(sinLat)), Lng: mathutil.ToDegrees(fromLng + dLng)}
}
