// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: user.proto

package protogen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Gender int32

const (
	Gender_Gender_UNDEFINED Gender = 0
	Gender_Gender_MALE      Gender = 1
	Gender_Gender_FEMALE    Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "Gender_UNDEFINED",
		1: "Gender_MALE",
		2: "Gender_FEMALE",
	}
	Gender_value = map[string]int32{
		"Gender_UNDEFINED": 0,
		"Gender_MALE":      1,
		"Gender_FEMALE":    2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	IsActive    bool       `protobuf:"varint,3,opt,name=is_active,proto3" json:"is_active,omitempty"`
	Password    []byte     `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Emails      []string   `protobuf:"bytes,5,rep,name=emails,proto3" json:"emails,omitempty"`
	Gender      Gender     `protobuf:"varint,6,opt,name=gender,proto3,enum=protogen.Gender" json:"gender,omitempty"`
	Addresses   *Address   `protobuf:"bytes,7,opt,name=addresses,proto3" json:"addresses,omitempty"`
	SocialMedia *anypb.Any `protobuf:"bytes,18,opt,name=social_media,json=socialMedia,proto3" json:"social_media,omitempty"`
	// Types that are assignable to AllowedSocmed:
	//
	//	*User_Socmed
	//	*User_InstantMsg
	AllowedSocmed    isUser_AllowedSocmed   `protobuf_oneof:"allowed_socmed"`
	SkillUser        map[string]*Skill      `protobuf:"bytes,21,rep,name=skill_user,json=skillUser,proto3" json:"skill_user,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastLoginUser    *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=last_login_user,json=lastLoginUser,proto3" json:"last_login_user,omitempty"`
	BirthDate        *Date                  `protobuf:"bytes,23,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	LastKnowLocation *LatLng                `protobuf:"bytes,24,opt,name=last_know_location,json=lastKnowLocation,proto3" json:"last_know_location,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *User) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_Gender_UNDEFINED
}

func (x *User) GetAddresses() *Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *User) GetSocialMedia() *anypb.Any {
	if x != nil {
		return x.SocialMedia
	}
	return nil
}

func (m *User) GetAllowedSocmed() isUser_AllowedSocmed {
	if m != nil {
		return m.AllowedSocmed
	}
	return nil
}

func (x *User) GetSocmed() *SocialMedia {
	if x, ok := x.GetAllowedSocmed().(*User_Socmed); ok {
		return x.Socmed
	}
	return nil
}

func (x *User) GetInstantMsg() *InstantMessaging {
	if x, ok := x.GetAllowedSocmed().(*User_InstantMsg); ok {
		return x.InstantMsg
	}
	return nil
}

func (x *User) GetSkillUser() map[string]*Skill {
	if x != nil {
		return x.SkillUser
	}
	return nil
}

func (x *User) GetLastLoginUser() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLoginUser
	}
	return nil
}

func (x *User) GetBirthDate() *Date {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *User) GetLastKnowLocation() *LatLng {
	if x != nil {
		return x.LastKnowLocation
	}
	return nil
}

type isUser_AllowedSocmed interface {
	isUser_AllowedSocmed()
}

type User_Socmed struct {
	Socmed *SocialMedia `protobuf:"bytes,19,opt,name=socmed,proto3,oneof"`
}

type User_InstantMsg struct {
	InstantMsg *InstantMessaging `protobuf:"bytes,20,opt,name=instant_msg,json=instantMsg,proto3,oneof"`
}

func (*User_Socmed) isUser_AllowedSocmed() {}

