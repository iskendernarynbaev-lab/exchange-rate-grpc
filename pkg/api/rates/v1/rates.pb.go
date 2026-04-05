// Code generated manually to match protoc output conventions. DO NOT EDIT.

package ratesv1

import (
	reflect "reflect"
	sync "sync"

	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetRatesRequest is request payload for dynamic calculation parameters.
type GetRatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	N      int32  `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	M      int32  `protobuf:"varint,3,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *GetRatesRequest) Reset() {
	*x = GetRatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_rates_v1_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesRequest) ProtoMessage() {}

func (x *GetRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_rates_v1_rates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetRatesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_rates_v1_rates_proto_rawDescGZIP(), []int{0}
}

func (x *GetRatesRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GetRatesRequest) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *GetRatesRequest) GetM() int32 {
	if x != nil {
		return x.M
	}
	return 0
}

// GetRatesResponse is a strongly typed gRPC response payload.
type GetRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ask        float64 `protobuf:"fixed64,1,opt,name=ask,proto3" json:"ask,omitempty"`
	Bid        float64 `protobuf:"fixed64,2,opt,name=bid,proto3" json:"bid,omitempty"`
	Timestamp  string  `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UnixMillis int64   `protobuf:"varint,4,opt,name=unix_millis,json=unixMillis,proto3" json:"unix_millis,omitempty"`
}

func (x *GetRatesResponse) Reset() {
	*x = GetRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_rates_v1_rates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesResponse) ProtoMessage() {}

func (x *GetRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_rates_v1_rates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_rates_v1_rates_proto_rawDescGZIP(), []int{1}
}

func (x *GetRatesResponse) GetAsk() float64 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *GetRatesResponse) GetBid() float64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *GetRatesResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetRatesResponse) GetUnixMillis() int64 {
	if x != nil {
		return x.UnixMillis
	}
	return 0
}

var File_pkg_api_rates_v1_rates_proto protoreflect.FileDescriptor

var file_pkg_api_rates_v1_rates_proto_rawDescOnce sync.Once
var file_pkg_api_rates_v1_rates_proto_rawDescData []byte

func file_pkg_api_rates_v1_rates_proto_rawDescGZIP() []byte {
	file_pkg_api_rates_v1_rates_proto_rawDescOnce.Do(func() {
		file_pkg_api_rates_v1_rates_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_rates_v1_rates_proto_rawDescData)
	})
	return file_pkg_api_rates_v1_rates_proto_rawDescData
}

var file_pkg_api_rates_v1_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_api_rates_v1_rates_proto_goTypes = []any{
	(*GetRatesRequest)(nil),  // 0: rates.v1.GetRatesRequest
	(*GetRatesResponse)(nil), // 1: rates.v1.GetRatesResponse
}
var file_pkg_api_rates_v1_rates_proto_depIdxs = []int32{
	0, // 0: rates.v1.RatesService.GetRates:input_type -> rates.v1.GetRatesRequest
	1, // 1: rates.v1.RatesService.GetRates:output_type -> rates.v1.GetRatesResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_rates_v1_rates_proto_init() }
func file_pkg_api_rates_v1_rates_proto_init() {
	if File_pkg_api_rates_v1_rates_proto != nil {
		return
	}

	fd := &descriptorpb.FileDescriptorProto{
		Syntax:  proto.String("proto3"),
		Name:    proto.String("pkg/api/rates/v1/rates.proto"),
		Package: proto.String("rates.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: proto.String("GetRatesRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     proto.String("method"),
						JsonName: proto.String("method"),
						Number:   proto.Int32(1),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
					{
						Name:     proto.String("n"),
						JsonName: proto.String("n"),
						Number:   proto.Int32(2),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
					},
					{
						Name:     proto.String("m"),
						JsonName: proto.String("m"),
						Number:   proto.Int32(3),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_INT32.Enum(),
					},
				},
			},
			{
				Name: proto.String("GetRatesResponse"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:     proto.String("ask"),
						JsonName: proto.String("ask"),
						Number:   proto.Int32(1),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
					},
					{
						Name:     proto.String("bid"),
						JsonName: proto.String("bid"),
						Number:   proto.Int32(2),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_DOUBLE.Enum(),
					},
					{
						Name:     proto.String("timestamp"),
						JsonName: proto.String("timestamp"),
						Number:   proto.Int32(3),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
					{
						Name:     proto.String("unix_millis"),
						JsonName: proto.String("unixMillis"),
						Number:   proto.Int32(4),
						Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:     descriptorpb.FieldDescriptorProto_TYPE_INT64.Enum(),
					},
				},
			},
		},
		Service: []*descriptorpb.ServiceDescriptorProto{
			{
				Name: proto.String("RatesService"),
				Method: []*descriptorpb.MethodDescriptorProto{
					{
						Name:       proto.String("GetRates"),
						InputType:  proto.String(".rates.v1.GetRatesRequest"),
						OutputType: proto.String(".rates.v1.GetRatesResponse"),
					},
				},
			},
		},
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/iskendernarynbaev-lab/exchange-rate-grpc/pkg/api/rates/v1;ratesv1"),
		},
	}

	rawDesc, _ := proto.Marshal(fd)
	file_pkg_api_rates_v1_rates_proto_rawDescData = rawDesc

	if !protoimpl.UnsafeEnabled {
		file_pkg_api_rates_v1_rates_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetRatesRequest); i {
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
		file_pkg_api_rates_v1_rates_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetRatesResponse); i {
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
			RawDescriptor: file_pkg_api_rates_v1_rates_proto_rawDescData,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_rates_v1_rates_proto_goTypes,
		DependencyIndexes: file_pkg_api_rates_v1_rates_proto_depIdxs,
		MessageInfos:      file_pkg_api_rates_v1_rates_proto_msgTypes,
	}.Build()

	File_pkg_api_rates_v1_rates_proto = out.File
	file_pkg_api_rates_v1_rates_proto_rawDescData = nil
	file_pkg_api_rates_v1_rates_proto_goTypes = nil
	file_pkg_api_rates_v1_rates_proto_depIdxs = nil
}
