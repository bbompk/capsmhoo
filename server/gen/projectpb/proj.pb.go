// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/proj.proto

package projectpb

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

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId   string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	TeamId      string `protobuf:"bytes,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	ProfessorId string `protobuf:"bytes,3,opt,name=professor_id,json=professorId,proto3" json:"professor_id,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Status      string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Label       string `protobuf:"bytes,7,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_proto_proj_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Project) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Project) GetProfessorId() string {
	if x != nil {
		return x.ProfessorId
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Project) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type ProjectList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ProjectList) Reset() {
	*x = ProjectList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectList) ProtoMessage() {}

func (x *ProjectList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectList.ProtoReflect.Descriptor instead.
func (*ProjectList) Descriptor() ([]byte, []int) {
	return file_proto_proj_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectList) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ProjectId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ProjectId) Reset() {
	*x = ProjectId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectId) ProtoMessage() {}

func (x *ProjectId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectId.ProtoReflect.Descriptor instead.
func (*ProjectId) Descriptor() ([]byte, []int) {
	return file_proto_proj_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectId) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type TeamId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId string `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *TeamId) Reset() {
	*x = TeamId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamId) ProtoMessage() {}

func (x *TeamId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[3]
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
	return file_proto_proj_proto_rawDescGZIP(), []int{3}
}

func (x *TeamId) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

type ProfessorId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfessorId string `protobuf:"bytes,1,opt,name=professor_id,json=professorId,proto3" json:"professor_id,omitempty"`
}

func (x *ProfessorId) Reset() {
	*x = ProfessorId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfessorId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfessorId) ProtoMessage() {}

func (x *ProfessorId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfessorId.ProtoReflect.Descriptor instead.
func (*ProfessorId) Descriptor() ([]byte, []int) {
	return file_proto_proj_proto_rawDescGZIP(), []int{4}
}

func (x *ProfessorId) GetProfessorId() string {
	if x != nil {
		return x.ProfessorId
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[5]
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
	return file_proto_proj_proto_rawDescGZIP(), []int{5}
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_proj_proto_rawDescGZIP(), []int{6}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectRequestId string `protobuf:"bytes,1,opt,name=project_request_id,json=projectRequestId,proto3" json:"project_request_id,omitempty"`
	ProjectId        string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	TeamId           string `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Message          string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Status           string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ProjectRequest) Reset() {
	*x = ProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectRequest) ProtoMessage() {}

func (x *ProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectRequest.ProtoReflect.Descriptor instead.
func (*ProjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_proj_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectRequest) GetProjectRequestId() string {
	if x != nil {
		return x.ProjectRequestId
	}
	return ""
}

func (x *ProjectRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *ProjectRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ProjectRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ProjectRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectRequests []*ProjectRequest `protobuf:"bytes,1,rep,name=project_requests,json=projectRequests,proto3" json:"project_requests,omitempty"`
}

func (x *ProjectRequestList) Reset() {
	*x = ProjectRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proj_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectRequestList) ProtoMessage() {}

func (x *ProjectRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proj_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectRequestList.ProtoReflect.Descriptor instead.
func (*ProjectRequestList) Descriptor() ([]byte, []int) {
	return file_proto_proj_proto_rawDescGZIP(), []int{8}
}

func (x *ProjectRequestList) GetProjectRequests() []*ProjectRequest {
	if x != nil {
		return x.ProjectRequests
	}
	return nil
}

var File_proto_proj_proto protoreflect.FileDescriptor

var file_proto_proj_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x22, 0xc8, 0x01,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3d, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x2a, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x06, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xa8,
	0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5a, 0x0a, 0x12, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x44, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x32, 0x95, 0x06, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x14, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x14,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proj_proto_rawDescOnce sync.Once
	file_proto_proj_proto_rawDescData = file_proto_proj_proto_rawDesc
)

func file_proto_proj_proto_rawDescGZIP() []byte {
	file_proto_proj_proto_rawDescOnce.Do(func() {
		file_proto_proj_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proj_proto_rawDescData)
	})
	return file_proto_proj_proto_rawDescData
}

