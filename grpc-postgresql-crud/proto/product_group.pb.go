// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product_group.proto

package protopb

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

type ProductGroup struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductGroup) Reset()         { *m = ProductGroup{} }
func (m *ProductGroup) String() string { return proto.CompactTextString(m) }
func (*ProductGroup) ProtoMessage()    {}
func (*ProductGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{0}
}

func (m *ProductGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductGroup.Unmarshal(m, b)
}
func (m *ProductGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductGroup.Marshal(b, m, deterministic)
}
func (m *ProductGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductGroup.Merge(m, src)
}
func (m *ProductGroup) XXX_Size() int {
	return xxx_messageInfo_ProductGroup.Size(m)
}
func (m *ProductGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ProductGroup proto.InternalMessageInfo

func (m *ProductGroup) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductGroup) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ProductGroup) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CreateProductReq struct {
	Product              *ProductGroup `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateProductReq) Reset()         { *m = CreateProductReq{} }
func (m *CreateProductReq) String() string { return proto.CompactTextString(m) }
func (*CreateProductReq) ProtoMessage()    {}
func (*CreateProductReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{1}
}

func (m *CreateProductReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductReq.Unmarshal(m, b)
}
func (m *CreateProductReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductReq.Marshal(b, m, deterministic)
}
func (m *CreateProductReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductReq.Merge(m, src)
}
func (m *CreateProductReq) XXX_Size() int {
	return xxx_messageInfo_CreateProductReq.Size(m)
}
func (m *CreateProductReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductReq proto.InternalMessageInfo

func (m *CreateProductReq) GetProduct() *ProductGroup {
	if m != nil {
		return m.Product
	}
	return nil
}

type CreateProductRes struct {
	Product              *ProductGroup `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateProductRes) Reset()         { *m = CreateProductRes{} }
func (m *CreateProductRes) String() string { return proto.CompactTextString(m) }
func (*CreateProductRes) ProtoMessage()    {}
func (*CreateProductRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{2}
}

func (m *CreateProductRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductRes.Unmarshal(m, b)
}
func (m *CreateProductRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductRes.Marshal(b, m, deterministic)
}
func (m *CreateProductRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductRes.Merge(m, src)
}
func (m *CreateProductRes) XXX_Size() int {
	return xxx_messageInfo_CreateProductRes.Size(m)
}
func (m *CreateProductRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductRes proto.InternalMessageInfo

func (m *CreateProductRes) GetProduct() *ProductGroup {
	if m != nil {
		return m.Product
	}
	return nil
}

type ReadProductReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadProductReq) Reset()         { *m = ReadProductReq{} }
func (m *ReadProductReq) String() string { return proto.CompactTextString(m) }
func (*ReadProductReq) ProtoMessage()    {}
func (*ReadProductReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{3}
}

