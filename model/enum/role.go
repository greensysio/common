package enum

import "strings"

// UserRole uses for all model
type UserRole int

const (
	// UserRoleUnknown is default status
	UserRoleUnknown UserRole = iota
	// UserRoleDriver is default status
	UserRoleDriver
	// UserRoleShipper uses when user was activated
	UserRoleShipper
	// UserRoleCarrier
	UserRoleCarrier
	// UserRoleAdmin
	UserRoleAdmin
	// UserRoleCustomerService
	UserRoleCustomerService
	// UserRoleViewer
	UserRoleViewer
)

func (s UserRole) Str() string {
	return []string{
		"",
		"Driver",
		"Shipper",
		"Carrier",
		"Admin",
		"CustomerService",
		"Viewer",
	}[s]
}

// Int parses enum to int
func (s UserRole) Int() int {
	return int(s)
}

// GetUserRoleEnum : return const by id
func GetUserRoleEnum(s string) UserRole {
	switch strings.ToLower(s) {
	case strings.ToLower(UserRoleDriver.Str()):
		return UserRoleDriver
	case strings.ToLower(UserRoleShipper.Str()):
		return UserRoleShipper
	case strings.ToLower(UserRoleCarrier.Str()):
		return UserRoleCarrier
	case strings.ToLower(UserRoleAdmin.Str()):
		return UserRoleAdmin
	case strings.ToLower(UserRoleCustomerService.Str()):
		return UserRoleCustomerService
	case strings.ToLower(UserRoleViewer.Str()):
		return UserRoleViewer
	default:
		return UserRoleUnknown
	}
}

// GetUserRoleEnumByInt : return const by id
func GetUserRoleEnumByInt(index int) UserRole {
	return []UserRole{
		UserRoleUnknown,
		UserRoleDriver,
		UserRoleShipper,
		UserRoleCarrier,
		UserRoleAdmin,
		UserRoleCustomerService,
		UserRoleViewer,
	}[index]
}