func (*User_InstantMsg) isUser_AllowedSocmed() {}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameSkill   string `protobuf:"bytes,1,opt,name=name_skill,json=nameSkill,proto3" json:"name_skill,omitempty"`
	RatingSkill uint32 `protobuf:"varint,2,opt,name=rating_skill,json=ratingSkill,proto3" json:"rating_skill,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *Skill) GetNameSkill() string {
	if x != nil {
		return x.NameSkill
	}
	return ""
}

func (x *Skill) GetRatingSkill() uint32 {
	if x != nil {
		return x.RatingSkill
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street     string              `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City       string              `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country    string              `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Coordinate *Address_Coordinate `protobuf:"bytes,16,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetCoordinate() *Address_Coordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type PaperMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperMail string `protobuf:"bytes,1,opt,name=paper_mail,json=paperMail,proto3" json:"paper_mail,omitempty"`
}

func (x *PaperMail) Reset() {
	*x = PaperMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaperMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaperMail) ProtoMessage() {}

func (x *PaperMail) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaperMail.ProtoReflect.Descriptor instead.
func (*PaperMail) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *PaperMail) GetPaperMail() string {
	if x != nil {
		return x.PaperMail
	}
	return ""
}

type SocialMedia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocialMedia string `protobuf:"bytes,1,opt,name=social_media,json=socialMedia,proto3" json:"social_media,omitempty"`
}

func (x *SocialMedia) Reset() {
	*x = SocialMedia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialMedia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialMedia) ProtoMessage() {}

func (x *SocialMedia) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialMedia.ProtoReflect.Descriptor instead.
func (*SocialMedia) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *SocialMedia) GetSocialMedia() string {
	if x != nil {
		return x.SocialMedia
	}
	return ""
}

type InstantMessaging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstantMessaging string `protobuf:"bytes,1,opt,name=instant_messaging,json=instantMessaging,proto3" json:"instant_messaging,omitempty"`
}

func (x *InstantMessaging) Reset() {
	*x = InstantMessaging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantMessaging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantMessaging) ProtoMessage() {}

func (x *InstantMessaging) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantMessaging.ProtoReflect.Descriptor instead.
func (*InstantMessaging) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *InstantMessaging) GetInstantMessaging() string {
	if x != nil {
		return x.InstantMessaging
	}
	return ""
}

type Address_Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Address_Coordinate) Reset() {
	*x = Address_Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address_Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address_Coordinate) ProtoMessage() {}

func (x *Address_Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address_Coordinate.ProtoReflect.Descriptor instead.
func (*Address_Coordinate) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Address_Coordinate) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Address_Coordinate) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x05, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x0b, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x6f, 0x63, 0x6d, 0x65, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x48, 0x00, 0x52, 0x06, 0x73, 0x6f, 0x63, 0x6d, 0x65, 0x64,
	0x12, 0x3d, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12,
	0x3c, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x15, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x30, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b, 0x6e, 0x6f, 0x77,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61,
	0x74, 0x4c, 0x6e, 0x67, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x4b, 0x6e, 0x6f, 0x77, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x4d, 0x0a, 0x0e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x73, 0x6f, 0x63, 0x6d, 0x65, 0x64, 0x22, 0x49, 0x0a, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x22, 0xd5, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x1a, 0x46, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x2a, 0x0a, 0x09, 0x50, 0x61,
	0x70, 0x65, 0x72, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x4d, 0x61, 0x69, 0x6c, 0x22, 0x30, 0x0a, 0x0b, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x3f, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x11,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2a, 0x42, 0x0a, 0x06, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x55, 0x4e,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x0b, 0x5a,
	0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(Gender)(0),                   // 0: protogen.Gender
	(*User)(nil),                  // 1: protogen.User
	(*Skill)(nil),                 // 2: protogen.Skill
	(*Address)(nil),               // 3: protogen.Address
	(*PaperMail)(nil),             // 4: protogen.PaperMail
	(*SocialMedia)(nil),           // 5: protogen.SocialMedia
	(*InstantMessaging)(nil),      // 6: protogen.InstantMessaging
	nil,                           // 7: protogen.User.SkillUserEntry
	(*Address_Coordinate)(nil),    // 8: protogen.Address.Coordinate
	(*anypb.Any)(nil),             // 9: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*Date)(nil),                  // 11: google.type.Date
	(*LatLng)(nil),                // 12: google.type.LatLng
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: protogen.User.gender:type_name -> protogen.Gender
	3,  // 1: protogen.User.addresses:type_name -> protogen.Address
	9,  // 2: protogen.User.social_media:type_name -> google.protobuf.Any
	5,  // 3: protogen.User.socmed:type_name -> protogen.SocialMedia
	6,  // 4: protogen.User.instant_msg:type_name -> protogen.InstantMessaging
	7,  // 5: protogen.User.skill_user:type_name -> protogen.User.SkillUserEntry
	10, // 6: protogen.User.last_login_user:type_name -> google.protobuf.Timestamp
	11, // 7: protogen.User.birth_date:type_name -> google.type.Date
	12, // 8: protogen.User.last_know_location:type_name -> google.type.LatLng
	8,  // 9: protogen.Address.coordinate:type_name -> protogen.Address.Coordinate
	2,  // 10: protogen.User.SkillUserEntry.value:type_name -> protogen.Skill
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_google_type_date_proto_init()
	file_google_type_latlng_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaperMail); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialMedia); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantMessaging); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address_Coordinate); i {
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
	file_user_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*User_Socmed)(nil),
		(*User_InstantMsg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
