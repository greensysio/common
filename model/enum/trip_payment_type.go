package enum

import "strings"

// PaymentType uses for all model
type PaymentType int

const (
	// PaymentTypeUnknown is default status
	PaymentTypeUnknown PaymentType = iota
	// PaymentTypeCash
	PaymentTypeCash
	// PaymentTypeTopup
	PaymentTypeTopup
)

func (s PaymentType) Str() string {
	return []string{"",
		"Cash",
		"Topup",
	}[s]
}

// Int parses enum to int
func (s PaymentType) Int() int {
	return int(s)
}

// GetPaymentTypeEnum : return const by id
func GetPaymentTypeEnum(s string) PaymentType {
	switch strings.ToLower(s) {
	case strings.ToLower(PaymentTypeCash.Str()):
		return PaymentTypeCash
	case strings.ToLower(PaymentTypeTopup.Str()):
		return PaymentTypeTopup

	default:
		return PaymentTypeUnknown
	}
}

// GetPaymentTypeEnumByInt : return const by id
func GetPaymentTypeEnumByInt(index int) PaymentType {
	return []PaymentType{
		PaymentTypeUnknown,
		PaymentTypeCash,
		PaymentTypeTopup,
	}[index]
}
