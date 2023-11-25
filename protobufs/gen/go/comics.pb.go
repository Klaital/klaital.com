// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: comics.proto

package protobufs

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateComicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64      `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	NewState *ComicState `protobuf:"bytes,2,opt,name=NewState,proto3" json:"NewState,omitempty"`
}

func (x *UpdateComicRequest) Reset() {
	*x = UpdateComicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateComicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateComicRequest) ProtoMessage() {}

func (x *UpdateComicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateComicRequest.ProtoReflect.Descriptor instead.
func (*UpdateComicRequest) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateComicRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateComicRequest) GetNewState() *ComicState {
	if x != nil {
		return x.NewState
	}
	return nil
}

type GetComicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetComicsRequest) Reset() {
	*x = GetComicsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetComicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetComicsRequest) ProtoMessage() {}

func (x *GetComicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetComicsRequest.ProtoReflect.Descriptor instead.
func (*GetComicsRequest) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{1}
}

func (x *GetComicsRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ComicState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComicId           uint64 `protobuf:"varint,1,opt,name=ComicId,proto3" json:"ComicId,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	HomeUrl           string `protobuf:"bytes,3,opt,name=HomeUrl,proto3" json:"HomeUrl,omitempty"`
	BookmarkUrl       string `protobuf:"bytes,4,opt,name=BookmarkUrl,proto3" json:"BookmarkUrl,omitempty"`
	LastReadTimestamp string `protobuf:"bytes,5,opt,name=LastReadTimestamp,proto3" json:"LastReadTimestamp,omitempty"`
	RssFeedUrl        string `protobuf:"bytes,6,opt,name=RssFeedUrl,proto3" json:"RssFeedUrl,omitempty"`
}

func (x *ComicState) Reset() {
	*x = ComicState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComicState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComicState) ProtoMessage() {}

func (x *ComicState) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComicState.ProtoReflect.Descriptor instead.
func (*ComicState) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{2}
}

func (x *ComicState) GetComicId() uint64 {
	if x != nil {
		return x.ComicId
	}
	return 0
}

func (x *ComicState) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComicState) GetHomeUrl() string {
	if x != nil {
		return x.HomeUrl
	}
	return ""
}

func (x *ComicState) GetBookmarkUrl() string {
	if x != nil {
		return x.BookmarkUrl
	}
	return ""
}

func (x *ComicState) GetLastReadTimestamp() string {
	if x != nil {
		return x.LastReadTimestamp
	}
	return ""
}

func (x *ComicState) GetRssFeedUrl() string {
	if x != nil {
		return x.RssFeedUrl
	}
	return ""
}

type GetComicsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comics []*ComicState `protobuf:"bytes,1,rep,name=Comics,proto3" json:"Comics,omitempty"`
}

func (x *GetComicsResponse) Reset() {
	*x = GetComicsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetComicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetComicsResponse) ProtoMessage() {}

func (x *GetComicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetComicsResponse.ProtoReflect.Descriptor instead.
func (*GetComicsResponse) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{3}
}

func (x *GetComicsResponse) GetComics() []*ComicState {
	if x != nil {
		return x.Comics
	}
	return nil
}

type MarkReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ComicId uint64  `protobuf:"varint,2,opt,name=ComicId,proto3" json:"ComicId,omitempty"`
	ReadAt  *string `protobuf:"bytes,3,opt,name=ReadAt,proto3,oneof" json:"ReadAt,omitempty"` // If omitted, use current time
}

func (x *MarkReadRequest) Reset() {
	*x = MarkReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkReadRequest) ProtoMessage() {}

func (x *MarkReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkReadRequest.ProtoReflect.Descriptor instead.
func (*MarkReadRequest) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{4}
}

func (x *MarkReadRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *MarkReadRequest) GetComicId() uint64 {
	if x != nil {
		return x.ComicId
	}
	return 0
}

func (x *MarkReadRequest) GetReadAt() string {
	if x != nil && x.ReadAt != nil {
		return *x.ReadAt
	}
	return ""
}

type MarkItemReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ComicId uint64  `protobuf:"varint,2,opt,name=ComicId,proto3" json:"ComicId,omitempty"`
	ReadAt  *string `protobuf:"bytes,3,opt,name=ReadAt,proto3,oneof" json:"ReadAt,omitempty"` // If omitted, use current time
	ItemID  uint64  `protobuf:"varint,4,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
}

func (x *MarkItemReadRequest) Reset() {
	*x = MarkItemReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkItemReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkItemReadRequest) ProtoMessage() {}

func (x *MarkItemReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkItemReadRequest.ProtoReflect.Descriptor instead.
func (*MarkItemReadRequest) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{5}
}

func (x *MarkItemReadRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *MarkItemReadRequest) GetComicId() uint64 {
	if x != nil {
		return x.ComicId
	}
	return 0
}

func (x *MarkItemReadRequest) GetReadAt() string {
	if x != nil && x.ReadAt != nil {
		return *x.ReadAt
	}
	return ""
}

func (x *MarkItemReadRequest) GetItemID() uint64 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

type MarkReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkReadResponse) Reset() {
	*x = MarkReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkReadResponse) ProtoMessage() {}

func (x *MarkReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkReadResponse.ProtoReflect.Descriptor instead.
func (*MarkReadResponse) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{6}
}

type DescribeComicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64   `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ComicIDs []uint64 `protobuf:"varint,2,rep,packed,name=ComicIDs,proto3" json:"ComicIDs,omitempty"`
}

func (x *DescribeComicsRequest) Reset() {
	*x = DescribeComicsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeComicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeComicsRequest) ProtoMessage() {}

func (x *DescribeComicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeComicsRequest.ProtoReflect.Descriptor instead.
func (*DescribeComicsRequest) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeComicsRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DescribeComicsRequest) GetComicIDs() []uint64 {
	if x != nil {
		return x.ComicIDs
	}
	return nil
}

type ComicDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComicId           uint64     `protobuf:"varint,1,opt,name=ComicId,proto3" json:"ComicId,omitempty"`
	Name              string     `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	HomeUrl           string     `protobuf:"bytes,3,opt,name=HomeUrl,proto3" json:"HomeUrl,omitempty"`
	BookmarkUrl       string     `protobuf:"bytes,4,opt,name=BookmarkUrl,proto3" json:"BookmarkUrl,omitempty"`
	LastReadTimestamp string     `protobuf:"bytes,5,opt,name=LastReadTimestamp,proto3" json:"LastReadTimestamp,omitempty"`
	RssFeedUrl        string     `protobuf:"bytes,6,opt,name=RssFeedUrl,proto3" json:"RssFeedUrl,omitempty"`
	Feed              []*RssItem `protobuf:"bytes,7,rep,name=Feed,proto3" json:"Feed,omitempty"`
}

func (x *ComicDetails) Reset() {
	*x = ComicDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComicDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComicDetails) ProtoMessage() {}

func (x *ComicDetails) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComicDetails.ProtoReflect.Descriptor instead.
func (*ComicDetails) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{8}
}

func (x *ComicDetails) GetComicId() uint64 {
	if x != nil {
		return x.ComicId
	}
	return 0
}

func (x *ComicDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComicDetails) GetHomeUrl() string {
	if x != nil {
		return x.HomeUrl
	}
	return ""
}

func (x *ComicDetails) GetBookmarkUrl() string {
	if x != nil {
		return x.BookmarkUrl
	}
	return ""
}

func (x *ComicDetails) GetLastReadTimestamp() string {
	if x != nil {
		return x.LastReadTimestamp
	}
	return ""
}

func (x *ComicDetails) GetRssFeedUrl() string {
	if x != nil {
		return x.RssFeedUrl
	}
	return ""
}

func (x *ComicDetails) GetFeed() []*RssItem {
	if x != nil {
		return x.Feed
	}
	return nil
}

type RssItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Link   string `protobuf:"bytes,2,opt,name=Link,proto3" json:"Link,omitempty"`
	Guid   string `protobuf:"bytes,3,opt,name=Guid,proto3" json:"Guid,omitempty"`
	IsRead bool   `protobuf:"varint,4,opt,name=IsRead,proto3" json:"IsRead,omitempty"`
	Title  string `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *RssItem) Reset() {
	*x = RssItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RssItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RssItem) ProtoMessage() {}

