package math

import (
	"bitbucket.org/greensys-tech/common/log"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundTo(t *testing.T) {
	// Init Logger
	log.InitLogger(true)

	values := []struct {
		Value  float64
		Place  int
		Result float64
	}{
		{Value: 5.168, Place: 1, Result: 5.200},
		{Value: 5.168, Place: 2, Result: 5.170},
		{Value: 3.004, Place: 1, Result: 3.000},
		{Value: 14.047, Place: 1, Result: 14.000},
		{Value: 15345, Place: -2, Result: 15300},
		{Value: 3281, Place: -2, Result: 3300},
	}
	for _, p := range values {
		rs := RoundTo(p.Value, p.Place)
		assert.Equal(t, p.Result, rs, fmt.Sprintf("Can't Round To: %f", p.Value))
	}
}

func TestRoundDown(t *testing.T) {
	// Init Logger
	log.InitLogger(true)

	values := []struct {
		Value  float64
		Place  int
		Result float64
	}{
		{Value: 5.168, Place: 1, Result: 5.100},
		{Value: 5.168, Place: 2, Result: 5.160},
		{Value: 3.004, Place: 1, Result: 3.000},
		{Value: 14.047, Place: 1, Result: 14.000},
	}
	for _, p := range values {
		rs := RoundDown(p.Value, p.Place)
		assert.Equal(t, p.Result, rs, fmt.Sprintf("Can't Round Down: %f", p.Value))
	}
}
