// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: company_login_service.proto

package company_v1

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

type CompanyLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CompanyLoginRequest) Reset() {
	*x = CompanyLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_login_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyLoginRequest) ProtoMessage() {}

func (x *CompanyLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_login_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyLoginRequest.ProtoReflect.Descriptor instead.
func (*CompanyLoginRequest) Descriptor() ([]byte, []int) {
	return file_company_login_service_proto_rawDescGZIP(), []int{0}
}

func (x *CompanyLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CompanyLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CompanyLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *CompanyResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Token    string           `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CompanyLoginResponse) Reset() {
	*x = CompanyLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_login_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyLoginResponse) ProtoMessage() {}

func (x *CompanyLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_company_login_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyLoginResponse.ProtoReflect.Descriptor instead.
func (*CompanyLoginResponse) Descriptor() ([]byte, []int) {
	return file_company_login_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyLoginResponse) GetResponse() *CompanyResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *CompanyLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CompanyRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CompanyRegisterRequest) Reset() {
	*x = CompanyRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_login_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyRegisterRequest) ProtoMessage() {}

func (x *CompanyRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_login_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyRegisterRequest.ProtoReflect.Descriptor instead.
func (*CompanyRegisterRequest) Descriptor() ([]byte, []int) {
	return file_company_login_service_proto_rawDescGZIP(), []int{2}
}

func (x *CompanyRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CompanyRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CompanyRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *CompanyResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Data     string           `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CompanyRegisterResponse) Reset() {
	*x = CompanyRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_login_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyRegisterResponse) ProtoMessage() {}

func (x *CompanyRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_company_login_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyRegisterResponse.ProtoReflect.Descriptor instead.
func (*CompanyRegisterResponse) Descriptor() ([]byte, []int) {
	return file_company_login_service_proto_rawDescGZIP(), []int{3}
}

func (x *CompanyRegisterResponse) GetResponse() *CompanyResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *CompanyRegisterResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_company_login_service_proto protoreflect.FileDescriptor

var file_company_login_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4d, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x65, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x50, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x66, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xba, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a,
	0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_login_service_proto_rawDescOnce sync.Once
	file_company_login_service_proto_rawDescData = file_company_login_service_proto_rawDesc
)

func file_company_login_service_proto_rawDescGZIP() []byte {
	file_company_login_service_proto_rawDescOnce.Do(func() {
		file_company_login_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_login_service_proto_rawDescData)
	})
	return file_company_login_service_proto_rawDescData
}

var file_company_login_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_company_login_service_proto_goTypes = []interface{}{
	(*CompanyLoginRequest)(nil),     // 0: company.v1.CompanyLoginRequest
	(*CompanyLoginResponse)(nil),    // 1: company.v1.CompanyLoginResponse
	(*CompanyRegisterRequest)(nil),  // 2: company.v1.CompanyRegisterRequest
	(*CompanyRegisterResponse)(nil), // 3: company.v1.CompanyRegisterResponse
	(*CompanyResponse)(nil),         // 4: company.v1.CompanyResponse
}
var file_company_login_service_proto_depIdxs = []int32{
	4, // 0: company.v1.CompanyLoginResponse.response:type_name -> company.v1.CompanyResponse
	4, // 1: company.v1.CompanyRegisterResponse.response:type_name -> company.v1.CompanyResponse
	0, // 2: company.v1.CompanyLoginService.Login:input_type -> company.v1.CompanyLoginRequest
	2, // 3: company.v1.CompanyLoginService.Register:input_type -> company.v1.CompanyRegisterRequest
	1, // 4: company.v1.CompanyLoginService.Login:output_type -> company.v1.CompanyLoginResponse
	3, // 5: company.v1.CompanyLoginService.Register:output_type -> company.v1.CompanyRegisterResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_company_login_service_proto_init() }
func file_company_login_service_proto_init() {
	if File_company_login_service_proto != nil {
		return
	}
	file_company_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_company_login_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyLoginRequest); i {
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
		file_company_login_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyLoginResponse); i {
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
		file_company_login_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyRegisterRequest); i {
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
		file_company_login_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyRegisterResponse); i {
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
			RawDescriptor: file_company_login_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_login_service_proto_goTypes,
		DependencyIndexes: file_company_login_service_proto_depIdxs,
		MessageInfos:      file_company_login_service_proto_msgTypes,
	}.Build()
	File_company_login_service_proto = out.File
	file_company_login_service_proto_rawDesc = nil
	file_company_login_service_proto_goTypes = nil
	file_company_login_service_proto_depIdxs = nil
}
