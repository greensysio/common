package assign_order_history

import "strings"

type AssignOrderHistoryStatus int

const (
	// UnknownOrderStatus is default status
	UnknownStatus AssignOrderHistoryStatus = iota
	// OnAssigned when assigned driver to order.
	OnAssigned
	// UnAssigned when unassigned driver.
	UnAssigned
	// WaitToChange, when the assigned driver to change driver cause orders has an accident.
	WaitToChange
)

func (s AssignOrderHistoryStatus) Str() string {
	return []string{"",
		"OnAssigned",
		"UnAssigned",
		"WaitToChange",}[s]
}

// Int parses enum to int
func (s AssignOrderHistoryStatus) Int() int {
	return int(s)
}

// GetOrderStatusEnum : return const by id
func GetAssignOrderHistoryStatusEnum(s string) AssignOrderHistoryStatus {
	switch strings.ToLower(s) {
	case strings.ToLower(OnAssigned.Str()):
		return OnAssigned
	case strings.ToLower(UnAssigned.Str()):
		return UnAssigned
	case strings.ToLower(WaitToChange.Str()):
		return WaitToChange
	default:
		return UnknownStatus
	}
}

// GetOrderStatusEnumByInt : return const by id
func GetAssignOrderHistoryStatusEnumByInt(index int) AssignOrderHistoryStatus {
	return []AssignOrderHistoryStatus{
		UnknownStatus,
		OnAssigned,
		UnAssigned,
		WaitToChange,
	}[index]
}

// GetOrderStatusEnumArray return array of OrderStatus from string
func GetAssignOrderHistoryStatusEnumArray(statusSrc []string) []AssignOrderHistoryStatus {
	results := []AssignOrderHistoryStatus{}
	for _, status := range statusSrc {
		results = append(results, GetAssignOrderHistoryStatusEnum(status))
	}
	return results
}

// FromOrderStatusToInt convert []OrderStatus to array of Int
func FromAssignOrderHistoryStatusToArrayStr(status []AssignOrderHistoryStatus) []string {
	results := []string{}
	for _, status := range status {
		results = append(results, status.Str())
	}
	return results
}

