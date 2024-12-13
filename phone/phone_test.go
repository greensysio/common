package phone

import (
	"fmt"
	"testing"

	"github.com/greensysio/common/log"
	"github.com/stretchr/testify/assert"
)

func TestValidatePhoneNumber(t *testing.T) {
	// Init Logger
	log.InitLogger(true)

	phones := []struct {
		Phone   string
		Country string
		Result  bool
	}{
		{Phone: "0938135110", Country: "VN", Result: true},
		{Phone: "0938135110", Country: "VN", Result: true},
		{Phone: "0938135110", Country: "", Result: true},
		{Phone: "0938135110", Country: "EN", Result: false},
		{Phone: "0938135110", Country: "en", Result: false},
		{Phone: "09381351101", Country: "VN", Result: false},
		{Phone: "09381351101", Country: "vn", Result: false},
		{Phone: "asdasdad", Country: "VN", Result: false},
		{Phone: "", Country: "VN", Result: false},
		{Phone: "43241324", Country: "", Result: false},
		{Phone: "sdasdas", Country: "", Result: false},
		{Phone: "", Country: "", Result: false},
		{Phone: "938135110", Country: "VN", Result: true},
		{Phone: "938135110", Country: "US", Result: false},
		{Phone: "1267851466", Country: "VN", Result: true},
	}
	for _, p := range phones {
		_, isValid := NormalizeDigitsOnly(p.Phone, p.Country)
		assert.Equal(t, p.Result, isValid, fmt.Sprintf("Can't normalize digits for phone %s", p.Phone))
	}
}

func TestGenerateVerifyCode(t *testing.T) {
	log.InitLogger(true)

	phones := []struct {
		OldVerifyCode string
	}{
		{OldVerifyCode: "8888"},
		{OldVerifyCode: "123456"},
		{OldVerifyCode: ""},
	}
	for _, p := range phones {
		rs := GenerateVerifyCode(p.OldVerifyCode)
		assert.NotEmpty(t, rs, "Can't generate verify code!")
		assert.Len(t, rs, 4, "Can't generate verify code!")
	}
}
