// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testpb/test.proto

package testpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "student-proto/studentpb"
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

type Test struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{0}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Test) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetTestRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTestRequest) Reset()         { *m = GetTestRequest{} }
func (m *GetTestRequest) String() string { return proto.CompactTextString(m) }
func (*GetTestRequest) ProtoMessage()    {}
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{1}
}

func (m *GetTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTestRequest.Unmarshal(m, b)
}
func (m *GetTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTestRequest.Marshal(b, m, deterministic)
}
func (m *GetTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTestRequest.Merge(m, src)
}
func (m *GetTestRequest) XXX_Size() int {
	return xxx_messageInfo_GetTestRequest.Size(m)
}
func (m *GetTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTestRequest proto.InternalMessageInfo

func (m *GetTestRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SetTestResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTestResponse) Reset()         { *m = SetTestResponse{} }
func (m *SetTestResponse) String() string { return proto.CompactTextString(m) }
func (*SetTestResponse) ProtoMessage()    {}
func (*SetTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{2}
}

func (m *SetTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTestResponse.Unmarshal(m, b)
}
func (m *SetTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTestResponse.Marshal(b, m, deterministic)
}
func (m *SetTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTestResponse.Merge(m, src)
}
func (m *SetTestResponse) XXX_Size() int {
	return xxx_messageInfo_SetTestResponse.Size(m)
}
func (m *SetTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetTestResponse proto.InternalMessageInfo

func (m *SetTestResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SetTestResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Question struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question             string   `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	TestId               string   `protobuf:"bytes,4,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{3}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Question) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *Question) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *Question) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

type QuestionsPerTest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Question             string   `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Ok                   bool     `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Current              int32    `protobuf:"varint,4,opt,name=current,proto3" json:"current,omitempty"`
	Total                int32    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuestionsPerTest) Reset()         { *m = QuestionsPerTest{} }
func (m *QuestionsPerTest) String() string { return proto.CompactTextString(m) }
func (*QuestionsPerTest) ProtoMessage()    {}
func (*QuestionsPerTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{4}
}

func (m *QuestionsPerTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuestionsPerTest.Unmarshal(m, b)
}
func (m *QuestionsPerTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuestionsPerTest.Marshal(b, m, deterministic)
}
func (m *QuestionsPerTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestionsPerTest.Merge(m, src)
}
func (m *QuestionsPerTest) XXX_Size() int {
	return xxx_messageInfo_QuestionsPerTest.Size(m)
}
func (m *QuestionsPerTest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestionsPerTest.DiscardUnknown(m)
}

var xxx_messageInfo_QuestionsPerTest proto.InternalMessageInfo

func (m *QuestionsPerTest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QuestionsPerTest) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *QuestionsPerTest) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *QuestionsPerTest) GetCurrent() int32 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *QuestionsPerTest) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type SetQuestionResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetQuestionResponse) Reset()         { *m = SetQuestionResponse{} }
func (m *SetQuestionResponse) String() string { return proto.CompactTextString(m) }
func (*SetQuestionResponse) ProtoMessage()    {}
func (*SetQuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{5}
}

