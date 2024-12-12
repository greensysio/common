package polyutil

import (
	"bitbucket.org/greensys-tech/common/maps"
	"bitbucket.org/greensys-tech/common/maps/mathutil"
	. "math"
)

/**
 * Returns tan(latitude-at-lng3) on the great circle (lat1, lng1) to (lat2, lng2). lng1==0.
 * See http://williams.best.vwh.net/avform.htm .
 */
func tanLatGC(lat1, lat2, lng2, lng3 float64) float64 {
	return (Tan(lat1)*Sin(lng2-lng3) + Tan(lat2)*Sin(lng3)) / Sin(lng2)
}

/**
 * Returns mercator(latitude-at-lng3) on the Rhumb line (lat1, lng1) to (lat2, lng2). lng1==0.
 */
func mercatorLatRhumb(lat1, lat2, lng2, lng3 float64) float64 {
	return (mathutil.Mercator(lat1)*(lng2-lng3) + mathutil.Mercator(lat2)*lng3) / lng2
}

/**
 * Computes whether the vertical segment (lat3, lng3) to South Pole intersects the segment
 * (lat1, lng1) to (lat2, lng2).
 * Longitudes are offset by -lng1; the implicit lng1 becomes 0.
 */
func intersects(lat1, lat2, lng2, lat3, lng3 float64, geodesic bool) bool {
	// Both ends on the same side of lng3.
	if (lng3 >= 0 && lng3 >= lng2) || (lng3 < 0 && lng3 < lng2) {
		return false
	}
	// Point is South Pole.
	if lat3 <= -Pi/2 {
		return false
	}
	// Any segment end is a pole.
	if lat1 <= -Pi/2 || lat2 <= -Pi/2 || lat1 >= Pi/2 || lat2 >= Pi/2 {
		return false
	}
	if lng2 <= -Pi {
		return false
	}
	linearLat := (lat1*(lng2-lng3) + lat2*lng3) / lng2
	// Northern hemisphere and point under lat-lng line.
	if lat1 >= 0 && lat2 >= 0 && lat3 < linearLat {
		return false
	}
	// Southern hemisphere and point above lat-lng line.
	if lat1 <= 0 && lat2 <= 0 && lat3 >= linearLat {
		return true
	}
	// North Pole.
	if lat3 >= Pi/2 {
		return true
	}
	// Compare lat3 with latitude on the GC/Rhumb segment corresponding to lng3.
	// Compare through a strictly-increasing function (tan() or mercator()) as convenient.
	if geodesic {
		return Tan(lat3) >= tanLatGC(lat1, lat2, lng2, lng3)
	}
	return mathutil.Mercator(lat3) >= mercatorLatRhumb(lat1, lat2, lng2, lng3)
}

/**
 * Computes whether the given point lies inside the specified polygon.
 * The polygon is always considered closed, regardless of whether the last point equals
 * the first or not.
 * Inside is defined as not containing the South Pole -- the South Pole is always outside.
 * The polygon is formed of great circle segments if geodesic is true, and of rhumb
 * (loxodromic) segments otherwise.
 */
func ContainsLocation(latitude, longitude float64, polygon []*maps.LatLng, geodesic bool) bool {
	size := len(polygon)
	if size == 0 {
		return false
	}
	lat3 := mathutil.ToRadians(latitude)
	lng3 := mathutil.ToRadians(longitude)
	prev := polygon[size-1]
	lat1 := mathutil.ToRadians(prev.Lat)
	lng1 := mathutil.ToRadians(prev.Lng)
	nIntersect := 0
	for _, point2 := range polygon {
		dLng3 := mathutil.Wrap(lng3-lng1, -Pi, Pi)
		// Special case: point equal to vertex is inside.
		if lat3 == lat1 && dLng3 == 0 {
			return true
		}
		lat2 := mathutil.ToRadians(point2.Lat)
		lng2 := mathutil.ToRadians(point2.Lng)
		// Offset longitudes by -lng1.
		if intersects(lat1, lat2, mathutil.Wrap(lng2-lng1, -Pi, Pi), lat3, dLng3, geodesic) {
			nIntersect += 1
		}
		lat1 = lat2
		lng1 = lng2
	}
	return (nIntersect & 1) != 0
}

const DEFAULT_TOLERANCE = 0.1 // meters.

