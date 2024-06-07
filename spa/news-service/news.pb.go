// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: protos/news.proto

package __

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

type GetNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
}

func (x *GetNewsRequest) Reset() {
	*x = GetNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsRequest) ProtoMessage() {}

func (x *GetNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsRequest.ProtoReflect.Descriptor instead.
func (*GetNewsRequest) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{0}
}

func (x *GetNewsRequest) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

type GetNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News *News `protobuf:"bytes,1,opt,name=news,proto3" json:"news,omitempty"`
}

func (x *GetNewsResponse) Reset() {
	*x = GetNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewsResponse) ProtoMessage() {}

func (x *GetNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewsResponse.ProtoReflect.Descriptor instead.
func (*GetNewsResponse) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{1}
}

func (x *GetNewsResponse) GetNews() *News {
	if x != nil {
		return x.News
	}
	return nil
}

type CreateNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateNewsRequest) Reset() {
	*x = CreateNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewsRequest) ProtoMessage() {}

func (x *CreateNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewsRequest.ProtoReflect.Descriptor instead.
func (*CreateNewsRequest) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNewsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateNewsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
}

func (x *CreateNewsResponse) Reset() {
	*x = CreateNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewsResponse) ProtoMessage() {}

func (x *CreateNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewsResponse.ProtoReflect.Descriptor instead.
func (*CreateNewsResponse) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNewsResponse) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

type UpdateNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId  string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateNewsRequest) Reset() {
	*x = UpdateNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNewsRequest) ProtoMessage() {}

func (x *UpdateNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNewsRequest.ProtoReflect.Descriptor instead.
func (*UpdateNewsRequest) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNewsRequest) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

func (x *UpdateNewsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateNewsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News *News `protobuf:"bytes,1,opt,name=news,proto3" json:"news,omitempty"`
}

func (x *UpdateNewsResponse) Reset() {
	*x = UpdateNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNewsResponse) ProtoMessage() {}

func (x *UpdateNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNewsResponse.ProtoReflect.Descriptor instead.
func (*UpdateNewsResponse) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNewsResponse) GetNews() *News {
	if x != nil {
		return x.News
	}
	return nil
}

type DeleteNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
}

func (x *DeleteNewsRequest) Reset() {
	*x = DeleteNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNewsRequest) ProtoMessage() {}

func (x *DeleteNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNewsRequest.ProtoReflect.Descriptor instead.
func (*DeleteNewsRequest) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteNewsRequest) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

type DeleteNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteNewsResponse) Reset() {
	*x = DeleteNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNewsResponse) ProtoMessage() {}

func (x *DeleteNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNewsResponse.ProtoReflect.Descriptor instead.
func (*DeleteNewsResponse) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteNewsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter   string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Sort     string `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	Page     int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListNewsRequest) Reset() {
	*x = ListNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNewsRequest) ProtoMessage() {}

func (x *ListNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNewsRequest.ProtoReflect.Descriptor instead.
func (*ListNewsRequest) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{8}
}

func (x *ListNewsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListNewsRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *ListNewsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListNewsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListNewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsList    []*News `protobuf:"bytes,1,rep,name=news_list,json=newsList,proto3" json:"news_list,omitempty"`
	TotalPages  int32   `protobuf:"varint,2,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	CurrentPage int32   `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
}

func (x *ListNewsResponse) Reset() {
	*x = ListNewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNewsResponse) ProtoMessage() {}

func (x *ListNewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNewsResponse.ProtoReflect.Descriptor instead.
func (*ListNewsResponse) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{9}
}

func (x *ListNewsResponse) GetNewsList() []*News {
	if x != nil {
		return x.NewsList
	}
	return nil
}

func (x *ListNewsResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *ListNewsResponse) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

type SendNewsToUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
}

func (x *SendNewsToUsersRequest) Reset() {
	*x = SendNewsToUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNewsToUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNewsToUsersRequest) ProtoMessage() {}

func (x *SendNewsToUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNewsToUsersRequest.ProtoReflect.Descriptor instead.
func (*SendNewsToUsersRequest) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{10}
}

func (x *SendNewsToUsersRequest) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

type SendNewsToUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendNewsToUsersResponse) Reset() {
	*x = SendNewsToUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNewsToUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNewsToUsersResponse) ProtoMessage() {}

func (x *SendNewsToUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNewsToUsersResponse.ProtoReflect.Descriptor instead.
func (*SendNewsToUsersResponse) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{11}
}

func (x *SendNewsToUsersResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type News struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId  string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *News) Reset() {
	*x = News{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_news_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *News) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*News) ProtoMessage() {}

func (x *News) ProtoReflect() protoreflect.Message {
	mi := &file_protos_news_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use News.ProtoReflect.Descriptor instead.
func (*News) Descriptor() ([]byte, []int) {
	return file_protos_news_proto_rawDescGZIP(), []int{12}
}

func (x *News) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

func (x *News) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *News) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_protos_news_proto protoreflect.FileDescriptor

