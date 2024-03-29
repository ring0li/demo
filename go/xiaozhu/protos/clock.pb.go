// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clock.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//早起
type GetUpRequest struct {
	OpenId               string   `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	ClockTime            string   `protobuf:"bytes,2,opt,name=clockTime,proto3" json:"clockTime,omitempty"`
	QrcodeUrl            string   `protobuf:"bytes,3,opt,name=qrcodeUrl,proto3" json:"qrcodeUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUpRequest) Reset()         { *m = GetUpRequest{} }
func (m *GetUpRequest) String() string { return proto.CompactTextString(m) }
func (*GetUpRequest) ProtoMessage()    {}
func (*GetUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{0}
}

func (m *GetUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpRequest.Unmarshal(m, b)
}
func (m *GetUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpRequest.Marshal(b, m, deterministic)
}
func (m *GetUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpRequest.Merge(m, src)
}
func (m *GetUpRequest) XXX_Size() int {
	return xxx_messageInfo_GetUpRequest.Size(m)
}
func (m *GetUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpRequest proto.InternalMessageInfo

func (m *GetUpRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *GetUpRequest) GetClockTime() string {
	if m != nil {
		return m.ClockTime
	}
	return ""
}

func (m *GetUpRequest) GetQrcodeUrl() string {
	if m != nil {
		return m.QrcodeUrl
	}
	return ""
}

type GetUpResponse struct {
	ImgUrl               string   `protobuf:"bytes,1,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	ContinueClockDays    int32    `protobuf:"varint,3,opt,name=continueClockDays,proto3" json:"continueClockDays,omitempty"`
	WaterText            string   `protobuf:"bytes,4,opt,name=waterText,proto3" json:"waterText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUpResponse) Reset()         { *m = GetUpResponse{} }
func (m *GetUpResponse) String() string { return proto.CompactTextString(m) }
func (*GetUpResponse) ProtoMessage()    {}
func (*GetUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{1}
}

func (m *GetUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUpResponse.Unmarshal(m, b)
}
func (m *GetUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUpResponse.Marshal(b, m, deterministic)
}
func (m *GetUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUpResponse.Merge(m, src)
}
func (m *GetUpResponse) XXX_Size() int {
	return xxx_messageInfo_GetUpResponse.Size(m)
}
func (m *GetUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUpResponse proto.InternalMessageInfo

func (m *GetUpResponse) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *GetUpResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *GetUpResponse) GetContinueClockDays() int32 {
	if m != nil {
		return m.ContinueClockDays
	}
	return 0
}

func (m *GetUpResponse) GetWaterText() string {
	if m != nil {
		return m.WaterText
	}
	return ""
}

//早睡
type SleepRequest struct {
	OpenId               string   `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	ClockTime            string   `protobuf:"bytes,2,opt,name=clockTime,proto3" json:"clockTime,omitempty"`
	QrcodeUrl            string   `protobuf:"bytes,3,opt,name=qrcodeUrl,proto3" json:"qrcodeUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SleepRequest) Reset()         { *m = SleepRequest{} }
func (m *SleepRequest) String() string { return proto.CompactTextString(m) }
func (*SleepRequest) ProtoMessage()    {}
func (*SleepRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{2}
}

func (m *SleepRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SleepRequest.Unmarshal(m, b)
}
func (m *SleepRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SleepRequest.Marshal(b, m, deterministic)
}
func (m *SleepRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SleepRequest.Merge(m, src)
}
func (m *SleepRequest) XXX_Size() int {
	return xxx_messageInfo_SleepRequest.Size(m)
}
func (m *SleepRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SleepRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SleepRequest proto.InternalMessageInfo

func (m *SleepRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *SleepRequest) GetClockTime() string {
	if m != nil {
		return m.ClockTime
	}
	return ""
}

func (m *SleepRequest) GetQrcodeUrl() string {
	if m != nil {
		return m.QrcodeUrl
	}
	return ""
}

type SleepResponse struct {
	ImgUrl               string   `protobuf:"bytes,1,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	ContinueClockDays    int32    `protobuf:"varint,3,opt,name=continueClockDays,proto3" json:"continueClockDays,omitempty"`
	WaterText            string   `protobuf:"bytes,4,opt,name=waterText,proto3" json:"waterText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SleepResponse) Reset()         { *m = SleepResponse{} }
func (m *SleepResponse) String() string { return proto.CompactTextString(m) }
func (*SleepResponse) ProtoMessage()    {}
func (*SleepResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{3}
}

func (m *SleepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SleepResponse.Unmarshal(m, b)
}
func (m *SleepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SleepResponse.Marshal(b, m, deterministic)
}
func (m *SleepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SleepResponse.Merge(m, src)
}
func (m *SleepResponse) XXX_Size() int {
	return xxx_messageInfo_SleepResponse.Size(m)
}
func (m *SleepResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SleepResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SleepResponse proto.InternalMessageInfo

func (m *SleepResponse) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *SleepResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *SleepResponse) GetContinueClockDays() int32 {
	if m != nil {
		return m.ContinueClockDays
	}
	return 0
}

func (m *SleepResponse) GetWaterText() string {
	if m != nil {
		return m.WaterText
	}
	return ""
}

//未打卡的用户列表
type NoClockUsersRequest struct {
	ClockType            int32    `protobuf:"varint,1,opt,name=clockType,proto3" json:"clockType,omitempty"`
	StartTime            string   `protobuf:"bytes,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              string   `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	LastUserId           int32    `protobuf:"varint,4,opt,name=lastUserId,proto3" json:"lastUserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoClockUsersRequest) Reset()         { *m = NoClockUsersRequest{} }
func (m *NoClockUsersRequest) String() string { return proto.CompactTextString(m) }
func (*NoClockUsersRequest) ProtoMessage()    {}
func (*NoClockUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{4}
}

func (m *NoClockUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoClockUsersRequest.Unmarshal(m, b)
}
func (m *NoClockUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoClockUsersRequest.Marshal(b, m, deterministic)
}
func (m *NoClockUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoClockUsersRequest.Merge(m, src)
}
func (m *NoClockUsersRequest) XXX_Size() int {
	return xxx_messageInfo_NoClockUsersRequest.Size(m)
}
func (m *NoClockUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NoClockUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NoClockUsersRequest proto.InternalMessageInfo

func (m *NoClockUsersRequest) GetClockType() int32 {
	if m != nil {
		return m.ClockType
	}
	return 0
}

func (m *NoClockUsersRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *NoClockUsersRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *NoClockUsersRequest) GetLastUserId() int32 {
	if m != nil {
		return m.LastUserId
	}
	return 0
}

type NoClockUsersResponse struct {
	List                 []*UserList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NoClockUsersResponse) Reset()         { *m = NoClockUsersResponse{} }
func (m *NoClockUsersResponse) String() string { return proto.CompactTextString(m) }
func (*NoClockUsersResponse) ProtoMessage()    {}
func (*NoClockUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{5}
}

func (m *NoClockUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoClockUsersResponse.Unmarshal(m, b)
}
func (m *NoClockUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoClockUsersResponse.Marshal(b, m, deterministic)
}
func (m *NoClockUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoClockUsersResponse.Merge(m, src)
}
func (m *NoClockUsersResponse) XXX_Size() int {
	return xxx_messageInfo_NoClockUsersResponse.Size(m)
}
func (m *NoClockUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NoClockUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NoClockUsersResponse proto.InternalMessageInfo

func (m *NoClockUsersResponse) GetList() []*UserList {
	if m != nil {
		return m.List
	}
	return nil
}

type UserList struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	OpenId               string   `protobuf:"bytes,2,opt,name=openId,proto3" json:"openId,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex                  int32    `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	IsBlock              int32    `protobuf:"varint,5,opt,name=isBlock,proto3" json:"isBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{6}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserList) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *UserList) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserList) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *UserList) GetIsBlock() int32 {
	if m != nil {
		return m.IsBlock
	}
	return 0
}

//未打卡的好友
type GetNoClockFriendsRequest struct {
	OpenId               string   `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	ClockType            int32    `protobuf:"varint,2,opt,name=clockType,proto3" json:"clockType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNoClockFriendsRequest) Reset()         { *m = GetNoClockFriendsRequest{} }
func (m *GetNoClockFriendsRequest) String() string { return proto.CompactTextString(m) }
func (*GetNoClockFriendsRequest) ProtoMessage()    {}
func (*GetNoClockFriendsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{7}
}

func (m *GetNoClockFriendsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNoClockFriendsRequest.Unmarshal(m, b)
}
func (m *GetNoClockFriendsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNoClockFriendsRequest.Marshal(b, m, deterministic)
}
func (m *GetNoClockFriendsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNoClockFriendsRequest.Merge(m, src)
}
func (m *GetNoClockFriendsRequest) XXX_Size() int {
	return xxx_messageInfo_GetNoClockFriendsRequest.Size(m)
}
func (m *GetNoClockFriendsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNoClockFriendsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNoClockFriendsRequest proto.InternalMessageInfo

func (m *GetNoClockFriendsRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *GetNoClockFriendsRequest) GetClockType() int32 {
	if m != nil {
		return m.ClockType
	}
	return 0
}

type GetNoClockFriendsResponse struct {
	List                 []*UserList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetNoClockFriendsResponse) Reset()         { *m = GetNoClockFriendsResponse{} }
func (m *GetNoClockFriendsResponse) String() string { return proto.CompactTextString(m) }
func (*GetNoClockFriendsResponse) ProtoMessage()    {}
func (*GetNoClockFriendsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{8}
}

func (m *GetNoClockFriendsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNoClockFriendsResponse.Unmarshal(m, b)
}
func (m *GetNoClockFriendsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNoClockFriendsResponse.Marshal(b, m, deterministic)
}
func (m *GetNoClockFriendsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNoClockFriendsResponse.Merge(m, src)
}
func (m *GetNoClockFriendsResponse) XXX_Size() int {
	return xxx_messageInfo_GetNoClockFriendsResponse.Size(m)
}
func (m *GetNoClockFriendsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNoClockFriendsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNoClockFriendsResponse proto.InternalMessageInfo

func (m *GetNoClockFriendsResponse) GetList() []*UserList {
	if m != nil {
		return m.List
	}
	return nil
}

//获得累计打卡天数
type GetTotalClockDaysRequest struct {
	OpenId               string   `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	ClockType            int32    `protobuf:"varint,2,opt,name=clockType,proto3" json:"clockType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTotalClockDaysRequest) Reset()         { *m = GetTotalClockDaysRequest{} }
func (m *GetTotalClockDaysRequest) String() string { return proto.CompactTextString(m) }
func (*GetTotalClockDaysRequest) ProtoMessage()    {}
func (*GetTotalClockDaysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{9}
}

func (m *GetTotalClockDaysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTotalClockDaysRequest.Unmarshal(m, b)
}
func (m *GetTotalClockDaysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTotalClockDaysRequest.Marshal(b, m, deterministic)
}
func (m *GetTotalClockDaysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTotalClockDaysRequest.Merge(m, src)
}
func (m *GetTotalClockDaysRequest) XXX_Size() int {
	return xxx_messageInfo_GetTotalClockDaysRequest.Size(m)
}
func (m *GetTotalClockDaysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTotalClockDaysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTotalClockDaysRequest proto.InternalMessageInfo

func (m *GetTotalClockDaysRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *GetTotalClockDaysRequest) GetClockType() int32 {
	if m != nil {
		return m.ClockType
	}
	return 0
}

type GetTotalClockDaysResponse struct {
	TotalClockDays       int32    `protobuf:"varint,1,opt,name=totalClockDays,proto3" json:"totalClockDays,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTotalClockDaysResponse) Reset()         { *m = GetTotalClockDaysResponse{} }
func (m *GetTotalClockDaysResponse) String() string { return proto.CompactTextString(m) }
func (*GetTotalClockDaysResponse) ProtoMessage()    {}
func (*GetTotalClockDaysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c0431c549fa44ce, []int{10}
}

func (m *GetTotalClockDaysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTotalClockDaysResponse.Unmarshal(m, b)
}
func (m *GetTotalClockDaysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTotalClockDaysResponse.Marshal(b, m, deterministic)
}
func (m *GetTotalClockDaysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTotalClockDaysResponse.Merge(m, src)
}
func (m *GetTotalClockDaysResponse) XXX_Size() int {
	return xxx_messageInfo_GetTotalClockDaysResponse.Size(m)
}
func (m *GetTotalClockDaysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTotalClockDaysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTotalClockDaysResponse proto.InternalMessageInfo

func (m *GetTotalClockDaysResponse) GetTotalClockDays() int32 {
	if m != nil {
		return m.TotalClockDays
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUpRequest)(nil), "protos.GetUpRequest")
	proto.RegisterType((*GetUpResponse)(nil), "protos.GetUpResponse")
	proto.RegisterType((*SleepRequest)(nil), "protos.SleepRequest")
	proto.RegisterType((*SleepResponse)(nil), "protos.SleepResponse")
	proto.RegisterType((*NoClockUsersRequest)(nil), "protos.NoClockUsersRequest")
	proto.RegisterType((*NoClockUsersResponse)(nil), "protos.NoClockUsersResponse")
	proto.RegisterType((*UserList)(nil), "protos.UserList")
	proto.RegisterType((*GetNoClockFriendsRequest)(nil), "protos.GetNoClockFriendsRequest")
	proto.RegisterType((*GetNoClockFriendsResponse)(nil), "protos.GetNoClockFriendsResponse")
	proto.RegisterType((*GetTotalClockDaysRequest)(nil), "protos.GetTotalClockDaysRequest")
	proto.RegisterType((*GetTotalClockDaysResponse)(nil), "protos.GetTotalClockDaysResponse")
}

func init() { proto.RegisterFile("clock.proto", fileDescriptor_7c0431c549fa44ce) }

var fileDescriptor_7c0431c549fa44ce = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0xdd, 0xec, 0x6e, 0x4a, 0x3b, 0x5d, 0x50, 0x6b, 0x0a, 0x0a, 0x4b, 0x85, 0x16, 0x0b, 0xa1,
	0x1e, 0x50, 0x0f, 0x45, 0xf4, 0xc4, 0x05, 0x8a, 0xa8, 0x2a, 0x10, 0x42, 0xe9, 0xee, 0x85, 0x5b,
	0x9a, 0x8c, 0x90, 0xd5, 0x34, 0x4e, 0x63, 0x2f, 0xb4, 0x57, 0x2e, 0x1c, 0xf8, 0x47, 0xfc, 0x3a,
	0x34, 0xb6, 0xf3, 0xd9, 0x08, 0xb1, 0x42, 0x88, 0xd3, 0x66, 0xde, 0xd8, 0xf3, 0xde, 0xf3, 0x8c,
	0xbd, 0xb0, 0x19, 0xa7, 0x32, 0x3e, 0xdf, 0xcf, 0x0b, 0xa9, 0x25, 0x5b, 0x33, 0x3f, 0x8a, 0x9f,
	0xc1, 0xe4, 0x18, 0xf5, 0x22, 0x0f, 0xf1, 0x72, 0x89, 0x4a, 0xb3, 0xfb, 0xb0, 0x26, 0x73, 0xcc,
	0x4e, 0x92, 0xc0, 0x9b, 0x79, 0x7b, 0x1b, 0xa1, 0x8b, 0xd8, 0x2e, 0x6c, 0x98, 0xed, 0x73, 0x71,
	0x81, 0xc1, 0xd0, 0xa4, 0x6a, 0x80, 0xb2, 0x97, 0x45, 0x2c, 0x13, 0x5c, 0x14, 0x69, 0x30, 0xb2,
	0xd9, 0x0a, 0xe0, 0xdf, 0x3d, 0xb8, 0xed, 0x48, 0x54, 0x2e, 0x33, 0x85, 0xc4, 0x22, 0x2e, 0x3e,
	0xd3, 0x62, 0xc7, 0x62, 0x23, 0xc6, 0x60, 0xac, 0xf1, 0x4a, 0x3b, 0x02, 0xf3, 0xcd, 0x9e, 0xc1,
	0x76, 0x2c, 0x33, 0x2d, 0xb2, 0x25, 0x1e, 0x11, 0xe1, 0x9b, 0xe8, 0x5a, 0x19, 0x0e, 0x3f, 0xbc,
	0x99, 0x20, 0x25, 0x5f, 0x23, 0x8d, 0xc5, 0x9c, 0xca, 0x8c, 0xad, 0x92, 0x0a, 0x20, 0xb7, 0xa7,
	0x29, 0xe2, 0x3f, 0x77, 0xeb, 0x48, 0xfe, 0xb3, 0xdb, 0x1f, 0x1e, 0xdc, 0xfd, 0x20, 0xcd, 0xea,
	0x85, 0xc2, 0x42, 0x95, 0xae, 0x2b, 0x77, 0xd7, 0x39, 0x1a, 0x49, 0x7e, 0x58, 0x03, 0x94, 0x55,
	0x3a, 0x2a, 0x74, 0xd3, 0x7b, 0x05, 0xb0, 0x00, 0x6e, 0x61, 0x96, 0x98, 0x9c, 0x75, 0x5e, 0x86,
	0xec, 0x11, 0x40, 0x1a, 0x29, 0x4d, 0x4c, 0x27, 0x89, 0x11, 0xe3, 0x87, 0x0d, 0x84, 0xbf, 0x84,
	0x9d, 0xb6, 0x18, 0x77, 0x3a, 0x4f, 0x60, 0x9c, 0x0a, 0xa5, 0x03, 0x6f, 0x36, 0xda, 0xdb, 0x3c,
	0xd8, 0xb2, 0xf3, 0xa9, 0xf6, 0x69, 0xd1, 0x7b, 0xa1, 0x74, 0x68, 0xb2, 0xfc, 0x9b, 0x07, 0xeb,
	0x25, 0x44, 0x07, 0xba, 0xb4, 0x34, 0x56, 0xbd, 0x8b, 0x1a, 0xed, 0x1c, 0xb6, 0xda, 0x39, 0x85,
	0xf5, 0x4c, 0xc4, 0xe7, 0x59, 0x54, 0xa9, 0xae, 0x62, 0xb6, 0x05, 0x23, 0x85, 0x57, 0x4e, 0x2f,
	0x7d, 0x92, 0x45, 0xa1, 0x5e, 0x93, 0xd0, 0xc0, 0x37, 0x68, 0x19, 0xf2, 0x8f, 0x10, 0x1c, 0xa3,
	0x76, 0x2e, 0xde, 0x16, 0x02, 0xb3, 0x44, 0xfd, 0xf1, 0x28, 0xd1, 0x61, 0x0f, 0x3b, 0x87, 0xcd,
	0x5f, 0xc1, 0x83, 0x9e, 0x8a, 0x2b, 0x9d, 0x8c, 0x15, 0x35, 0x97, 0x3a, 0x4a, 0xab, 0xc1, 0xf8,
	0x3b, 0x51, 0x47, 0x46, 0x54, 0xb7, 0xa2, 0x13, 0xf5, 0x14, 0xee, 0xe8, 0x56, 0xc6, 0xf5, 0xa0,
	0x83, 0x1e, 0xfc, 0x1c, 0xc1, 0xc4, 0x44, 0xa7, 0x58, 0x7c, 0x11, 0x31, 0xb2, 0x43, 0xf0, 0xcd,
	0x23, 0xc0, 0x76, 0x4a, 0x23, 0xcd, 0x87, 0x67, 0x7a, 0xaf, 0x83, 0x5a, 0x3a, 0x3e, 0x60, 0x2f,
	0x60, 0x3c, 0x27, 0x2f, 0x2b, 0x6e, 0x3b, 0x04, 0xdf, 0xdc, 0xc2, 0x7a, 0x5f, 0xf3, 0xe6, 0xd7,
	0xfb, 0x5a, 0x57, 0x95, 0x0f, 0xd8, 0x3b, 0x98, 0x34, 0xc7, 0x94, 0x3d, 0x2c, 0x17, 0xf6, 0xdc,
	0xa4, 0xe9, 0x6e, 0x7f, 0xb2, 0x2a, 0xf6, 0x09, 0xb6, 0x6f, 0xb4, 0x97, 0xcd, 0x1a, 0x92, 0x7b,
	0x67, 0x69, 0xfa, 0xf8, 0x37, 0x2b, 0x3a, 0xb5, 0xdb, 0x5d, 0x6a, 0xd5, 0xee, 0x1d, 0x89, 0x56,
	0xed, 0xfe, 0x16, 0xf3, 0xc1, 0x99, 0xfd, 0x77, 0x78, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0x15, 0x8d, 0xf1, 0x33, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClockServiceClient is the client API for ClockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClockServiceClient interface {
	GetUp(ctx context.Context, in *GetUpRequest, opts ...grpc.CallOption) (*GetUpResponse, error)
	Test(ctx context.Context, in *GetUpRequest, opts ...grpc.CallOption) (*GetUpResponse, error)
	Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponse, error)
	NoClockUsers(ctx context.Context, in *NoClockUsersRequest, opts ...grpc.CallOption) (*NoClockUsersResponse, error)
	GetNoClockFriends(ctx context.Context, in *GetNoClockFriendsRequest, opts ...grpc.CallOption) (*GetNoClockFriendsResponse, error)
	GetTotalClockDays(ctx context.Context, in *GetTotalClockDaysRequest, opts ...grpc.CallOption) (*GetTotalClockDaysResponse, error)
}

type clockServiceClient struct {
	cc *grpc.ClientConn
}

func NewClockServiceClient(cc *grpc.ClientConn) ClockServiceClient {
	return &clockServiceClient{cc}
}

func (c *clockServiceClient) GetUp(ctx context.Context, in *GetUpRequest, opts ...grpc.CallOption) (*GetUpResponse, error) {
	out := new(GetUpResponse)
	err := c.cc.Invoke(ctx, "/protos.ClockService/GetUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockServiceClient) Test(ctx context.Context, in *GetUpRequest, opts ...grpc.CallOption) (*GetUpResponse, error) {
	out := new(GetUpResponse)
	err := c.cc.Invoke(ctx, "/protos.ClockService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockServiceClient) Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponse, error) {
	out := new(SleepResponse)
	err := c.cc.Invoke(ctx, "/protos.ClockService/Sleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockServiceClient) NoClockUsers(ctx context.Context, in *NoClockUsersRequest, opts ...grpc.CallOption) (*NoClockUsersResponse, error) {
	out := new(NoClockUsersResponse)
	err := c.cc.Invoke(ctx, "/protos.ClockService/NoClockUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockServiceClient) GetNoClockFriends(ctx context.Context, in *GetNoClockFriendsRequest, opts ...grpc.CallOption) (*GetNoClockFriendsResponse, error) {
	out := new(GetNoClockFriendsResponse)
	err := c.cc.Invoke(ctx, "/protos.ClockService/GetNoClockFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clockServiceClient) GetTotalClockDays(ctx context.Context, in *GetTotalClockDaysRequest, opts ...grpc.CallOption) (*GetTotalClockDaysResponse, error) {
	out := new(GetTotalClockDaysResponse)
	err := c.cc.Invoke(ctx, "/protos.ClockService/GetTotalClockDays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClockServiceServer is the server API for ClockService service.
type ClockServiceServer interface {
	GetUp(context.Context, *GetUpRequest) (*GetUpResponse, error)
	Test(context.Context, *GetUpRequest) (*GetUpResponse, error)
	Sleep(context.Context, *SleepRequest) (*SleepResponse, error)
	NoClockUsers(context.Context, *NoClockUsersRequest) (*NoClockUsersResponse, error)
	GetNoClockFriends(context.Context, *GetNoClockFriendsRequest) (*GetNoClockFriendsResponse, error)
	GetTotalClockDays(context.Context, *GetTotalClockDaysRequest) (*GetTotalClockDaysResponse, error)
}

// UnimplementedClockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClockServiceServer struct {
}

func (*UnimplementedClockServiceServer) GetUp(ctx context.Context, req *GetUpRequest) (*GetUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUp not implemented")
}
func (*UnimplementedClockServiceServer) Test(ctx context.Context, req *GetUpRequest) (*GetUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (*UnimplementedClockServiceServer) Sleep(ctx context.Context, req *SleepRequest) (*SleepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sleep not implemented")
}
func (*UnimplementedClockServiceServer) NoClockUsers(ctx context.Context, req *NoClockUsersRequest) (*NoClockUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoClockUsers not implemented")
}
func (*UnimplementedClockServiceServer) GetNoClockFriends(ctx context.Context, req *GetNoClockFriendsRequest) (*GetNoClockFriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoClockFriends not implemented")
}
func (*UnimplementedClockServiceServer) GetTotalClockDays(ctx context.Context, req *GetTotalClockDaysRequest) (*GetTotalClockDaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalClockDays not implemented")
}

func RegisterClockServiceServer(s *grpc.Server, srv ClockServiceServer) {
	s.RegisterService(&_ClockService_serviceDesc, srv)
}

func _ClockService_GetUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).GetUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ClockService/GetUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).GetUp(ctx, req.(*GetUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClockService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ClockService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).Test(ctx, req.(*GetUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClockService_Sleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SleepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).Sleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ClockService/Sleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).Sleep(ctx, req.(*SleepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClockService_NoClockUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoClockUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).NoClockUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ClockService/NoClockUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).NoClockUsers(ctx, req.(*NoClockUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClockService_GetNoClockFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoClockFriendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).GetNoClockFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ClockService/GetNoClockFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).GetNoClockFriends(ctx, req.(*GetNoClockFriendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClockService_GetTotalClockDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalClockDaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockServiceServer).GetTotalClockDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ClockService/GetTotalClockDays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockServiceServer).GetTotalClockDays(ctx, req.(*GetTotalClockDaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ClockService",
	HandlerType: (*ClockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUp",
			Handler:    _ClockService_GetUp_Handler,
		},
		{
			MethodName: "Test",
			Handler:    _ClockService_Test_Handler,
		},
		{
			MethodName: "Sleep",
			Handler:    _ClockService_Sleep_Handler,
		},
		{
			MethodName: "NoClockUsers",
			Handler:    _ClockService_NoClockUsers_Handler,
		},
		{
			MethodName: "GetNoClockFriends",
			Handler:    _ClockService_GetNoClockFriends_Handler,
		},
		{
			MethodName: "GetTotalClockDays",
			Handler:    _ClockService_GetTotalClockDays_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clock.proto",
}
