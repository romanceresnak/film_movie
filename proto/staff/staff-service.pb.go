// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/staff/staff-service.proto

package staff_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Staff struct {
	//unique identifier of staff
	StaffId int64 `protobuf:"varint,1,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	//first name
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	//last name
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	//address id of store
	AddressId string `protobuf:"bytes,4,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	//picture, it will be represent as blob
	Picture string `protobuf:"bytes,5,opt,name=picture,proto3" json:"picture,omitempty"`
	//email of the store
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	//unique identifier of store
	StoreId int64 `protobuf:"varint,7,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	//user is active 0 alse 1
	Active int64 `protobuf:"varint,8,opt,name=active,proto3" json:"active,omitempty"`
	//user name
	Username string `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	//password of user
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	//Attribute about last update
	LastUpdate           *timestamp.Timestamp `protobuf:"bytes,11,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Staff) Reset()         { *m = Staff{} }
func (m *Staff) String() string { return proto.CompactTextString(m) }
func (*Staff) ProtoMessage()    {}
func (*Staff) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{0}
}

func (m *Staff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Staff.Unmarshal(m, b)
}
func (m *Staff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Staff.Marshal(b, m, deterministic)
}
func (m *Staff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Staff.Merge(m, src)
}
func (m *Staff) XXX_Size() int {
	return xxx_messageInfo_Staff.Size(m)
}
func (m *Staff) XXX_DiscardUnknown() {
	xxx_messageInfo_Staff.DiscardUnknown(m)
}

var xxx_messageInfo_Staff proto.InternalMessageInfo

func (m *Staff) GetStaffId() int64 {
	if m != nil {
		return m.StaffId
	}
	return 0
}

func (m *Staff) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Staff) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Staff) GetAddressId() string {
	if m != nil {
		return m.AddressId
	}
	return ""
}

func (m *Staff) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *Staff) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Staff) GetStoreId() int64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Staff) GetActive() int64 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *Staff) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Staff) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Staff) GetLastUpdate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

//Request data to create new staff
type CreateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Staff entity
	Staff                *Staff   `protobuf:"bytes,2,opt,name=staff,proto3" json:"staff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetStaff() *Staff {
	if m != nil {
		return m.Staff
	}
	return nil
}

//Response data of created new staff
type CreateResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Id of created Staff
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetCountryId() int64 {
	if m != nil {
		return m.CountryId
	}
	return 0
}

//Request data to read Staff by id
type ReadRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Get staff by id
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetCountryId() int64 {
	if m != nil {
		return m.CountryId
	}
	return 0
}

//Response data by id
type ReadResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Store text by id
	StaffId              *Staff   `protobuf:"bytes,2,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetStaffId() *Staff {
	if m != nil {
		return m.StaffId
	}
	return nil
}

//Request data to update exists staff
type UpdateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Store staff entity
	Staff                *Staff   `protobuf:"bytes,2,opt,name=staff,proto3" json:"staff,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetStaff() *Staff {
	if m != nil {
		return m.Staff
	}
	return nil
}

//Response data of updated staff
type UpdateResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have been updated
	// Equals 1 in case of successful update
	Updated              int64    `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

//Request to delete staff by id
type DeleteRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//id of staff which we want to delete
	StaffId              int64    `protobuf:"varint,2,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetStaffId() int64 {
	if m != nil {
		return m.StaffId
	}
	return 0
}

//Response for delete request
type DeleteResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have been deleted
	// Equals 1 in case of successful delete
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

//Request for reading all staff entities
type ReadAllRequest struct {
	//api version
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{9}
}

