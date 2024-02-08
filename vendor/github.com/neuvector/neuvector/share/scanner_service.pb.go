// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scanner_service.proto

package share

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScanImageRequest struct {
	Registry           string                 `protobuf:"bytes,1,opt,name=Registry" json:"Registry,omitempty"`
	Username           string                 `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	Password           string                 `protobuf:"bytes,3,opt,name=Password" json:"Password,omitempty"`
	Repository         string                 `protobuf:"bytes,4,opt,name=Repository" json:"Repository,omitempty"`
	Tag                string                 `protobuf:"bytes,5,opt,name=Tag" json:"Tag,omitempty"`
	Proxy              string                 `protobuf:"bytes,6,opt,name=Proxy" json:"Proxy,omitempty"`
	ScanLayers         bool                   `protobuf:"varint,7,opt,name=ScanLayers" json:"ScanLayers,omitempty"`
	ScanSecrets        bool                   `protobuf:"varint,8,opt,name=ScanSecrets" json:"ScanSecrets,omitempty"`
	BaseImage          string                 `protobuf:"bytes,9,opt,name=BaseImage" json:"BaseImage,omitempty"`
	RootsOfTrust       []*SigstoreRootOfTrust `protobuf:"bytes,10,rep,name=RootsOfTrust" json:"RootsOfTrust,omitempty"`
	Token              string                 `protobuf:"bytes,11,opt,name=Token" json:"Token,omitempty"`
	ScanTypesRequested *ScanTypeMap           `protobuf:"bytes,12,opt,name=ScanTypesRequested" json:"ScanTypesRequested,omitempty"`
}

func (m *ScanImageRequest) Reset()                    { *m = ScanImageRequest{} }
func (m *ScanImageRequest) String() string            { return proto.CompactTextString(m) }
func (*ScanImageRequest) ProtoMessage()               {}
func (*ScanImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ScanImageRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *ScanImageRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ScanImageRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ScanImageRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ScanImageRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ScanImageRequest) GetProxy() string {
	if m != nil {
		return m.Proxy
	}
	return ""
}

func (m *ScanImageRequest) GetScanLayers() bool {
	if m != nil {
		return m.ScanLayers
	}
	return false
}

func (m *ScanImageRequest) GetScanSecrets() bool {
	if m != nil {
		return m.ScanSecrets
	}
	return false
}

func (m *ScanImageRequest) GetBaseImage() string {
	if m != nil {
		return m.BaseImage
	}
	return ""
}

func (m *ScanImageRequest) GetRootsOfTrust() []*SigstoreRootOfTrust {
	if m != nil {
		return m.RootsOfTrust
	}
	return nil
}

func (m *ScanImageRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ScanImageRequest) GetScanTypesRequested() *ScanTypeMap {
	if m != nil {
		return m.ScanTypesRequested
	}
	return nil
}

type SigstoreRootOfTrust struct {
	Name                 string              `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	RekorPublicKey       string              `protobuf:"bytes,2,opt,name=RekorPublicKey" json:"RekorPublicKey,omitempty"`
	RootCert             string              `protobuf:"bytes,3,opt,name=RootCert" json:"RootCert,omitempty"`
	SCTPublicKey         string              `protobuf:"bytes,4,opt,name=SCTPublicKey" json:"SCTPublicKey,omitempty"`
	Verifiers            []*SigstoreVerifier `protobuf:"bytes,5,rep,name=Verifiers" json:"Verifiers,omitempty"`
	RootlessKeypairsOnly bool                `protobuf:"varint,6,opt,name=RootlessKeypairsOnly" json:"RootlessKeypairsOnly,omitempty"`
}

