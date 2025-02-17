// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: einride/saga/extend/book/v1beta1/booking.proto

package bookv1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The booking type.
type Booking_BookingType int32

const (
	// Unspecified booking type
	Booking_BOOKING_TYPE_UNSPECIFIED Booking_BookingType = 0
	// The booking is provisional and can be updated.
	Booking_PROVISIONAL Booking_BookingType = 1
	// The booking is confirmed and can not be updated.
	Booking_CONFIRMED Booking_BookingType = 2
)

// Enum value maps for Booking_BookingType.
var (
	Booking_BookingType_name = map[int32]string{
		0: "BOOKING_TYPE_UNSPECIFIED",
		1: "PROVISIONAL",
		2: "CONFIRMED",
	}
	Booking_BookingType_value = map[string]int32{
		"BOOKING_TYPE_UNSPECIFIED": 0,
		"PROVISIONAL":              1,
		"CONFIRMED":                2,
	}
)

func (x Booking_BookingType) Enum() *Booking_BookingType {
	p := new(Booking_BookingType)
	*p = x
	return p
}

func (x Booking_BookingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_BookingType) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[0].Descriptor()
}

func (Booking_BookingType) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[0]
}

func (x Booking_BookingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_BookingType.Descriptor instead.
func (Booking_BookingType) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP(), []int{0, 0}
}

// The state of the booking.
type Booking_State int32

const (
	// Unspecified state
	Booking_STATE_UNSPECIFIED Booking_State = 0
	// The booking is received. Awaiting accept/reject.
	Booking_PENDING Booking_State = 1
	// The booking is accepted.
	Booking_ACCEPTED Booking_State = 2
	// The booking is rejected.
	Booking_REJECTED Booking_State = 3
)

// Enum value maps for Booking_State.
var (
	Booking_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PENDING",
		2: "ACCEPTED",
		3: "REJECTED",
	}
	Booking_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PENDING":           1,
		"ACCEPTED":          2,
		"REJECTED":          3,
	}
)

func (x Booking_State) Enum() *Booking_State {
	p := new(Booking_State)
	*p = x
	return p
}

func (x Booking_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_State) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[1].Descriptor()
}

func (Booking_State) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[1]
}

func (x Booking_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_State.Descriptor instead.
func (Booking_State) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP(), []int{0, 1}
}

// The service type of the booking.
type Booking_ServiceType int32

const (
	// Unspecified service type.
	Booking_SERVICE_TYPE_UNSPECIFIED Booking_ServiceType = 0
	// Service type FTL.
	Booking_FULL_TRUCK_LOAD Booking_ServiceType = 1
	// Service type distribution.
	Booking_DISTRIBUTION Booking_ServiceType = 2
	// Service type shuttle.
	Booking_SHUTTLE Booking_ServiceType = 3
	// Service type milk run.
	Booking_MILK_RUN Booking_ServiceType = 4
)

// Enum value maps for Booking_ServiceType.
var (
	Booking_ServiceType_name = map[int32]string{
		0: "SERVICE_TYPE_UNSPECIFIED",
		1: "FULL_TRUCK_LOAD",
		2: "DISTRIBUTION",
		3: "SHUTTLE",
		4: "MILK_RUN",
	}
	Booking_ServiceType_value = map[string]int32{
		"SERVICE_TYPE_UNSPECIFIED": 0,
		"FULL_TRUCK_LOAD":          1,
		"DISTRIBUTION":             2,
		"SHUTTLE":                  3,
		"MILK_RUN":                 4,
	}
)

func (x Booking_ServiceType) Enum() *Booking_ServiceType {
	p := new(Booking_ServiceType)
	*p = x
	return p
}

func (x Booking_ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[2].Descriptor()
}

func (Booking_ServiceType) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[2]
}

func (x Booking_ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_ServiceType.Descriptor instead.
func (Booking_ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP(), []int{0, 2}
}

// The rule that is applied when a confirmed booking is accepted.
type Booking_AutomationRule int32

const (
	// Unspecified auto rule.
	Booking_AUTOMATION_RULE_UNSPECIFIED Booking_AutomationRule = 0
	// When the confirmed booking is accepted, shipments will be created automatically from its stops and units.
	Booking_CREATE_SHIPMENTS Booking_AutomationRule = 1
	// When the confirmed booking is accepted, shipments will be created automatically from its stops and units
	// and then released.
	Booking_CREATE_AND_RELEASE_SHIPMENTS Booking_AutomationRule = 2
)

// Enum value maps for Booking_AutomationRule.
var (
	Booking_AutomationRule_name = map[int32]string{
		0: "AUTOMATION_RULE_UNSPECIFIED",
		1: "CREATE_SHIPMENTS",
		2: "CREATE_AND_RELEASE_SHIPMENTS",
	}
	Booking_AutomationRule_value = map[string]int32{
		"AUTOMATION_RULE_UNSPECIFIED":  0,
		"CREATE_SHIPMENTS":             1,
		"CREATE_AND_RELEASE_SHIPMENTS": 2,
	}
)

func (x Booking_AutomationRule) Enum() *Booking_AutomationRule {
	p := new(Booking_AutomationRule)
	*p = x
	return p
}

func (x Booking_AutomationRule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_AutomationRule) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[3].Descriptor()
}

