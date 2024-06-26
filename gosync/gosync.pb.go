// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: gosync.proto

package gosync

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

type Status int32

const (
	Status_Ok       Status = 0
	Status_Continue Status = 1
	Status_Error    Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Ok",
		1: "Continue",
		2: "Error",
	}
	Status_value = map[string]int32{
		"Ok":       0,
		"Continue": 1,
		"Error":    2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_gosync_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_gosync_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_gosync_proto_rawDescGZIP(), []int{0}
}

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk       []byte `protobuf:"bytes,1,opt,name=Chunk,proto3" json:"Chunk,omitempty"`
	Offset      int64  `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	MD5Checksum string `protobuf:"bytes,3,opt,name=MD5Checksum,proto3" json:"MD5Checksum,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_gosync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_gosync_proto_rawDescGZIP(), []int{0}
}

func (x *FileChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *FileChunk) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FileChunk) GetMD5Checksum() string {
	if x != nil {
		return x.MD5Checksum
	}
	return ""
}

type FileDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    string `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	MD5Checksum string `protobuf:"bytes,2,opt,name=MD5Checksum,proto3" json:"MD5Checksum,omitempty"`
	Size        int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (x *FileDetails) Reset() {
	*x = FileDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDetails) ProtoMessage() {}

func (x *FileDetails) ProtoReflect() protoreflect.Message {
	mi := &file_gosync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDetails.ProtoReflect.Descriptor instead.
func (*FileDetails) Descriptor() ([]byte, []int) {
	return file_gosync_proto_rawDescGZIP(), []int{1}
}

func (x *FileDetails) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileDetails) GetMD5Checksum() string {
	if x != nil {
		return x.MD5Checksum
	}
	return ""
}

func (x *FileDetails) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkSize  int32 `protobuf:"varint,1,opt,name=ChunkSize,proto3" json:"ChunkSize,omitempty"`
	Encryption bool  `protobuf:"varint,2,opt,name=Encryption,proto3" json:"Encryption,omitempty"`
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_gosync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_gosync_proto_rawDescGZIP(), []int{2}
}

func (x *FileOptions) GetChunkSize() int32 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

func (x *FileOptions) GetEncryption() bool {
	if x != nil {
		return x.Encryption
	}
	return false
}

type FilePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      Status       `protobuf:"varint,1,opt,name=Status,proto3,enum=Status" json:"Status,omitempty"`
	FileDetails *FileDetails `protobuf:"bytes,2,opt,name=FileDetails,proto3" json:"FileDetails,omitempty"`
	FileOptions *FileOptions `protobuf:"bytes,3,opt,name=FileOptions,proto3" json:"FileOptions,omitempty"`
	FileChunk   *FileChunk   `protobuf:"bytes,4,opt,name=FileChunk,proto3" json:"FileChunk,omitempty"`
	FileChunks  []*FileChunk `protobuf:"bytes,5,rep,name=FileChunks,proto3" json:"FileChunks,omitempty"`
}

func (x *FilePayload) Reset() {
	*x = FilePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePayload) ProtoMessage() {}

func (x *FilePayload) ProtoReflect() protoreflect.Message {
	mi := &file_gosync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePayload.ProtoReflect.Descriptor instead.
func (*FilePayload) Descriptor() ([]byte, []int) {
	return file_gosync_proto_rawDescGZIP(), []int{3}
}

func (x *FilePayload) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Ok
}

func (x *FilePayload) GetFileDetails() *FileDetails {
	if x != nil {
		return x.FileDetails
	}
	return nil
}

func (x *FilePayload) GetFileOptions() *FileOptions {
	if x != nil {
		return x.FileOptions
	}
	return nil
}

func (x *FilePayload) GetFileChunk() *FileChunk {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

func (x *FilePayload) GetFileChunks() []*FileChunk {
	if x != nil {
		return x.FileChunks
	}
	return nil
}

var File_gosync_proto protoreflect.FileDescriptor

var file_gosync_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b,
	0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x44, 0x35,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4d, 0x44, 0x35, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22, 0x5f, 0x0a, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x44, 0x35, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x44, 0x35,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4b, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe4, 0x01, 0x0a, 0x0b, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x46, 0x69,
	0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x09, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x2a, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x2a, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x32, 0xcf, 0x01, 0x0a, 0x0d,
	0x47, 0x6f, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0c, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x2e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a,
	0x0c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0c, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0c, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2e, 0x0a,
	0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0c, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0c, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1e, 0x5a,
	0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6a, 0x73, 0x61,
	0x66, 0x72, 0x61, 0x6e, 0x65, 0x6b, 0x2f, 0x67, 0x6f, 0x73, 0x79, 0x6e, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gosync_proto_rawDescOnce sync.Once
	file_gosync_proto_rawDescData = file_gosync_proto_rawDesc
)

func file_gosync_proto_rawDescGZIP() []byte {
	file_gosync_proto_rawDescOnce.Do(func() {
		file_gosync_proto_rawDescData = protoimpl.X.CompressGZIP(file_gosync_proto_rawDescData)
	})
	return file_gosync_proto_rawDescData
}

var file_gosync_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gosync_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gosync_proto_goTypes = []interface{}{
	(Status)(0),         // 0: Status
	(*FileChunk)(nil),   // 1: FileChunk
	(*FileDetails)(nil), // 2: FileDetails
	(*FileOptions)(nil), // 3: FileOptions
	(*FilePayload)(nil), // 4: FilePayload
}
var file_gosync_proto_depIdxs = []int32{
	0, // 0: FilePayload.Status:type_name -> Status
	2, // 1: FilePayload.FileDetails:type_name -> FileDetails
	3, // 2: FilePayload.FileOptions:type_name -> FileOptions
	1, // 3: FilePayload.FileChunk:type_name -> FileChunk
	1, // 4: FilePayload.FileChunks:type_name -> FileChunk
	4, // 5: GoSyncService.Authenticate:input_type -> FilePayload
	4, // 6: GoSyncService.GetFileDetails:input_type -> FilePayload
	4, // 7: GoSyncService.UploadFile:input_type -> FilePayload
	4, // 8: GoSyncService.DownloadFile:input_type -> FilePayload
	4, // 9: GoSyncService.Authenticate:output_type -> FilePayload
	4, // 10: GoSyncService.GetFileDetails:output_type -> FilePayload
	4, // 11: GoSyncService.UploadFile:output_type -> FilePayload
	4, // 12: GoSyncService.DownloadFile:output_type -> FilePayload
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gosync_proto_init() }
func file_gosync_proto_init() {
	if File_gosync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gosync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileChunk); i {
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
		file_gosync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDetails); i {
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
		file_gosync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOptions); i {
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
		file_gosync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilePayload); i {
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
			RawDescriptor: file_gosync_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gosync_proto_goTypes,
		DependencyIndexes: file_gosync_proto_depIdxs,
		EnumInfos:         file_gosync_proto_enumTypes,
		MessageInfos:      file_gosync_proto_msgTypes,
	}.Build()
	File_gosync_proto = out.File
	file_gosync_proto_rawDesc = nil
	file_gosync_proto_goTypes = nil
	file_gosync_proto_depIdxs = nil
}