func (x *RssItem) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RssItem.ProtoReflect.Descriptor instead.
func (*RssItem) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{9}
}

func (x *RssItem) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RssItem) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *RssItem) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *RssItem) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *RssItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type DescribeComicsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comics []*ComicDetails `protobuf:"bytes,1,rep,name=Comics,proto3" json:"Comics,omitempty"`
}

func (x *DescribeComicsResponse) Reset() {
	*x = DescribeComicsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeComicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeComicsResponse) ProtoMessage() {}

func (x *DescribeComicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeComicsResponse.ProtoReflect.Descriptor instead.
func (*DescribeComicsResponse) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{10}
}

func (x *DescribeComicsResponse) GetComics() []*ComicDetails {
	if x != nil {
		return x.Comics
	}
	return nil
}

type RefreshRssFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshRssFeedResponse) Reset() {
	*x = RefreshRssFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comics_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshRssFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRssFeedResponse) ProtoMessage() {}

func (x *RefreshRssFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comics_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRssFeedResponse.ProtoReflect.Descriptor instead.
func (*RefreshRssFeedResponse) Descriptor() ([]byte, []int) {
	return file_comics_proto_rawDescGZIP(), []int{11}
}

var File_comics_proto protoreflect.FileDescriptor

var file_comics_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x4e,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x43, 0x6f, 0x6d, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x08, 0x4e, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xc4, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x48, 0x6f, 0x6d, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x48, 0x6f, 0x6d, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x4c, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x73, 0x73, 0x46, 0x65,
	0x65, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x73, 0x73,
	0x46, 0x65, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x38, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06,
	0x43, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43,
	0x6f, 0x6d, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x43, 0x6f, 0x6d, 0x69, 0x63,
	0x73, 0x22, 0x6b, 0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x43,
	0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x22, 0x87,
	0x01, 0x0a, 0x13, 0x4d, 0x61, 0x72, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x52, 0x65, 0x61, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x0a, 0x15,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x08, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x44, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x0c, 0x43, 0x6f,
	0x6d, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x69, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x43, 0x6f, 0x6d,
	0x69, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f, 0x6d, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48, 0x6f, 0x6d, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x55, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x55, 0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x73, 0x73, 0x46, 0x65, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x73, 0x73, 0x46, 0x65, 0x65, 0x64, 0x55,
	0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x52, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x46, 0x65, 0x65, 0x64,
	0x22, 0x6f, 0x0a, 0x07, 0x52, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4c,
	0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x47, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x47,
	0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x3f, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x43,
	0x6f, 0x6d, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f,
	0x6d, 0x69, 0x63, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x06, 0x43, 0x6f, 0x6d, 0x69,
	0x63, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x73, 0x73,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb1, 0x04, 0x0a,
	0x06, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x12, 0x47, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x69, 0x63, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73,
	0x12, 0x5a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x46, 0x65, 0x65, 0x64,
	0x73, 0x12, 0x16, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x12, 0x61, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x12, 0x13, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x1a,
	0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x7b, 0x4e, 0x65,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x64, 0x7d, 0x12,
	0x58, 0x0a, 0x0a, 0x4d, 0x61, 0x72, 0x6b, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x7b, 0x43, 0x6f, 0x6d, 0x69,
	0x63, 0x49, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x6b, 0x0a, 0x0c, 0x4d, 0x61, 0x72,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x1a, 0x27, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x7b, 0x43, 0x6f, 0x6d, 0x69,
	0x63, 0x49, 0x64, 0x7d, 0x2f, 0x72, 0x73, 0x73, 0x2f, 0x7b, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44,
	0x7d, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x12, 0x58, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x52, 0x73, 0x73, 0x46, 0x65, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x73, 0x73, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22,
	0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x72, 0x73, 0x73,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comics_proto_rawDescOnce sync.Once
	file_comics_proto_rawDescData = file_comics_proto_rawDesc
)

