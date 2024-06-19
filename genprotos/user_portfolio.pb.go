// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: user_portfolio.proto

package genprotos

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

type UserPortfolio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects    []*Project    `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	Skills      []*Skill      `protobuf:"bytes,2,rep,name=skills,proto3" json:"skills,omitempty"`
	Experiences []*Experience `protobuf:"bytes,3,rep,name=experiences,proto3" json:"experiences,omitempty"`
	Educations  []*Education  `protobuf:"bytes,4,rep,name=educations,proto3" json:"educations,omitempty"`
}

func (x *UserPortfolio) Reset() {
	*x = UserPortfolio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_portfolio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPortfolio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPortfolio) ProtoMessage() {}

func (x *UserPortfolio) ProtoReflect() protoreflect.Message {
	mi := &file_user_portfolio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPortfolio.ProtoReflect.Descriptor instead.
func (*UserPortfolio) Descriptor() ([]byte, []int) {
	return file_user_portfolio_proto_rawDescGZIP(), []int{0}
}

func (x *UserPortfolio) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *UserPortfolio) GetSkills() []*Skill {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *UserPortfolio) GetExperiences() []*Experience {
	if x != nil {
		return x.Experiences
	}
	return nil
}

func (x *UserPortfolio) GetEducations() []*Education {
	if x != nil {
		return x.Educations
	}
	return nil
}

var File_user_portfolio_proto protoreflect.FileDescriptor

var file_user_portfolio_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x10, 0x65, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80,
	0x02, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x12, 0x38, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x73,
	0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x41,
	0x0a, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f,
	0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_portfolio_proto_rawDescOnce sync.Once
	file_user_portfolio_proto_rawDescData = file_user_portfolio_proto_rawDesc
)

func file_user_portfolio_proto_rawDescGZIP() []byte {
	file_user_portfolio_proto_rawDescOnce.Do(func() {
		file_user_portfolio_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_portfolio_proto_rawDescData)
	})
	return file_user_portfolio_proto_rawDescData
}

var file_user_portfolio_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_portfolio_proto_goTypes = []interface{}{
	(*UserPortfolio)(nil), // 0: portfolio_submodule.UserPortfolio
	(*Project)(nil),       // 1: portfolio_submodule.Project
	(*Skill)(nil),         // 2: portfolio_submodule.Skill
	(*Experience)(nil),    // 3: portfolio_submodule.Experience
	(*Education)(nil),     // 4: portfolio_submodule.Education
}
var file_user_portfolio_proto_depIdxs = []int32{
	1, // 0: portfolio_submodule.UserPortfolio.projects:type_name -> portfolio_submodule.Project
	2, // 1: portfolio_submodule.UserPortfolio.skills:type_name -> portfolio_submodule.Skill
	3, // 2: portfolio_submodule.UserPortfolio.experiences:type_name -> portfolio_submodule.Experience
	4, // 3: portfolio_submodule.UserPortfolio.educations:type_name -> portfolio_submodule.Education
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_portfolio_proto_init() }
func file_user_portfolio_proto_init() {
	if File_user_portfolio_proto != nil {
		return
	}
	file_educations_proto_init()
	file_experiences_proto_init()
	file_projects_proto_init()
	file_skills_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_portfolio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPortfolio); i {
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
			RawDescriptor: file_user_portfolio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_portfolio_proto_goTypes,
		DependencyIndexes: file_user_portfolio_proto_depIdxs,
		MessageInfos:      file_user_portfolio_proto_msgTypes,
	}.Build()
	File_user_portfolio_proto = out.File
	file_user_portfolio_proto_rawDesc = nil
	file_user_portfolio_proto_goTypes = nil
	file_user_portfolio_proto_depIdxs = nil
}
