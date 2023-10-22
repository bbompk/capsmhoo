// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/teamjoinrequest.proto

package team_join_request_pb

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

type TeamJoinReqeustEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TeamJoinReqeustEmpty) Reset() {
	*x = TeamJoinReqeustEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teamjoinrequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamJoinReqeustEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamJoinReqeustEmpty) ProtoMessage() {}

func (x *TeamJoinReqeustEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teamjoinrequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamJoinReqeustEmpty.ProtoReflect.Descriptor instead.
func (*TeamJoinReqeustEmpty) Descriptor() ([]byte, []int) {
	return file_proto_teamjoinrequest_proto_rawDescGZIP(), []int{0}
}

type TeamJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId    string `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	StudentId string `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *TeamJoinRequest) Reset() {
	*x = TeamJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teamjoinrequest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamJoinRequest) ProtoMessage() {}

func (x *TeamJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teamjoinrequest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamJoinRequest.ProtoReflect.Descriptor instead.
func (*TeamJoinRequest) Descriptor() ([]byte, []int) {
	return file_proto_teamjoinrequest_proto_rawDescGZIP(), []int{1}
}

func (x *TeamJoinRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TeamJoinRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *TeamJoinRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type TeamJoinRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JoinRequests []*TeamJoinRequest `protobuf:"bytes,1,rep,name=join_requests,json=joinRequests,proto3" json:"join_requests,omitempty"`
}

func (x *TeamJoinRequestList) Reset() {
	*x = TeamJoinRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teamjoinrequest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamJoinRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamJoinRequestList) ProtoMessage() {}

func (x *TeamJoinRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teamjoinrequest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamJoinRequestList.ProtoReflect.Descriptor instead.
func (*TeamJoinRequestList) Descriptor() ([]byte, []int) {
	return file_proto_teamjoinrequest_proto_rawDescGZIP(), []int{2}
}

func (x *TeamJoinRequestList) GetJoinRequests() []*TeamJoinRequest {
	if x != nil {
		return x.JoinRequests
	}
	return nil
}

type TeamJoinRequestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeamJoinRequestId) Reset() {
	*x = TeamJoinRequestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teamjoinrequest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamJoinRequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamJoinRequestId) ProtoMessage() {}

func (x *TeamJoinRequestId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teamjoinrequest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamJoinRequestId.ProtoReflect.Descriptor instead.
func (*TeamJoinRequestId) Descriptor() ([]byte, []int) {
	return file_proto_teamjoinrequest_proto_rawDescGZIP(), []int{3}
}

func (x *TeamJoinRequestId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_teamjoinrequest_proto protoreflect.FileDescriptor

var file_proto_teamjoinrequest_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x6a, 0x6f, 0x69, 0x6e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a,
	0x14, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x59, 0x0a, 0x0f, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x4c, 0x0a, 0x13, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0c, 0x6a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x23,
	0x0a, 0x11, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xbc, 0x03, 0x0a, 0x16, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x12, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x12, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e,
	0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x1a, 0x10, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x6a, 0x6f, 0x69,
	0x6e, 0x2d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_teamjoinrequest_proto_rawDescOnce sync.Once
	file_proto_teamjoinrequest_proto_rawDescData = file_proto_teamjoinrequest_proto_rawDesc
)

func file_proto_teamjoinrequest_proto_rawDescGZIP() []byte {
	file_proto_teamjoinrequest_proto_rawDescOnce.Do(func() {
		file_proto_teamjoinrequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_teamjoinrequest_proto_rawDescData)
	})
	return file_proto_teamjoinrequest_proto_rawDescData
}

var file_proto_teamjoinrequest_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_teamjoinrequest_proto_goTypes = []interface{}{
	(*TeamJoinReqeustEmpty)(nil), // 0: TeamJoinReqeustEmpty
	(*TeamJoinRequest)(nil),      // 1: TeamJoinRequest
	(*TeamJoinRequestList)(nil),  // 2: TeamJoinRequestList
	(*TeamJoinRequestId)(nil),    // 3: TeamJoinRequestId
}
var file_proto_teamjoinrequest_proto_depIdxs = []int32{
	1, // 0: TeamJoinRequestList.join_requests:type_name -> TeamJoinRequest
	0, // 1: TeamJoinRequestService.GetAllJoinRequests:input_type -> TeamJoinReqeustEmpty
	3, // 2: TeamJoinRequestService.GetJoinRequestById:input_type -> TeamJoinRequestId
	1, // 3: TeamJoinRequestService.CreateJoinRequest:input_type -> TeamJoinRequest
	1, // 4: TeamJoinRequestService.UpdateJoinRequest:input_type -> TeamJoinRequest
	3, // 5: TeamJoinRequestService.DeleteJoinRequest:input_type -> TeamJoinRequestId
	3, // 6: TeamJoinRequestService.ApproveJoinRequest:input_type -> TeamJoinRequestId
	3, // 7: TeamJoinRequestService.DeclineJoinRequest:input_type -> TeamJoinRequestId
	2, // 8: TeamJoinRequestService.GetAllJoinRequests:output_type -> TeamJoinRequestList
	1, // 9: TeamJoinRequestService.GetJoinRequestById:output_type -> TeamJoinRequest
	1, // 10: TeamJoinRequestService.CreateJoinRequest:output_type -> TeamJoinRequest
	1, // 11: TeamJoinRequestService.UpdateJoinRequest:output_type -> TeamJoinRequest
	1, // 12: TeamJoinRequestService.DeleteJoinRequest:output_type -> TeamJoinRequest
	1, // 13: TeamJoinRequestService.ApproveJoinRequest:output_type -> TeamJoinRequest
	1, // 14: TeamJoinRequestService.DeclineJoinRequest:output_type -> TeamJoinRequest
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_teamjoinrequest_proto_init() }
func file_proto_teamjoinrequest_proto_init() {
	if File_proto_teamjoinrequest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_teamjoinrequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamJoinReqeustEmpty); i {
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
		file_proto_teamjoinrequest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamJoinRequest); i {
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
		file_proto_teamjoinrequest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamJoinRequestList); i {
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
		file_proto_teamjoinrequest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamJoinRequestId); i {
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
			RawDescriptor: file_proto_teamjoinrequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_teamjoinrequest_proto_goTypes,
		DependencyIndexes: file_proto_teamjoinrequest_proto_depIdxs,
		MessageInfos:      file_proto_teamjoinrequest_proto_msgTypes,
	}.Build()
	File_proto_teamjoinrequest_proto = out.File
	file_proto_teamjoinrequest_proto_rawDesc = nil
	file_proto_teamjoinrequest_proto_goTypes = nil
	file_proto_teamjoinrequest_proto_depIdxs = nil
}