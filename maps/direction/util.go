package direction

import "github.com/greensysio/common/maps"

func latLngsToString(latLngs []maps.LatLng) (result []string) {
	lenlatLngs := len(latLngs)
	if lenlatLngs == 0 {
		return result
	}
	result = make([]string, lenlatLngs)
	for i := 0; i < lenlatLngs; i++ {
		result[i] = latLngs[i].String()
	}
	return result
}
