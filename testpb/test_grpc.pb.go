// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	studentpb "student-proto/studentpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*Test, error)
	SetTest(ctx context.Context, in *Test, opts ...grpc.CallOption) (*SetTestResponse, error)
	SetQuestions(ctx context.Context, opts ...grpc.CallOption) (TestService_SetQuestionsClient, error)
	EnrollStudents(ctx context.Context, opts ...grpc.CallOption) (TestService_EnrollStudentsClient, error)
	GetStudentsPerTest(ctx context.Context, in *GetStudentsPerTestRequest, opts ...grpc.CallOption) (TestService_GetStudentsPerTestClient, error)
	TakeTest(ctx context.Context, opts ...grpc.CallOption) (TestService_TakeTestClient, error)
	GetTestScore(ctx context.Context, in *GetScoreTestRequest, opts ...grpc.CallOption) (*TestScore, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) GetTest(ctx context.Context, in *GetTestRequest, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/test.TestService/GetTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) SetTest(ctx context.Context, in *Test, opts ...grpc.CallOption) (*SetTestResponse, error) {
	out := new(SetTestResponse)
	err := c.cc.Invoke(ctx, "/test.TestService/SetTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) SetQuestions(ctx context.Context, opts ...grpc.CallOption) (TestService_SetQuestionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], "/test.TestService/SetQuestions", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceSetQuestionsClient{stream}
	return x, nil
}

type TestService_SetQuestionsClient interface {
	Send(*Question) error
	CloseAndRecv() (*SetQuestionResponse, error)
	grpc.ClientStream
}

type testServiceSetQuestionsClient struct {
	grpc.ClientStream
}

func (x *testServiceSetQuestionsClient) Send(m *Question) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceSetQuestionsClient) CloseAndRecv() (*SetQuestionResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SetQuestionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) EnrollStudents(ctx context.Context, opts ...grpc.CallOption) (TestService_EnrollStudentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[1], "/test.TestService/EnrollStudents", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceEnrollStudentsClient{stream}
	return x, nil
}

type TestService_EnrollStudentsClient interface {
	Send(*EnrollmentRequest) error
	CloseAndRecv() (*SetQuestionResponse, error)
	grpc.ClientStream
}

type testServiceEnrollStudentsClient struct {
	grpc.ClientStream
}

func (x *testServiceEnrollStudentsClient) Send(m *EnrollmentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceEnrollStudentsClient) CloseAndRecv() (*SetQuestionResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SetQuestionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) GetStudentsPerTest(ctx context.Context, in *GetStudentsPerTestRequest, opts ...grpc.CallOption) (TestService_GetStudentsPerTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[2], "/test.TestService/GetStudentsPerTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceGetStudentsPerTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_GetStudentsPerTestClient interface {
	Recv() (*studentpb.Student, error)
	grpc.ClientStream
}

type testServiceGetStudentsPerTestClient struct {
	grpc.ClientStream
}

func (x *testServiceGetStudentsPerTestClient) Recv() (*studentpb.Student, error) {
	m := new(studentpb.Student)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) TakeTest(ctx context.Context, opts ...grpc.CallOption) (TestService_TakeTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[3], "/test.TestService/TakeTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTakeTestClient{stream}
	return x, nil
}

type TestService_TakeTestClient interface {
	Send(*TakeTestRequest) error
	Recv() (*QuestionsPerTest, error)
	grpc.ClientStream
}

type testServiceTakeTestClient struct {
	grpc.ClientStream
}

