package phone

import (
	"github.com/greensysio/common/log"
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	log.InitLogger(false)
}

// ------- EXPORT FUNC -------

// NormalizeDigitsOnly : Format number as 84938123456
func NormalizeDigitsOnly(phone string, locale string) (string, bool) {
	if len(phone) == 0 {
		return phone, false
	}
	if locale == "" {
		locale = "VN"
	}
	parsedPhonenumber, isValid := ValidatePhoneNumber(phone, strings.ToUpper(locale))
	if parsedPhonenumber != nil {
		phone = strconv.FormatUint(*parsedPhonenumber.NationalNumber, 10)
	} else {
		phone = ""
	}
	return phone, isValid
}

// GenerateVerifyCode : Generate 4 digist verifycode
func GenerateVerifyCode(oldCode string) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	verifyCode := fmt.Sprintf("%04d", r.Intn(9999))
	if oldCode != "" {
		for verifyCode == oldCode {
			verifyCode = fmt.Sprintf("%04d", r.Intn(9999))
		}
	}
	return verifyCode
}

func ValidatePhoneNumber(phonenumber string, countryCode string) (*phonenumbers.PhoneNumber, bool) {
	if len(phonenumber) == 0 || len(countryCode) == 0 {
		return nil, false
	}
	parsedPhonenumber, err := phonenumbers.Parse(phonenumber, strings.ToUpper(countryCode))
	if err != nil {
		log.Errorf("Phone number: %s. CountryCode: %s. Error: %+v", phonenumber, countryCode, err)
		return nil, false
	}
	return parsedPhonenumber, phonenumbers.IsValidNumber(parsedPhonenumber)
}

func FormatForMobileDialing(phone string, region string) (parsedPhone string) {
	if len(phone) == 0 || len(region) == 0 {
		return phone
	}
	num, err := phonenumbers.Parse(phone, region)
	if err != nil {
		log.Errorf("Can not parse number %s with region %s! Err: %+v", phone, region, err)
		return phone
	}
	parsedPhone = phonenumbers.FormatNumberForMobileDialing(num, strings.ToUpper(region), false)
	return parsedPhone
}
