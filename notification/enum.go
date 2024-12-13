package notification

import "strings"

// NotiType uses for all model
type NotiType int

const (
	// UnknownNotiType is default status
	UnknownNotiType NotiType = iota
	TruckOperatorCreateOrdersForShipperNotiType
	// AcceptTripNotiType for accept
	AcceptTripNotiType
	// StartTripNotiType for start
	StartTripNotiType
	// CancelTripNotiType for start
	CancelTripNotiType
	// SystemCancelTripNotiType notify type for admin cancel trip
	SystemCancelTripNotiType
	// CompleteOrderNotiType for start
	CompleteOrderNotiType
	// SOSOrderNotiType for driver
	SOSOrderNotiType
	// DriverArriveNotiType for driver
	DriverArriveNotiType
	// DriverContinueNotiType for driver
	DriverContinueNotiType
	// AdminApproveDriverDocNotiType for driver
	AdminApproveDriverDocNotiType
	// AdminRejectDriverDocNotiType for driver
	AdminRejectDriverDocNotiType
	// UpcomingScheduleNotiType
	UpcomingScheduleNotiType
	// 	TruckOperatorInviteDriverNotiType for invite driver
	TruckOperatorInviteDriverNotiType
	// 	TruckOperatorDeleteDriverNotiType for delete driver of list
	TruckOperatorDeleteDriverNotiType
	// 	TruckOperatorAssignNotiType when Truck operator assign driver
	TruckOperatorAssignNotiType
	// 	TruckOperatorUnassignNotiType when Truck operator unassign driver
	TruckOperatorUnassignNotiType
	// CreateBookingNotiType when booking a trip
	CreateBookingNotiType
	// AdminRejectOrderNotiType when admin exgo reject order
	AdminRejectOrderNotiType
	// CarrierConfirmTripNotiType when carrier confirm accept trip
	//2.1
	ShippingCancelTripNotiType
	AdminRejectCancelNotiType
	AdminAssignCarrierNotiType
	TruckOperatorApproveCancelTripNotiType
	TruckOperatorRejectCancelTripNotiType
	TruckOperatorCancelSuccessNotiType
	ShipperCancelSuccessNotiType
	ShipperCancelFailNotiType
	//3.1
	TruckOperatorCancelTripNotiType
	AdminApproveCancelTripNotiType
	//4.1
	TruckOperatorCancelDriverNotiType
	// Edit Trip notification type
	EditTripNotiType
	AlertVehicleWrongWayNotiType
	AlertCompleteTripOutsideAreaNotiType
	AlertArrivedScheduleOutsideAreaNotiType

	// FOR ACCIDENT
	CarrierAssignedForAccident
	HandOverToDriverForAccident
	// Password expiry
	PasswordWillExpireAlertNotiType

	// Sub contractor manager.
	SubcontractorInvitationNewNotiType
	SubcontractorAcceptInvitationNotiType
	SubcontractorCanceledInvitationNotiType

	//ETA
	ETANotiType
)

var notiTypeStrArr = []string{"",
	"TruckOperatorCreateOrdersForShipper",
	"AcceptTrip",
	"StartTrip",
	"CancelTrip",
	"SystemCancelTrip",
	"CompleteOrder",
	"SOSOrder",
	"DriverArrive",
	"DriverContinue",
	"AdminApproveDriverDoc",
	"AdminRejectDriverDoc",
	"UpcomingSchedule",
	"TruckOperatorInviteDriver",
	"TruckOperatorDeleteDriver",
	"TruckOperatorAssign",
	"TruckOperatorUnassign",
	"CreateBooking",
	"AdminRejectOrder",
	"ShippingCancelTrip",
	"AdminRejectCancel",
	"AdminAssignCarrier",
	"TruckOperatorApproveCancelTrip",
	"TruckOperatorRejectCancelTrip",
	"TruckOperatorCancelSuccess",
	"ShipperCancelSuccess",
	"ShipperCancelFail",
	"TruckOperatorCancelTrip",
	"AdminApproveCancelTrip",
	"TruckOperatorCancelDriver",
	"EditTrip",
	"AlertVehicleWrongWay",
	"AlertCompleteTripOutsideArea",
	"AlertArrivedScheduleOutsideArea",
	"CarrierAssignedForAccident",
	"HandOverToDriverForAccident",
	"PasswordWillExpireAlert",
	"SubcontractorInvitationNew",
	"SubcontractorAcceptInvitation",
	"SubcontractorCanceledInvitation",
	"ETA",
}

var notiTypeArr = []NotiType{
	UnknownNotiType,
	TruckOperatorCreateOrdersForShipperNotiType,
	AcceptTripNotiType,
	StartTripNotiType,
	CancelTripNotiType,
	SystemCancelTripNotiType,
	CompleteOrderNotiType,
	SOSOrderNotiType,
	DriverArriveNotiType,
	DriverContinueNotiType,
	AdminApproveDriverDocNotiType,
	AdminRejectDriverDocNotiType,
	UpcomingScheduleNotiType,
	TruckOperatorDeleteDriverNotiType,
	TruckOperatorInviteDriverNotiType,
	TruckOperatorAssignNotiType,
	TruckOperatorUnassignNotiType,
	CreateBookingNotiType,
	AdminRejectOrderNotiType,
	ShippingCancelTripNotiType,
	AdminRejectCancelNotiType,
	AdminAssignCarrierNotiType,
	TruckOperatorApproveCancelTripNotiType,
	TruckOperatorRejectCancelTripNotiType,
	TruckOperatorCancelSuccessNotiType,
	ShipperCancelSuccessNotiType,
	ShipperCancelFailNotiType,
	TruckOperatorCancelTripNotiType,
	AdminApproveCancelTripNotiType,
	TruckOperatorCancelDriverNotiType,
	EditTripNotiType,
	AlertVehicleWrongWayNotiType,
	AlertCompleteTripOutsideAreaNotiType,
	AlertArrivedScheduleOutsideAreaNotiType,
	CarrierAssignedForAccident,
	HandOverToDriverForAccident,
	PasswordWillExpireAlertNotiType,
	SubcontractorInvitationNewNotiType,
	SubcontractorAcceptInvitationNotiType,
	SubcontractorCanceledInvitationNotiType,
	ETANotiType,
}

