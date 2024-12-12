package enum

import "strings"

// TripStatus is type when create order
type TripStatus int

const (
	// TripStatusUnknown is default status
	TripStatusUnknown TripStatus = iota
	// TripStatusMissingSchedule when create new order
	TripStatusMissingSchedule
	// TripStatusMissingConfirm after create schedules for order
	TripStatusMissingConfirm
	// TripStatusMatching after confirm complete order
	TripStatusMatching
	// TripStatusDriverAccepted after finish order
	TripStatusDriverAccepted
	// TripStatusStartedTrip when driver click Start button
	TripStatusStartedTrip
	// TripStatusOnTrip when driver arrived first schedule
	TripStatusOnTrip
	// TripStatusCompleted after finish order
	TripStatusCompleted
	// TripStatusCanceled when cancel order
	TripStatusCanceled
	// TripStatusDeleted when delete order
	TripStatusDeleted
	// TripStatusExpired when expired order
	TripStatusExpired
	// TripStatusCarrierAccepted after finish order
	TripStatusCarrierAccepted
	// TripStatusPlanning after create booking plan
	TripStatusPlanning
)

func (s TripStatus) Str() string {
	return []string{"",
		"MissingSchedule",
		"MissingConfirmTrip",
		"Matching",
		"DriverAccepted",
		"StartedTrip",
		"OnTrip",
		"Completed",
		"Canceled",
		"Deleted",
		"Expired",
		"CarrierAccepted",
		"Planning",
	}[s]
}

// Int parses enum to int
func (s TripStatus) Int() int {
	return int(s)
}

// GetTripStatusEnum : return const by id
func GetTripStatusEnum(s string) TripStatus {
	switch strings.ToLower(s) {
	case strings.ToLower(TripStatusMissingSchedule.Str()):
		return TripStatusMissingSchedule
	case strings.ToLower(TripStatusMissingConfirm.Str()):
		return TripStatusMissingConfirm
	case strings.ToLower(TripStatusMatching.Str()):
		return TripStatusMatching
	case strings.ToLower(TripStatusDriverAccepted.Str()):
		return TripStatusDriverAccepted
	case strings.ToLower(TripStatusStartedTrip.Str()):
		return TripStatusStartedTrip
	case strings.ToLower(TripStatusOnTrip.Str()):
		return TripStatusOnTrip
	case strings.ToLower(TripStatusCompleted.Str()):
		return TripStatusCompleted
	case strings.ToLower(TripStatusCanceled.Str()):
		return TripStatusCanceled
	case strings.ToLower(TripStatusDeleted.Str()):
		return TripStatusDeleted
	case strings.ToLower(TripStatusExpired.Str()):
		return TripStatusExpired
	case strings.ToLower(TripStatusCarrierAccepted.Str()):
		return TripStatusCarrierAccepted
	case strings.ToLower(TripStatusPlanning.Str()):
		return TripStatusPlanning
	default:
		return TripStatusUnknown
	}
}

// GetTripStatusEnumByInt : return const by id
func GetTripStatusEnumByInt(index int) TripStatus {
	return []TripStatus{
		TripStatusUnknown,
		TripStatusMissingSchedule,
		TripStatusMissingConfirm,
		TripStatusMatching,
		TripStatusDriverAccepted,
		TripStatusStartedTrip,
		TripStatusOnTrip,
		TripStatusCompleted,
		TripStatusCanceled,
		TripStatusDeleted,
		TripStatusExpired,
		TripStatusCarrierAccepted,
		TripStatusPlanning,
	}[index]
}

// GetTripStatusEnumArray return array of TripStatus from string
func GetTripStatusEnumArray(statusSrc []string) []TripStatus {
	var results []TripStatus
	for _, status := range statusSrc {
		results = append(results, GetTripStatusEnum(status))
	}
	return results
}

// FromTripStatusToInt convert []TripStatus to array of Int
func FromTripStatusToArrayStr(status []TripStatus) []string {
	var results []string
	for _, status := range status {
		results = append(results, status.Str())
	}
	return results
}