func (x *testServiceTakeTestClient) Send(m *TakeTestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceTakeTestClient) Recv() (*QuestionsPerTest, error) {
	m := new(QuestionsPerTest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) GetTestScore(ctx context.Context, in *GetScoreTestRequest, opts ...grpc.CallOption) (*TestScore, error) {
	out := new(TestScore)
	err := c.cc.Invoke(ctx, "/test.TestService/GetTestScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	GetTest(context.Context, *GetTestRequest) (*Test, error)
	SetTest(context.Context, *Test) (*SetTestResponse, error)
	SetQuestions(TestService_SetQuestionsServer) error
	EnrollStudents(TestService_EnrollStudentsServer) error
	GetStudentsPerTest(*GetStudentsPerTestRequest, TestService_GetStudentsPerTestServer) error
	TakeTest(TestService_TakeTestServer) error
	GetTestScore(context.Context, *GetScoreTestRequest) (*TestScore, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) GetTest(context.Context, *GetTestRequest) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTest not implemented")
}
func (UnimplementedTestServiceServer) SetTest(context.Context, *Test) (*SetTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTest not implemented")
}
func (UnimplementedTestServiceServer) SetQuestions(TestService_SetQuestionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SetQuestions not implemented")
}
func (UnimplementedTestServiceServer) EnrollStudents(TestService_EnrollStudentsServer) error {
	return status.Errorf(codes.Unimplemented, "method EnrollStudents not implemented")
}
func (UnimplementedTestServiceServer) GetStudentsPerTest(*GetStudentsPerTestRequest, TestService_GetStudentsPerTestServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStudentsPerTest not implemented")
}
func (UnimplementedTestServiceServer) TakeTest(TestService_TakeTestServer) error {
	return status.Errorf(codes.Unimplemented, "method TakeTest not implemented")
}
func (UnimplementedTestServiceServer) GetTestScore(context.Context, *GetScoreTestRequest) (*TestScore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestScore not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_GetTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/GetTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTest(ctx, req.(*GetTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_SetTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SetTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/SetTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SetTest(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_SetQuestions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).SetQuestions(&testServiceSetQuestionsServer{stream})
}

type TestService_SetQuestionsServer interface {
	SendAndClose(*SetQuestionResponse) error
	Recv() (*Question, error)
	grpc.ServerStream
}

type testServiceSetQuestionsServer struct {
	grpc.ServerStream
}

func (x *testServiceSetQuestionsServer) SendAndClose(m *SetQuestionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceSetQuestionsServer) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_EnrollStudents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).EnrollStudents(&testServiceEnrollStudentsServer{stream})
}

type TestService_EnrollStudentsServer interface {
	SendAndClose(*SetQuestionResponse) error
	Recv() (*EnrollmentRequest, error)
	grpc.ServerStream
}

type testServiceEnrollStudentsServer struct {
	grpc.ServerStream
}

func (x *testServiceEnrollStudentsServer) SendAndClose(m *SetQuestionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceEnrollStudentsServer) Recv() (*EnrollmentRequest, error) {
	m := new(EnrollmentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_GetStudentsPerTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStudentsPerTestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).GetStudentsPerTest(m, &testServiceGetStudentsPerTestServer{stream})
}

type TestService_GetStudentsPerTestServer interface {
	Send(*studentpb.Student) error
	grpc.ServerStream
}

type testServiceGetStudentsPerTestServer struct {
	grpc.ServerStream
}

func (x *testServiceGetStudentsPerTestServer) Send(m *studentpb.Student) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_TakeTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).TakeTest(&testServiceTakeTestServer{stream})
}

type TestService_TakeTestServer interface {
	Send(*QuestionsPerTest) error
	Recv() (*TakeTestRequest, error)
	grpc.ServerStream
}

type testServiceTakeTestServer struct {
	grpc.ServerStream
}

func (x *testServiceTakeTestServer) Send(m *QuestionsPerTest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceTakeTestServer) Recv() (*TakeTestRequest, error) {
	m := new(TakeTestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_GetTestScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTestScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/GetTestScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTestScore(ctx, req.(*GetScoreTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTest",
			Handler:    _TestService_GetTest_Handler,
		},
		{
			MethodName: "SetTest",
			Handler:    _TestService_SetTest_Handler,
		},
		{
			MethodName: "GetTestScore",
			Handler:    _TestService_GetTestScore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetQuestions",
			Handler:       _TestService_SetQuestions_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "EnrollStudents",
			Handler:       _TestService_EnrollStudents_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetStudentsPerTest",
			Handler:       _TestService_GetStudentsPerTest_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TakeTest",
			Handler:       _TestService_TakeTest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "testpb/test.proto",
}