var file_proto_proj_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_proj_proto_goTypes = []interface{}{
	(*Project)(nil),            // 0: projectpb.Project
	(*ProjectList)(nil),        // 1: projectpb.ProjectList
	(*ProjectId)(nil),          // 2: projectpb.ProjectId
	(*TeamId)(nil),             // 3: projectpb.TeamId
	(*ProfessorId)(nil),        // 4: projectpb.ProfessorId
	(*Empty)(nil),              // 5: projectpb.Empty
	(*SuccessResponse)(nil),    // 6: projectpb.SuccessResponse
	(*ProjectRequest)(nil),     // 7: projectpb.ProjectRequest
	(*ProjectRequestList)(nil), // 8: projectpb.ProjectRequestList
}
var file_proto_proj_proto_depIdxs = []int32{
	0,  // 0: projectpb.ProjectList.projects:type_name -> projectpb.Project
	7,  // 1: projectpb.ProjectRequestList.project_requests:type_name -> projectpb.ProjectRequest
	2,  // 2: projectpb.ProjectService.GetProjectById:input_type -> projectpb.ProjectId
	3,  // 3: projectpb.ProjectService.GetProjectByTeamId:input_type -> projectpb.TeamId
	4,  // 4: projectpb.ProjectService.GetProjectByProfessorId:input_type -> projectpb.ProfessorId
	5,  // 5: projectpb.ProjectService.GetAllProjects:input_type -> projectpb.Empty
	0,  // 6: projectpb.ProjectService.CreateProject:input_type -> projectpb.Project
	0,  // 7: projectpb.ProjectService.UpdateProject:input_type -> projectpb.Project
	2,  // 8: projectpb.ProjectService.DeleteProject:input_type -> projectpb.ProjectId
	2,  // 9: projectpb.ProjectService.GetProjectRequestsByProjectId:input_type -> projectpb.ProjectId
	7,  // 10: projectpb.ProjectService.CreateProjectRequest:input_type -> projectpb.ProjectRequest
	7,  // 11: projectpb.ProjectService.AcceptProjectRequest:input_type -> projectpb.ProjectRequest
	7,  // 12: projectpb.ProjectService.RejectProjectRequest:input_type -> projectpb.ProjectRequest
	0,  // 13: projectpb.ProjectService.GetProjectById:output_type -> projectpb.Project
	0,  // 14: projectpb.ProjectService.GetProjectByTeamId:output_type -> projectpb.Project
	1,  // 15: projectpb.ProjectService.GetProjectByProfessorId:output_type -> projectpb.ProjectList
	1,  // 16: projectpb.ProjectService.GetAllProjects:output_type -> projectpb.ProjectList
	0,  // 17: projectpb.ProjectService.CreateProject:output_type -> projectpb.Project
	0,  // 18: projectpb.ProjectService.UpdateProject:output_type -> projectpb.Project
	0,  // 19: projectpb.ProjectService.DeleteProject:output_type -> projectpb.Project
	8,  // 20: projectpb.ProjectService.GetProjectRequestsByProjectId:output_type -> projectpb.ProjectRequestList
	7,  // 21: projectpb.ProjectService.CreateProjectRequest:output_type -> projectpb.ProjectRequest
	6,  // 22: projectpb.ProjectService.AcceptProjectRequest:output_type -> projectpb.SuccessResponse
	6,  // 23: projectpb.ProjectService.RejectProjectRequest:output_type -> projectpb.SuccessResponse
	13, // [13:24] is the sub-list for method output_type
	2,  // [2:13] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_proj_proto_init() }
func file_proto_proj_proto_init() {
	if File_proto_proj_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proj_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_proto_proj_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectList); i {
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
		file_proto_proj_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectId); i {
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
		file_proto_proj_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_proj_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfessorId); i {
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
		file_proto_proj_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_proj_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_proto_proj_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectRequest); i {
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
		file_proto_proj_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectRequestList); i {
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
			RawDescriptor: file_proto_proj_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proj_proto_goTypes,
		DependencyIndexes: file_proto_proj_proto_depIdxs,
		MessageInfos:      file_proto_proj_proto_msgTypes,
	}.Build()
	File_proto_proj_proto = out.File
	file_proto_proj_proto_rawDesc = nil
	file_proto_proj_proto_goTypes = nil
	file_proto_proj_proto_depIdxs = nil
}
