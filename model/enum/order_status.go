package enum

import "strings"

// OrderStatus is type when create order
type OrderStatus int

const (
	// UnknownOrderStatus is default status
	UnknownOrderStatus OrderStatus = iota
	// MissingScheduleOrderStatus when create new order
	MissingScheduleOrderStatus
	// MissingConfirmOrderStatus after create schedules for order
	MissingConfirmOrderStatus
	// MatchingOrderStatus after confirm complete order
	MatchingOrderStatus
	// DriverAcceptedOrderStatus after finish order
	DriverAcceptedOrderStatus
	// StartedTripOrderStatus when driver click Start button
	StartedTripOrderStatus
	// OnTripOrderStatus when driver arrived first schedule
	OnTripOrderStatus
	// CompletedOrderStatus after finish order
	CompletedOrderStatus
	// CanceledOrderStatus when cancel order
	CanceledOrderStatus
	// DeletedOrderStatus when delete order
	DeletedOrderStatus
	// ExpiredOrderStatus when expired order
	ExpiredOrderStatus
	// TruckOperatorAcceptedOrderStatus after finish order
	TruckOperatorAcceptedOrderStatus
	// Accident when has an accident.
	AccidentOrderStatus
	// WaitForNewTruck when the carrier assigns to new drivers.
	WaitForNewTruckOrderStatus
)

func (s OrderStatus) Str() string {
	return []string{"",
		"MissingSchedule",
		"MissingConfirmOrder",
		"Matching",
		"DriverAccepted",
		"StartedTrip",
		"OnTrip",
		"Completed",
		"Canceled",
		"Deleted",
		"Expired",
		"TruckOperatorAccepted",
		"Accident",
		"WaitForNewTruck"}[s]

}

// Int parses enum to int
func (s OrderStatus) Int() int {
	return int(s)
}

// GetOrderStatusEnum : return const by id
func GetOrderStatusEnum(s string) OrderStatus {
	switch strings.ToLower(s) {
	case strings.ToLower(MissingScheduleOrderStatus.Str()):
		return MissingScheduleOrderStatus
	case strings.ToLower(MissingConfirmOrderStatus.Str()):
		return MissingConfirmOrderStatus
	case strings.ToLower(MatchingOrderStatus.Str()):
		return MatchingOrderStatus
	case strings.ToLower(DriverAcceptedOrderStatus.Str()):
		return DriverAcceptedOrderStatus
	case strings.ToLower(StartedTripOrderStatus.Str()):
		return StartedTripOrderStatus
	case strings.ToLower(OnTripOrderStatus.Str()):
		return OnTripOrderStatus
	case strings.ToLower(CompletedOrderStatus.Str()):
		return CompletedOrderStatus
	case strings.ToLower(CanceledOrderStatus.Str()):
		return CanceledOrderStatus
	case strings.ToLower(DeletedOrderStatus.Str()):
		return DeletedOrderStatus
	case strings.ToLower(ExpiredOrderStatus.Str()):
		return ExpiredOrderStatus
	case strings.ToLower(TruckOperatorAcceptedOrderStatus.Str()):
		return TruckOperatorAcceptedOrderStatus
	case strings.ToLower(AccidentOrderStatus.Str()):
		return AccidentOrderStatus
	case strings.ToLower(WaitForNewTruckOrderStatus.Str()):
		return WaitForNewTruckOrderStatus
	default:
		return UnknownOrderStatus
	}
}

// GetOrderStatusEnumByInt : return const by id
func GetOrderStatusEnumByInt(index int) OrderStatus {
	return []OrderStatus{
		UnknownOrderStatus,
		MissingScheduleOrderStatus,
		MissingConfirmOrderStatus,
		MatchingOrderStatus,
		DriverAcceptedOrderStatus,
		StartedTripOrderStatus,
		OnTripOrderStatus,
		CompletedOrderStatus,
		CanceledOrderStatus,
		DeletedOrderStatus,
		ExpiredOrderStatus,
		TruckOperatorAcceptedOrderStatus,
		AccidentOrderStatus,
		WaitForNewTruckOrderStatus,
	}[index]
}

// GetOrderStatusEnumArray return array of OrderStatus from string
func GetOrderStatusEnumArray(statusSrc []string) []OrderStatus {
	var results []OrderStatus
	for _, status := range statusSrc {
		odStatus := GetOrderStatusEnum(status)
		if odStatus != UnknownOrderStatus {
			results = append(results, odStatus)
		}
	}
	return results
}

// FromOrderStatusToInt convert []OrderStatus to array of Int
func FromOrderStatusToArrayStr(status []OrderStatus) []string {
	var results []string
	for _, status := range status {
		results = append(results, status.Str())
	}
	return results
}
