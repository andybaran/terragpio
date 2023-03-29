// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: terragpio.proto

package terragpio

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PinSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PinNumber string `protobuf:"bytes,1,opt,name=pinNumber,proto3" json:"pinNumber,omitempty"`
}

func (x *PinSetResponse) Reset() {
	*x = PinSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PinSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PinSetResponse) ProtoMessage() {}

func (x *PinSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PinSetResponse.ProtoReflect.Descriptor instead.
func (*PinSetResponse) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{0}
}

func (x *PinSetResponse) GetPinNumber() string {
	if x != nil {
		return x.PinNumber
	}
	return ""
}

type PWMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pin       string `protobuf:"bytes,1,opt,name=pin,proto3" json:"pin,omitempty"`             // expect the pin to be expressed in terms of GPIO standard (i.e. GPIO6)
	Dutycycle string `protobuf:"bytes,2,opt,name=dutycycle,proto3" json:"dutycycle,omitempty"` // format is "nn%" where nn is 00 - 100
	Frequency string `protobuf:"bytes,3,opt,name=frequency,proto3" json:"frequency,omitempty"` // format is "nM" where n is Mega Hertz
}

func (x *PWMRequest) Reset() {
	*x = PWMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PWMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PWMRequest) ProtoMessage() {}

func (x *PWMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PWMRequest.ProtoReflect.Descriptor instead.
func (*PWMRequest) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{1}
}

func (x *PWMRequest) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

func (x *PWMRequest) GetDutycycle() string {
	if x != nil {
		return x.Dutycycle
	}
	return ""
}

func (x *PWMRequest) GetFrequency() string {
	if x != nil {
		return x.Frequency
	}
	return ""
}

type PWMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dutycycle string `protobuf:"bytes,1,opt,name=dutycycle,proto3" json:"dutycycle,omitempty"`
	Frequency string `protobuf:"bytes,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (x *PWMResponse) Reset() {
	*x = PWMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PWMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PWMResponse) ProtoMessage() {}

func (x *PWMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PWMResponse.ProtoReflect.Descriptor instead.
func (*PWMResponse) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{2}
}

func (x *PWMResponse) GetDutycycle() string {
	if x != nil {
		return x.Dutycycle
	}
	return ""
}

func (x *PWMResponse) GetFrequency() string {
	if x != nil {
		return x.Frequency
	}
	return ""
}

type BME280Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I2Cbus  string `protobuf:"bytes,1,opt,name=I2Cbus,proto3" json:"I2Cbus,omitempty"`
	I2Caddr string `protobuf:"bytes,2,opt,name=I2Caddr,proto3" json:"I2Caddr,omitempty"` // format is 0x76
}

func (x *BME280Request) Reset() {
	*x = BME280Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BME280Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BME280Request) ProtoMessage() {}

func (x *BME280Request) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BME280Request.ProtoReflect.Descriptor instead.
func (*BME280Request) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{3}
}

func (x *BME280Request) GetI2Cbus() string {
	if x != nil {
		return x.I2Cbus
	}
	return ""
}

func (x *BME280Request) GetI2Caddr() string {
	if x != nil {
		return x.I2Caddr
	}
	return ""
}

type BME280Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature string `protobuf:"bytes,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    string `protobuf:"bytes,2,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Pressure    string `protobuf:"bytes,3,opt,name=pressure,proto3" json:"pressure,omitempty"`
}

func (x *BME280Response) Reset() {
	*x = BME280Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BME280Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BME280Response) ProtoMessage() {}

func (x *BME280Response) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BME280Response.ProtoReflect.Descriptor instead.
func (*BME280Response) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{4}
}

func (x *BME280Response) GetTemperature() string {
	if x != nil {
		return x.Temperature
	}
	return ""
}

func (x *BME280Response) GetHumidity() string {
	if x != nil {
		return x.Humidity
	}
	return ""
}

func (x *BME280Response) GetPressure() string {
	if x != nil {
		return x.Pressure
	}
	return ""
}

type PinSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PinNumber string `protobuf:"bytes,1,opt,name=pinNumber,proto3" json:"pinNumber,omitempty"`
}

func (x *PinSetRequest) Reset() {
	*x = PinSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PinSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PinSetRequest) ProtoMessage() {}

func (x *PinSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PinSetRequest.ProtoReflect.Descriptor instead.
func (*PinSetRequest) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{5}
}

func (x *PinSetRequest) GetPinNumber() string {
	if x != nil {
		return x.PinNumber
	}
	return ""
}

type FanControllerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeInterval    uint64 `protobuf:"varint,1,opt,name=timeInterval,proto3" json:"timeInterval,omitempty"`
	BME280DevicePin string `protobuf:"bytes,2,opt,name=BME280DevicePin,proto3" json:"BME280DevicePin,omitempty"`
	TemperatureMax  uint64 `protobuf:"varint,3,opt,name=temperatureMax,proto3" json:"temperatureMax,omitempty"`
	TemperatureMin  uint64 `protobuf:"varint,4,opt,name=temperatureMin,proto3" json:"temperatureMin,omitempty"`
	FanDevicePin    string `protobuf:"bytes,5,opt,name=fanDevicePin,proto3" json:"fanDevicePin,omitempty"`
	DutyCycleMax    uint64 `protobuf:"varint,6,opt,name=dutyCycleMax,proto3" json:"dutyCycleMax,omitempty"`
	DutyCycleMin    uint64 `protobuf:"varint,7,opt,name=dutyCycleMin,proto3" json:"dutyCycleMin,omitempty"`
}

func (x *FanControllerRequest) Reset() {
	*x = FanControllerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanControllerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanControllerRequest) ProtoMessage() {}

func (x *FanControllerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanControllerRequest.ProtoReflect.Descriptor instead.
func (*FanControllerRequest) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{6}
}

func (x *FanControllerRequest) GetTimeInterval() uint64 {
	if x != nil {
		return x.TimeInterval
	}
	return 0
}

func (x *FanControllerRequest) GetBME280DevicePin() string {
	if x != nil {
		return x.BME280DevicePin
	}
	return ""
}

func (x *FanControllerRequest) GetTemperatureMax() uint64 {
	if x != nil {
		return x.TemperatureMax
	}
	return 0
}

func (x *FanControllerRequest) GetTemperatureMin() uint64 {
	if x != nil {
		return x.TemperatureMin
	}
	return 0
}

func (x *FanControllerRequest) GetFanDevicePin() string {
	if x != nil {
		return x.FanDevicePin
	}
	return ""
}

func (x *FanControllerRequest) GetDutyCycleMax() uint64 {
	if x != nil {
		return x.DutyCycleMax
	}
	return 0
}

func (x *FanControllerRequest) GetDutyCycleMin() uint64 {
	if x != nil {
		return x.DutyCycleMin
	}
	return 0
}

type FanControllerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PinCombo string `protobuf:"bytes,1,opt,name=pinCombo,proto3" json:"pinCombo,omitempty"`
}

func (x *FanControllerResponse) Reset() {
	*x = FanControllerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terragpio_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FanControllerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FanControllerResponse) ProtoMessage() {}

func (x *FanControllerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terragpio_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FanControllerResponse.ProtoReflect.Descriptor instead.
func (*FanControllerResponse) Descriptor() ([]byte, []int) {
	return file_terragpio_proto_rawDescGZIP(), []int{7}
}

func (x *FanControllerResponse) GetPinCombo() string {
	if x != nil {
		return x.PinCombo
	}
	return ""
}

var File_terragpio_proto protoreflect.FileDescriptor

var file_terragpio_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x22, 0x2e, 0x0a, 0x0e,
	0x50, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x0a,
	0x50, 0x57, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x75, 0x74, 0x79, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x75, 0x74, 0x79, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x49, 0x0a, 0x0b, 0x50, 0x57, 0x4d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x75, 0x74, 0x79, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x75, 0x74, 0x79,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0x41, 0x0a, 0x0d, 0x42, 0x4d, 0x45, 0x32, 0x38, 0x30, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x32, 0x43, 0x62, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x32, 0x43, 0x62, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x49, 0x32, 0x43, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49,
	0x32, 0x43, 0x61, 0x64, 0x64, 0x72, 0x22, 0x6a, 0x0a, 0x0e, 0x42, 0x4d, 0x45, 0x32, 0x38, 0x30,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x22, 0x2d, 0x0a, 0x0d, 0x50, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0xa0, 0x02, 0x0a, 0x14, 0x46, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x28,
	0x0a, 0x0f, 0x42, 0x4d, 0x45, 0x32, 0x38, 0x30, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x42, 0x4d, 0x45, 0x32, 0x38, 0x30, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x78,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x61, 0x6e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x61, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x64, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x64, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x78,
	0x12, 0x22, 0x0a, 0x0c, 0x64, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x4d, 0x69, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x4d, 0x69, 0x6e, 0x22, 0x33, 0x0a, 0x15, 0x46, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x62, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x62, 0x6f, 0x32, 0xfc, 0x02, 0x0a, 0x07, 0x73, 0x65,
	0x74, 0x67, 0x70, 0x69, 0x6f, 0x12, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x74, 0x50, 0x57, 0x4d, 0x12,
	0x15, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e, 0x50, 0x57, 0x4d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70,
	0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x42, 0x4d, 0x45, 0x32, 0x38, 0x30,
	0x12, 0x18, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e, 0x42, 0x4d, 0x45,
	0x32, 0x38, 0x30, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x73, 0x65,
	0x42, 0x4d, 0x45, 0x32, 0x38, 0x30, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70,
	0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e, 0x42, 0x4d, 0x45,
	0x32, 0x38, 0x30, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x57, 0x4d, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e,
	0x50, 0x57, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a,
	0x22, 0x50, 0x57, 0x4d, 0x44, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x5f, 0x42, 0x4d, 0x45, 0x32, 0x38, 0x30, 0x54, 0x65, 0x6d, 0x70, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2e,
	0x46, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f,
	0x2e, 0x46, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x67, 0x70, 0x69, 0x6f, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x67, 0x70, 0x69,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_terragpio_proto_rawDescOnce sync.Once
	file_terragpio_proto_rawDescData = file_terragpio_proto_rawDesc
)

