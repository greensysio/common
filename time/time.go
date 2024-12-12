package time

import (
	"strconv"
	"time"

	"bitbucket.org/greensys-tech/common/log"
)

func init() {
	log.InitLogger(false)
}

// ParseTimeUTC : Uses time.RFC3339 for parsing
func ParseTimeUTC(t string) (*time.Time, bool) {
	if t == "" {
		return nil, true
	}
	timeParsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Errorf("Can not parse this string %s !", t)
		return nil, false
	}
	UTCTime := timeParsed.UTC()
	return &UTCTime, true
}

// ParseTimeLoc : Uses time.RFC3339 for parsing
func ParseTimeLoc(t *time.Time) (*time.Time, bool) {
	if t == nil {
		return nil, true
	}

	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Error("Get error when parse time!", err)
	}

	locTime := t.In(location)
	return &locTime, true
}

// ParseTimeLocFromStr : Uses time.RFC3339 for parsing
func ParseTimeLocFromStr(t string) (*time.Time, bool) {
	if t == "" {
		return nil, true
	}
	timeParsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Errorf("Can not parse this string %s !", t)
		return nil, false
	}

	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		log.Error("Get error when parse time!", err)
	}

	locTime := timeParsed.In(location)
	return &locTime, true
}

// ParsePointerTime parse time variable to pointer
func ParsePointerTime(t time.Time) *time.Time {
	return &t
}

// Document: https://www.myonlinetraininghub.com/excel-date-and-time
// https://stackoverflow.com/questions/53013156/java-convert-excel-date-serial-number-to-datetime
func ParseDateTimeFromExcel(dateTimeExcel string, dateTimeLayout string, timeLoc *time.Location) (rs *time.Time, ok bool) {
	rootTime, _ := time.Parse(time.RFC3339, "1899-12-30T00:00:00+07:00")
	dateTime, err := strconv.ParseFloat(dateTimeExcel, 64)
	if err != nil {
		parsedTimeFromString, errTry2 := time.ParseInLocation(dateTimeLayout, dateTimeExcel, timeLoc)
		if errTry2 == nil {
			return &parsedTimeFromString, true
		}
		log.Error("Get error when parse time!", err, errTry2)
		return nil, false
	}
	days := int(dateTime)
	rootTime = rootTime.Add(time.Hour * 24 * time.Duration(days))
	fraction := dateTime - float64(days)
	nanos := int(fraction * float64(24*time.Hour.Nanoseconds()))
	rootTime = rootTime.Add(time.Duration(nanos)).Truncate(time.Second)

	return &rootTime, true
}

// Format date
const (
	ISO8601 = "2006-01-02T15:04:05Z"
)