package enum

import "strings"

// InvitedStatus uses for all model
type InvitedStatus int

const (
	// InvitedStatusUnknown is default status
	InvitedStatusUnknown InvitedStatus = iota
	// InvitedStatusWaitingAdminAccept
	InvitedStatusWaitingAdminAccept
	// InvitedStatusWaitingCarrierAccept
	InvitedStatusWaitingCarrierAccept
	// InvitedStatusAccepted
	InvitedStatusAccepted
	// InvitedStatusDenied
	InvitedStatusDenied
	// InvitedStatusDeleted
	InvitedStatusDeleted
)

// InvitedStatus in string
func (s InvitedStatus) Str() string {
	return []string{"", "WaitingAdminAccept", "WaitingCarrierAccept", "Accepted", "Denied", "Deleted"}[s]
}

// GetInvitedDriverStatusEnum : return const by id
func GetInvitedDriverStatusEnum(s string) InvitedStatus {
	switch strings.ToLower(s) {
	case strings.ToLower(InvitedStatusAccepted.Str()):
		return InvitedStatusAccepted
	case strings.ToLower(InvitedStatusDenied.Str()):
		return InvitedStatusDenied
	case strings.ToLower(InvitedStatusDeleted.Str()):
		return InvitedStatusDeleted
	case strings.ToLower(InvitedStatusWaitingAdminAccept.Str()):
		return InvitedStatusWaitingAdminAccept
	case strings.ToLower(InvitedStatusWaitingCarrierAccept.Str()):
		return InvitedStatusWaitingCarrierAccept
	default:
		return InvitedStatusUnknown
	}
}
