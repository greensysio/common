package enum

import "strings"

// TripType is type when create order
type TripType int

const (
	// TripTypeUnknown is ID type of order by VAN
	TripTypeUnknown TripType = iota
	// TripTypeVan is ID type of order by VAN
	TripTypeVan
	// TripTypeExport : Mã số order
	TripTypeExport
	// TripTypeImport : Mã số container
	TripTypeImport
	// TripTypeChangeWarehouse : Mã phiếu xuất kho
	TripTypeChangeWarehouse
)

func (s TripType) Str() string {
	return []string{"", "Van", "Export", "Import", "ChangeWarehouse"}[s]
}

// Int return int of enum
func (s TripType) Int() int {
	return int(s)
}

// GetTripTypeEnum : return const by id
func GetTripTypeEnum(s string) TripType {
	switch strings.ToLower(s) {
	case strings.ToLower(TripTypeVan.Str()):
		return TripTypeVan
	case strings.ToLower(TripTypeExport.Str()):
		return TripTypeExport
	case strings.ToLower(TripTypeImport.Str()):
		return TripTypeImport
	case strings.ToLower(TripTypeChangeWarehouse.Str()):
		return TripTypeChangeWarehouse
	default:
		return TripTypeUnknown
	}
}

// GetTripTypeEnumByInt : return const by id
func GetTripTypeEnumByInt(index int) TripType {
	return []TripType{
		TripTypeUnknown,
		TripTypeVan,
		TripTypeExport,
		TripTypeImport,
		TripTypeChangeWarehouse,
	}[index]
}