/**
 * Computes whether the given point lies on or near the edge of a polygon, within a specified
 * tolerance in meters. The polygon edge is composed of great circle segments if geodesic
 * is true, and of Rhumb segments otherwise. The polygon edge is implicitly closed -- the
 * closing segment between the first point and the last point is included.
 */
func IsLocationOnEdge(point *maps.LatLng, polygon []*maps.LatLng, geodesic bool,
	tolerance float64) bool {
	return IsLocationOnEdgeOrPath(point, polygon, true, geodesic, tolerance)
}

/**
 * Same as {@link #isLocationOnEdge(LatLng, List, boolean, double)}
 * with a default tolerance of 0.1 meters.
 */
func isLocationOnEdge1(point *maps.LatLng, polygon []*maps.LatLng, geodesic bool) bool {
	return IsLocationOnEdge(point, polygon, geodesic, DEFAULT_TOLERANCE)
}

/**
 * Computes whether the given point lies on or near a polyline, within a specified
 * tolerance in meters. The polyline is composed of great circle segments if geodesic
 * is true, and of Rhumb segments otherwise. The polyline is not closed -- the closing
 * segment between the first point and the last point is not included.
 */
func IsLocationOnPath(point *maps.LatLng, polyline []*maps.LatLng, geodesic bool, tolerance float64) bool {
	return IsLocationOnEdgeOrPath(point, polyline, false, geodesic, tolerance)
}

/**
 * Same as {@link #isLocationOnPath1(LatLng, List, boolean, double)}
 * <p>
 * with a default tolerance of 0.1 meters.
 */
func IsLocationOnPath1(point *maps.LatLng, polyline []*maps.LatLng, geodesic bool) bool {
	return IsLocationOnPath(point, polyline, geodesic, DEFAULT_TOLERANCE)
}

func IsLocationOnEdgeOrPath(point *maps.LatLng, poly []*maps.LatLng, closed, geodesic bool, toleranceEarth float64) bool {
	idx := LocationIndexOnEdgeOrPath(point, poly, closed, geodesic, toleranceEarth)
	return idx >= 0
}

/**
 * Computes whether (and where) a given point lies on or near a polyline, within a specified tolerance.
 * If closed, the closing segment between the last and first points of the polyline is not considered.
 *
 * @param point          our needle
 * @param poly           our haystack
 * @param closed         whether the polyline should be considered closed by a segment connecting the last point back to the first one
 * @param geodesic       the polyline is composed of great circle segments if geodesic
 *                       is true, and of Rhumb segments otherwise
 * @param toleranceEarth tolerance (in meters)
 * @return -1 if point does not lie on or near the polyline.
 * 0 if point is between poly[0] and poly[1] (inclusive),
 * 1 if between poly[1] and poly[2],
 * ...,
 * poly.size()-2 if between poly[poly.size() - 2] and poly[poly.size() - 1]
 */
