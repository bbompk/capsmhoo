// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: team.proto

package proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{0}
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Profile   string `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	ProjectId string `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{1}
}

func (x *Team) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *Team) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type TeamList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*Team `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *TeamList) Reset() {
	*x = TeamList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamList) ProtoMessage() {}

func (x *TeamList) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamList.ProtoReflect.Descriptor instead.
func (*TeamList) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{2}
}

func (x *TeamList) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

type StudentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StudentId) Reset() {
	*x = StudentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentId) ProtoMessage() {}

func (x *StudentId) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentId.ProtoReflect.Descriptor instead.
func (*StudentId) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{3}
}

func (x *StudentId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TeamId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeamId) Reset() {
	*x = TeamId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamId) ProtoMessage() {}

func (x *TeamId) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamId.ProtoReflect.Descriptor instead.
func (*TeamId) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{4}
}

func (x *TeamId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TeamAndStudentID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId    string `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	StudentId string `protobuf:"bytes,2,opt,name=studentId,proto3" json:"studentId,omitempty"`
}

func (x *TeamAndStudentID) Reset() {
	*x = TeamAndStudentID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamAndStudentID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamAndStudentID) ProtoMessage() {}

func (x *TeamAndStudentID) ProtoReflect() protoreflect.Message {
	mi := &file_team_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamAndStudentID.ProtoReflect.Descriptor instead.
func (*TeamAndStudentID) Descriptor() ([]byte, []int) {
	return file_team_proto_rawDescGZIP(), []int{5}
}

func (x *TeamAndStudentID) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *TeamAndStudentID) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

var File_team_proto protoreflect.FileDescriptor

var file_team_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x63, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x08, 0x54, 0x65,
	0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x18, 0x0a, 0x06, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x10, 0x54, 0x65,
	0x61, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x32, 0x87, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x42, 0x79, 0x49, 0x64, 0x12, 0x07, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x1a, 0x05,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x1a, 0x05, 0x2e, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x1a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12,
	0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x1a, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x07, 0x2e, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x1a, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x2d, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x12,
	0x11, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_team_proto_rawDescOnce sync.Once
	file_team_proto_rawDescData = file_team_proto_rawDesc
)

func file_team_proto_rawDescGZIP() []byte {
	file_team_proto_rawDescOnce.Do(func() {
		file_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_team_proto_rawDescData)
	})
	return file_team_proto_rawDescData
}

var file_team_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_team_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: Empty
	(*Team)(nil),             // 1: Team
	(*TeamList)(nil),         // 2: TeamList
	(*StudentId)(nil),        // 3: StudentId
	(*TeamId)(nil),           // 4: TeamId
	(*TeamAndStudentID)(nil), // 5: TeamAndStudentID
}
var file_team_proto_depIdxs = []int32{
	1, // 0: TeamList.teams:type_name -> Team
	0, // 1: TeamService.GetAllTeams:input_type -> Empty
	4, // 2: TeamService.GetTeamById:input_type -> TeamId
	1, // 3: TeamService.CreateTeam:input_type -> Team
	1, // 4: TeamService.UpdateTeam:input_type -> Team
	4, // 5: TeamService.DeleteTeam:input_type -> TeamId
	5, // 6: TeamService.AddStudentToTeam:input_type -> TeamAndStudentID
	5, // 7: TeamService.RemoveStudentFromTeam:input_type -> TeamAndStudentID
	2, // 8: TeamService.GetAllTeams:output_type -> TeamList
	1, // 9: TeamService.GetTeamById:output_type -> Team
	1, // 10: TeamService.CreateTeam:output_type -> Team
	1, // 11: TeamService.UpdateTeam:output_type -> Team
	1, // 12: TeamService.DeleteTeam:output_type -> Team
	0, // 13: TeamService.AddStudentToTeam:output_type -> Empty
	0, // 14: TeamService.RemoveStudentFromTeam:output_type -> Empty
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_team_proto_init() }
func file_team_proto_init() {
	if File_team_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamList); i {
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
		file_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentId); i {
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
		file_team_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamId); i {
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
		file_team_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamAndStudentID); i {
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
			RawDescriptor: file_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_team_proto_goTypes,
		DependencyIndexes: file_team_proto_depIdxs,
		MessageInfos:      file_team_proto_msgTypes,
	}.Build()
	File_team_proto = out.File
	file_team_proto_rawDesc = nil
	file_team_proto_goTypes = nil
	file_team_proto_depIdxs = nil
}