func (m *ReadProductReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadProductReq.Unmarshal(m, b)
}
func (m *ReadProductReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadProductReq.Marshal(b, m, deterministic)
}
func (m *ReadProductReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadProductReq.Merge(m, src)
}
func (m *ReadProductReq) XXX_Size() int {
	return xxx_messageInfo_ReadProductReq.Size(m)
}
func (m *ReadProductReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadProductReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadProductReq proto.InternalMessageInfo

func (m *ReadProductReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadProductRes struct {
	Product              *ProductGroup `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadProductRes) Reset()         { *m = ReadProductRes{} }
func (m *ReadProductRes) String() string { return proto.CompactTextString(m) }
func (*ReadProductRes) ProtoMessage()    {}
func (*ReadProductRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{4}
}

func (m *ReadProductRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadProductRes.Unmarshal(m, b)
}
func (m *ReadProductRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadProductRes.Marshal(b, m, deterministic)
}
func (m *ReadProductRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadProductRes.Merge(m, src)
}
func (m *ReadProductRes) XXX_Size() int {
	return xxx_messageInfo_ReadProductRes.Size(m)
}
func (m *ReadProductRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadProductRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadProductRes proto.InternalMessageInfo

func (m *ReadProductRes) GetProduct() *ProductGroup {
	if m != nil {
		return m.Product
	}
	return nil
}

type UpdateProductReq struct {
	Product              *ProductGroup `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateProductReq) Reset()         { *m = UpdateProductReq{} }
func (m *UpdateProductReq) String() string { return proto.CompactTextString(m) }
func (*UpdateProductReq) ProtoMessage()    {}
func (*UpdateProductReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{5}
}

func (m *UpdateProductReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProductReq.Unmarshal(m, b)
}
func (m *UpdateProductReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProductReq.Marshal(b, m, deterministic)
}
func (m *UpdateProductReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProductReq.Merge(m, src)
}
func (m *UpdateProductReq) XXX_Size() int {
	return xxx_messageInfo_UpdateProductReq.Size(m)
}
func (m *UpdateProductReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProductReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProductReq proto.InternalMessageInfo

func (m *UpdateProductReq) GetProduct() *ProductGroup {
	if m != nil {
		return m.Product
	}
	return nil
}

type UpdateProductRes struct {
	Product              *ProductGroup `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateProductRes) Reset()         { *m = UpdateProductRes{} }
func (m *UpdateProductRes) String() string { return proto.CompactTextString(m) }
func (*UpdateProductRes) ProtoMessage()    {}
func (*UpdateProductRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{6}
}

func (m *UpdateProductRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProductRes.Unmarshal(m, b)
}
func (m *UpdateProductRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProductRes.Marshal(b, m, deterministic)
}
func (m *UpdateProductRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProductRes.Merge(m, src)
}
func (m *UpdateProductRes) XXX_Size() int {
	return xxx_messageInfo_UpdateProductRes.Size(m)
}
func (m *UpdateProductRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProductRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProductRes proto.InternalMessageInfo

func (m *UpdateProductRes) GetProduct() *ProductGroup {
	if m != nil {
		return m.Product
	}
	return nil
}

type DeleteProductReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductReq) Reset()         { *m = DeleteProductReq{} }
func (m *DeleteProductReq) String() string { return proto.CompactTextString(m) }
func (*DeleteProductReq) ProtoMessage()    {}
func (*DeleteProductReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{7}
}

func (m *DeleteProductReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductReq.Unmarshal(m, b)
}
func (m *DeleteProductReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductReq.Marshal(b, m, deterministic)
}
func (m *DeleteProductReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductReq.Merge(m, src)
}
func (m *DeleteProductReq) XXX_Size() int {
	return xxx_messageInfo_DeleteProductReq.Size(m)
}
func (m *DeleteProductReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductReq proto.InternalMessageInfo

func (m *DeleteProductReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteProductRes struct {
	Del                  bool     `protobuf:"varint,1,opt,name=del,proto3" json:"del,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductRes) Reset()         { *m = DeleteProductRes{} }
func (m *DeleteProductRes) String() string { return proto.CompactTextString(m) }
func (*DeleteProductRes) ProtoMessage()    {}
func (*DeleteProductRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{8}
}

func (m *DeleteProductRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductRes.Unmarshal(m, b)
}
func (m *DeleteProductRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductRes.Marshal(b, m, deterministic)
}
func (m *DeleteProductRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductRes.Merge(m, src)
}
func (m *DeleteProductRes) XXX_Size() int {
	return xxx_messageInfo_DeleteProductRes.Size(m)
}
func (m *DeleteProductRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductRes proto.InternalMessageInfo

func (m *DeleteProductRes) GetDel() bool {
	if m != nil {
		return m.Del
	}
	return false
}

type ListProductReq struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProductReq) Reset()         { *m = ListProductReq{} }
func (m *ListProductReq) String() string { return proto.CompactTextString(m) }
func (*ListProductReq) ProtoMessage()    {}
func (*ListProductReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{9}
}

func (m *ListProductReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductReq.Unmarshal(m, b)
}
func (m *ListProductReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductReq.Marshal(b, m, deterministic)
}
func (m *ListProductReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductReq.Merge(m, src)
}
func (m *ListProductReq) XXX_Size() int {
	return xxx_messageInfo_ListProductReq.Size(m)
}
func (m *ListProductReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductReq proto.InternalMessageInfo

func (m *ListProductReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListProductReq) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListProductRes struct {
	Products             []*ProductGroup `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListProductRes) Reset()         { *m = ListProductRes{} }
func (m *ListProductRes) String() string { return proto.CompactTextString(m) }
func (*ListProductRes) ProtoMessage()    {}
func (*ListProductRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca955ae7f1b8fb1c, []int{10}
}

func (m *ListProductRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductRes.Unmarshal(m, b)
}
func (m *ListProductRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductRes.Marshal(b, m, deterministic)
}
func (m *ListProductRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductRes.Merge(m, src)
}
func (m *ListProductRes) XXX_Size() int {
	return xxx_messageInfo_ListProductRes.Size(m)
}
func (m *ListProductRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductRes proto.InternalMessageInfo

func (m *ListProductRes) GetProducts() []*ProductGroup {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductGroup)(nil), "proto.ProductGroup")
	proto.RegisterType((*CreateProductReq)(nil), "proto.CreateProductReq")
	proto.RegisterType((*CreateProductRes)(nil), "proto.CreateProductRes")
	proto.RegisterType((*ReadProductReq)(nil), "proto.ReadProductReq")
	proto.RegisterType((*ReadProductRes)(nil), "proto.ReadProductRes")
	proto.RegisterType((*UpdateProductReq)(nil), "proto.UpdateProductReq")
	proto.RegisterType((*UpdateProductRes)(nil), "proto.UpdateProductRes")
	proto.RegisterType((*DeleteProductReq)(nil), "proto.DeleteProductReq")
	proto.RegisterType((*DeleteProductRes)(nil), "proto.DeleteProductRes")
	proto.RegisterType((*ListProductReq)(nil), "proto.ListProductReq")
	proto.RegisterType((*ListProductRes)(nil), "proto.ListProductRes")
}

func init() { proto.RegisterFile("proto/product_group.proto", fileDescriptor_ca955ae7f1b8fb1c) }

var fileDescriptor_ca955ae7f1b8fb1c = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcf, 0x4b, 0xc3, 0x30,
	0x18, 0x5d, 0xbb, 0x75, 0x3f, 0x3e, 0xa5, 0x94, 0xf8, 0xab, 0x16, 0x84, 0x12, 0x3c, 0xec, 0xe2,
	0x06, 0xf3, 0xae, 0x74, 0x0a, 0xee, 0xe0, 0x41, 0x22, 0x5e, 0xbc, 0x8c, 0x6e, 0xc9, 0xa4, 0xb0,
	0xd9, 0xda, 0x64, 0xfe, 0xe9, 0x9e, 0xa5, 0x69, 0x26, 0x4d, 0x6c, 0x0f, 0xd3, 0xd3, 0x92, 0xef,
	0xbd, 0xbd, 0xbc, 0xd7, 0xef, 0xc1, 0x79, 0x96, 0xa7, 0x22, 0x1d, 0x67, 0x79, 0x4a, 0xb7, 0x4b,
	0x31, 0x7f, 0xcb, 0xd3, 0x6d, 0x36, 0x92, 0x33, 0xe4, 0xc8, 0x1f, 0x9c, 0xc1, 0xe1, 0x53, 0x89,
	0x3e, 0x14, 0x20, 0x72, 0xc1, 0x4e, 0xa8, 0x6f, 0x85, 0xd6, 0xd0, 0x21, 0x76, 0x42, 0x11, 0x82,
	0xce, 0x7b, 0xbc, 0x61, 0xbe, 0x1d, 0x5a, 0xc3, 0x01, 0x91, 0x67, 0x74, 0x01, 0xb0, 0xcc, 0x59,
	0x2c, 0x18, 0x9d, 0xc7, 0xc2, 0x6f, 0x4b, 0x64, 0xa0, 0x26, 0x91, 0x28, 0xe0, 0x6d, 0x46, 0x77,
	0x70, 0xa7, 0x84, 0xd5, 0x24, 0x12, 0x38, 0x02, 0xef, 0x4e, 0x72, 0xd5, 0xbb, 0x84, 0x7d, 0xa0,
	0x2b, 0xe8, 0x29, 0x8f, 0xf2, 0xe9, 0x83, 0xc9, 0x51, 0xe9, 0x72, 0x54, 0xf5, 0x46, 0x76, 0x9c,
	0x1a, 0x09, 0xbe, 0xaf, 0x44, 0x08, 0x2e, 0x61, 0x31, 0xad, 0x78, 0x30, 0x92, 0xe3, 0x5b, 0x83,
	0xc1, 0xff, 0xe0, 0xf2, 0x45, 0xa6, 0xfe, 0x57, 0x50, 0x43, 0x62, 0x6f, 0x17, 0x18, 0xbc, 0x7b,
	0xb6, 0x66, 0x9a, 0x0b, 0x33, 0xea, 0xe5, 0x2f, 0x0e, 0x47, 0x1e, 0xb4, 0x29, 0x5b, 0x4b, 0x52,
	0x9f, 0x14, 0x47, 0x7c, 0x03, 0xee, 0x63, 0xc2, 0x45, 0x45, 0xe7, 0x18, 0x9c, 0x75, 0xb2, 0x49,
	0x84, 0x92, 0x2a, 0x2f, 0xe8, 0x14, 0xba, 0xe9, 0x6a, 0xc5, 0x99, 0x90, 0xa5, 0x71, 0x88, 0xba,
	0xe1, 0xc8, 0xf8, 0x3f, 0x47, 0x63, 0xe8, 0x2b, 0x9b, 0xdc, 0xb7, 0xc2, 0x76, 0x53, 0x96, 0x1f,
	0xd2, 0xe4, 0xcb, 0x06, 0x57, 0x41, 0xcf, 0x2c, 0xff, 0x4c, 0x96, 0x0c, 0xcd, 0x00, 0x69, 0x5d,
	0x28, 0x6b, 0x7c, 0xa6, 0x74, 0xcc, 0xa6, 0x05, 0x0d, 0x00, 0xc7, 0x2d, 0x34, 0x05, 0xaf, 0xb2,
	0xf0, 0x52, 0xe7, 0x44, 0xd1, 0xf5, 0xae, 0x04, 0xb5, 0xe3, 0x42, 0x63, 0x06, 0x48, 0x5b, 0x98,
	0xee, 0xc6, 0xac, 0x43, 0xd0, 0x00, 0x28, 0x25, 0x6d, 0x27, 0xba, 0x92, 0xb9, 0xd2, 0xa0, 0x01,
	0x50, 0xb9, 0x2a, 0xdf, 0x5d, 0xcf, 0xa5, 0x2f, 0x34, 0xa8, 0x1d, 0x73, 0xdc, 0x9a, 0x0e, 0x5e,
	0x7b, 0x12, 0xc9, 0x16, 0x8b, 0xae, 0x3c, 0x5c, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xc2,
	0xa9, 0xe4, 0x5c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProductGroup(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductRes, error)
	ReadProductGroup(ctx context.Context, in *ReadProductReq, opts ...grpc.CallOption) (*ReadProductRes, error)
	UpdateProductGroup(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductRes, error)
	DeleteProductGroup(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductRes, error)
	ListProductGroup(ctx context.Context, in *ListProductReq, opts ...grpc.CallOption) (*ListProductRes, error)
}

type productServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductServiceClient(cc *grpc.ClientConn) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProductGroup(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductRes, error) {
	out := new(CreateProductRes)
	err := c.cc.Invoke(ctx, "/proto.ProductService/CreateProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ReadProductGroup(ctx context.Context, in *ReadProductReq, opts ...grpc.CallOption) (*ReadProductRes, error) {
	out := new(ReadProductRes)
	err := c.cc.Invoke(ctx, "/proto.ProductService/ReadProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProductGroup(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductRes, error) {
	out := new(UpdateProductRes)
	err := c.cc.Invoke(ctx, "/proto.ProductService/UpdateProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProductGroup(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductRes, error) {
	out := new(DeleteProductRes)
	err := c.cc.Invoke(ctx, "/proto.ProductService/DeleteProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProductGroup(ctx context.Context, in *ListProductReq, opts ...grpc.CallOption) (*ListProductRes, error) {
	out := new(ListProductRes)
	err := c.cc.Invoke(ctx, "/proto.ProductService/ListProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	CreateProductGroup(context.Context, *CreateProductReq) (*CreateProductRes, error)
	ReadProductGroup(context.Context, *ReadProductReq) (*ReadProductRes, error)
	UpdateProductGroup(context.Context, *UpdateProductReq) (*UpdateProductRes, error)
	DeleteProductGroup(context.Context, *DeleteProductReq) (*DeleteProductRes, error)
	ListProductGroup(context.Context, *ListProductReq) (*ListProductRes, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) CreateProductGroup(ctx context.Context, req *CreateProductReq) (*CreateProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductGroup not implemented")
}
func (*UnimplementedProductServiceServer) ReadProductGroup(ctx context.Context, req *ReadProductReq) (*ReadProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProductGroup not implemented")
}
func (*UnimplementedProductServiceServer) UpdateProductGroup(ctx context.Context, req *UpdateProductReq) (*UpdateProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductGroup not implemented")
}
func (*UnimplementedProductServiceServer) DeleteProductGroup(ctx context.Context, req *DeleteProductReq) (*DeleteProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductGroup not implemented")
}
func (*UnimplementedProductServiceServer) ListProductGroup(ctx context.Context, req *ListProductReq) (*ListProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductGroup not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_CreateProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/CreateProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProductGroup(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ReadProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ReadProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/ReadProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ReadProductGroup(ctx, req.(*ReadProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/UpdateProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProductGroup(ctx, req.(*UpdateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/DeleteProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProductGroup(ctx, req.(*DeleteProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/ListProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProductGroup(ctx, req.(*ListProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProductGroup",
			Handler:    _ProductService_CreateProductGroup_Handler,
		},
		{
			MethodName: "ReadProductGroup",
			Handler:    _ProductService_ReadProductGroup_Handler,
		},
		{
			MethodName: "UpdateProductGroup",
			Handler:    _ProductService_UpdateProductGroup_Handler,
		},
		{
			MethodName: "DeleteProductGroup",
			Handler:    _ProductService_DeleteProductGroup_Handler,
		},
		{
			MethodName: "ListProductGroup",
			Handler:    _ProductService_ListProductGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/product_group.proto",
}