func (m *SigstoreRootOfTrust) Reset()                    { *m = SigstoreRootOfTrust{} }
func (m *SigstoreRootOfTrust) String() string            { return proto.CompactTextString(m) }
func (*SigstoreRootOfTrust) ProtoMessage()               {}
func (*SigstoreRootOfTrust) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *SigstoreRootOfTrust) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetRekorPublicKey() string {
	if m != nil {
		return m.RekorPublicKey
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetRootCert() string {
	if m != nil {
		return m.RootCert
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetSCTPublicKey() string {
	if m != nil {
		return m.SCTPublicKey
	}
	return ""
}

func (m *SigstoreRootOfTrust) GetVerifiers() []*SigstoreVerifier {
	if m != nil {
		return m.Verifiers
	}
	return nil
}

func (m *SigstoreRootOfTrust) GetRootlessKeypairsOnly() bool {
	if m != nil {
		return m.RootlessKeypairsOnly
	}
	return false
}

type SigstoreVerifier struct {
	Name           string                  `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Type           string                  `protobuf:"bytes,2,opt,name=Type" json:"Type,omitempty"`
	KeypairOptions *SigstoreKeypairOptions `protobuf:"bytes,3,opt,name=KeypairOptions" json:"KeypairOptions,omitempty"`
	KeylessOptions *SigstoreKeylessOptions `protobuf:"bytes,4,opt,name=KeylessOptions" json:"KeylessOptions,omitempty"`
}

func (m *SigstoreVerifier) Reset()                    { *m = SigstoreVerifier{} }
func (m *SigstoreVerifier) String() string            { return proto.CompactTextString(m) }
func (*SigstoreVerifier) ProtoMessage()               {}
func (*SigstoreVerifier) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *SigstoreVerifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SigstoreVerifier) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SigstoreVerifier) GetKeypairOptions() *SigstoreKeypairOptions {
	if m != nil {
		return m.KeypairOptions
	}
	return nil
}

func (m *SigstoreVerifier) GetKeylessOptions() *SigstoreKeylessOptions {
	if m != nil {
		return m.KeylessOptions
	}
	return nil
}

type SigstoreKeypairOptions struct {
	PublicKey string `protobuf:"bytes,1,opt,name=PublicKey" json:"PublicKey,omitempty"`
}

func (m *SigstoreKeypairOptions) Reset()                    { *m = SigstoreKeypairOptions{} }
func (m *SigstoreKeypairOptions) String() string            { return proto.CompactTextString(m) }
func (*SigstoreKeypairOptions) ProtoMessage()               {}
func (*SigstoreKeypairOptions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *SigstoreKeypairOptions) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type SigstoreKeylessOptions struct {
	CertIssuer  string `protobuf:"bytes,1,opt,name=CertIssuer" json:"CertIssuer,omitempty"`
	CertSubject string `protobuf:"bytes,2,opt,name=CertSubject" json:"CertSubject,omitempty"`
}

func (m *SigstoreKeylessOptions) Reset()                    { *m = SigstoreKeylessOptions{} }
func (m *SigstoreKeylessOptions) String() string            { return proto.CompactTextString(m) }
func (*SigstoreKeylessOptions) ProtoMessage()               {}
func (*SigstoreKeylessOptions) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *SigstoreKeylessOptions) GetCertIssuer() string {
	if m != nil {
		return m.CertIssuer
	}
	return ""
}

func (m *SigstoreKeylessOptions) GetCertSubject() string {
	if m != nil {
		return m.CertSubject
	}
	return ""
}

func init() {
	proto.RegisterType((*ScanImageRequest)(nil), "share.ScanImageRequest")
	proto.RegisterType((*SigstoreRootOfTrust)(nil), "share.SigstoreRootOfTrust")
	proto.RegisterType((*SigstoreVerifier)(nil), "share.SigstoreVerifier")
	proto.RegisterType((*SigstoreKeypairOptions)(nil), "share.SigstoreKeypairOptions")
	proto.RegisterType((*SigstoreKeylessOptions)(nil), "share.SigstoreKeylessOptions")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ScannerService service

type ScannerServiceClient interface {
	ScanRunning(ctx context.Context, in *ScanRunningRequest, opts ...grpc.CallOption) (*ScanResult, error)
	ScanImageData(ctx context.Context, in *ScanData, opts ...grpc.CallOption) (*ScanResult, error)
	ScanImage(ctx context.Context, in *ScanImageRequest, opts ...grpc.CallOption) (*ScanResult, error)
	ScanAppPackage(ctx context.Context, in *ScanAppRequest, opts ...grpc.CallOption) (*ScanResult, error)
	Ping(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*RPCVoid, error)
	ScanAwsLambda(ctx context.Context, in *ScanAwsLambdaRequest, opts ...grpc.CallOption) (*ScanResult, error)
}

type scannerServiceClient struct {
	cc *grpc.ClientConn
}

func NewScannerServiceClient(cc *grpc.ClientConn) ScannerServiceClient {
	return &scannerServiceClient{cc}
}

func (c *scannerServiceClient) ScanRunning(ctx context.Context, in *ScanRunningRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanRunning", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanImageData(ctx context.Context, in *ScanData, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanImageData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanImage(ctx context.Context, in *ScanImageRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanAppPackage(ctx context.Context, in *ScanAppRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanAppPackage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) Ping(ctx context.Context, in *RPCVoid, opts ...grpc.CallOption) (*RPCVoid, error) {
	out := new(RPCVoid)
	err := grpc.Invoke(ctx, "/share.ScannerService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scannerServiceClient) ScanAwsLambda(ctx context.Context, in *ScanAwsLambdaRequest, opts ...grpc.CallOption) (*ScanResult, error) {
	out := new(ScanResult)
	err := grpc.Invoke(ctx, "/share.ScannerService/ScanAwsLambda", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScannerService service

type ScannerServiceServer interface {
	ScanRunning(context.Context, *ScanRunningRequest) (*ScanResult, error)
	ScanImageData(context.Context, *ScanData) (*ScanResult, error)
	ScanImage(context.Context, *ScanImageRequest) (*ScanResult, error)
	ScanAppPackage(context.Context, *ScanAppRequest) (*ScanResult, error)
	Ping(context.Context, *RPCVoid) (*RPCVoid, error)
	ScanAwsLambda(context.Context, *ScanAwsLambdaRequest) (*ScanResult, error)
}

func RegisterScannerServiceServer(s *grpc.Server, srv ScannerServiceServer) {
	s.RegisterService(&_ScannerService_serviceDesc, srv)
}

func _ScannerService_ScanRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanRunningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanRunning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanRunning(ctx, req.(*ScanRunningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanImageData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanImageData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanImageData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanImageData(ctx, req.(*ScanData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanImage(ctx, req.(*ScanImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanAppPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanAppPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanAppPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanAppPackage(ctx, req.(*ScanAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).Ping(ctx, req.(*RPCVoid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScannerService_ScanAwsLambda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanAwsLambdaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).ScanAwsLambda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/share.ScannerService/ScanAwsLambda",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).ScanAwsLambda(ctx, req.(*ScanAwsLambdaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScannerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "share.ScannerService",
	HandlerType: (*ScannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanRunning",
			Handler:    _ScannerService_ScanRunning_Handler,
		},
		{
			MethodName: "ScanImageData",
			Handler:    _ScannerService_ScanImageData_Handler,
		},
		{
			MethodName: "ScanImage",
			Handler:    _ScannerService_ScanImage_Handler,
		},
		{
			MethodName: "ScanAppPackage",
			Handler:    _ScannerService_ScanAppPackage_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ScannerService_Ping_Handler,
		},
		{
			MethodName: "ScanAwsLambda",
			Handler:    _ScannerService_ScanAwsLambda_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanner_service.proto",
}

func init() { proto.RegisterFile("scanner_service.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x56, 0x9a, 0xa4, 0x24, 0x93, 0x10, 0x8a, 0x69, 0xe9, 0x12, 0x7e, 0x14, 0xe5, 0x50, 0xe5,
	0xd4, 0x43, 0x2a, 0x40, 0x02, 0x09, 0xd4, 0x16, 0x0e, 0x55, 0x0b, 0x5d, 0x79, 0x43, 0x0f, 0x5c,
	0x90, 0xb3, 0x99, 0x86, 0xa5, 0xc9, 0x7a, 0xb1, 0xbd, 0x94, 0x7d, 0x23, 0x1e, 0x85, 0xc7, 0xe0,
	0x41, 0x38, 0x20, 0xdb, 0xbb, 0x8d, 0xb3, 0xda, 0x72, 0xf3, 0x7c, 0xdf, 0x37, 0xe3, 0xf9, 0xb1,
	0x07, 0x76, 0x64, 0xc8, 0xe2, 0x18, 0xc5, 0x17, 0x89, 0xe2, 0x47, 0x14, 0xe2, 0x7e, 0x22, 0xb8,
	0xe2, 0xa4, 0x29, 0xbf, 0x32, 0x81, 0xfd, 0x6e, 0xc8, 0x97, 0x4b, 0x1e, 0x5b, 0xb0, 0x0f, 0x5a,
	0x6b, 0xcf, 0xc3, 0x5f, 0x75, 0xd8, 0x0a, 0x42, 0x16, 0x9f, 0x2c, 0xd9, 0x1c, 0x29, 0x7e, 0x4f,
	0x51, 0x2a, 0xd2, 0x87, 0x16, 0xc5, 0x79, 0x24, 0x95, 0xc8, 0xbc, 0xda, 0xa0, 0x36, 0x6a, 0xd3,
	0x1b, 0x5b, 0x73, 0x9f, 0x24, 0x8a, 0x98, 0x2d, 0xd1, 0xdb, 0xb0, 0x5c, 0x61, 0x6b, 0xce, 0x67,
	0x52, 0x5e, 0x73, 0x31, 0xf3, 0xea, 0x96, 0x2b, 0x6c, 0xf2, 0x0c, 0x80, 0x62, 0xc2, 0x65, 0xa4,
	0xb8, 0xc8, 0xbc, 0x86, 0x61, 0x1d, 0x84, 0x6c, 0x41, 0x7d, 0xc2, 0xe6, 0x5e, 0xd3, 0x10, 0xfa,
	0x48, 0xb6, 0xa1, 0xe9, 0x0b, 0xfe, 0x33, 0xf3, 0x36, 0x0d, 0x66, 0x0d, 0x1d, 0x47, 0xe7, 0x7b,
	0xc6, 0x32, 0x14, 0xd2, 0xbb, 0x33, 0xa8, 0x8d, 0x5a, 0xd4, 0x41, 0xc8, 0x00, 0x3a, 0xda, 0x0a,
	0x30, 0x14, 0xa8, 0xa4, 0xd7, 0x32, 0x02, 0x17, 0x22, 0x4f, 0xa0, 0x7d, 0xc4, 0x24, 0x9a, 0x8a,
	0xbd, 0xb6, 0x89, 0xbd, 0x02, 0xc8, 0x1b, 0xe8, 0x52, 0xce, 0x95, 0x3c, 0xbf, 0x9c, 0x88, 0x54,
	0x2a, 0x0f, 0x06, 0xf5, 0x51, 0x67, 0xdc, 0xdf, 0x37, 0x8d, 0xdc, 0x0f, 0xa2, 0xb9, 0x54, 0x5c,
	0xa0, 0x96, 0xe4, 0x0a, 0xba, 0xa6, 0xd7, 0x59, 0x4f, 0xf8, 0x15, 0xc6, 0x5e, 0xc7, 0x66, 0x6d,
	0x0c, 0x72, 0x04, 0x44, 0xa7, 0x30, 0xc9, 0x12, 0x94, 0x79, 0x97, 0x71, 0xe6, 0x75, 0x07, 0xb5,
	0x51, 0x67, 0x4c, 0x8a, 0xd8, 0xb9, 0xe0, 0x03, 0x4b, 0x68, 0x85, 0x7a, 0xf8, 0xb7, 0x06, 0x0f,
	0x2a, 0xee, 0x27, 0x04, 0x1a, 0x1f, 0xf5, 0x34, 0xec, 0xa4, 0xcc, 0x99, 0xec, 0x41, 0x8f, 0xe2,
	0x15, 0x17, 0x7e, 0x3a, 0x5d, 0x44, 0xe1, 0x29, 0x66, 0xf9, 0xac, 0x4a, 0xa8, 0x99, 0x34, 0xe7,
	0xea, 0x18, 0x85, 0x2a, 0x26, 0x56, 0xd8, 0x64, 0x08, 0xdd, 0xe0, 0x78, 0xb2, 0x8a, 0x60, 0x67,
	0xb6, 0x86, 0x91, 0xe7, 0xd0, 0xbe, 0x40, 0x11, 0x5d, 0x46, 0x7a, 0x18, 0x4d, 0xd3, 0xaa, 0xdd,
	0x52, 0xab, 0x0a, 0x9e, 0xae, 0x94, 0x64, 0x0c, 0xdb, 0xfa, 0x9a, 0x05, 0x4a, 0x79, 0x8a, 0x59,
	0xc2, 0x22, 0x21, 0xcf, 0xe3, 0x85, 0x9d, 0x74, 0x8b, 0x56, 0x72, 0xc3, 0xdf, 0x35, 0xd8, 0x2a,
	0xc7, 0xac, 0xac, 0x9d, 0x40, 0x43, 0x77, 0x2e, 0xaf, 0xd8, 0x9c, 0xc9, 0x7b, 0xe8, 0xe5, 0xc1,
	0xce, 0x13, 0x15, 0xf1, 0x58, 0x9a, 0x6a, 0x3b, 0xe3, 0xa7, 0xa5, 0x64, 0xd7, 0x45, 0xb4, 0xe4,
	0x94, 0x87, 0xd1, 0xa9, 0x15, 0x61, 0x1a, 0xb7, 0x85, 0x71, 0x44, 0xb4, 0xe4, 0x34, 0x7c, 0x01,
	0x0f, 0xab, 0x2f, 0xd4, 0x6f, 0x73, 0xd5, 0x70, 0x5b, 0xd4, 0x0a, 0x18, 0x7e, 0x5e, 0xf3, 0x73,
	0x22, 0xea, 0x5f, 0xa1, 0x67, 0x76, 0x22, 0x65, 0x8a, 0x22, 0x77, 0x74, 0x10, 0xfd, 0x2b, 0xb4,
	0x15, 0xa4, 0xd3, 0x6f, 0x18, 0xaa, 0xbc, 0x35, 0x2e, 0x34, 0xfe, 0xb3, 0x01, 0xbd, 0xc0, 0xee,
	0x90, 0xc0, 0xae, 0x10, 0xf2, 0xda, 0x7e, 0x25, 0x9a, 0xc6, 0x71, 0x14, 0xcf, 0xc9, 0x23, 0xe7,
	0x9d, 0xe6, 0x58, 0xfe, 0x38, 0xfb, 0xf7, 0x5d, 0x0a, 0x65, 0xba, 0x50, 0xe4, 0x00, 0xee, 0xde,
	0xec, 0x95, 0x77, 0x4c, 0x31, 0x72, 0xcf, 0xd1, 0x68, 0xa0, 0xca, 0xe9, 0x25, 0xb4, 0x6f, 0x9c,
	0xc8, 0xae, 0xc3, 0xbb, 0xeb, 0xa9, 0xca, 0xf1, 0x95, 0x4d, 0xfe, 0x30, 0x49, 0x7c, 0x16, 0x5e,
	0x69, 0xef, 0x1d, 0x47, 0x74, 0x98, 0x24, 0xff, 0xf1, 0xdd, 0x83, 0x86, 0xaf, 0xeb, 0xeb, 0xe5,
	0x14, 0xf5, 0x8f, 0x2f, 0x78, 0x34, 0xeb, 0x97, 0x6c, 0xf2, 0xd6, 0x56, 0x74, 0x78, 0x2d, 0xcf,
	0xd8, 0x72, 0x3a, 0x63, 0xe4, 0xb1, 0x7b, 0x45, 0x81, 0xde, 0x7e, 0xd1, 0x74, 0xd3, 0xac, 0xdc,
	0x83, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x83, 0x42, 0xc0, 0xb2, 0xac, 0x05, 0x00, 0x00,
}
