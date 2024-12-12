package enum

import "strings"

// GeneralStatus uses for all model
type GeneralStatus int

const (
	// GeneralStatusUnknown is default status
	GeneralStatusUnknown GeneralStatus = iota
	// StatusInactive is default status
	StatusInactive
	// StatusActive uses when user was activated
	StatusActive
	// StatusPending uses when user does not verify
	StatusPending
	// StatusBlocked is used to represent blocked user account status
	StatusBlocked
)

func (s GeneralStatus) Str() string {
	return []string{"", "Inactive", "Active", "Pending", "Blocked"}[s]
}

// Int parses enum to int
func (s GeneralStatus) Int() int {
	return int(s)
}

// GetGeneralStatusEnum : return const by id
func GetGeneralStatusEnum(s string) GeneralStatus {
	switch strings.ToLower(s) {
	case strings.ToLower(StatusInactive.Str()):
		return StatusInactive
	case strings.ToLower(StatusActive.Str()):
		return StatusActive
	case strings.ToLower(StatusPending.Str()):
		return StatusPending
	case strings.ToLower(StatusBlocked.Str()):
		return StatusBlocked
	default:
		return GeneralStatusUnknown
	}
}

// GetGeneralStatusEnums : return const by ids
func GetGeneralStatusEnums(statusSrc []string, excludeUnknown bool) []GeneralStatus {
	results := []GeneralStatus{}
	for _, status := range statusSrc {
		status := GetGeneralStatusEnum(status)
		if status != GeneralStatusUnknown || !excludeUnknown {
			results = append(results, status)
		}
	}
	return results
}

// GetGeneralStatusEnumByInt : return const by id
func GetGeneralStatusEnumByInt(index int) GeneralStatus {
	return []GeneralStatus{
		GeneralStatusUnknown,
		StatusInactive,
		StatusActive,
		StatusPending,
		StatusBlocked,
	}[index]
}

// FromGeneralStatusToStrArray convert []GeneralStatus to array of string, exclude unknown status
func FromGeneralStatusToStrArray(status []GeneralStatus) []string {
	results := []string{}
	for _, status := range status {
		if status == GeneralStatusUnknown {
			continue
		}
		results = append(results, status.Str())
	}
	return results
}