func (Booking_AutomationRule) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[3]
}

func (x Booking_AutomationRule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_AutomationRule.Descriptor instead.
func (Booking_AutomationRule) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP(), []int{0, 3}
}

// Type describes the reason for the stop. E.g. a stop to deliver or pickup.
type Booking_Stop_Type int32

const (
	// Unknown type.
	Booking_Stop_TYPE_UNSPECIFIED Booking_Stop_Type = 0
	// Stop to pick up goods.
	Booking_Stop_PICKUP Booking_Stop_Type = 1
	// Stop to deliver goods.
	Booking_Stop_DELIVER Booking_Stop_Type = 2
)

// Enum value maps for Booking_Stop_Type.
var (
	Booking_Stop_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "PICKUP",
		2: "DELIVER",
	}
	Booking_Stop_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"PICKUP":           1,
		"DELIVER":          2,
	}
)

func (x Booking_Stop_Type) Enum() *Booking_Stop_Type {
	p := new(Booking_Stop_Type)
	*p = x
	return p
}

func (x Booking_Stop_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Booking_Stop_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[4].Descriptor()
}

func (Booking_Stop_Type) Type() protoreflect.EnumType {
	return &file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes[4]
}

func (x Booking_Stop_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Booking_Stop_Type.Descriptor instead.
func (Booking_Stop_Type) EnumDescriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP(), []int{0, 0, 0}
}

// A booking represents an estimated demand to deliver goods.
type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the booking.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Resource name of the space that owns the booking.
	Space string `protobuf:"bytes,2,opt,name=space,proto3" json:"space,omitempty"`
	// Resource name of the sender of the booking.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// The creation timestamp of the booking.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the booking.
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the booking. Set if the booking is deleted.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Booking type. Defaults to PROVISIONAL.
	Type Booking_BookingType `protobuf:"varint,7,opt,name=type,proto3,enum=einride.saga.extend.book.v1beta1.Booking_BookingType" json:"type,omitempty"`
	// The type of service to create a booking for.
	ServiceType Booking_ServiceType `protobuf:"varint,8,opt,name=service_type,json=serviceType,proto3,enum=einride.saga.extend.book.v1beta1.Booking_ServiceType" json:"service_type,omitempty"`
	// External reference ID is a unique identifier that can be set by the client.
	ExternalReferenceId string `protobuf:"bytes,9,opt,name=external_reference_id,json=externalReferenceId,proto3" json:"external_reference_id,omitempty"`
	// The rule that is applied when a confirmed booking is accepted. Defaults to no auto rule.
	AutomationRule Booking_AutomationRule `protobuf:"varint,10,opt,name=automation_rule,json=automationRule,proto3,enum=einride.saga.extend.book.v1beta1.Booking_AutomationRule" json:"automation_rule,omitempty"`
	// The units included in this booking.
	Units []*Unit `protobuf:"bytes,11,rep,name=units,proto3" json:"units,omitempty"`
	// The stops included in this booking.
	Stops []*Booking_Stop `protobuf:"bytes,12,rep,name=stops,proto3" json:"stops,omitempty"`
	// Tracking ID for the booking.
	TrackingId string `protobuf:"bytes,13,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// The state of the booking. Set to PENDING upon creation.
	// A booking is then either accepted or rejected.
	State Booking_State `protobuf:"varint,14,opt,name=state,proto3,enum=einride.saga.extend.book.v1beta1.Booking_State" json:"state,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Booking) GetSpace() string {
	if x != nil {
		return x.Space
	}
	return ""
}

func (x *Booking) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Booking) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Booking) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Booking) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Booking) GetType() Booking_BookingType {
	if x != nil {
		return x.Type
	}
	return Booking_BOOKING_TYPE_UNSPECIFIED
}

func (x *Booking) GetServiceType() Booking_ServiceType {
	if x != nil {
		return x.ServiceType
	}
	return Booking_SERVICE_TYPE_UNSPECIFIED
}

func (x *Booking) GetExternalReferenceId() string {
	if x != nil {
		return x.ExternalReferenceId
	}
	return ""
}

func (x *Booking) GetAutomationRule() Booking_AutomationRule {
	if x != nil {
		return x.AutomationRule
	}
	return Booking_AUTOMATION_RULE_UNSPECIFIED
}

