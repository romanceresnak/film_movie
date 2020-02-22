// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/store/store-service.proto

package v1

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

type Store struct {
	//unique identifier of store
	StoreId int64 `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	//film manager staff
	ManagerStaffId string `protobuf:"bytes,2,opt,name=manager_staff_id,json=managerStaffId,proto3" json:"manager_staff_id,omitempty"`
	//address id
	AddressId string `protobuf:"bytes,3,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	//Attribute about last update
	LastUpdate           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Store) Reset()         { *m = Store{} }
func (m *Store) String() string { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()    {}
func (*Store) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{0}
}

func (m *Store) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Store.Unmarshal(m, b)
}
func (m *Store) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Store.Marshal(b, m, deterministic)
}
func (m *Store) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Store.Merge(m, src)
}
func (m *Store) XXX_Size() int {
	return xxx_messageInfo_Store.Size(m)
}
func (m *Store) XXX_DiscardUnknown() {
	xxx_messageInfo_Store.DiscardUnknown(m)
}

var xxx_messageInfo_Store proto.InternalMessageInfo

func (m *Store) GetStoreId() int64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

func (m *Store) GetManagerStaffId() string {
	if m != nil {
		return m.ManagerStaffId
	}
	return ""
}

func (m *Store) GetAddressId() string {
	if m != nil {
		return m.AddressId
	}
	return ""
}

func (m *Store) GetLastUpdate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

//Request data to create new film text
type CreateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Store entity
	Country              *Store   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{1}
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

func (m *CreateRequest) GetCountry() *Store {
	if m != nil {
		return m.Country
	}
	return nil
}

//Response data of created new film text
type CreateResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Id of created Store
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{2}
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

//Request data to read Store by id
type ReadRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Get store text by id
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{3}
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
	FilmText             *Store   `protobuf:"bytes,2,opt,name=filmText,proto3" json:"filmText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{4}
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

func (m *ReadResponse) GetFilmText() *Store {
	if m != nil {
		return m.FilmText
	}
	return nil
}

//Request data to update exists store text
type UpdateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Store text entity
	FilmText             *Store   `protobuf:"bytes,2,opt,name=filmText,proto3" json:"filmText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{5}
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

func (m *UpdateRequest) GetFilmText() *Store {
	if m != nil {
		return m.FilmText
	}
	return nil
}

//Response data of updated store text
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
	return fileDescriptor_d7484550df16d0e2, []int{6}
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

//Request to delete store by id
type DeleteRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//id of film text which we want to delete
	StoreId              int64    `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{7}
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

func (m *DeleteRequest) GetStoreId() int64 {
	if m != nil {
		return m.StoreId
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
	return fileDescriptor_d7484550df16d0e2, []int{8}
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

//Request for reading all store entities
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
	return fileDescriptor_d7484550df16d0e2, []int{9}
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
	//all film
	Stores               []*Store `protobuf:"bytes,2,rep,name=stores,proto3" json:"stores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7484550df16d0e2, []int{10}
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

func (m *ReadAllResponse) GetStores() []*Store {
	if m != nil {
		return m.Stores
	}
	return nil
}