var notifyTypeMap = map[string]NotiType{
	strings.ToLower(TruckOperatorCreateOrdersForShipperNotiType.Str()): TruckOperatorCreateOrdersForShipperNotiType,
	strings.ToLower(AcceptTripNotiType.Str()):                          AcceptTripNotiType,
	strings.ToLower(StartTripNotiType.Str()):                           StartTripNotiType,
	strings.ToLower(CancelTripNotiType.Str()):                          CancelTripNotiType,
	strings.ToLower(SystemCancelTripNotiType.Str()):                    SystemCancelTripNotiType,

	strings.ToLower(CompleteOrderNotiType.Str()):         CompleteOrderNotiType,
	strings.ToLower(SOSOrderNotiType.Str()):              SOSOrderNotiType,
	strings.ToLower(DriverArriveNotiType.Str()):          DriverArriveNotiType,
	strings.ToLower(DriverContinueNotiType.Str()):        DriverContinueNotiType,
	strings.ToLower(AdminApproveDriverDocNotiType.Str()): AdminApproveDriverDocNotiType,

	strings.ToLower(AdminRejectDriverDocNotiType.Str()):      AdminRejectDriverDocNotiType,
	strings.ToLower(UpcomingScheduleNotiType.Str()):          UpcomingScheduleNotiType,
	strings.ToLower(TruckOperatorInviteDriverNotiType.Str()): TruckOperatorInviteDriverNotiType,
	strings.ToLower(TruckOperatorDeleteDriverNotiType.Str()): TruckOperatorDeleteDriverNotiType,
	strings.ToLower(TruckOperatorAssignNotiType.Str()):       TruckOperatorAssignNotiType,

	strings.ToLower(TruckOperatorUnassignNotiType.Str()): TruckOperatorUnassignNotiType,
	strings.ToLower(CreateBookingNotiType.Str()):         CreateBookingNotiType,
	strings.ToLower(AdminRejectOrderNotiType.Str()):      AdminRejectOrderNotiType,
	strings.ToLower(ShippingCancelTripNotiType.Str()):    ShippingCancelTripNotiType,
	strings.ToLower(AdminRejectCancelNotiType.Str()):     AdminRejectCancelNotiType,

	strings.ToLower(AdminAssignCarrierNotiType.Str()):             AdminAssignCarrierNotiType,
	strings.ToLower(TruckOperatorApproveCancelTripNotiType.Str()): TruckOperatorApproveCancelTripNotiType,
	strings.ToLower(TruckOperatorRejectCancelTripNotiType.Str()):  TruckOperatorRejectCancelTripNotiType,
	strings.ToLower(TruckOperatorCancelSuccessNotiType.Str()):     TruckOperatorCancelSuccessNotiType,
	strings.ToLower(ShipperCancelSuccessNotiType.Str()):           ShipperCancelSuccessNotiType,

	strings.ToLower(ShipperCancelFailNotiType.Str()):         ShipperCancelFailNotiType,
	strings.ToLower(TruckOperatorCancelTripNotiType.Str()):   TruckOperatorCancelTripNotiType,
	strings.ToLower(AdminApproveCancelTripNotiType.Str()):    AdminApproveCancelTripNotiType,
	strings.ToLower(TruckOperatorCancelDriverNotiType.Str()): TruckOperatorCancelDriverNotiType,
	strings.ToLower(EditTripNotiType.Str()):                  EditTripNotiType,

	strings.ToLower(AlertVehicleWrongWayNotiType.Str()):            AlertVehicleWrongWayNotiType,
	strings.ToLower(AlertCompleteTripOutsideAreaNotiType.Str()):    AlertCompleteTripOutsideAreaNotiType,
	strings.ToLower(AlertArrivedScheduleOutsideAreaNotiType.Str()): AlertArrivedScheduleOutsideAreaNotiType,
	strings.ToLower(CarrierAssignedForAccident.Str()):              CarrierAssignedForAccident,
	strings.ToLower(HandOverToDriverForAccident.Str()):             HandOverToDriverForAccident,
	strings.ToLower(PasswordWillExpireAlertNotiType.Str()):         PasswordWillExpireAlertNotiType,
	strings.ToLower(SubcontractorInvitationNewNotiType.Str()):      SubcontractorInvitationNewNotiType,
	strings.ToLower(SubcontractorAcceptInvitationNotiType.Str()): SubcontractorAcceptInvitationNotiType,
	strings.ToLower(SubcontractorCanceledInvitationNotiType.Str()): SubcontractorCanceledInvitationNotiType,

	strings.ToLower(ETANotiType.Str()): ETANotiType,
}

func (s NotiType) Str() string {
	return notiTypeStrArr[s]
}

// GetNotiTypeEnumByInt : return const by id
func GetNotiTypeEnumByInt(index int) NotiType {
	return notiTypeArr[index]
}

// GetNotiTypeEnum : return const by id
func GetNotiTypeEnum(s string) NotiType {
	return notifyTypeMap[strings.ToLower(s)]
}

// Int parses enum to int
func (s NotiType) Int() int {
	return int(s)
}