var file_protos_news_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x65, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65,
	0x77, 0x73, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x22, 0x43, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x22, 0x5c, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x22,
	0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x22, 0x2e, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x6e, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7f, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x08, 0x6e, 0x65, 0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x31,
	0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49,
	0x64, 0x22, 0x33, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x54, 0x6f, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x93, 0x03, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x77, 0x73, 0x12, 0x14, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x17, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x17,
	0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12,
	0x17, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x15,
	0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x54, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a,
	0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_news_proto_rawDescOnce sync.Once
	file_protos_news_proto_rawDescData = file_protos_news_proto_rawDesc
)

func file_protos_news_proto_rawDescGZIP() []byte {
	file_protos_news_proto_rawDescOnce.Do(func() {
		file_protos_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_news_proto_rawDescData)
	})
	return file_protos_news_proto_rawDescData
}

var file_protos_news_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protos_news_proto_goTypes = []interface{}{
	(*GetNewsRequest)(nil),          // 0: news.GetNewsRequest
	(*GetNewsResponse)(nil),         // 1: news.GetNewsResponse
	(*CreateNewsRequest)(nil),       // 2: news.CreateNewsRequest
	(*CreateNewsResponse)(nil),      // 3: news.CreateNewsResponse
	(*UpdateNewsRequest)(nil),       // 4: news.UpdateNewsRequest
	(*UpdateNewsResponse)(nil),      // 5: news.UpdateNewsResponse
	(*DeleteNewsRequest)(nil),       // 6: news.DeleteNewsRequest
	(*DeleteNewsResponse)(nil),      // 7: news.DeleteNewsResponse
	(*ListNewsRequest)(nil),         // 8: news.ListNewsRequest
	(*ListNewsResponse)(nil),        // 9: news.ListNewsResponse
	(*SendNewsToUsersRequest)(nil),  // 10: news.SendNewsToUsersRequest
	(*SendNewsToUsersResponse)(nil), // 11: news.SendNewsToUsersResponse
	(*News)(nil),                    // 12: news.News
}
var file_protos_news_proto_depIdxs = []int32{
	12, // 0: news.GetNewsResponse.news:type_name -> news.News
	12, // 1: news.UpdateNewsResponse.news:type_name -> news.News
	12, // 2: news.ListNewsResponse.news_list:type_name -> news.News
	0,  // 3: news.NewsService.GetNews:input_type -> news.GetNewsRequest
	2,  // 4: news.NewsService.CreateNews:input_type -> news.CreateNewsRequest
	4,  // 5: news.NewsService.UpdateNews:input_type -> news.UpdateNewsRequest
	6,  // 6: news.NewsService.DeleteNews:input_type -> news.DeleteNewsRequest
	8,  // 7: news.NewsService.ListNews:input_type -> news.ListNewsRequest
	10, // 8: news.NewsService.SendNewsToUsers:input_type -> news.SendNewsToUsersRequest
	1,  // 9: news.NewsService.GetNews:output_type -> news.GetNewsResponse
	3,  // 10: news.NewsService.CreateNews:output_type -> news.CreateNewsResponse
	5,  // 11: news.NewsService.UpdateNews:output_type -> news.UpdateNewsResponse
	7,  // 12: news.NewsService.DeleteNews:output_type -> news.DeleteNewsResponse
	9,  // 13: news.NewsService.ListNews:output_type -> news.ListNewsResponse
	11, // 14: news.NewsService.SendNewsToUsers:output_type -> news.SendNewsToUsersResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protos_news_proto_init() }
func file_protos_news_proto_init() {
	if File_protos_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsRequest); i {
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
		file_protos_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewsResponse); i {
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
		file_protos_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewsRequest); i {
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
		file_protos_news_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewsResponse); i {
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
		file_protos_news_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNewsRequest); i {
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
		file_protos_news_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNewsResponse); i {
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
		file_protos_news_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNewsRequest); i {
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
		file_protos_news_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNewsResponse); i {
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
		file_protos_news_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNewsRequest); i {
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
		file_protos_news_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNewsResponse); i {
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
		file_protos_news_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNewsToUsersRequest); i {
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
		file_protos_news_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNewsToUsersResponse); i {
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
		file_protos_news_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*News); i {
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
			RawDescriptor: file_protos_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_news_proto_goTypes,
		DependencyIndexes: file_protos_news_proto_depIdxs,
		MessageInfos:      file_protos_news_proto_msgTypes,
	}.Build()
	File_protos_news_proto = out.File
	file_protos_news_proto_rawDesc = nil
	file_protos_news_proto_goTypes = nil
	file_protos_news_proto_depIdxs = nil
}
