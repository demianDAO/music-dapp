// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: songService.proto

package pb

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

type SongModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ArtistAddress string `protobuf:"bytes,3,opt,name=artist_address,json=artistAddress,proto3" json:"artist_address,omitempty"`
	Overview      string `protobuf:"bytes,4,opt,name=overview,proto3" json:"overview,omitempty"`
	NftAddress    string `protobuf:"bytes,5,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
	TokenId       uint64 `protobuf:"varint,6,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	TxId          string `protobuf:"bytes,7,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
}

func (x *SongModel) Reset() {
	*x = SongModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongModel) ProtoMessage() {}

func (x *SongModel) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongModel.ProtoReflect.Descriptor instead.
func (*SongModel) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{0}
}

func (x *SongModel) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SongModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SongModel) GetArtistAddress() string {
	if x != nil {
		return x.ArtistAddress
	}
	return ""
}

func (x *SongModel) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

func (x *SongModel) GetNftAddress() string {
	if x != nil {
		return x.NftAddress
	}
	return ""
}

func (x *SongModel) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *SongModel) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

type CreateSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song    *SongModel `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
	Content []byte     `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateSongRequest) Reset() {
	*x = CreateSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongRequest) ProtoMessage() {}

func (x *CreateSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongRequest.ProtoReflect.Descriptor instead.
func (*CreateSongRequest) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSongRequest) GetSong() *SongModel {
	if x != nil {
		return x.Song
	}
	return nil
}

func (x *CreateSongRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSongResponse) Reset() {
	*x = CreateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongResponse) ProtoMessage() {}

func (x *CreateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongResponse.ProtoReflect.Descriptor instead.
func (*CreateSongResponse) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{2}
}

type FindSongsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtistAddress string `protobuf:"bytes,1,opt,name=artist_address,json=artistAddress,proto3" json:"artist_address,omitempty"`
}

func (x *FindSongsRequest) Reset() {
	*x = FindSongsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSongsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSongsRequest) ProtoMessage() {}

func (x *FindSongsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSongsRequest.ProtoReflect.Descriptor instead.
func (*FindSongsRequest) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{3}
}

func (x *FindSongsRequest) GetArtistAddress() string {
	if x != nil {
		return x.ArtistAddress
	}
	return ""
}

type FindSongsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*SongModel `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *FindSongsResponse) Reset() {
	*x = FindSongsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSongsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSongsResponse) ProtoMessage() {}

func (x *FindSongsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSongsResponse.ProtoReflect.Descriptor instead.
func (*FindSongsResponse) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{4}
}

func (x *FindSongsResponse) GetSongs() []*SongModel {
	if x != nil {
		return x.Songs
	}
	return nil
}

type DownloadSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId string `protobuf:"bytes,1,opt,name=txId,proto3" json:"txId,omitempty"`
}

func (x *DownloadSongRequest) Reset() {
	*x = DownloadSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadSongRequest) ProtoMessage() {}

func (x *DownloadSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadSongRequest.ProtoReflect.Descriptor instead.
func (*DownloadSongRequest) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadSongRequest) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

type DownloadSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongBytes []byte `protobuf:"bytes,1,opt,name=songBytes,proto3" json:"songBytes,omitempty"`
}

func (x *DownloadSongResponse) Reset() {
	*x = DownloadSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadSongResponse) ProtoMessage() {}

func (x *DownloadSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadSongResponse.ProtoReflect.Descriptor instead.
func (*DownloadSongResponse) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadSongResponse) GetSongBytes() []byte {
	if x != nil {
		return x.SongBytes
	}
	return nil
}

var File_songService_proto protoreflect.FileDescriptor

var file_songService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x6f, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xc5, 0x01,
	0x0a, 0x09, 0x53, 0x6f, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x66, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x66, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64,
	0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x78, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x73, 0x6f,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x73,
	0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x14, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3e,
	0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6f,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x29,
	0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x14, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x6e, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x6f, 0x6e, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x32,
	0xeb, 0x01, 0x0a, 0x0b, 0x53, 0x6f, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x47, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64,
	0x53, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_songService_proto_rawDescOnce sync.Once
	file_songService_proto_rawDescData = file_songService_proto_rawDesc
)

func file_songService_proto_rawDescGZIP() []byte {
	file_songService_proto_rawDescOnce.Do(func() {
		file_songService_proto_rawDescData = protoimpl.X.CompressGZIP(file_songService_proto_rawDescData)
	})
	return file_songService_proto_rawDescData
}

var file_songService_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_songService_proto_goTypes = []interface{}{
	(*SongModel)(nil),            // 0: services.SongModel
	(*CreateSongRequest)(nil),    // 1: services.CreateSongRequest
	(*CreateSongResponse)(nil),   // 2: services.CreateSongResponse
	(*FindSongsRequest)(nil),     // 3: services.FindSongsRequest
	(*FindSongsResponse)(nil),    // 4: services.FindSongsResponse
	(*DownloadSongRequest)(nil),  // 5: services.DownloadSongRequest
	(*DownloadSongResponse)(nil), // 6: services.DownloadSongResponse
}
var file_songService_proto_depIdxs = []int32{
	0, // 0: services.CreateSongRequest.song:type_name -> services.SongModel
	0, // 1: services.FindSongsResponse.songs:type_name -> services.SongModel
	1, // 2: services.SongService.UploadSong:input_type -> services.CreateSongRequest
	3, // 3: services.SongService.FindSongs:input_type -> services.FindSongsRequest
	5, // 4: services.SongService.DownloadSong:input_type -> services.DownloadSongRequest
	2, // 5: services.SongService.UploadSong:output_type -> services.CreateSongResponse
	4, // 6: services.SongService.FindSongs:output_type -> services.FindSongsResponse
	6, // 7: services.SongService.DownloadSong:output_type -> services.DownloadSongResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_songService_proto_init() }
func file_songService_proto_init() {
	if File_songService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_songService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongModel); i {
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
		file_songService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongRequest); i {
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
		file_songService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongResponse); i {
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
		file_songService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSongsRequest); i {
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
		file_songService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSongsResponse); i {
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
		file_songService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadSongRequest); i {
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
		file_songService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadSongResponse); i {
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
			RawDescriptor: file_songService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_songService_proto_goTypes,
		DependencyIndexes: file_songService_proto_depIdxs,
		MessageInfos:      file_songService_proto_msgTypes,
	}.Build()
	File_songService_proto = out.File
	file_songService_proto_rawDesc = nil
	file_songService_proto_goTypes = nil
	file_songService_proto_depIdxs = nil
}