func file_terragpio_proto_rawDescGZIP() []byte {
	file_terragpio_proto_rawDescOnce.Do(func() {
		file_terragpio_proto_rawDescData = protoimpl.X.CompressGZIP(file_terragpio_proto_rawDescData)
	})
	return file_terragpio_proto_rawDescData
}

var file_terragpio_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_terragpio_proto_goTypes = []interface{}{
	(*PinSetResponse)(nil),        // 0: terragpio.PinSetResponse
	(*PWMRequest)(nil),            // 1: terragpio.PWMRequest
	(*PWMResponse)(nil),           // 2: terragpio.PWMResponse
	(*BME280Request)(nil),         // 3: terragpio.BME280Request
	(*BME280Response)(nil),        // 4: terragpio.BME280Response
	(*PinSetRequest)(nil),         // 5: terragpio.PinSetRequest
	(*FanControllerRequest)(nil),  // 6: terragpio.FanControllerRequest
	(*FanControllerResponse)(nil), // 7: terragpio.FanControllerResponse
}
var file_terragpio_proto_depIdxs = []int32{
	1, // 0: terragpio.setgpio.SetPWM:input_type -> terragpio.PWMRequest
	3, // 1: terragpio.setgpio.SetBME280:input_type -> terragpio.BME280Request
	5, // 2: terragpio.setgpio.SenseBME280:input_type -> terragpio.PinSetRequest
	5, // 3: terragpio.setgpio.SensePWM:input_type -> terragpio.PinSetRequest
	6, // 4: terragpio.setgpio.PWMDutyCycleOutput_BME280TempInput:input_type -> terragpio.FanControllerRequest
	0, // 5: terragpio.setgpio.SetPWM:output_type -> terragpio.PinSetResponse
	0, // 6: terragpio.setgpio.SetBME280:output_type -> terragpio.PinSetResponse
	4, // 7: terragpio.setgpio.SenseBME280:output_type -> terragpio.BME280Response
	2, // 8: terragpio.setgpio.SensePWM:output_type -> terragpio.PWMResponse
	7, // 9: terragpio.setgpio.PWMDutyCycleOutput_BME280TempInput:output_type -> terragpio.FanControllerResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_terragpio_proto_init() }
func file_terragpio_proto_init() {
	if File_terragpio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_terragpio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PinSetResponse); i {
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
		file_terragpio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PWMRequest); i {
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
		file_terragpio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PWMResponse); i {
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
		file_terragpio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BME280Request); i {
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
		file_terragpio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BME280Response); i {
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
		file_terragpio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PinSetRequest); i {
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
		file_terragpio_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanControllerRequest); i {
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
		file_terragpio_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FanControllerResponse); i {
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
			RawDescriptor: file_terragpio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_terragpio_proto_goTypes,
		DependencyIndexes: file_terragpio_proto_depIdxs,
		MessageInfos:      file_terragpio_proto_msgTypes,
	}.Build()
	File_terragpio_proto = out.File
	file_terragpio_proto_rawDesc = nil
	file_terragpio_proto_goTypes = nil
	file_terragpio_proto_depIdxs = nil
}
