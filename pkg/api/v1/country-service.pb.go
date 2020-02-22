// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/country/country-service.proto

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

type Country struct {
	//Unique identifier of each country
	CountryId int64 `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	//Name of the country
	Country              string               `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	LastUpdate           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Country) Reset()         { *m = Country{} }
func (m *Country) String() string { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()    {}
func (*Country) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{0}
}

func (m *Country) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Country.Unmarshal(m, b)
}
func (m *Country) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Country.Marshal(b, m, deterministic)
}
func (m *Country) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Country.Merge(m, src)
}
func (m *Country) XXX_Size() int {
	return xxx_messageInfo_Country.Size(m)
}
func (m *Country) XXX_DiscardUnknown() {
	xxx_messageInfo_Country.DiscardUnknown(m)
}

var xxx_messageInfo_Country proto.InternalMessageInfo

func (m *Country) GetCountryId() int64 {
	if m != nil {
		return m.CountryId
	}
	return 0
}

func (m *Country) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Country) GetLastUpdate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

//Request data to create new country
type CreateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Country entity
	Country              *Country `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{1}
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

func (m *CreateRequest) GetCountry() *Country {
	if m != nil {
		return m.Country
	}
	return nil
}

//Response data of created new country
type CreateResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Id of created country
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{2}
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

//Request data to read country by id
type ReadRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Get country by id
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{3}
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
	//Country by id
	Country              *Country `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{4}
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

func (m *ReadResponse) GetCountry() *Country {
	if m != nil {
		return m.Country
	}
	return nil
}

//Request data to update exists country
type UpdateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Country entity
	Country              *Country `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{5}
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

func (m *UpdateRequest) GetCountry() *Country {
	if m != nil {
		return m.Country
	}
	return nil
}

//Response data of updated country
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
	return fileDescriptor_8dc6626b99715c9f, []int{6}
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

//Request to delete country by id
type DeleteRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//id of country which we want to delete
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{7}
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

func (m *DeleteRequest) GetCountryId() int64 {
	if m != nil {
		return m.CountryId
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
	return fileDescriptor_8dc6626b99715c9f, []int{8}
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

//Request for reading all country entities
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
	return fileDescriptor_8dc6626b99715c9f, []int{9}
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

//Response for read all country
type ReadAllResponse struct {
	//api version
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//all country
	Country              []*Country `protobuf:"bytes,2,rep,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dc6626b99715c9f, []int{10}
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

func (m *ReadAllResponse) GetCountry() []*Country {
	if m != nil {
		return m.Country
	}
	return nil
}

