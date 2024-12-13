package enum

import "strings"

// VehicleType uses for all model
type VehicleType int

const (
	// VehicleTypeUnknown is default status
	VehicleTypeUnknown VehicleType = iota
	// VehicleTypeVan is default status
	VehicleTypeVan
	// VehicleTypeContainer uses when user was activated
	VehicleTypeContainer
)

func (s VehicleType) Str() string {
	return []string{"", "Van", "Container"}[s]
}

// Int parses enum to int
func (s VehicleType) Int() int {
	return int(s)
}

// GetVehicleTypeEnum : return const by id
func GetVehicleTypeEnum(s string) VehicleType {
	switch strings.ToLower(s) {
	case strings.ToLower(VehicleTypeVan.Str()):
		return VehicleTypeVan
	case strings.ToLower(VehicleTypeContainer.Str()):
		return VehicleTypeContainer
	default:
		return VehicleTypeUnknown
	}
}

// GetVehicleTypeEnumByInt : return const by id
func GetVehicleTypeEnumByInt(index int) VehicleType {
	return []VehicleType{
		VehicleTypeUnknown,
		VehicleTypeVan,
		VehicleTypeContainer,
	}[index]
}

// GetVehicleTypeFromTripType will return VehicleType from TripType (TripType)
func GetVehicleTypeFromTripType(t TripType) VehicleType {
	switch t {
	case TripTypeVan:
		return VehicleTypeVan
	case TripTypeExport, TripTypeImport, TripTypeChangeWarehouse:
		return VehicleTypeContainer
	default:
		return VehicleTypeUnknown
	}
}

// GetVehicleTypesFromString return array of VehicleType from string
func GetVehicleTypesFromString(vhTypeSrc []string) (vhTypes []VehicleType) {
	for _, status := range vhTypeSrc {
		vhType := GetVehicleTypeEnum(status)
		if vhType != VehicleTypeUnknown {
			vhTypes = append(vhTypes, vhType)
		}
	}
	return vhTypes
}

// VehicleTypeToStrArray convert []VehicleType to array of string, excluding unknown status
func VehicleTypesToStrArray(status []VehicleType) []string {
	results := []string{}
	for _, status := range status {
		if status == VehicleTypeUnknown {
			continue
		}
		results = append(results, status.Str())
	}
	return results
}
