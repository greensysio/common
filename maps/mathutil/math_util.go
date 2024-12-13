package mathutil

import "math"

/**
 * The earth's radius, in meters.
 * Mean radius as defined by IUGG.
 */
const EarthRadius = 6371009

func ToRadians(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func ToDegrees(radians float64) float64 {
	return (radians * 180) / math.Pi
}

/**
 * Restrict x to the range [low, high].
 */
func Clamp(x, low, high float64) float64 {
	if x < low {
		return low
	}
	if x > high {
		return high
	}
	return x
}

/**
 * Wraps the given value into the inclusive-exclusive interval between min and max.
 *
 * @param n   The value to wrap.
 * @param min The minimum.
 * @param max The maximum.
 */
func Wrap(n, min, max float64) float64 {
	if n >= min && n < max {
		return n
	}
	return Mod(n-min, max-min) + min
}

/**
 * Returns the non-negative remainder of x / m.
 *
 * @param x The operand.
 * @param m The modulus.
 */
func Mod(x, m float64) float64 {
	xI, mI := int(x), int(m)
	return float64(((xI % mI) + mI) % mI)
}

/**
 * Returns mercator Y corresponding to latitude.
 * See http://en.wikipedia.org/wiki/Mercator_projection .
 */
func Mercator(lat float64) float64 {
	return math.Log(math.Tan(lat*0.5 + math.Pi/4))
}

/**
 * Returns latitude from mercator Y.
 */
func InverseMercator(y float64) float64 {
	return 2*math.Atan(math.Exp(y)) - math.Pi/2
}

/**
 * Returns haversine(angle-in-radians).
 * hav(x) == (1 - cos(x)) / 2 == sin(x / 2)^2.
 */
func Hav(x float64) float64 {
	sinHalf := math.Sin(x * 0.5)
	return sinHalf * sinHalf
}

/**
 * Computes inverse haversine. Has good numerical stability around 0.
 * arcHav(x) == acos(1 - 2 * x) == 2 * asin(sqrt(x)).
 * The argument must be in [0, 1], and the result is positive.
 */
func ArcHav(x float64) float64 {
	return 2 * math.Asin(math.Sqrt(x))
}

// Given h==hav(x), returns sin(abs(x)).
func SinFromHav(h float64) float64 {
	return 2 * math.Sqrt(h*(1-h))
}

// Returns hav(asin(x)).
func HavFromSin(x float64) float64 {
	x2 := x * x
	return x2 / (1 + math.Sqrt(1-x2)) * .5
}

// Returns sin(arcHav(x) + arcHav(y)).
func SinSumFromHav(x, y float64) float64 {
	a := math.Sqrt(x * (1 - x))
	b := math.Sqrt(y * (1 - y))
	return 2 * (a + b - 2*(a*y+b*x))
}

/**
 * Returns hav() of distance from (lat1, lng1) to (lat2, lng2) on the unit sphere.
 */
func HavDistance(lat1, lat2, dLng float64) float64 {
	return Hav(lat1-lat2) + Hav(dLng)*math.Cos(lat1)*math.Cos(lat2)
}

// haversin(θ) function
func Hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
