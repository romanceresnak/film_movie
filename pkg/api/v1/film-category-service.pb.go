// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/film_category/film-category-service.proto

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

type FilmCategory struct {
	//id of film category id actor
	FilmCategoryId int64 `protobuf:"varint,1,opt,name=film_category_id,json=filmCategoryId,proto3" json:"film_category_id,omitempty"`
	//id of film
	FilmId int64 `protobuf:"varint,2,opt,name=film_id,json=filmId,proto3" json:"film_id,omitempty"`
	//category id as primary foreign key
	CategoryId int64 `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	//Time when the entity was update last time
	LastUpdate           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FilmCategory) Reset()         { *m = FilmCategory{} }
func (m *FilmCategory) String() string { return proto.CompactTextString(m) }
func (*FilmCategory) ProtoMessage()    {}
func (*FilmCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{0}
}

func (m *FilmCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilmCategory.Unmarshal(m, b)
}
func (m *FilmCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilmCategory.Marshal(b, m, deterministic)
}
func (m *FilmCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilmCategory.Merge(m, src)
}
func (m *FilmCategory) XXX_Size() int {
	return xxx_messageInfo_FilmCategory.Size(m)
}
func (m *FilmCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_FilmCategory.DiscardUnknown(m)
}

var xxx_messageInfo_FilmCategory proto.InternalMessageInfo

func (m *FilmCategory) GetFilmCategoryId() int64 {
	if m != nil {
		return m.FilmCategoryId
	}
	return 0
}

func (m *FilmCategory) GetFilmId() int64 {
	if m != nil {
		return m.FilmId
	}
	return 0
}

func (m *FilmCategory) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *FilmCategory) GetLastUpdate() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdate
	}
	return nil
}

//Request data to create new film category entity
type CreateRequest struct {
	//Api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Film category entity to add to database
	FilmCategory         *FilmCategory `protobuf:"bytes,2,opt,name=film_category,json=filmCategory,proto3" json:"film_category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{1}
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

func (m *CreateRequest) GetFilmCategory() *FilmCategory {
	if m != nil {
		return m.FilmCategory
	}
	return nil
}

//Response message of created film category
type CreateResponse struct {
	//APi version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//We get id of film category
	FilmActorId          int64    `protobuf:"varint,2,opt,name=film_actor_id,json=filmActorId,proto3" json:"film_actor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{2}
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

func (m *CreateResponse) GetFilmActorId() int64 {
	if m != nil {
		return m.FilmActorId
	}
	return 0
}

//Request message request for get data about film category
type ReadMessage struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Get film category by id
	FilmCategoryId       int64    `protobuf:"varint,2,opt,name=film_category_id,json=filmCategoryId,proto3" json:"film_category_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMessage) Reset()         { *m = ReadMessage{} }
func (m *ReadMessage) String() string { return proto.CompactTextString(m) }
func (*ReadMessage) ProtoMessage()    {}
func (*ReadMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{3}
}

func (m *ReadMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMessage.Unmarshal(m, b)
}
func (m *ReadMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMessage.Marshal(b, m, deterministic)
}
func (m *ReadMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMessage.Merge(m, src)
}
func (m *ReadMessage) XXX_Size() int {
	return xxx_messageInfo_ReadMessage.Size(m)
}
func (m *ReadMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMessage proto.InternalMessageInfo

func (m *ReadMessage) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadMessage) GetFilmCategoryId() int64 {
	if m != nil {
		return m.FilmCategoryId
	}
	return 0
}

//Response message about film category
type ReadResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Film category entity read by id
	FilmCategory         *FilmCategory `protobuf:"bytes,2,opt,name=film_category,json=filmCategory,proto3" json:"film_category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{4}
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

func (m *ReadResponse) GetFilmCategory() *FilmCategory {
	if m != nil {
		return m.FilmCategory
	}
	return nil
}

//Request for update film category
type UpdateRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//Update entity
	FilmCategory         *FilmCategory `protobuf:"bytes,2,opt,name=film_category,json=filmCategory,proto3" json:"film_category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{5}
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

func (m *UpdateRequest) GetFilmCategory() *FilmCategory {
	if m != nil {
		return m.FilmCategory
	}
	return nil
}

//Response to update film category
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
	return fileDescriptor_4c3cc50599255003, []int{6}
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

//Delete request for film category
type DeleteRequest struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Contains number of entities have been deleted
	// Equals 1 in case of successful delete
	Deleted              int64    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{7}
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

func (m *DeleteRequest) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

//Response to delete film category
type DeleteResponse struct {
	//api version
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{8}
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

//Request for read all film category
type ReadAllRequest struct {
	//api version
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{9}
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

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

//Response for get all the film category
type ReadAllResponse struct {
	//api version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	//List of all film category
	FilmCategory         []*FilmCategory `protobuf:"bytes,2,rep,name=film_category,json=filmCategory,proto3" json:"film_category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c3cc50599255003, []int{10}
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

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetFilmCategory() []*FilmCategory {
	if m != nil {
		return m.FilmCategory
	}
	return nil
}

func init() {
	proto.RegisterType((*FilmCategory)(nil), "FilmCategory")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*ReadMessage)(nil), "ReadMessage")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "ReadAllResponse")
}

func init() {
	proto.RegisterFile("proto/film_category/film-category-service.proto", fileDescriptor_4c3cc50599255003)
}

var fileDescriptor_4c3cc50599255003 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6b, 0x9c, 0x40,
	0x14, 0xc7, 0x31, 0x1b, 0x5c, 0xfa, 0x5c, 0xdd, 0x65, 0x7a, 0xa8, 0x78, 0x49, 0x18, 0x28, 0xec,
	0x42, 0x33, 0x0b, 0xf6, 0x98, 0x5e, 0x42, 0x4a, 0xc0, 0x43, 0x2f, 0x36, 0x4b, 0x8f, 0xcb, 0xc4,
	0x99, 0x15, 0x41, 0xab, 0x75, 0x66, 0x0b, 0xfd, 0x54, 0xfd, 0x60, 0xfd, 0x12, 0x65, 0x66, 0x9c,
	0x44, 0xad, 0x9b, 0xd2, 0x92, 0x9b, 0x4f, 0xff, 0xef, 0xef, 0xff, 0xfd, 0xe6, 0x0d, 0x6c, 0x9b,
	0xb6, 0x96, 0xf5, 0xf6, 0x50, 0x94, 0xd5, 0x3e, 0xa3, 0x92, 0xe7, 0x75, 0xfb, 0x43, 0x57, 0x57,
	0xb6, 0xba, 0x12, 0xbc, 0xfd, 0x5e, 0x64, 0x9c, 0x68, 0x65, 0x74, 0x91, 0xd7, 0x75, 0x5e, 0x72,
	0xd3, 0xf7, 0x70, 0x3c, 0x6c, 0x65, 0x51, 0x71, 0x21, 0x69, 0xd5, 0x18, 0x01, 0xfe, 0xe9, 0xc0,
	0xe2, 0xae, 0x28, 0xab, 0xdb, 0xae, 0x1f, 0xad, 0x61, 0x35, 0xb0, 0xdf, 0x17, 0x2c, 0x74, 0x2e,
	0x9d, 0xf5, 0x2c, 0x0d, 0x0e, 0x3d, 0x5d, 0xc2, 0xd0, 0x1b, 0x98, 0x6b, 0x65, 0xc1, 0xc2, 0x33,
	0x2d, 0x70, 0x55, 0x99, 0x30, 0x74, 0x01, 0x5e, 0xbf, 0x7b, 0xa6, 0x3f, 0x42, 0xf6, 0xd4, 0x79,
	0x0d, 0x5e, 0x49, 0x85, 0xdc, 0x1f, 0x1b, 0x46, 0x25, 0x0f, 0xcf, 0x2f, 0x9d, 0xb5, 0x17, 0x47,
	0xc4, 0x64, 0x25, 0x36, 0x2b, 0xb9, 0xb7, 0x59, 0x53, 0x50, 0xf2, 0x9d, 0x56, 0xe3, 0x1d, 0xf8,
	0xb7, 0x2d, 0xa7, 0x92, 0xa7, 0xfc, 0xdb, 0x91, 0x0b, 0x89, 0x56, 0x30, 0xa3, 0x4d, 0xa1, 0x43,
	0xbe, 0x4a, 0xd5, 0x23, 0x8a, 0xc1, 0x1f, 0xcc, 0xa0, 0xf3, 0x79, 0xb1, 0x4f, 0xfa, 0x93, 0xa6,
	0x8b, 0xfe, 0x3c, 0xf8, 0x0e, 0x02, 0x6b, 0x2b, 0x9a, 0xfa, 0xab, 0xe0, 0x13, 0xbe, 0xb8, 0xf3,
	0xa5, 0x99, 0xac, 0xdb, 0xa7, 0xb9, 0x3d, 0xf5, 0xf2, 0x46, 0xbd, 0x4b, 0x18, 0x4e, 0xc0, 0x4b,
	0x39, 0x65, 0x9f, 0xb8, 0x10, 0x34, 0x9f, 0x32, 0x99, 0x02, 0x7c, 0x36, 0x05, 0x18, 0xdf, 0xc3,
	0x42, 0x59, 0x3d, 0x13, 0xe8, 0x7f, 0x06, 0xdd, 0x81, 0x6f, 0x48, 0xbe, 0x2c, 0xbf, 0x0f, 0x10,
	0x58, 0xdb, 0x93, 0x71, 0x43, 0x98, 0x9b, 0x23, 0xb7, 0x13, 0xdb, 0x12, 0x5f, 0x83, 0xff, 0x91,
	0x97, 0xfc, 0xb9, 0x50, 0x21, 0xcc, 0x99, 0x96, 0x3c, 0x36, 0x77, 0x25, 0xc6, 0x10, 0xd8, 0xe6,
	0x53, 0xbf, 0x56, 0x1a, 0xc5, 0xf2, 0xa6, 0x2c, 0x4f, 0xfe, 0x01, 0x7f, 0x81, 0xe5, 0xa3, 0xe6,
	0x5f, 0x90, 0xcf, 0xfe, 0xc2, 0x26, 0xfe, 0xe5, 0xc0, 0xeb, 0xfe, 0xe7, 0xcf, 0xe6, 0x8e, 0xa2,
	0x0d, 0xb8, 0x66, 0xe7, 0x50, 0x40, 0x06, 0x3b, 0x1d, 0x2d, 0xc9, 0x68, 0x19, 0xdf, 0xc2, 0xb9,
	0xca, 0x86, 0x16, 0xa4, 0xb7, 0x5d, 0x7f, 0xca, 0x36, 0xe0, 0x9a, 0x53, 0x98, 0x70, 0x1c, 0x1d,
	0xcf, 0x06, 0x5c, 0x43, 0x0d, 0x05, 0x64, 0xc0, 0x3e, 0x5a, 0x92, 0x11, 0xce, 0x77, 0x30, 0xef,
	0xc0, 0xa0, 0x25, 0x19, 0x62, 0x8c, 0x56, 0x64, 0xc4, 0xec, 0xc1, 0xd5, 0x17, 0xf8, 0xfd, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x67, 0xf8, 0x0a, 0xad, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FilmCategoryServiceClient is the client API for FilmCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilmCategoryServiceClient interface {
	//Create new film category
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	//Read film category
	Read(ctx context.Context, in *ReadMessage, opts ...grpc.CallOption) (*CreateResponse, error)
	//Update film category
	Update(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	//Delete film category
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	//Read all film category
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type filmCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilmCategoryServiceClient(cc grpc.ClientConnInterface) FilmCategoryServiceClient {
	return &filmCategoryServiceClient{cc}
}

func (c *filmCategoryServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/FilmCategoryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmCategoryServiceClient) Read(ctx context.Context, in *ReadMessage, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/FilmCategoryService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmCategoryServiceClient) Update(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/FilmCategoryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmCategoryServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/FilmCategoryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filmCategoryServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/FilmCategoryService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilmCategoryServiceServer is the server API for FilmCategoryService service.
type FilmCategoryServiceServer interface {
	//Create new film category
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	//Read film category
	Read(context.Context, *ReadMessage) (*CreateResponse, error)
	//Update film category
	Update(context.Context, *CreateRequest) (*UpdateResponse, error)
	//Delete film category
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	//Read all film category
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
}

// UnimplementedFilmCategoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFilmCategoryServiceServer struct {
}

func (*UnimplementedFilmCategoryServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedFilmCategoryServiceServer) Read(ctx context.Context, req *ReadMessage) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedFilmCategoryServiceServer) Update(ctx context.Context, req *CreateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedFilmCategoryServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedFilmCategoryServiceServer) ReadAll(ctx context.Context, req *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}

func RegisterFilmCategoryServiceServer(s *grpc.Server, srv FilmCategoryServiceServer) {
	s.RegisterService(&_FilmCategoryService_serviceDesc, srv)
}

func _FilmCategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmCategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FilmCategoryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmCategoryServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmCategoryService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmCategoryServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FilmCategoryService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmCategoryServiceServer).Read(ctx, req.(*ReadMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmCategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmCategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FilmCategoryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmCategoryServiceServer).Update(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmCategoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmCategoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FilmCategoryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmCategoryServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilmCategoryService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilmCategoryServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FilmCategoryService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilmCategoryServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FilmCategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FilmCategoryService",
	HandlerType: (*FilmCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FilmCategoryService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _FilmCategoryService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FilmCategoryService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FilmCategoryService_Delete_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _FilmCategoryService_ReadAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/film_category/film-category-service.proto",
}