func LocationIndexOnEdgeOrPath(point *maps.LatLng, poly []*maps.LatLng, closed, geodesic bool, toleranceEarth float64) int {
	size := len(poly)
	if size == 0 {
		return -1
	}
	tolerance := toleranceEarth / mathutil.EarthRadius
	havTolerance := mathutil.Hav(tolerance)
	lat3 := mathutil.ToRadians(point.Lat)
	lng3 := mathutil.ToRadians(point.Lng)
	prevIndex := 0
	if closed {
		prevIndex = size - 1
	}
	prev := poly[prevIndex]
	lat1 := mathutil.ToRadians(prev.Lat)
	lng1 := mathutil.ToRadians(prev.Lng)
	idx := 0
	if geodesic {
		for _, point2 := range poly {
			lat2 := mathutil.ToRadians(point2.Lat)
			lng2 := mathutil.ToRadians(point2.Lng)
			if isOnSegmentGC(lat1, lng1, lat2, lng2, lat3, lng3, havTolerance) {
				return int(Max(0.0, float64(idx-1)))
			}
			lat1 = lat2
			lng1 = lng2
			idx++
		}
	} else {
		// We project the points to mercator space, where the Rhumb segment is a straight line,
		// and compute the geodesic distance between point3 and the closest point on the
		// segment. This method is an approximation, because it uses "closest" in mercator
		// space which is not "closest" on the sphere -- but the error is small because
		// "tolerance" is small.
		minAcceptable := lat3 - tolerance
		maxAcceptable := lat3 + tolerance
		y1 := mathutil.Mercator(lat1)
		y3 := mathutil.Mercator(lat3)
		xTry := make([]float64, 3)
		for _, point2 := range poly {
			lat2 := mathutil.ToRadians(point2.Lat)
			y2 := mathutil.Mercator(lat2)
			lng2 := mathutil.ToRadians(point2.Lng)
			if Max(lat1, lat2) >= minAcceptable && Min(lat1, lat2) <= maxAcceptable {
				// We offset longitudes by -lng1; the implicit x1 is 0.
				x2 := mathutil.Wrap(lng2-lng1, -Pi, Pi)
				x3Base := mathutil.Wrap(lng3-lng1, -Pi, Pi)
				xTry[0] = x3Base
				// Also explore wrapping of x3Base around the world in both directions.
				xTry[1] = x3Base + 2*Pi
				xTry[2] = x3Base - 2*Pi
				for _, x3 := range xTry {
					dy := y2 - y1
					len2 := x2*x2 + dy*dy
					t := 0.0
					if len2 > 0 {
						t = mathutil.Clamp((x3*x2+(y3-y1)*dy)/len2, 0, 1)
					}
					xClosest := t * x2
					yClosest := y1 + t*dy
					latClosest := mathutil.InverseMercator(yClosest)
					havDist := mathutil.HavDistance(lat3, latClosest, x3-xClosest)
					if havDist < havTolerance {
						return int(Max(0.0, float64(idx-1)))
					}
				}
			}
			lat1 = lat2
			lng1 = lng2
			y1 = y2
			idx++
		}
	}
	return -1
}

/**
 * Returns sin(initial bearing from (lat1,lng1) to (lat3,lng3) minus initial bearing
 * from (lat1, lng1) to (lat2,lng2)).
 */
func sinDeltaBearing(lat1, lng1, lat2, lng2, lat3, lng3 float64) float64 {
	sinLat1 := Sin(lat1)
	cosLat2 := Cos(lat2)
	cosLat3 := Cos(lat3)
	lat31 := lat3 - lat1
	lng31 := lng3 - lng1
	lat21 := lat2 - lat1
	lng21 := lng2 - lng1
	a := Sin(lng31) * cosLat3
	c := Sin(lng21) * cosLat2
	b := Sin(lat31) + 2*sinLat1*cosLat3*mathutil.Hav(lng31)
	d := Sin(lat21) + 2*sinLat1*cosLat2*mathutil.Hav(lng21)
	denom := (a*a + b*b) * (c*c + d*d)
	if denom <= 0 {
		return 1
	}
	return (a*d - b*c) / Sqrt(denom)
}

func isOnSegmentGC(lat1 float64, lng1 float64, lat2 float64, lng2 float64, lat3 float64, lng3 float64, havTolerance float64) bool {
	havDist13 := mathutil.HavDistance(lat1, lat3, lng1-lng3)
	if havDist13 <= havTolerance {
		return true
	}
	havDist23 := mathutil.HavDistance(lat2, lat3, lng2-lng3)
	if havDist23 <= havTolerance {
		return true
	}
	sinBearing := sinDeltaBearing(lat1, lng1, lat2, lng2, lat3, lng3)
	sinDist13 := mathutil.SinFromHav(havDist13)
	havCrossTrack := mathutil.HavFromSin(sinDist13 * sinBearing)
	if havCrossTrack > havTolerance {
		return false
	}
	havDist12 := mathutil.HavDistance(lat1, lat2, lng1-lng2)
	term := havDist12 + havCrossTrack*(1-2*havDist12)
	if havDist13 > term || havDist23 > term {
		return false
	}
	if havDist12 < 0.74 {
		return true
	}
	cosCrossTrack := 1 - 2*havCrossTrack
	havAlongTrack13 := (havDist13 - havCrossTrack) / cosCrossTrack
	havAlongTrack23 := (havDist23 - havCrossTrack) / cosCrossTrack
	sinSumAlongTrack := mathutil.SinSumFromHav(havAlongTrack13, havAlongTrack23)
	return sinSumAlongTrack > 0 // Compare with half-circle == PI using sign of sin().
}