func init() {
	proto.RegisterType((*Country)(nil), "Country")
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

func init() {
	proto.RegisterFile("proto/country/country-service.proto", fileDescriptor_8dc6626b99715c9f)
}

var fileDescriptor_8dc6626b99715c9f = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0x76, 0xb1, 0xeb, 0x49, 0x2c, 0x07, 0x9d, 0x8c, 0xa1, 0xd4, 0x28, 0x97, 0x04, 0x5a,
	0x05, 0xd2, 0x5b, 0x09, 0xa5, 0x21, 0xc9, 0xa1, 0x57, 0x77, 0xf7, 0xbc, 0x38, 0xb1, 0x36, 0x18,
	0x9c, 0xd8, 0x6b, 0xcb, 0x0b, 0x7b, 0xd9, 0x5f, 0xbb, 0x3f, 0x64, 0x89, 0x25, 0x85, 0xc8, 0xe4,
	0x63, 0x61, 0x4f, 0xc9, 0xc8, 0x4f, 0xf3, 0xde, 0xbc, 0x79, 0x82, 0x61, 0x59, 0x15, 0xbc, 0x98,
	0x6c, 0x8a, 0x66, 0xcf, 0xab, 0x17, 0xf5, 0xfb, 0xb3, 0x66, 0xd5, 0x73, 0xb6, 0x61, 0xb4, 0xfd,
	0x1a, 0x7e, 0xdf, 0x16, 0xc5, 0x36, 0x67, 0x93, 0xb6, 0x5a, 0x37, 0x8f, 0x13, 0x9e, 0xed, 0x58,
	0xcd, 0x93, 0x5d, 0x29, 0x00, 0xe4, 0x15, 0x9c, 0x85, 0xb8, 0x89, 0xbf, 0x01, 0xc8, 0x26, 0x0f,
	0x59, 0x1a, 0x18, 0x91, 0x31, 0xb2, 0x62, 0x57, 0x9e, 0xfc, 0x4b, 0x71, 0x00, 0x8e, 0x2c, 0x02,
	0x33, 0x32, 0x46, 0x6e, 0xac, 0x4a, 0xfc, 0x1b, 0x20, 0x4f, 0x6a, 0x7e, 0x5f, 0xa6, 0x09, 0x67,
	0x81, 0x15, 0x19, 0xa3, 0xde, 0x34, 0xa4, 0x82, 0x99, 0x2a, 0x66, 0x7a, 0xa7, 0x98, 0xe3, 0x13,
	0x34, 0x59, 0x81, 0xb7, 0xa8, 0x58, 0xc2, 0x59, 0xcc, 0x9e, 0x1a, 0x56, 0x73, 0x3c, 0x00, 0x2b,
	0x29, 0xb3, 0x96, 0xde, 0x8d, 0x0f, 0x7f, 0x31, 0xd1, 0x89, 0x7b, 0xd3, 0xaf, 0x54, 0x4a, 0x3e,
	0x4a, 0x20, 0x73, 0x40, 0xaa, 0x4d, 0x5d, 0x16, 0xfb, 0x9a, 0x9d, 0xe9, 0xa3, 0xcf, 0x67, 0x76,
	0xe6, 0x23, 0x7f, 0xa0, 0x17, 0xb3, 0x24, 0xbd, 0xac, 0xe3, 0xc6, 0xfd, 0x25, 0xf4, 0xc5, 0xfd,
	0x8b, 0x02, 0x3e, 0x32, 0xc8, 0x0a, 0x3c, 0xe1, 0xcc, 0xe7, 0xfc, 0x98, 0x01, 0x52, 0x6d, 0x2e,
	0xca, 0x09, 0xc0, 0x69, 0x5a, 0x8c, 0x1a, 0x46, 0x95, 0xe4, 0x2f, 0x78, 0x4b, 0x96, 0xb3, 0x6b,
	0x22, 0x6e, 0x98, 0x31, 0x03, 0xa4, 0x3a, 0x5c, 0xe3, 0x4f, 0x5b, 0xcc, 0x91, 0x5f, 0x96, 0x24,
	0x02, 0x74, 0xb0, 0x72, 0x9e, 0xe7, 0x4a, 0x00, 0x02, 0x53, 0x66, 0xd2, 0x8d, 0xcd, 0x2c, 0x25,
	0x2b, 0xf0, 0x8f, 0x08, 0x49, 0xd0, 0x81, 0xe8, 0x36, 0x59, 0x67, 0x6d, 0x9a, 0xbe, 0x19, 0x80,
	0xe4, 0xe1, 0x7f, 0xf1, 0x6e, 0xf0, 0x18, 0x6c, 0x91, 0x24, 0x8c, 0xa8, 0x96, 0xcc, 0xd0, 0xa7,
	0x9d, 0x88, 0x0d, 0xe1, 0xcb, 0x41, 0x04, 0xee, 0xd3, 0x93, 0xe0, 0x84, 0x1e, 0xd5, 0x62, 0x30,
	0x06, 0x5b, 0x6c, 0x02, 0x23, 0xaa, 0x6d, 0x36, 0xf4, 0x69, 0x67, 0x45, 0x63, 0xb0, 0x85, 0x69,
	0x18, 0x51, 0xcd, 0xff, 0xd0, 0xa7, 0x1d, 0x37, 0x7f, 0x80, 0x23, 0xe7, 0xc7, 0x3e, 0xd5, 0xbd,
	0x0a, 0x07, 0xb4, 0x63, 0xcd, 0xda, 0x6e, 0x1f, 0xe1, 0xaf, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf1, 0x19, 0xef, 0xd7, 0x33, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CountryServiceClient is the client API for CountryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountryServiceClient interface {
	//Create new country
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	//Read country
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	//Update country
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	//Delete country
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	//Read all country
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type countryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountryServiceClient(cc grpc.ClientConnInterface) CountryServiceClient {
	return &countryServiceClient{cc}
}

func (c *countryServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/CountryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/CountryService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/CountryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/CountryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countryServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/CountryService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryServiceServer is the server API for CountryService service.
type CountryServiceServer interface {
	//Create new country
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	//Read country
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	//Update country
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	//Delete country
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	//Read all country
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedCountryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCountryServiceServer struct {
}

func (*UnimplementedCountryServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCountryServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedCountryServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCountryServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCountryServiceServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterCountryServiceServer(s *grpc.Server, srv CountryServiceServer) {
	s.RegisterService(&_CountryService_serviceDesc, srv)
}

func _CountryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CountryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CountryService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CountryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CountryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountryService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CountryService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CountryService",
	HandlerType: (*CountryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CountryService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _CountryService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CountryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CountryService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _CountryService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/country/country-service.proto",
}