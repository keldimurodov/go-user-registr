// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: recommendation-service/recommendation.proto

package recommendation

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Status struct {
	Liked                bool     `protobuf:"varint,1,opt,name=liked,proto3" json:"liked"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0a93de7db22fcf5, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetLiked() bool {
	if m != nil {
		return m.Liked
	}
	return false
}

type PostLikeOwners struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	PostId               int64    `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostLikeOwners) Reset()         { *m = PostLikeOwners{} }
func (m *PostLikeOwners) String() string { return proto.CompactTextString(m) }
func (*PostLikeOwners) ProtoMessage()    {}
func (*PostLikeOwners) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0a93de7db22fcf5, []int{1}
}
func (m *PostLikeOwners) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostLikeOwners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostLikeOwners.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostLikeOwners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostLikeOwners.Merge(m, src)
}
func (m *PostLikeOwners) XXX_Size() int {
	return m.Size()
}
func (m *PostLikeOwners) XXX_DiscardUnknown() {
	xxx_messageInfo_PostLikeOwners.DiscardUnknown(m)
}

var xxx_messageInfo_PostLikeOwners proto.InternalMessageInfo

func (m *PostLikeOwners) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PostLikeOwners) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "recommendation.Status")
	proto.RegisterType((*PostLikeOwners)(nil), "recommendation.PostLikeOwners")
}

func init() {
	proto.RegisterFile("recommendation-service/recommendation.proto", fileDescriptor_a0a93de7db22fcf5)
}

var fileDescriptor_a0a93de7db22fcf5 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x4a, 0x4d, 0xce,
	0xcf, 0xcd, 0x4d, 0xcd, 0x4b, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x47, 0x15, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x43, 0x15, 0x55,
	0x92, 0xe3, 0x62, 0x0b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x16, 0x12, 0xe1, 0x62, 0xcd, 0xc9, 0xcc,
	0x4e, 0x4d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x70, 0x94, 0x9c, 0xb8, 0xf8, 0x02,
	0xf2, 0x8b, 0x4b, 0x7c, 0x32, 0xb3, 0x53, 0xfd, 0xcb, 0xf3, 0x52, 0x8b, 0x8a, 0x85, 0xc4, 0xb9,
	0xd8, 0x4b, 0x8b, 0x53, 0x8b, 0xe2, 0x33, 0x21, 0x2a, 0x39, 0x83, 0xd8, 0x40, 0x5c, 0xcf, 0x14,
	0x90, 0x44, 0x41, 0x7e, 0x71, 0x09, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x88, 0x0d, 0xc4,
	0xf5, 0x4c, 0x31, 0x8a, 0xe6, 0x12, 0x0d, 0x42, 0xb1, 0x35, 0x18, 0xe2, 0x42, 0x21, 0x27, 0x2e,
	0x0e, 0x90, 0xc1, 0x20, 0x0b, 0x84, 0xe4, 0xf4, 0xd0, 0xdc, 0x8b, 0x6a, 0xad, 0x94, 0x18, 0xba,
	0x3c, 0xc4, 0xd9, 0x4e, 0x8a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x51, 0xfc, 0x68, 0x3e, 0x4f, 0x62, 0x03, 0x7b, 0xdd, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x2f, 0xec, 0xd0, 0x29, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecommendationServiceClient is the client API for RecommendationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommendationServiceClient interface {
	LikePost(ctx context.Context, in *PostLikeOwners, opts ...grpc.CallOption) (*Status, error)
}

type recommendationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecommendationServiceClient(cc *grpc.ClientConn) RecommendationServiceClient {
	return &recommendationServiceClient{cc}
}

func (c *recommendationServiceClient) LikePost(ctx context.Context, in *PostLikeOwners, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/recommendation.RecommendationService/LikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceServer is the server API for RecommendationService service.
type RecommendationServiceServer interface {
	LikePost(context.Context, *PostLikeOwners) (*Status, error)
}

// UnimplementedRecommendationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecommendationServiceServer struct {
}

func (*UnimplementedRecommendationServiceServer) LikePost(ctx context.Context, req *PostLikeOwners) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}

func RegisterRecommendationServiceServer(s *grpc.Server, srv RecommendationServiceServer) {
	s.RegisterService(&_RecommendationService_serviceDesc, srv)
}

func _RecommendationService_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostLikeOwners)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recommendation.RecommendationService/LikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).LikePost(ctx, req.(*PostLikeOwners))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecommendationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recommendation.RecommendationService",
	HandlerType: (*RecommendationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LikePost",
			Handler:    _RecommendationService_LikePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommendation-service/recommendation.proto",
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Liked {
		i--
		if m.Liked {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PostLikeOwners) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostLikeOwners) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostLikeOwners) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PostId != 0 {
		i = encodeVarintRecommendation(dAtA, i, uint64(m.PostId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.UserId) > 0 {
		i -= len(m.UserId)
		copy(dAtA[i:], m.UserId)
		i = encodeVarintRecommendation(dAtA, i, uint64(len(m.UserId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecommendation(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecommendation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Liked {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PostLikeOwners) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovRecommendation(uint64(l))
	}
	if m.PostId != 0 {
		n += 1 + sovRecommendation(uint64(m.PostId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRecommendation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecommendation(x uint64) (n int) {
	return sovRecommendation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecommendation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liked", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecommendation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Liked = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRecommendation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecommendation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PostLikeOwners) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecommendation
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PostLikeOwners: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostLikeOwners: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecommendation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRecommendation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRecommendation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
			}
			m.PostId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecommendation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecommendation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecommendation
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRecommendation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecommendation
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecommendation
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecommendation
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRecommendation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecommendation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecommendation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecommendation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecommendation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecommendation = fmt.Errorf("proto: unexpected end of group")
)
