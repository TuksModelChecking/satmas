// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: experiment/experimentStateController.proto

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

type SaveExperimentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experiment *Experiment `protobuf:"bytes,1,opt,name=experiment,proto3" json:"experiment,omitempty"`
}

func (x *SaveExperimentRequest) Reset() {
	*x = SaveExperimentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveExperimentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveExperimentRequest) ProtoMessage() {}

func (x *SaveExperimentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveExperimentRequest.ProtoReflect.Descriptor instead.
func (*SaveExperimentRequest) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{0}
}

func (x *SaveExperimentRequest) GetExperiment() *Experiment {
	if x != nil {
		return x.Experiment
	}
	return nil
}

type SaveExperimentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experiment *Experiment `protobuf:"bytes,1,opt,name=experiment,proto3" json:"experiment,omitempty"`
}

func (x *SaveExperimentResponse) Reset() {
	*x = SaveExperimentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveExperimentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveExperimentResponse) ProtoMessage() {}

func (x *SaveExperimentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveExperimentResponse.ProtoReflect.Descriptor instead.
func (*SaveExperimentResponse) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{1}
}

func (x *SaveExperimentResponse) GetExperiment() *Experiment {
	if x != nil {
		return x.Experiment
	}
	return nil
}

type RunExperimentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experiment *Experiment `protobuf:"bytes,1,opt,name=experiment,proto3" json:"experiment,omitempty"`
}

func (x *RunExperimentRequest) Reset() {
	*x = RunExperimentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunExperimentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunExperimentRequest) ProtoMessage() {}

func (x *RunExperimentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunExperimentRequest.ProtoReflect.Descriptor instead.
func (*RunExperimentRequest) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{2}
}

func (x *RunExperimentRequest) GetExperiment() *Experiment {
	if x != nil {
		return x.Experiment
	}
	return nil
}

type RunExperimentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RunExperimentResponse) Reset() {
	*x = RunExperimentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunExperimentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunExperimentResponse) ProtoMessage() {}

func (x *RunExperimentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunExperimentResponse.ProtoReflect.Descriptor instead.
func (*RunExperimentResponse) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{3}
}

type MarkExperimentFailedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MarkExperimentFailedRequest) Reset() {
	*x = MarkExperimentFailedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkExperimentFailedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkExperimentFailedRequest) ProtoMessage() {}

func (x *MarkExperimentFailedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkExperimentFailedRequest.ProtoReflect.Descriptor instead.
func (*MarkExperimentFailedRequest) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{4}
}

func (x *MarkExperimentFailedRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MarkExperimentFailedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkExperimentFailedResponse) Reset() {
	*x = MarkExperimentFailedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkExperimentFailedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkExperimentFailedResponse) ProtoMessage() {}

func (x *MarkExperimentFailedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkExperimentFailedResponse.ProtoReflect.Descriptor instead.
func (*MarkExperimentFailedResponse) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{5}
}

type ResourceStateAtTimestemp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceIDs    []string `protobuf:"bytes,1,rep,name=resourceIDs,proto3" json:"resourceIDs,omitempty"`
	ResourceStates []int32  `protobuf:"varint,2,rep,packed,name=resourceStates,proto3" json:"resourceStates,omitempty"`
}

func (x *ResourceStateAtTimestemp) Reset() {
	*x = ResourceStateAtTimestemp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceStateAtTimestemp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceStateAtTimestemp) ProtoMessage() {}

func (x *ResourceStateAtTimestemp) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceStateAtTimestemp.ProtoReflect.Descriptor instead.
func (*ResourceStateAtTimestemp) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{6}
}

func (x *ResourceStateAtTimestemp) GetResourceIDs() []string {
	if x != nil {
		return x.ResourceIDs
	}
	return nil
}

func (x *ResourceStateAtTimestemp) GetResourceStates() []int32 {
	if x != nil {
		return x.ResourceStates
	}
	return nil
}

