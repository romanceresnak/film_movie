// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rental/rental-service.proto

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

type Rental struct {
	//unique identifier of rental
	RentalId int64 `protobuf:"varint,1,opt,name=rental_id,json=rentalId,proto3" json:"rental_id,omitempty"`
	//rental date
	RentalDate string `protobuf:"bytes,2,opt,name=rental_date,json=rentalDate,proto3" json:"rental_date,omitempty"`
	//primary key from inventory table
	InventoryId string `protobuf:"bytes,3,opt,name=inventory_id,json=inventoryId,proto3" json:"inventory_id,omitempty"`
	//primary key from customer entity
	CustomerId string `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	//return date of rent
	ReturnDate string `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	//primary key from staff entity
	StaffId string `protobuf:"bytes,6,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	//Attribute about last update
	LastUpdate           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Rental) Reset()         { *m = Rental{} }
func (m *Rental) String() string { return proto.CompactTextString(m) }
func (*Rental) ProtoMessage()    {}
func (*Rental) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{0}
}

func (m *Rental) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rental.Unmarshal(m, b)
}
func (m *Rental) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rental.Marshal(b, m, deterministic)
}
func (m *Rental) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rental.Merge(m, src)
}
func (m *Rental) XXX_Size() int {
	return xxx_messageInfo_Rental.Size(m)
}
func (m *Rental) XXX_DiscardUnknown() {
	xxx_messageInfo_Rental.DiscardUnknown(m)
}

var xxx_messageInfo_Rental proto.InternalMessageInfo

func (m *Rental) GetRentalId() int64 {
	if m != nil {
		return m.RentalId
	}
	return 0
}

func (m *Rental) GetRentalDate() string {
	if m != nil {
		return m.RentalDate
	}
	return ""
}

func (m *Rental) GetInventoryId() string {
	if m != nil {
		return m.InventoryId
	}
	return ""
}

func (m *Rental) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Rental) GetReturnDate() string {
	if m != nil {
		return m.ReturnDate
	}
	return ""
}

func (m *Rental) GetStaffId() string {
	if m != nil {
		return m.StaffId
	}
	return ""
}

func (m *Rental) GetLastUpdate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

//Request data to create new rental
type CreateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Rental entity
	Rental               *Rental  `protobuf:"bytes,2,opt,name=rental,proto3" json:"rental,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{1}
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

func (m *CreateRequest) GetRental() *Rental {
	if m != nil {
		return m.Rental
	}
	return nil
}

//Response data of created new rental
type CreateResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Id of created rental
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{2}
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

//Request data to read rental by id
type ReadRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Get rental by id
	CountryId            int64    `protobuf:"varint,2,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{3}
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
	//Store rental by id
	RentalId             *Rental  `protobuf:"bytes,2,opt,name=rental_id,json=rentalId,proto3" json:"rental_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{4}
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

func (m *ReadResponse) GetRentalId() *Rental {
	if m != nil {
		return m.RentalId
	}
	return nil
}

//Request data to update exists rental
type UpdateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Rental entity
	Rental               *Rental  `protobuf:"bytes,2,opt,name=rental,proto3" json:"rental,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{5}
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

func (m *UpdateRequest) GetRental() *Rental {
	if m != nil {
		return m.Rental
	}
	return nil
}

//Response data of updated rental
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
	return fileDescriptor_3af188db2d5632a8, []int{6}
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

//Request to delete rental by id
type DeleteRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//id of rentals which we want to delete
	RentalId             int64    `protobuf:"varint,2,opt,name=rental_id,json=rentalId,proto3" json:"rental_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{7}
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

func (m *DeleteRequest) GetRentalId() int64 {
	if m != nil {
		return m.RentalId
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
	return fileDescriptor_3af188db2d5632a8, []int{8}
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

//Request for reading all rental entities
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
	return fileDescriptor_3af188db2d5632a8, []int{9}
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

//Response for read all rentals
type ReadAllResponse struct {
	//api version
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//all rentals
	Rentals              []*Rental `protobuf:"bytes,2,rep,name=rentals,proto3" json:"rentals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af188db2d5632a8, []int{10}
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

func (m *ReadAllResponse) GetRentals() []*Rental {
	if m != nil {
		return m.Rentals
	}
	return nil
}

func init() {
	proto.RegisterType((*Rental)(nil), "Rental")
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

func init() { proto.RegisterFile("proto/rental/rental-service.proto", fileDescriptor_3af188db2d5632a8) }

var fileDescriptor_3af188db2d5632a8 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x8f, 0x9b, 0x3c,
	0x10, 0x15, 0x64, 0x3f, 0x58, 0x86, 0x04, 0x56, 0x3e, 0xe5, 0x63, 0x55, 0x85, 0xd0, 0x1e, 0x12,
	0xa9, 0x75, 0xa4, 0xf4, 0x56, 0x55, 0x2b, 0x6d, 0x1b, 0x55, 0xca, 0x95, 0xb6, 0xe7, 0xca, 0x1b,
	0x9c, 0x15, 0x12, 0x01, 0x8a, 0xcd, 0x4a, 0xbd, 0xf6, 0xaf, 0xf6, 0x8f, 0x54, 0xf6, 0xe0, 0x08,
	0x50, 0xb3, 0x3d, 0xf4, 0x04, 0x7e, 0x9e, 0x99, 0x37, 0xef, 0xcd, 0x18, 0x96, 0x75, 0x53, 0xc9,
	0x6a, 0xd3, 0xf0, 0x52, 0xb2, 0xa2, 0xfb, 0xbc, 0x11, 0xbc, 0x79, 0xca, 0x0f, 0x9c, 0xea, 0xbb,
	0x68, 0xf1, 0x58, 0x55, 0x8f, 0x05, 0xdf, 0xe8, 0xd3, 0x43, 0x7b, 0xdc, 0xc8, 0xfc, 0xc4, 0x85,
	0x64, 0xa7, 0x1a, 0x03, 0x92, 0x9f, 0x36, 0x38, 0xa9, 0xce, 0x24, 0xb7, 0xe0, 0x61, 0x8d, 0x6f,
	0x79, 0x36, 0xb7, 0x62, 0x6b, 0x35, 0x49, 0xaf, 0x11, 0xd8, 0x67, 0x64, 0x01, 0x7e, 0x77, 0x99,
	0x31, 0xc9, 0xe7, 0x76, 0x6c, 0xad, 0xbc, 0x14, 0x10, 0xda, 0x31, 0xc9, 0xc9, 0x12, 0xa6, 0x79,
	0xf9, 0xc4, 0x4b, 0x59, 0x35, 0x3f, 0x54, 0x81, 0x89, 0x8e, 0xf0, 0xcf, 0x18, 0xd6, 0x38, 0xb4,
	0x42, 0x56, 0x27, 0xde, 0xa8, 0x88, 0x2b, 0xac, 0x61, 0x20, 0x43, 0x22, 0xdb, 0xa6, 0x44, 0x92,
	0xff, 0x0c, 0x89, 0x82, 0x34, 0xc9, 0xff, 0x70, 0x2d, 0x24, 0x3b, 0x1e, 0x55, 0xba, 0xa3, 0x6f,
	0x5d, 0x7d, 0xde, 0x67, 0xe4, 0x1d, 0x40, 0xc1, 0x84, 0xfc, 0x5a, 0xeb, 0x54, 0x37, 0xb6, 0x56,
	0xfe, 0x36, 0xa2, 0x28, 0x9f, 0x1a, 0xf9, 0xf4, 0x8b, 0x91, 0x9f, 0xf6, 0xa2, 0x93, 0x0f, 0x30,
	0xfb, 0xd8, 0x70, 0x26, 0x79, 0xca, 0xbf, 0xb7, 0x5c, 0x48, 0x72, 0x03, 0x13, 0x56, 0xe7, 0xda,
	0x04, 0x2f, 0x55, 0xbf, 0x64, 0x01, 0x0e, 0x8a, 0xd5, 0xd2, 0xfd, 0xad, 0x4b, 0xd1, 0xb5, 0xb4,
	0x83, 0x93, 0x7b, 0x08, 0x4c, 0x0d, 0x51, 0x57, 0xa5, 0xe0, 0x7f, 0x28, 0xf2, 0x02, 0xe0, 0x50,
	0xb5, 0xa5, 0x44, 0x87, 0x6c, 0x6d, 0xb1, 0xd7, 0x21, 0xfb, 0x2c, 0xb9, 0x03, 0x3f, 0xe5, 0x2c,
	0xbb, 0xdc, 0xc4, 0x5f, 0xf2, 0x3f, 0xc1, 0x14, 0xf3, 0x2f, 0x36, 0xf0, 0xaa, 0x3f, 0xe2, 0x91,
	0x90, 0xf3, 0xac, 0x95, 0x1d, 0x68, 0xcc, 0x3f, 0xd8, 0xf1, 0x1e, 0x02, 0x53, 0xe3, 0x62, 0x37,
	0x73, 0x70, 0x5b, 0x1d, 0x63, 0xb4, 0x98, 0x63, 0x72, 0x07, 0xb3, 0x1d, 0x2f, 0xf8, 0x73, 0x1d,
	0xdc, 0x8e, 0xa5, 0xf4, 0xb6, 0x55, 0xb1, 0x9b, 0xfc, 0xe7, 0xd8, 0x33, 0x1d, 0x73, 0x66, 0xef,
	0x8e, 0x49, 0x0c, 0x81, 0xf2, 0xf1, 0xbe, 0x28, 0x0c, 0x7d, 0x00, 0x76, 0xf7, 0x26, 0xbc, 0xd4,
	0xce, 0xb3, 0x64, 0x07, 0xe1, 0x39, 0xa2, 0x23, 0x18, 0x85, 0x90, 0x25, 0xb8, 0xd8, 0x8e, 0x98,
	0xdb, 0xf1, 0xa4, 0x6f, 0x91, 0xc1, 0xb7, 0xbf, 0x2c, 0x98, 0x21, 0xf6, 0x19, 0x1f, 0x2d, 0x59,
	0x83, 0x83, 0x4b, 0x44, 0x02, 0x3a, 0xd8, 0xc8, 0x28, 0xa4, 0xa3, 0xed, 0x7a, 0x09, 0x57, 0xaa,
	0x05, 0x32, 0xa5, 0xbd, 0x9d, 0x89, 0x66, 0x74, 0xb0, 0x01, 0x6b, 0x70, 0x70, 0x0a, 0x24, 0xa0,
	0x83, 0x91, 0x46, 0x21, 0x1d, 0x8d, 0x67, 0x0d, 0x0e, 0x5a, 0x46, 0x02, 0x3a, 0xf0, 0x3e, 0x0a,
	0xe9, 0xc8, 0xcb, 0xd7, 0xe0, 0x76, 0xea, 0x49, 0x48, 0x87, 0x4e, 0x45, 0x37, 0x74, 0x64, 0xcc,
	0x83, 0xa3, 0x1f, 0xdf, 0xdb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x84, 0x73, 0xea, 0xae,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RentalServiceClient is the client API for RentalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RentalServiceClient interface {
	//Create new rental
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	//Read rental
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	//Update rental
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	//Delete rental
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	//Read all rental
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type rentalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRentalServiceClient(cc grpc.ClientConnInterface) RentalServiceClient {
	return &rentalServiceClient{cc}
}

func (c *rentalServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/RentalService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/RentalService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/RentalService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/RentalService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/RentalService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RentalServiceServer is the server API for RentalService service.
type RentalServiceServer interface {
	//Create new rental
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	//Read rental
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	//Update rental
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	//Delete rental
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	//Read all rental
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedRentalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRentalServiceServer struct {
}

func (*UnimplementedRentalServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRentalServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedRentalServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedRentalServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRentalServiceServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterRentalServiceServer(s *grpc.Server, srv RentalServiceServer) {
	s.RegisterService(&_RentalService_serviceDesc, srv)
}

func _RentalService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RentalService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RentalService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RentalService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RentalService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RentalService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RentalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RentalService",
	HandlerType: (*RentalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RentalService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _RentalService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RentalService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RentalService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _RentalService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rental/rental-service.proto",
}
