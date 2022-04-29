// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: coop/drivers/geo/v1beta1/geocode.proto

package v1beta1

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

type GeocodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceId string `protobuf:"bytes,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *GeocodeRequest) Reset() {
	*x = GeocodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeocodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeocodeRequest) ProtoMessage() {}

func (x *GeocodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeocodeRequest.ProtoReflect.Descriptor instead.
func (*GeocodeRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_geo_v1beta1_geocode_proto_rawDescGZIP(), []int{0}
}

func (x *GeocodeRequest) GetPlaceId() string {
	if x != nil {
		return x.PlaceId
	}
	return ""
}

type GeocodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatLng *LatLng `protobuf:"bytes,1,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
}

func (x *GeocodeResponse) Reset() {
	*x = GeocodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeocodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeocodeResponse) ProtoMessage() {}

func (x *GeocodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeocodeResponse.ProtoReflect.Descriptor instead.
func (*GeocodeResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_geo_v1beta1_geocode_proto_rawDescGZIP(), []int{1}
}

func (x *GeocodeResponse) GetLatLng() *LatLng {
	if x != nil {
		return x.LatLng
	}
	return nil
}

type ReverseGeocodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatLng *LatLng `protobuf:"bytes,1,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
}

func (x *ReverseGeocodeRequest) Reset() {
	*x = ReverseGeocodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseGeocodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseGeocodeRequest) ProtoMessage() {}

func (x *ReverseGeocodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseGeocodeRequest.ProtoReflect.Descriptor instead.
func (*ReverseGeocodeRequest) Descriptor() ([]byte, []int) {
	return file_coop_drivers_geo_v1beta1_geocode_proto_rawDescGZIP(), []int{2}
}

func (x *ReverseGeocodeRequest) GetLatLng() *LatLng {
	if x != nil {
		return x.LatLng
	}
	return nil
}

type ReverseGeocodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaceId string `protobuf:"bytes,1,opt,name=place_id,json=placeId,proto3" json:"place_id,omitempty"`
}

func (x *ReverseGeocodeResponse) Reset() {
	*x = ReverseGeocodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReverseGeocodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReverseGeocodeResponse) ProtoMessage() {}

func (x *ReverseGeocodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReverseGeocodeResponse.ProtoReflect.Descriptor instead.
func (*ReverseGeocodeResponse) Descriptor() ([]byte, []int) {
	return file_coop_drivers_geo_v1beta1_geocode_proto_rawDescGZIP(), []int{3}
}

func (x *ReverseGeocodeResponse) GetPlaceId() string {
	if x != nil {
		return x.PlaceId
	}
	return ""
}

var File_coop_drivers_geo_v1beta1_geocode_proto protoreflect.FileDescriptor

var file_coop_drivers_geo_v1beta1_geocode_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x65, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x25, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x2f, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61, 0x74,
	0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x6f,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x6c, 0x61, 0x74,
	0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6f,
	0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61,
	0x74, 0x4c, 0x6e, 0x67, 0x22, 0x52, 0x0a, 0x15, 0x52, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x47,
	0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x63, 0x6f, 0x6f, 0x70, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x65,
	0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67,
	0x52, 0x06, 0x6c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x22, 0x33, 0x0a, 0x16, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x47, 0x65, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x4b, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x69,
	0x6e, 0x6d, 0x69, 0x63, 0x68, 0x61, 0x65, 0x6c, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x65, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x64,
	0x6c, 0x2f, 0x63, 0x6f, 0x6f, 0x70, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x65, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_coop_drivers_geo_v1beta1_geocode_proto_rawDescOnce sync.Once
	file_coop_drivers_geo_v1beta1_geocode_proto_rawDescData = file_coop_drivers_geo_v1beta1_geocode_proto_rawDesc
)

func file_coop_drivers_geo_v1beta1_geocode_proto_rawDescGZIP() []byte {
	file_coop_drivers_geo_v1beta1_geocode_proto_rawDescOnce.Do(func() {
		file_coop_drivers_geo_v1beta1_geocode_proto_rawDescData = protoimpl.X.CompressGZIP(file_coop_drivers_geo_v1beta1_geocode_proto_rawDescData)
	})
	return file_coop_drivers_geo_v1beta1_geocode_proto_rawDescData
}

var file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_coop_drivers_geo_v1beta1_geocode_proto_goTypes = []interface{}{
	(*GeocodeRequest)(nil),         // 0: coop.drivers.geo.v1beta1.GeocodeRequest
	(*GeocodeResponse)(nil),        // 1: coop.drivers.geo.v1beta1.GeocodeResponse
	(*ReverseGeocodeRequest)(nil),  // 2: coop.drivers.geo.v1beta1.ReverseGeocodeRequest
	(*ReverseGeocodeResponse)(nil), // 3: coop.drivers.geo.v1beta1.ReverseGeocodeResponse
	(*LatLng)(nil),                 // 4: coop.drivers.geo.v1beta1.LatLng
}
var file_coop_drivers_geo_v1beta1_geocode_proto_depIdxs = []int32{
	4, // 0: coop.drivers.geo.v1beta1.GeocodeResponse.lat_lng:type_name -> coop.drivers.geo.v1beta1.LatLng
	4, // 1: coop.drivers.geo.v1beta1.ReverseGeocodeRequest.lat_lng:type_name -> coop.drivers.geo.v1beta1.LatLng
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coop_drivers_geo_v1beta1_geocode_proto_init() }
func file_coop_drivers_geo_v1beta1_geocode_proto_init() {
	if File_coop_drivers_geo_v1beta1_geocode_proto != nil {
		return
	}
	file_coop_drivers_geo_v1beta1_latlng_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeocodeRequest); i {
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
		file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeocodeResponse); i {
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
		file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseGeocodeRequest); i {
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
		file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReverseGeocodeResponse); i {
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
			RawDescriptor: file_coop_drivers_geo_v1beta1_geocode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coop_drivers_geo_v1beta1_geocode_proto_goTypes,
		DependencyIndexes: file_coop_drivers_geo_v1beta1_geocode_proto_depIdxs,
		MessageInfos:      file_coop_drivers_geo_v1beta1_geocode_proto_msgTypes,
	}.Build()
	File_coop_drivers_geo_v1beta1_geocode_proto = out.File
	file_coop_drivers_geo_v1beta1_geocode_proto_rawDesc = nil
	file_coop_drivers_geo_v1beta1_geocode_proto_goTypes = nil
	file_coop_drivers_geo_v1beta1_geocode_proto_depIdxs = nil
}