type ActionListAtTimestep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentIDs []int32  `protobuf:"varint,1,rep,packed,name=agentIDs,proto3" json:"agentIDs,omitempty"`
	Actions  []string `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *ActionListAtTimestep) Reset() {
	*x = ActionListAtTimestep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionListAtTimestep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionListAtTimestep) ProtoMessage() {}

func (x *ActionListAtTimestep) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionListAtTimestep.ProtoReflect.Descriptor instead.
func (*ActionListAtTimestep) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{7}
}

func (x *ActionListAtTimestep) GetAgentIDs() []int32 {
	if x != nil {
		return x.AgentIDs
	}
	return nil
}

func (x *ActionListAtTimestep) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type ExperimentResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceStates []*ResourceStateAtTimestemp `protobuf:"bytes,1,rep,name=resourceStates,proto3" json:"resourceStates,omitempty"`
	ActionList     []*ActionListAtTimestep     `protobuf:"bytes,2,rep,name=actionList,proto3" json:"actionList,omitempty"`
}

func (x *ExperimentResult) Reset() {
	*x = ExperimentResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentResult) ProtoMessage() {}

func (x *ExperimentResult) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentResult.ProtoReflect.Descriptor instead.
func (*ExperimentResult) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{8}
}

func (x *ExperimentResult) GetResourceStates() []*ResourceStateAtTimestemp {
	if x != nil {
		return x.ResourceStates
	}
	return nil
}

func (x *ExperimentResult) GetActionList() []*ActionListAtTimestep {
	if x != nil {
		return x.ActionList
	}
	return nil
}

type MarkExperimentSuccessfulRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Result *ExperimentResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MarkExperimentSuccessfulRequest) Reset() {
	*x = MarkExperimentSuccessfulRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkExperimentSuccessfulRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkExperimentSuccessfulRequest) ProtoMessage() {}

func (x *MarkExperimentSuccessfulRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkExperimentSuccessfulRequest.ProtoReflect.Descriptor instead.
func (*MarkExperimentSuccessfulRequest) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{9}
}

func (x *MarkExperimentSuccessfulRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MarkExperimentSuccessfulRequest) GetResult() *ExperimentResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type MarkExperimentSuccessfulResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkExperimentSuccessfulResponse) Reset() {
	*x = MarkExperimentSuccessfulResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_experimentStateController_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkExperimentSuccessfulResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkExperimentSuccessfulResponse) ProtoMessage() {}

func (x *MarkExperimentSuccessfulResponse) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_experimentStateController_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkExperimentSuccessfulResponse.ProtoReflect.Descriptor instead.
func (*MarkExperimentSuccessfulResponse) Descriptor() ([]byte, []int) {
	return file_experiment_experimentStateController_proto_rawDescGZIP(), []int{10}
}

var File_experiment_experimentStateController_proto protoreflect.FileDescriptor

var file_experiment_experimentStateController_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x52, 0x75, 0x6e, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x75, 0x6e, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2d, 0x0a, 0x1b, 0x4d, 0x61, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1e, 0x0a, 0x1c, 0x4d, 0x61, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x64, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x44, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x70, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4c, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x65, 0x6d, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x70, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x1f, 0x4d, 0x61, 0x72,
	0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x22, 0x0a, 0x20, 0x4d, 0x61, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb4, 0x03, 0x0a, 0x19, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x56, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x75,
	0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x75, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12,
	0x27, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x18, 0x4d, 0x61, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x12, 0x2b, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61,
	0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x75, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_experiment_experimentStateController_proto_rawDescOnce sync.Once
	file_experiment_experimentStateController_proto_rawDescData = file_experiment_experimentStateController_proto_rawDesc
)

func file_experiment_experimentStateController_proto_rawDescGZIP() []byte {
	file_experiment_experimentStateController_proto_rawDescOnce.Do(func() {
		file_experiment_experimentStateController_proto_rawDescData = protoimpl.X.CompressGZIP(file_experiment_experimentStateController_proto_rawDescData)
	})
	return file_experiment_experimentStateController_proto_rawDescData
}

var file_experiment_experimentStateController_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_experiment_experimentStateController_proto_goTypes = []any{
	(*SaveExperimentRequest)(nil),            // 0: experiment.SaveExperimentRequest
	(*SaveExperimentResponse)(nil),           // 1: experiment.SaveExperimentResponse
	(*RunExperimentRequest)(nil),             // 2: experiment.RunExperimentRequest
	(*RunExperimentResponse)(nil),            // 3: experiment.RunExperimentResponse
	(*MarkExperimentFailedRequest)(nil),      // 4: experiment.MarkExperimentFailedRequest
	(*MarkExperimentFailedResponse)(nil),     // 5: experiment.MarkExperimentFailedResponse
	(*ResourceStateAtTimestemp)(nil),         // 6: experiment.ResourceStateAtTimestemp
	(*ActionListAtTimestep)(nil),             // 7: experiment.ActionListAtTimestep
	(*ExperimentResult)(nil),                 // 8: experiment.ExperimentResult
	(*MarkExperimentSuccessfulRequest)(nil),  // 9: experiment.MarkExperimentSuccessfulRequest
	(*MarkExperimentSuccessfulResponse)(nil), // 10: experiment.MarkExperimentSuccessfulResponse
	(*Experiment)(nil),                       // 11: experiment.Experiment
}
var file_experiment_experimentStateController_proto_depIdxs = []int32{
	11, // 0: experiment.SaveExperimentRequest.experiment:type_name -> experiment.Experiment
	11, // 1: experiment.SaveExperimentResponse.experiment:type_name -> experiment.Experiment
	11, // 2: experiment.RunExperimentRequest.experiment:type_name -> experiment.Experiment
	6,  // 3: experiment.ExperimentResult.resourceStates:type_name -> experiment.ResourceStateAtTimestemp
	7,  // 4: experiment.ExperimentResult.actionList:type_name -> experiment.ActionListAtTimestep
	8,  // 5: experiment.MarkExperimentSuccessfulRequest.result:type_name -> experiment.ExperimentResult
	0,  // 6: experiment.ExperimentStateController.SaveExperiment:input_type -> experiment.SaveExperimentRequest
	2,  // 7: experiment.ExperimentStateController.RunExperiment:input_type -> experiment.RunExperimentRequest
	4,  // 8: experiment.ExperimentStateController.MarkExperimentFailed:input_type -> experiment.MarkExperimentFailedRequest
	9,  // 9: experiment.ExperimentStateController.MarkExperimentSuccessful:input_type -> experiment.MarkExperimentSuccessfulRequest
	1,  // 10: experiment.ExperimentStateController.SaveExperiment:output_type -> experiment.SaveExperimentResponse
	3,  // 11: experiment.ExperimentStateController.RunExperiment:output_type -> experiment.RunExperimentResponse
	5,  // 12: experiment.ExperimentStateController.MarkExperimentFailed:output_type -> experiment.MarkExperimentFailedResponse
	10, // 13: experiment.ExperimentStateController.MarkExperimentSuccessful:output_type -> experiment.MarkExperimentSuccessfulResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_experiment_experimentStateController_proto_init() }
func file_experiment_experimentStateController_proto_init() {
	if File_experiment_experimentStateController_proto != nil {
		return
	}
	file_experiment_experiment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_experiment_experimentStateController_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SaveExperimentRequest); i {
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
		file_experiment_experimentStateController_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SaveExperimentResponse); i {
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
		file_experiment_experimentStateController_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RunExperimentRequest); i {
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
		file_experiment_experimentStateController_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RunExperimentResponse); i {
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
		file_experiment_experimentStateController_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MarkExperimentFailedRequest); i {
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
		file_experiment_experimentStateController_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MarkExperimentFailedResponse); i {
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
		file_experiment_experimentStateController_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceStateAtTimestemp); i {
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
		file_experiment_experimentStateController_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ActionListAtTimestep); i {
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
		file_experiment_experimentStateController_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ExperimentResult); i {
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
		file_experiment_experimentStateController_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*MarkExperimentSuccessfulRequest); i {
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
		file_experiment_experimentStateController_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*MarkExperimentSuccessfulResponse); i {
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
			RawDescriptor: file_experiment_experimentStateController_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_experiment_experimentStateController_proto_goTypes,
		DependencyIndexes: file_experiment_experimentStateController_proto_depIdxs,
		MessageInfos:      file_experiment_experimentStateController_proto_msgTypes,
	}.Build()
	File_experiment_experimentStateController_proto = out.File
	file_experiment_experimentStateController_proto_rawDesc = nil
	file_experiment_experimentStateController_proto_goTypes = nil
	file_experiment_experimentStateController_proto_depIdxs = nil
}
