package time

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDateTimeSerialNumber(t *testing.T) {
	testCases := []struct {
		ExcelTime string
		Result    string
	}{
		{ExcelTime: "43626.15625", Result: "2019-06-10 03:45:00 +0700 +0700"},
		{ExcelTime: "43626.165972222225", Result: "2019-06-10 03:59:00 +0700 +0700"},
	}
	for _, tc := range testCases {
		parsedTime, ok := ParseDateTimeFromExcel(tc.ExcelTime, "", nil)
		assert.True(t, ok, fmt.Sprintf("Can't parse value: %s", tc.ExcelTime))
		assert.Equal(t, tc.Result, fmt.Sprintf("%s", parsedTime), fmt.Sprintf("Time after parsed is not correct! Excel time: %s", tc.ExcelTime))
	}
}