func (x *Booking) GetUnits() []*Unit {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *Booking) GetStops() []*Booking_Stop {
	if x != nil {
		return x.Stops
	}
	return nil
}

func (x *Booking) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *Booking) GetState() Booking_State {
	if x != nil {
		return x.State
	}
	return Booking_STATE_UNSPECIFIED
}

// A stop is an address where an event is going to happen. The event could be
// to deliver goods, charge the vehicle etc.
type Booking_Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The address of this stop.
	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The type of stop.
	StopType Booking_Stop_Type `protobuf:"varint,2,opt,name=stop_type,json=stopType,proto3,enum=einride.saga.extend.book.v1beta1.Booking_Stop_Type" json:"stop_type,omitempty"`
	// Instructions for this stop.
	Instructions string `protobuf:"bytes,3,opt,name=instructions,proto3" json:"instructions,omitempty"`
	// The earliest time to arrive this stop.
	RequestedStartTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=requested_start_time,json=requestedStartTime,proto3" json:"requested_start_time,omitempty"`
	// The latest time to depart this stop.
	RequestedEndTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=requested_end_time,json=requestedEndTime,proto3" json:"requested_end_time,omitempty"`
	// The unit external reference id associated with this stop. E.g. for a pickup stop this will hold the units to pickup.
	UnitExternalReferenceIds []string `protobuf:"bytes,6,rep,name=unit_external_reference_ids,json=unitExternalReferenceIds,proto3" json:"unit_external_reference_ids,omitempty"`
}

func (x *Booking_Stop) Reset() {
	*x = Booking_Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking_Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking_Stop) ProtoMessage() {}

func (x *Booking_Stop) ProtoReflect() protoreflect.Message {
	mi := &file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking_Stop.ProtoReflect.Descriptor instead.
func (*Booking_Stop) Descriptor() ([]byte, []int) {
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Booking_Stop) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Booking_Stop) GetStopType() Booking_Stop_Type {
	if x != nil {
		return x.StopType
	}
	return Booking_Stop_TYPE_UNSPECIFIED
}

func (x *Booking_Stop) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *Booking_Stop) GetRequestedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedStartTime
	}
	return nil
}

func (x *Booking_Stop) GetRequestedEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedEndTime
	}
	return nil
}

func (x *Booking_Stop) GetUnitExternalReferenceIds() []string {
	if x != nil {
		return x.UnitExternalReferenceIds
	}
	return nil
}

var File_einride_saga_extend_book_v1beta1_booking_proto protoreflect.FileDescriptor

var file_einride_saga_extend_book_v1beta1_booking_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x0e, 0x0a,
	0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x05,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xe2, 0x41, 0x01,
	0x03, 0xfa, 0x41, 0x1d, 0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41,
	0x1e, 0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x49, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x5e, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x61, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52,
	0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x70, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e,
	0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05, 0x73, 0x74, 0x6f,
	0x70, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0xe1, 0x03, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12,
	0x49, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x04, 0xe2, 0x41, 0x01,
	0x02, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x56, 0x0a, 0x09, 0x73, 0x74,
	0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e,
	0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43,
	0x0a, 0x1b, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x18, 0x75, 0x6e, 0x69, 0x74, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x10, 0x02, 0x22, 0x4b, 0x0a, 0x0b, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x4f, 0x4f,
	0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x56, 0x49,
	0x53, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x22, 0x47, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x22, 0x6d, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x54, 0x52, 0x55, 0x43, 0x4b, 0x5f, 0x4c, 0x4f, 0x41, 0x44,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x48, 0x55, 0x54, 0x54, 0x4c, 0x45, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x49, 0x4c, 0x4b, 0x5f, 0x52, 0x55, 0x4e, 0x10, 0x04, 0x22,
	0x69, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x52, 0x55, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x48, 0x49,
	0x50, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x53,
	0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x02, 0x3a, 0x58, 0xea, 0x41, 0x55, 0x0a,
	0x1d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x21,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x7d, 0x2a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x32, 0x07, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x42, 0xac, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x73, 0x61, 0x67, 0x61, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x73, 0x61, 0x67, 0x61,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xa2, 0x02, 0x04, 0x45, 0x53, 0x45, 0x42, 0xaa, 0x02, 0x20, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x53, 0x61, 0x67, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2c,
	0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x53, 0x61, 0x67, 0x61, 0x5c, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x45,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x53, 0x61, 0x67, 0x61, 0x3a, 0x3a, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_saga_extend_book_v1beta1_booking_proto_rawDescOnce sync.Once
	file_einride_saga_extend_book_v1beta1_booking_proto_rawDescData = file_einride_saga_extend_book_v1beta1_booking_proto_rawDesc
)