func (m *SetQuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetQuestionResponse.Unmarshal(m, b)
}
func (m *SetQuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetQuestionResponse.Marshal(b, m, deterministic)
}
func (m *SetQuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetQuestionResponse.Merge(m, src)
}
func (m *SetQuestionResponse) XXX_Size() int {
	return xxx_messageInfo_SetQuestionResponse.Size(m)
}
func (m *SetQuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetQuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetQuestionResponse proto.InternalMessageInfo

func (m *SetQuestionResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type EnrollmentRequest struct {
	StudentId            string   `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TestId               string   `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnrollmentRequest) Reset()         { *m = EnrollmentRequest{} }
func (m *EnrollmentRequest) String() string { return proto.CompactTextString(m) }
func (*EnrollmentRequest) ProtoMessage()    {}
func (*EnrollmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{6}
}

func (m *EnrollmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnrollmentRequest.Unmarshal(m, b)
}
func (m *EnrollmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnrollmentRequest.Marshal(b, m, deterministic)
}
func (m *EnrollmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnrollmentRequest.Merge(m, src)
}
func (m *EnrollmentRequest) XXX_Size() int {
	return xxx_messageInfo_EnrollmentRequest.Size(m)
}
func (m *EnrollmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnrollmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnrollmentRequest proto.InternalMessageInfo

func (m *EnrollmentRequest) GetStudentId() string {
	if m != nil {
		return m.StudentId
	}
	return ""
}

func (m *EnrollmentRequest) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

type GetStudentsPerTestRequest struct {
	TestId               string   `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentsPerTestRequest) Reset()         { *m = GetStudentsPerTestRequest{} }
func (m *GetStudentsPerTestRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentsPerTestRequest) ProtoMessage()    {}
func (*GetStudentsPerTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{7}
}

func (m *GetStudentsPerTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStudentsPerTestRequest.Unmarshal(m, b)
}
func (m *GetStudentsPerTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStudentsPerTestRequest.Marshal(b, m, deterministic)
}
func (m *GetStudentsPerTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentsPerTestRequest.Merge(m, src)
}
func (m *GetStudentsPerTestRequest) XXX_Size() int {
	return xxx_messageInfo_GetStudentsPerTestRequest.Size(m)
}
func (m *GetStudentsPerTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentsPerTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentsPerTestRequest proto.InternalMessageInfo

func (m *GetStudentsPerTestRequest) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

type TakeTestRequest struct {
	Answer               string   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	TestId               string   `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	StudentId            string   `protobuf:"bytes,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeTestRequest) Reset()         { *m = TakeTestRequest{} }
func (m *TakeTestRequest) String() string { return proto.CompactTextString(m) }
func (*TakeTestRequest) ProtoMessage()    {}
func (*TakeTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{8}
}

func (m *TakeTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeTestRequest.Unmarshal(m, b)
}
func (m *TakeTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeTestRequest.Marshal(b, m, deterministic)
}
func (m *TakeTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeTestRequest.Merge(m, src)
}
func (m *TakeTestRequest) XXX_Size() int {
	return xxx_messageInfo_TakeTestRequest.Size(m)
}
func (m *TakeTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TakeTestRequest proto.InternalMessageInfo

func (m *TakeTestRequest) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *TakeTestRequest) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *TakeTestRequest) GetStudentId() string {
	if m != nil {
		return m.StudentId
	}
	return ""
}

type GetScoreTestRequest struct {
	TestId               string   `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	StudentId            string   `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetScoreTestRequest) Reset()         { *m = GetScoreTestRequest{} }
func (m *GetScoreTestRequest) String() string { return proto.CompactTextString(m) }
func (*GetScoreTestRequest) ProtoMessage()    {}
func (*GetScoreTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{9}
}

func (m *GetScoreTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetScoreTestRequest.Unmarshal(m, b)
}
func (m *GetScoreTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetScoreTestRequest.Marshal(b, m, deterministic)
}
func (m *GetScoreTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetScoreTestRequest.Merge(m, src)
}
func (m *GetScoreTestRequest) XXX_Size() int {
	return xxx_messageInfo_GetScoreTestRequest.Size(m)
}
func (m *GetScoreTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetScoreTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetScoreTestRequest proto.InternalMessageInfo

func (m *GetScoreTestRequest) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *GetScoreTestRequest) GetStudentId() string {
	if m != nil {
		return m.StudentId
	}
	return ""
}

type TestScore struct {
	TestId               string   `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	StudentId            string   `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Ok                   int32    `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Ko                   int32    `protobuf:"varint,4,opt,name=ko,proto3" json:"ko,omitempty"`
	Total                int32    `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	Score                int32    `protobuf:"varint,6,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestScore) Reset()         { *m = TestScore{} }
func (m *TestScore) String() string { return proto.CompactTextString(m) }
func (*TestScore) ProtoMessage()    {}
func (*TestScore) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c67e33ca9d1f26, []int{10}
}

func (m *TestScore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestScore.Unmarshal(m, b)
}
func (m *TestScore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestScore.Marshal(b, m, deterministic)
}
func (m *TestScore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestScore.Merge(m, src)
}
func (m *TestScore) XXX_Size() int {
	return xxx_messageInfo_TestScore.Size(m)
}
func (m *TestScore) XXX_DiscardUnknown() {
	xxx_messageInfo_TestScore.DiscardUnknown(m)
}

var xxx_messageInfo_TestScore proto.InternalMessageInfo

func (m *TestScore) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *TestScore) GetStudentId() string {
	if m != nil {
		return m.StudentId
	}
	return ""
}

func (m *TestScore) GetOk() int32 {
	if m != nil {
		return m.Ok
	}
	return 0
}

func (m *TestScore) GetKo() int32 {
	if m != nil {
		return m.Ko
	}
	return 0
}

func (m *TestScore) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *TestScore) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*Test)(nil), "test.Test")
	proto.RegisterType((*GetTestRequest)(nil), "test.GetTestRequest")
	proto.RegisterType((*SetTestResponse)(nil), "test.SetTestResponse")
	proto.RegisterType((*Question)(nil), "test.Question")
	proto.RegisterType((*QuestionsPerTest)(nil), "test.QuestionsPerTest")
	proto.RegisterType((*SetQuestionResponse)(nil), "test.SetQuestionResponse")
	proto.RegisterType((*EnrollmentRequest)(nil), "test.EnrollmentRequest")
	proto.RegisterType((*GetStudentsPerTestRequest)(nil), "test.GetStudentsPerTestRequest")
	proto.RegisterType((*TakeTestRequest)(nil), "test.TakeTestRequest")
	proto.RegisterType((*GetScoreTestRequest)(nil), "test.GetScoreTestRequest")
	proto.RegisterType((*TestScore)(nil), "test.TestScore")
}

