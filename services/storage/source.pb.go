// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReadSource struct {
	// Database identifies which database to query.
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// RetentionPolicy identifies which retention policy to query.
	RetentionPolicy string `protobuf:"bytes,2,opt,name=retention_policy,json=retentionPolicy,proto3" json:"retention_policy,omitempty"`
}

func (m *ReadSource) Reset()                    { *m = ReadSource{} }
func (m *ReadSource) String() string            { return proto.CompactTextString(m) }
func (*ReadSource) ProtoMessage()               {}
func (*ReadSource) Descriptor() ([]byte, []int) { return fileDescriptorStorage, []int{0} }

func init() {
	proto.RegisterType((*ReadSource)(nil), "com.github.influxdata.influxdb.services.storage.ReadSource")
}
func (m *ReadSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Database) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Database)))
		i += copy(dAtA[i:], m.Database)
	}
	if len(m.RetentionPolicy) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStorage(dAtA, i, uint64(len(m.RetentionPolicy)))
		i += copy(dAtA[i:], m.RetentionPolicy)
	}
	return i, nil
}

func encodeFixed64Storage(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Storage(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReadSource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Database)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	l = len(m.RetentionPolicy)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}

func sovStorage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReadSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Database = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetentionPolicy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetentionPolicy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
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
					return 0, ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStorage
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStorage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStorage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStorage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStorage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage.proto", fileDescriptorStorage) }

var fileDescriptorStorage = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x4f, 0xce, 0xcf, 0xd5, 0x4b,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0xcb, 0xcc, 0x4b, 0xcb, 0x29, 0xad, 0x48, 0x49, 0x2c, 0x49,
	0x84, 0x31, 0x93, 0xf4, 0x8a, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x8b, 0xf5, 0xa0, 0xda, 0xa4,
	0x74, 0xa1, 0x8a, 0x93, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xe6, 0x24, 0x95,
	0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0x31, 0x5f, 0x29, 0x83, 0x8b, 0x2b, 0x28, 0x35, 0x31,
	0x25, 0x38, 0xbf, 0xb4, 0x28, 0x39, 0x55, 0x48, 0x8a, 0x8b, 0x03, 0x64, 0x7c, 0x52, 0x62, 0x71,
	0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0x64, 0xc7, 0x25, 0x50, 0x94, 0x5a,
	0x92, 0x9a, 0x57, 0x92, 0x99, 0x9f, 0x17, 0x5f, 0x90, 0x9f, 0x93, 0x99, 0x5c, 0x29, 0xc1, 0x04,
	0x52, 0xe3, 0x24, 0xfc, 0xe8, 0x9e, 0x3c, 0x7f, 0x10, 0x4c, 0x2e, 0x00, 0x2c, 0x15, 0xc4, 0x5f,
	0x84, 0x2a, 0xe0, 0x24, 0x7b, 0xe2, 0xa1, 0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9,
	0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43, 0x14, 0x3b, 0xd4, 0xdd, 0x49, 0x6c,
	0x60, 0xf7, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x44, 0xf3, 0x2d, 0x00, 0x01, 0x00,
	0x00,
}