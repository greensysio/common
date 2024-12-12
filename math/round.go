package math

import "math"

func Round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}

func RoundTo(input float64, places int) (newVal float64) {
	// return math.Round(n*float64(places)) / float64(places)
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Round(digit)
	newVal = round / pow
	return
}

func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}