func file_einride_saga_extend_book_v1beta1_booking_proto_rawDescGZIP() []byte {
	file_einride_saga_extend_book_v1beta1_booking_proto_rawDescOnce.Do(func() {
		file_einride_saga_extend_book_v1beta1_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_saga_extend_book_v1beta1_booking_proto_rawDescData)
	})
	return file_einride_saga_extend_book_v1beta1_booking_proto_rawDescData
}

var file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_einride_saga_extend_book_v1beta1_booking_proto_goTypes = []interface{}{
	(Booking_BookingType)(0),      // 0: einride.saga.extend.book.v1beta1.Booking.BookingType
	(Booking_State)(0),            // 1: einride.saga.extend.book.v1beta1.Booking.State
	(Booking_ServiceType)(0),      // 2: einride.saga.extend.book.v1beta1.Booking.ServiceType
	(Booking_AutomationRule)(0),   // 3: einride.saga.extend.book.v1beta1.Booking.AutomationRule
	(Booking_Stop_Type)(0),        // 4: einride.saga.extend.book.v1beta1.Booking.Stop.Type
	(*Booking)(nil),               // 5: einride.saga.extend.book.v1beta1.Booking
	(*Booking_Stop)(nil),          // 6: einride.saga.extend.book.v1beta1.Booking.Stop
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*Unit)(nil),                  // 8: einride.saga.extend.book.v1beta1.Unit
	(*Address)(nil),               // 9: einride.saga.extend.book.v1beta1.Address
}
var file_einride_saga_extend_book_v1beta1_booking_proto_depIdxs = []int32{
	7,  // 0: einride.saga.extend.book.v1beta1.Booking.create_time:type_name -> google.protobuf.Timestamp
	7,  // 1: einride.saga.extend.book.v1beta1.Booking.update_time:type_name -> google.protobuf.Timestamp
	7,  // 2: einride.saga.extend.book.v1beta1.Booking.delete_time:type_name -> google.protobuf.Timestamp
	0,  // 3: einride.saga.extend.book.v1beta1.Booking.type:type_name -> einride.saga.extend.book.v1beta1.Booking.BookingType
	2,  // 4: einride.saga.extend.book.v1beta1.Booking.service_type:type_name -> einride.saga.extend.book.v1beta1.Booking.ServiceType
	3,  // 5: einride.saga.extend.book.v1beta1.Booking.automation_rule:type_name -> einride.saga.extend.book.v1beta1.Booking.AutomationRule
	8,  // 6: einride.saga.extend.book.v1beta1.Booking.units:type_name -> einride.saga.extend.book.v1beta1.Unit
	6,  // 7: einride.saga.extend.book.v1beta1.Booking.stops:type_name -> einride.saga.extend.book.v1beta1.Booking.Stop
	1,  // 8: einride.saga.extend.book.v1beta1.Booking.state:type_name -> einride.saga.extend.book.v1beta1.Booking.State
	9,  // 9: einride.saga.extend.book.v1beta1.Booking.Stop.address:type_name -> einride.saga.extend.book.v1beta1.Address
	4,  // 10: einride.saga.extend.book.v1beta1.Booking.Stop.stop_type:type_name -> einride.saga.extend.book.v1beta1.Booking.Stop.Type
	7,  // 11: einride.saga.extend.book.v1beta1.Booking.Stop.requested_start_time:type_name -> google.protobuf.Timestamp
	7,  // 12: einride.saga.extend.book.v1beta1.Booking.Stop.requested_end_time:type_name -> google.protobuf.Timestamp
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_einride_saga_extend_book_v1beta1_booking_proto_init() }
func file_einride_saga_extend_book_v1beta1_booking_proto_init() {
	if File_einride_saga_extend_book_v1beta1_booking_proto != nil {
		return
	}
	file_einride_saga_extend_book_v1beta1_address_proto_init()
	file_einride_saga_extend_book_v1beta1_unit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking_Stop); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_saga_extend_book_v1beta1_booking_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_saga_extend_book_v1beta1_booking_proto_goTypes,
		DependencyIndexes: file_einride_saga_extend_book_v1beta1_booking_proto_depIdxs,
		EnumInfos:         file_einride_saga_extend_book_v1beta1_booking_proto_enumTypes,
		MessageInfos:      file_einride_saga_extend_book_v1beta1_booking_proto_msgTypes,
	}.Build()
	File_einride_saga_extend_book_v1beta1_booking_proto = out.File
	file_einride_saga_extend_book_v1beta1_booking_proto_rawDesc = nil
	file_einride_saga_extend_book_v1beta1_booking_proto_goTypes = nil
	file_einride_saga_extend_book_v1beta1_booking_proto_depIdxs = nil
}