func init() { proto.RegisterFile("testpb/test.proto", fileDescriptor_41c67e33ca9d1f26) }

var fileDescriptor_41c67e33ca9d1f26 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x96, 0x37, 0xef, 0xa1, 0x72, 0xda, 0x69, 0x48, 0x1c, 0x4b, 0x88, 0xca, 0x12, 0x52, 0xc4,
	0x23, 0xad, 0x78, 0x5c, 0x40, 0x5c, 0x90, 0xa0, 0x8a, 0x10, 0x12, 0x24, 0x3d, 0x71, 0xa9, 0xdc,
	0x78, 0x85, 0x2c, 0xa7, 0xde, 0xd4, 0xbb, 0x81, 0x1b, 0x7f, 0x81, 0x1f, 0xc3, 0x1f, 0x44, 0xfb,
	0x72, 0xec, 0x4d, 0x40, 0x11, 0x27, 0xef, 0xcc, 0xce, 0x7c, 0xf3, 0xcd, 0xce, 0x37, 0x86, 0x13,
	0x41, 0xb9, 0x58, 0xdf, 0x9c, 0xcb, 0xcf, 0x74, 0x5d, 0x30, 0xc1, 0xb0, 0x29, 0xcf, 0xe1, 0x88,
	0x8b, 0x4d, 0x42, 0x73, 0x79, 0x67, 0x4e, 0xfa, 0x3a, 0x7a, 0x0c, 0xcd, 0x2b, 0xca, 0x05, 0xfa,
	0x40, 0xd2, 0x24, 0xf0, 0xce, 0xbc, 0x49, 0x6f, 0x4e, 0xd2, 0x04, 0x11, 0x9a, 0x79, 0x7c, 0x4b,
	0x03, 0xa2, 0x3c, 0xea, 0x1c, 0x9d, 0x81, 0x7f, 0x49, 0x85, 0x0c, 0x9f, 0xd3, 0xbb, 0xcd, 0x9e,
	0xac, 0xe8, 0x15, 0xf4, 0x17, 0x36, 0x82, 0xaf, 0x59, 0xce, 0xe9, 0x41, 0xc0, 0xdf, 0xa0, 0xfb,
	0x45, 0xe2, 0xa5, 0x2c, 0xdf, 0x89, 0x1f, 0x42, 0x3b, 0xce, 0xf9, 0x0f, 0x5a, 0x98, 0x0c, 0x63,
	0x61, 0x08, 0xdd, 0x3b, 0x93, 0x13, 0x34, 0xd4, 0x4d, 0x69, 0xe3, 0x08, 0x3a, 0xb2, 0xeb, 0xeb,
	0x34, 0x09, 0x9a, 0x3a, 0x49, 0x9a, 0xb3, 0x24, 0xfa, 0x09, 0xc7, 0xb6, 0x10, 0xff, 0x4c, 0x8b,
	0xbd, 0x9d, 0x57, 0x81, 0x89, 0x03, 0xec, 0x03, 0x61, 0x99, 0x2a, 0xd7, 0x9d, 0x13, 0x96, 0x61,
	0x00, 0x9d, 0xe5, 0xa6, 0x28, 0x68, 0x2e, 0x54, 0xa1, 0xd6, 0xdc, 0x9a, 0x38, 0x80, 0x96, 0x60,
	0x22, 0x5e, 0x05, 0x2d, 0xe5, 0xd7, 0x46, 0xf4, 0x08, 0x4e, 0x17, 0x54, 0x58, 0x0a, 0xd5, 0x37,
	0x62, 0x99, 0xa2, 0xa0, 0x60, 0xa3, 0x8f, 0x70, 0xf2, 0x3e, 0x2f, 0xd8, 0x6a, 0x75, 0x4b, 0xf3,
	0xf2, 0xad, 0x1f, 0x00, 0x98, 0xd1, 0x5d, 0x97, 0x7c, 0x7b, 0xc6, 0x33, 0x4b, 0xaa, 0x3d, 0x93,
	0x5a, 0xcf, 0x2f, 0x61, 0x7c, 0x49, 0xc5, 0x42, 0x07, 0xda, 0xae, 0x2d, 0x68, 0x25, 0xcb, 0xab,
	0x65, 0xc5, 0xd0, 0xbf, 0x8a, 0x33, 0x5a, 0x8d, 0xdd, 0x4e, 0xc2, 0xab, 0x4d, 0xe2, 0x6f, 0x95,
	0x1d, 0xc6, 0x0d, 0x87, 0x71, 0xf4, 0x09, 0x4e, 0x25, 0xb1, 0x25, 0x2b, 0xe8, 0x21, 0x94, 0x1c,
	0x38, 0xe2, 0xc2, 0xfd, 0xf2, 0xa0, 0x27, 0x71, 0x14, 0xe0, 0xff, 0xa2, 0x54, 0x26, 0xdc, 0x52,
	0x13, 0xf6, 0x81, 0x64, 0xcc, 0x0c, 0x97, 0x64, 0x6c, 0xff, 0x5c, 0xa5, 0x97, 0xcb, 0xb2, 0x41,
	0x5b, 0x7b, 0x95, 0xf1, 0xfc, 0x77, 0x03, 0xee, 0x29, 0x46, 0xb4, 0xf8, 0x9e, 0x2e, 0x29, 0x3e,
	0x81, 0x8e, 0xd9, 0x1f, 0x1c, 0x4c, 0xd5, 0x8a, 0xd6, 0xd7, 0x29, 0x04, 0xed, 0x55, 0x11, 0x4f,
	0xa1, 0x63, 0x56, 0x09, 0x2b, 0xee, 0xf0, 0xbe, 0x3e, 0xbb, 0x5b, 0xf6, 0x06, 0x8e, 0x2a, 0xc2,
	0xe2, 0xe8, 0xeb, 0x30, 0xeb, 0x08, 0xc7, 0x65, 0x9a, 0x2b, 0xbe, 0x89, 0x87, 0x1f, 0xc0, 0xd7,
	0x72, 0xb3, 0x22, 0xc1, 0x91, 0x0e, 0xdf, 0x11, 0xe1, 0xbf, 0x71, 0x66, 0x80, 0xbb, 0x4a, 0xc3,
	0x87, 0x65, 0xab, 0xfb, 0x35, 0x18, 0x1e, 0x4f, 0xed, 0x2f, 0xc9, 0x04, 0x5c, 0x78, 0xf8, 0x16,
	0xba, 0x56, 0x7e, 0x68, 0x5a, 0x76, 0xe4, 0x18, 0x0e, 0xeb, 0x2d, 0x5a, 0xd4, 0x89, 0x77, 0xe1,
	0xe1, 0x6b, 0x38, 0x32, 0x4f, 0xab, 0xd5, 0x30, 0xde, 0x72, 0x70, 0xe4, 0x16, 0xf6, 0xb7, 0x8f,
	0xab, 0xee, 0xde, 0x0d, 0xbf, 0x0e, 0x0c, 0x9f, 0x67, 0xea, 0x17, 0x79, 0xae, 0xff, 0xa9, 0x37,
	0x6d, 0x65, 0xbd, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xda, 0x63, 0xdb, 0xf3, 0x64, 0x05, 0x00,
	0x00,
}