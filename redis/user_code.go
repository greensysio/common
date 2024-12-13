package redis

import (
	"github.com/greensysio/common/model/enum"
	redis "github.com/greensysio/redis"
	"errors"
)

func GetKeyOfUserIDByCode(role enum.UserRole, userCode string) (string, error) {
	switch role {
	case enum.UserRoleShipper:
		return ShipperCodeId + userCode, nil
	case enum.UserRoleCarrier:
		return CarrierCodeId + userCode, nil
	}
	return "", errors.New("not supported user role")
}

func GetUserIDByCode(role enum.UserRole, userCode string) (string, error) {
	key, err := GetKeyOfUserIDByCode(role, userCode)
	if err != nil {
		return "", err
	}
	redisValueGet, err := redis.GetInstance().Get(key)
	return redisValueGet, err
}
