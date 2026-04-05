// Code generated manually to match protoc output conventions. DO NOT EDIT.

package ratesv1

import (
	reflect "reflect"
	sync "sync"

	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

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
		mi := &file_pkg_api_rates_v1_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesResponse) ProtoMessage() {}

func (x *GetRatesResponse) ProtoReflect() protoreflect.Message {
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

func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_rates_v1_rates_proto_rawDescGZIP(), []int{0}
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

var file_pkg_api_rates_v1_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_api_rates_v1_rates_proto_goTypes = []any{
	(*GetRatesResponse)(nil), // 0: rates.v1.GetRatesResponse
}
var file_pkg_api_rates_v1_rates_proto_depIdxs = []int32{}

func init() { file_pkg_api_rates_v1_rates_proto_init() }
func file_pkg_api_rates_v1_rates_proto_init() {
	if File_pkg_api_rates_v1_rates_proto != nil {
		return
	}

	fd := &descriptorpb.FileDescriptorProto{
		Syntax:  proto.String("proto3"),
		Name:    proto.String("pkg/api/rates/v1/rates.proto"),
		Package: proto.String("rates.v1"),
		Dependency: []string{
			"google/protobuf/empty.proto",
		},
		MessageType: []*descriptorpb.DescriptorProto{
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
						InputType:  proto.String(".google.protobuf.Empty"),
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
			NumMessages:   1,
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