func file_comics_proto_rawDescGZIP() []byte {
	file_comics_proto_rawDescOnce.Do(func() {
		file_comics_proto_rawDescData = protoimpl.X.CompressGZIP(file_comics_proto_rawDescData)
	})
	return file_comics_proto_rawDescData
}

var file_comics_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_comics_proto_goTypes = []interface{}{
	(*UpdateComicRequest)(nil),     // 0: UpdateComicRequest
	(*GetComicsRequest)(nil),       // 1: GetComicsRequest
	(*ComicState)(nil),             // 2: ComicState
	(*GetComicsResponse)(nil),      // 3: GetComicsResponse
	(*MarkReadRequest)(nil),        // 4: MarkReadRequest
	(*MarkItemReadRequest)(nil),    // 5: MarkItemReadRequest
	(*MarkReadResponse)(nil),       // 6: MarkReadResponse
	(*DescribeComicsRequest)(nil),  // 7: DescribeComicsRequest
	(*ComicDetails)(nil),           // 8: ComicDetails
	(*RssItem)(nil),                // 9: RssItem
	(*DescribeComicsResponse)(nil), // 10: DescribeComicsResponse
	(*RefreshRssFeedResponse)(nil), // 11: RefreshRssFeedResponse
}
var file_comics_proto_depIdxs = []int32{
	2,  // 0: UpdateComicRequest.NewState:type_name -> ComicState
	2,  // 1: GetComicsResponse.Comics:type_name -> ComicState
	9,  // 2: ComicDetails.Feed:type_name -> RssItem
	8,  // 3: DescribeComicsResponse.Comics:type_name -> ComicDetails
	1,  // 4: Comics.GetComics:input_type -> GetComicsRequest
	7,  // 5: Comics.GetComicFeeds:input_type -> DescribeComicsRequest
	0,  // 6: Comics.UpdateComic:input_type -> UpdateComicRequest
	4,  // 7: Comics.MarkAsRead:input_type -> MarkReadRequest
	5,  // 8: Comics.MarkItemRead:input_type -> MarkItemReadRequest
	1,  // 9: Comics.RefreshRssFeed:input_type -> GetComicsRequest
	3,  // 10: Comics.GetComics:output_type -> GetComicsResponse
	10, // 11: Comics.GetComicFeeds:output_type -> DescribeComicsResponse
	3,  // 12: Comics.UpdateComic:output_type -> GetComicsResponse
	6,  // 13: Comics.MarkAsRead:output_type -> MarkReadResponse
	6,  // 14: Comics.MarkItemRead:output_type -> MarkReadResponse
	11, // 15: Comics.RefreshRssFeed:output_type -> RefreshRssFeedResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_comics_proto_init() }
func file_comics_proto_init() {
	if File_comics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateComicRequest); i {
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
		file_comics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetComicsRequest); i {
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
		file_comics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComicState); i {
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
		file_comics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetComicsResponse); i {
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
		file_comics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkReadRequest); i {
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
		file_comics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkItemReadRequest); i {
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
		file_comics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkReadResponse); i {
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
		file_comics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeComicsRequest); i {
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
		file_comics_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComicDetails); i {
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
		file_comics_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RssItem); i {
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
		file_comics_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeComicsResponse); i {
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
		file_comics_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshRssFeedResponse); i {
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
	file_comics_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_comics_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comics_proto_goTypes,
		DependencyIndexes: file_comics_proto_depIdxs,
		MessageInfos:      file_comics_proto_msgTypes,
	}.Build()
	File_comics_proto = out.File
	file_comics_proto_rawDesc = nil
	file_comics_proto_goTypes = nil
	file_comics_proto_depIdxs = nil
}
