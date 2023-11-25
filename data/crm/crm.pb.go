// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: crm.proto

package crm

import (
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

type OpenActivitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferenceType string `protobuf:"bytes,1,opt,name=reference_type,json=referenceType,proto3" json:"reference_type,omitempty"`
	ReferenceName string `protobuf:"bytes,2,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
}

func (x *OpenActivitiesRequest) Reset() {
	*x = OpenActivitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenActivitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenActivitiesRequest) ProtoMessage() {}

func (x *OpenActivitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenActivitiesRequest.ProtoReflect.Descriptor instead.
func (*OpenActivitiesRequest) Descriptor() ([]byte, []int) {
	return file_crm_proto_rawDescGZIP(), []int{0}
}

func (x *OpenActivitiesRequest) GetReferenceType() string {
	if x != nil {
		return x.ReferenceType
	}
	return ""
}

func (x *OpenActivitiesRequest) GetReferenceName() string {
	if x != nil {
		return x.ReferenceName
	}
	return ""
}

type OpenActivitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
	Tasks  []*ToDo  `protobuf:"bytes,2,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
}

func (x *OpenActivitiesResponse) Reset() {
	*x = OpenActivitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenActivitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenActivitiesResponse) ProtoMessage() {}

func (x *OpenActivitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenActivitiesResponse.ProtoReflect.Descriptor instead.
func (*OpenActivitiesResponse) Descriptor() ([]byte, []int) {
	return file_crm_proto_rawDescGZIP(), []int{1}
}

func (x *OpenActivitiesResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *OpenActivitiesResponse) GetTasks() []*ToDo {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject       string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	EventCategory string                 `protobuf:"bytes,3,opt,name=event_category,json=eventCategory,proto3" json:"event_category,omitempty"`
	StartsOn      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=starts_on,json=startsOn,proto3" json:"starts_on,omitempty"`
	EndsOn        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ends_on,json=endsOn,proto3" json:"ends_on,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_crm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_crm_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Event) GetEventCategory() string {
	if x != nil {
		return x.EventCategory
	}
	return ""
}

func (x *Event) GetStartsOn() *timestamppb.Timestamp {
	if x != nil {
		return x.StartsOn
	}
	return nil
}

func (x *Event) GetEndsOn() *timestamppb.Timestamp {
	if x != nil {
		return x.EndsOn
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AllocatedTo string                 `protobuf:"bytes,3,opt,name=allocated_to,json=allocatedTo,proto3" json:"allocated_to,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_crm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_crm_proto_rawDescGZIP(), []int{3}
}

func (x *ToDo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToDo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo) GetAllocatedTo() string {
	if x != nil {
		return x.AllocatedTo
	}
	return ""
}

func (x *ToDo) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_crm_proto protoreflect.FileDescriptor

var file_crm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x72, 0x6d,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x65, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x72, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x72, 0x6d, 0x2e, 0x54, 0x6f, 0x44, 0x6f,
	0x52, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x4f, 0x6e, 0x12, 0x33, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x32, 0x5c, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4e,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x63, 0x72, 0x6d, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x63, 0x72, 0x6d, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x63, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crm_proto_rawDescOnce sync.Once
	file_crm_proto_rawDescData = file_crm_proto_rawDesc
)

func file_crm_proto_rawDescGZIP() []byte {
	file_crm_proto_rawDescOnce.Do(func() {
		file_crm_proto_rawDescData = protoimpl.X.CompressGZIP(file_crm_proto_rawDescData)
	})
	return file_crm_proto_rawDescData
}

var file_crm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crm_proto_goTypes = []interface{}{
	(*OpenActivitiesRequest)(nil),  // 0: crm.OpenActivitiesRequest
	(*OpenActivitiesResponse)(nil), // 1: crm.OpenActivitiesResponse
	(*Event)(nil),                  // 2: crm.Event
	(*ToDo)(nil),                   // 3: crm.ToDo
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_crm_proto_depIdxs = []int32{
	2, // 0: crm.OpenActivitiesResponse.Events:type_name -> crm.Event
	3, // 1: crm.OpenActivitiesResponse.Tasks:type_name -> crm.ToDo
	4, // 2: crm.Event.starts_on:type_name -> google.protobuf.Timestamp
	4, // 3: crm.Event.ends_on:type_name -> google.protobuf.Timestamp
	4, // 4: crm.ToDo.date:type_name -> google.protobuf.Timestamp
	0, // 5: crm.Activities.GetOpenActivities:input_type -> crm.OpenActivitiesRequest
	1, // 6: crm.Activities.GetOpenActivities:output_type -> crm.OpenActivitiesResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_crm_proto_init() }
func file_crm_proto_init() {
	if File_crm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenActivitiesRequest); i {
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
		file_crm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenActivitiesResponse); i {
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
		file_crm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_crm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToDo); i {
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
			RawDescriptor: file_crm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crm_proto_goTypes,
		DependencyIndexes: file_crm_proto_depIdxs,
		MessageInfos:      file_crm_proto_msgTypes,
	}.Build()
	File_crm_proto = out.File
	file_crm_proto_rawDesc = nil
	file_crm_proto_goTypes = nil
	file_crm_proto_depIdxs = nil
}
