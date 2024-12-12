package transit

import "strings"

type Status int

const (
	UnknownStatus Status = iota
	Waiting
	HasTransferred
	Finished
)

func (s Status) Str() string {
	return []string{
		"UnknownStatus",
		"Waiting",
		"HasTransferred",
		"Finished",}[s]
}

// Int parses enum to int
func (s Status) Int() int {
	return int(s)
}

// GetOrderStatusEnum : return const by id
func GetEnum(s string) Status {
	switch strings.ToLower(s) {
	case strings.ToLower(Waiting.Str()):
		return Waiting
	case strings.ToLower(HasTransferred.Str()):
		return HasTransferred
	case strings.ToLower(Finished.Str()):
		return Finished
	default:
		return UnknownStatus
	}
}

// GetOrderStatusEnumByInt : return const by id
func GetEnumByInt(index int) Status {
	return []Status{
		UnknownStatus,
		Waiting,
		HasTransferred,
		Finished,
	}[index]
}

// GetOrderStatusEnumArray return array of OrderStatus from string
func GetEnumArray(statusSrc []string) []Status {
	var results []Status
	for _, status := range statusSrc {
		results = append(results, GetEnum(status))
	}
	return results
}

// FromOrderStatusToInt convert []OrderStatus to array of Int
func ToArrayStr(status []Status) []string {
	var results []string
	for _, status := range status {
		results = append(results, status.Str())
	}
	return results
}