func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (m *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(m, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

//Response for read all stores
type ReadAllResponse struct {
	//api version
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//all staffs
	Staffs               []*Staff `protobuf:"bytes,2,rep,name=staffs,proto3" json:"staffs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01e7cb58d95696ff, []int{10}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadAllResponse) GetStaffs() []*Staff {
	if m != nil {
		return m.Staffs
	}
	return nil
}

func init() {
	proto.RegisterType((*Staff)(nil), "Staff")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "ReadAllResponse")
}

func init() { proto.RegisterFile("proto/staff/staff-service.proto", fileDescriptor_01e7cb58d95696ff) }

var fileDescriptor_01e7cb58d95696ff = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x8e, 0xd2, 0x4e,
	0x14, 0xc7, 0x43, 0x59, 0x0a, 0x3d, 0x40, 0xd9, 0x4c, 0x7e, 0xf9, 0xa5, 0x5b, 0xc5, 0xc5, 0x7a,
	0x03, 0x89, 0x0e, 0x09, 0xde, 0x99, 0x8d, 0x86, 0xac, 0x37, 0xdc, 0x78, 0xd1, 0xd5, 0xeb, 0xcd,
	0x2c, 0x73, 0xd8, 0x34, 0x29, 0xb4, 0xce, 0x4c, 0xd7, 0xf8, 0x9e, 0x3e, 0x81, 0x4f, 0x62, 0x66,
	0xa6, 0xb3, 0xd2, 0x46, 0xd6, 0xc4, 0x1b, 0xc2, 0xf7, 0xfc, 0x3f, 0x9f, 0x39, 0x85, 0xcb, 0x52,
	0x14, 0xaa, 0x58, 0x4a, 0xc5, 0x76, 0x3b, 0xfb, 0xfb, 0x46, 0xa2, 0x78, 0xc8, 0xb6, 0x48, 0x8d,
	0x27, 0xbe, 0xbc, 0x2f, 0x8a, 0xfb, 0x1c, 0x97, 0x46, 0xdd, 0x55, 0xbb, 0xa5, 0xca, 0xf6, 0x28,
	0x15, 0xdb, 0x97, 0x36, 0x20, 0xf9, 0xe1, 0x41, 0xef, 0x46, 0x27, 0x92, 0x0b, 0x18, 0x98, 0x0a,
	0xb7, 0x19, 0x8f, 0x3a, 0xb3, 0xce, 0xbc, 0x9b, 0xf6, 0x8d, 0xde, 0x70, 0x32, 0x05, 0xd8, 0x65,
	0x42, 0xaa, 0xdb, 0x03, 0xdb, 0x63, 0xe4, 0xcd, 0x3a, 0xf3, 0x20, 0x0d, 0x8c, 0xe5, 0x13, 0xdb,
	0x23, 0x79, 0x06, 0x41, 0xce, 0x9c, 0xb7, 0x6b, 0xbc, 0x03, 0x6d, 0x30, 0xce, 0x29, 0x00, 0xe3,
	0x5c, 0xa0, 0x94, 0xba, 0xf0, 0x99, 0xcd, 0xad, 0x2d, 0x1b, 0x4e, 0x22, 0xe8, 0x97, 0xd9, 0x56,
	0x55, 0x02, 0xa3, 0x9e, 0xf1, 0x39, 0x49, 0xfe, 0x83, 0x1e, 0xee, 0x59, 0x96, 0x47, 0xbe, 0xb1,
	0x5b, 0x61, 0xa7, 0x2c, 0x04, 0xea, 0x62, 0x7d, 0x37, 0x65, 0x21, 0x70, 0xc3, 0xc9, 0xff, 0xe0,
	0xb3, 0xad, 0xca, 0x1e, 0x30, 0x1a, 0x18, 0x47, 0xad, 0x48, 0x0c, 0x83, 0x4a, 0xa2, 0x30, 0xd3,
	0x05, 0x76, 0x3a, 0xa7, 0xb5, 0xaf, 0x64, 0x52, 0x7e, 0x2b, 0x04, 0x8f, 0xc0, 0xfa, 0x9c, 0x26,
	0xef, 0x00, 0xf4, 0x16, 0x5f, 0x4a, 0xce, 0x14, 0x46, 0xc3, 0x59, 0x67, 0x3e, 0x5c, 0xc5, 0xd4,
	0x02, 0xa5, 0x0e, 0x28, 0xfd, 0xec, 0x80, 0xa6, 0x47, 0xd1, 0xc9, 0x07, 0x18, 0x5f, 0x0b, 0x64,
	0x0a, 0x53, 0xfc, 0x5a, 0xa1, 0x54, 0xe4, 0x1c, 0xba, 0xac, 0xcc, 0x0c, 0xd8, 0x20, 0xd5, 0x7f,
	0xc9, 0x73, 0xe8, 0x19, 0xbe, 0x86, 0xe7, 0x70, 0xe5, 0x53, 0xf3, 0x0c, 0xa9, 0x35, 0x26, 0x6b,
	0x08, 0x5d, 0x01, 0x59, 0x16, 0x07, 0x89, 0x7f, 0xa8, 0x30, 0x05, 0xd8, 0x16, 0xd5, 0x41, 0x89,
	0xef, 0x9a, 0x86, 0x67, 0x96, 0x0e, 0x6a, 0xcb, 0x86, 0x27, 0xef, 0x61, 0x98, 0x22, 0xe3, 0xa7,
	0x27, 0xf8, 0x4b, 0xfe, 0x35, 0x8c, 0x6c, 0xfe, 0xc9, 0x01, 0x5e, 0x1e, 0x9d, 0x4c, 0x73, 0x0b,
	0x77, 0x3a, 0x1a, 0x84, 0x45, 0xf2, 0xaf, 0x20, 0xae, 0x20, 0x74, 0x05, 0x4e, 0xce, 0x11, 0x41,
	0xbf, 0x32, 0x31, 0x6e, 0x0b, 0x27, 0x93, 0x2b, 0x18, 0x7f, 0xc4, 0x1c, 0x9f, 0x6a, 0x7f, 0xd1,
	0x5a, 0xe2, 0xf7, 0xdd, 0xeb, 0xde, 0x2e, 0xfb, 0xa9, 0xde, 0xdc, 0xc4, 0x3c, 0x66, 0xd7, 0x32,
	0x99, 0x41, 0xa8, 0xf9, 0xad, 0xf3, 0xdc, 0x35, 0x0f, 0xc1, 0xab, 0x3f, 0xae, 0x20, 0xf5, 0x32,
	0x9e, 0xac, 0x61, 0xf2, 0x18, 0x51, 0x37, 0x68, 0x85, 0x90, 0x17, 0xe0, 0x9b, 0x69, 0x64, 0xe4,
	0xcd, 0xba, 0x47, 0x74, 0x6a, 0xeb, 0xea, 0x67, 0x07, 0x46, 0xc6, 0x72, 0x63, 0xbf, 0x7b, 0xb2,
	0x00, 0xdf, 0x1e, 0x0e, 0x09, 0x69, 0xe3, 0x04, 0xe3, 0x09, 0x6d, 0x5d, 0xd4, 0x2b, 0x38, 0xd3,
	0xed, 0xc9, 0x88, 0x1e, 0xdd, 0x49, 0x3c, 0xa6, 0x8d, 0x57, 0x5f, 0x80, 0x6f, 0xf9, 0x93, 0x90,
	0x36, 0x5e, 0x32, 0x9e, 0xd0, 0xd6, 0xc3, 0x2c, 0xc0, 0xb7, 0xb8, 0x48, 0x48, 0x1b, 0xd4, 0xe3,
	0x09, 0x6d, 0x71, 0x7c, 0x0d, 0xfd, 0x7a, 0x73, 0x32, 0xa1, 0x4d, 0x4a, 0xf1, 0x39, 0x6d, 0x41,
	0xb9, 0xf3, 0xcd, 0xd7, 0xf6, 0xf6, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xba, 0xff, 0x6e,
	0xef, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StaffServiceClient interface {
	//Create new staff
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	//Read staff
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	//Update staff
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	//Delete staff
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	//Read all staff
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/StaffService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/StaffService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/StaffService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/StaffService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/StaffService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
type StaffServiceServer interface {
	//Create new staff
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	//Read staff
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	//Update staff
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	//Delete staff
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	//Read all staff
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedStaffServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (*UnimplementedStaffServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedStaffServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedStaffServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedStaffServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedStaffServiceServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterStaffServiceServer(s *grpc.Server, srv StaffServiceServer) {
	s.RegisterService(&_StaffService_serviceDesc, srv)
}

func _StaffService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StaffService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StaffService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StaffService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _StaffService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _StaffService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/staff/staff-service.proto",
}