func init() {
	proto.RegisterType((*Store)(nil), "Store")
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

func init() { proto.RegisterFile("proto/store/store-service.proto", fileDescriptor_d7484550df16d0e2) }

var fileDescriptor_d7484550df16d0e2 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x55, 0x92, 0x2e, 0x34, 0x17, 0x48, 0x90, 0x9f, 0x58, 0xa4, 0xae, 0x91, 0xf7, 0x02, 0xd2,
	0x66, 0x24, 0xf6, 0x36, 0x55, 0x93, 0x50, 0xbb, 0x07, 0x5e, 0x4d, 0xf7, 0x5c, 0xb9, 0xb5, 0x41,
	0x91, 0x02, 0xc9, 0x62, 0x33, 0x6d, 0xbf, 0x68, 0xff, 0x6b, 0xbf, 0x64, 0x8a, 0x3f, 0x10, 0x89,
	0x16, 0xd6, 0x17, 0x84, 0x4f, 0xce, 0xbd, 0xf7, 0xdc, 0xe3, 0x63, 0xb8, 0xad, 0xea, 0x52, 0x95,
	0x0b, 0xa9, 0xca, 0x5a, 0x98, 0xdf, 0x8f, 0x52, 0xd4, 0x3f, 0xf2, 0x17, 0x41, 0xf4, 0x97, 0xf4,
	0x76, 0x57, 0x96, 0xbb, 0x42, 0x2c, 0xf4, 0xe9, 0xf9, 0xb8, 0x5d, 0xa8, 0x7c, 0x2f, 0xa4, 0x62,
	0xfb, 0xca, 0x10, 0xf0, 0x6f, 0x0f, 0xde, 0x6c, 0x9a, 0x42, 0xf4, 0x16, 0xae, 0x75, 0x87, 0xa7,
	0x9c, 0x4f, 0xbd, 0xcc, 0x9b, 0x05, 0x74, 0xa0, 0xcf, 0x6b, 0x8e, 0x66, 0x30, 0xd9, 0xb3, 0x03,
	0xdb, 0x89, 0xfa, 0x49, 0x2a, 0xb6, 0xdd, 0x36, 0x14, 0x3f, 0xf3, 0x66, 0x11, 0x8d, 0x2d, 0xbe,
	0x69, 0xe0, 0x35, 0x47, 0x37, 0x00, 0x8c, 0xf3, 0x5a, 0x48, 0xd9, 0x70, 0x02, 0xcd, 0x89, 0x2c,
	0xb2, 0xe6, 0xe8, 0x33, 0x40, 0xc1, 0xa4, 0xfa, 0x56, 0x71, 0xa6, 0xc4, 0xf4, 0x2a, 0xf3, 0x66,
	0xc3, 0x65, 0x4a, 0x8c, 0x46, 0xe2, 0x34, 0x92, 0x47, 0xa7, 0x91, 0x9e, 0xb1, 0xf1, 0x3d, 0x8c,
	0xef, 0x6b, 0xc1, 0x94, 0xa0, 0xe2, 0xfb, 0x51, 0x48, 0x85, 0x26, 0x10, 0xb0, 0x2a, 0xd7, 0x5a,
	0x23, 0xda, 0xfc, 0x45, 0x19, 0x0c, 0x5e, 0xca, 0xe3, 0x41, 0xd5, 0xbf, 0xb4, 0xbc, 0xe1, 0x32,
	0x24, 0x7a, 0x37, 0xea, 0x60, 0xbc, 0x82, 0xd8, 0x35, 0x91, 0x55, 0x79, 0x90, 0xe2, 0x1f, 0x5d,
	0x6e, 0x00, 0x2c, 0xdd, 0xed, 0x19, 0xd0, 0xc8, 0x22, 0x6b, 0x8e, 0xbf, 0xc0, 0x90, 0x0a, 0xc6,
	0xfb, 0x55, 0xfc, 0xa7, 0xfe, 0x01, 0x46, 0xa6, 0xbe, 0x57, 0x00, 0x86, 0xeb, 0x6d, 0x5e, 0xec,
	0x1f, 0xc5, 0x4f, 0xd5, 0xd9, 0xe3, 0x84, 0xe3, 0xaf, 0x30, 0x36, 0xbe, 0xf4, 0xeb, 0x78, 0x4d,
	0x9b, 0x3b, 0x88, 0x5d, 0x9b, 0x5e, 0x39, 0x53, 0x18, 0x1c, 0x35, 0xc7, 0x2d, 0xe3, 0x8e, 0xf8,
	0x0e, 0xc6, 0x0f, 0xa2, 0x10, 0x97, 0x44, 0x9c, 0xa7, 0xca, 0x6f, 0xa5, 0xaa, 0x99, 0xed, 0xaa,
	0x2f, 0xcd, 0xe6, 0x9a, 0x73, 0xaa, 0xb6, 0x47, 0x9c, 0x41, 0xdc, 0xd8, 0xb8, 0x2a, 0x0a, 0x37,
	0x3c, 0x06, 0xdf, 0x46, 0x37, 0xa2, 0x7e, 0xce, 0xf1, 0x0a, 0x92, 0x13, 0xc3, 0x0e, 0xe8, 0x50,
	0xd0, 0x3b, 0x08, 0xb5, 0x1a, 0x39, 0xf5, 0xb3, 0xe0, 0xcc, 0x20, 0x8b, 0x2e, 0xff, 0x78, 0x30,
	0xd2, 0xc8, 0xc6, 0xbc, 0x2a, 0x34, 0x87, 0xd0, 0xe4, 0x07, 0xc5, 0xa4, 0x95, 0xc6, 0x34, 0x21,
	0x9d, 0x60, 0xbd, 0x87, 0xab, 0x66, 0x3c, 0x1a, 0x91, 0xb3, 0xb8, 0xa4, 0x63, 0xd2, 0xba, 0xfc,
	0x39, 0x84, 0xc6, 0x7f, 0x14, 0x93, 0xd6, 0x7d, 0xa6, 0x09, 0xe9, 0x5c, 0xcc, 0x1c, 0x42, 0x63,
	0x17, 0x8a, 0x49, 0xcb, 0xf5, 0x34, 0x21, 0x1d, 0x1f, 0x3f, 0xc0, 0xc0, 0x6e, 0x8e, 0x12, 0xd2,
	0x76, 0x29, 0x9d, 0x90, 0x8e, 0x29, 0xcf, 0xa1, 0x7e, 0x78, 0x9f, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x50, 0xd3, 0x77, 0x13, 0x4d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	//Create new store text
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	//Read store text
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	//Update store text
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	//Delete store text
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	//Read all store text
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/StoreService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/StoreService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/StoreService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/StoreService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/StoreService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	//Create new store text
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	//Read store text
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	//Update store text
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	//Delete store text
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	//Read all store text
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (*UnimplementedStoreServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedStoreServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedStoreServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedStoreServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedStoreServiceServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoreService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoreService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoreService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoreService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoreService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StoreService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _StoreService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StoreService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StoreService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _StoreService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/store/store-service.proto",
}