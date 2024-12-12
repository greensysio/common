package redis

import (
	"bitbucket.org/greensys-tech/common/model/enum"
	redis "bitbucket.org/greensys-tech/redis"
)

// GetRedisUserStatusKeyFromRole return redis key
func GetRedisUserStatusKeyFromRole(userID string, role enum.UserRole) string {
	switch role {
	case enum.UserRoleDriver:
		return DriverStatus + userID
	case enum.UserRoleShipper:
		return ShipperStatus + userID
	case enum.UserRoleCarrier:
		return CarrierStatus + userID
	}
	return ""
}

// GetStatus return user status in redis
func GetStatus(userID string, role enum.UserRole) (string, error) {
	return redis.GetInstance().Get(GetRedisUserStatusKeyFromRole(userID, role))
}

// IsActive return true if user is Active
func IsActive(userID string, role enum.UserRole) bool {
	redisValueGet, _ := redis.GetInstance().Get(GetRedisUserStatusKeyFromRole(userID, role))
	if redisValueGet == enum.StatusActive.Str() {
		return true
	}
	return false
}

// IsActiveOrPending return true if user is Active or Pending
func IsActiveOrPending(userID string, role enum.UserRole) bool {
	// Admin user is not yet implement status.
	if role != enum.UserRoleDriver && role != enum.UserRoleCarrier && role != enum.UserRoleShipper {
		return true
	}
	// Check status user.
	// User has status is active or pending can login.
	redisValueGet, _ := redis.GetInstance().Get(GetRedisUserStatusKeyFromRole(userID, role))
	if redisValueGet == enum.StatusActive.Str() || redisValueGet == enum.StatusPending.Str() {
		return true
	}
	return false
}
