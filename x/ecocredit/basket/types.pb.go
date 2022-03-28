// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/basket/v1/types.proto

// Revision 1

package basket

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BasketCredit represents the information for a credit batch inside a basket.
type BasketCredit struct {
	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the number of credits being put into or taken out of the basket.
	// Decimal values are acceptable within the precision of the corresponding
	//  credit type for this batch.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *BasketCredit) Reset()         { *m = BasketCredit{} }
func (m *BasketCredit) String() string { return proto.CompactTextString(m) }
func (*BasketCredit) ProtoMessage()    {}
func (*BasketCredit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c256e957c69c4d, []int{0}
}
func (m *BasketCredit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasketCredit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasketCredit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasketCredit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketCredit.Merge(m, src)
}
func (m *BasketCredit) XXX_Size() int {
	return m.Size()
}
func (m *BasketCredit) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketCredit.DiscardUnknown(m)
}

var xxx_messageInfo_BasketCredit proto.InternalMessageInfo

func (m *BasketCredit) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *BasketCredit) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

// DateCriteria represents the information for credit acceptance in a basket.
// At most, only one of the values should be set.
// NOTE: gogo proto `oneof` is not compatible with Amino signing, hence we
// directly define both `start_date_window` and `min_start_date`. In the future,
// with pulsar, this should change and we should use `oneof`.
type DateCriteria struct {
	// min_start_date (optional) is the earliest start date for batches of credits
	// allowed into the basket. At most only one of `start_date_window`,
	// `min_start_date`, and `years_in_the_past` can be set for a basket.
	MinStartDate *types.Timestamp `protobuf:"bytes,1,opt,name=min_start_date,json=minStartDate,proto3" json:"min_start_date,omitempty"`
	// start_date_window (optional) is a duration of time measured into the past
	// which sets a cutoff for batch start dates when adding new credits to the
	// basket. Based on the current block timestamp, credits whose start date is
	// before `block_timestamp - start_date_window` will not be allowed into the
	// basket. At most only one of `start_date_window`, `min_start_date`, and
	// `years_in_the_past` can be set for a basket.
	StartDateWindow *types.Duration `protobuf:"bytes,2,opt,name=start_date_window,json=startDateWindow,proto3" json:"start_date_window,omitempty"`
	// years_in_the_past (optional) is the number of years into the past which
	// sets a cutoff for the batch start dates when adding new credits to the
	// basket. Based on the current block timestamp, credits whose start date year
	// is less than `block_timestamp_year - years_in_the_past` will not be allowed
	// into the basket. At most only one of `start_date_window`, `min_start_date`,
	// and `years_in_the_past` can be set for a basket.
	//
	// Since Revision 1
	YearsInThePast uint32 `protobuf:"varint,3,opt,name=years_in_the_past,json=yearsInThePast,proto3" json:"years_in_the_past,omitempty"`
}

func (m *DateCriteria) Reset()         { *m = DateCriteria{} }
func (m *DateCriteria) String() string { return proto.CompactTextString(m) }
func (*DateCriteria) ProtoMessage()    {}
func (*DateCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c256e957c69c4d, []int{1}
}
func (m *DateCriteria) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DateCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DateCriteria.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DateCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateCriteria.Merge(m, src)
}
func (m *DateCriteria) XXX_Size() int {
	return m.Size()
}
func (m *DateCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_DateCriteria.DiscardUnknown(m)
}

var xxx_messageInfo_DateCriteria proto.InternalMessageInfo

func (m *DateCriteria) GetMinStartDate() *types.Timestamp {
	if m != nil {
		return m.MinStartDate
	}
	return nil
}

func (m *DateCriteria) GetStartDateWindow() *types.Duration {
	if m != nil {
		return m.StartDateWindow
	}
	return nil
}

func (m *DateCriteria) GetYearsInThePast() uint32 {
	if m != nil {
		return m.YearsInThePast
	}
	return 0
}

func init() {
	proto.RegisterType((*BasketCredit)(nil), "regen.ecocredit.basket.v1.BasketCredit")
	proto.RegisterType((*DateCriteria)(nil), "regen.ecocredit.basket.v1.DateCriteria")
}

func init() {
	proto.RegisterFile("regen/ecocredit/basket/v1/types.proto", fileDescriptor_e6c256e957c69c4d)
}

var fileDescriptor_e6c256e957c69c4d = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x6b, 0xe2, 0x40,
	0x14, 0xc6, 0xcd, 0x2e, 0x08, 0x3b, 0xba, 0x2e, 0xe6, 0xb0, 0xa8, 0x87, 0x28, 0xc2, 0x82, 0x7b,
	0xe8, 0x0c, 0xda, 0x4b, 0x8f, 0x45, 0x2d, 0xa5, 0xb7, 0x92, 0x0a, 0x85, 0x5e, 0xc2, 0x24, 0x79,
	0x4d, 0x06, 0xcd, 0x4c, 0x98, 0x79, 0xd1, 0xfa, 0x5f, 0xf4, 0x6f, 0xea, 0xa9, 0x47, 0x8f, 0x3d,
	0x16, 0xfd, 0x47, 0x8a, 0x13, 0x6d, 0xa1, 0x1e, 0xdf, 0xfb, 0xbe, 0xef, 0x37, 0xdf, 0xf0, 0xc8,
	0x3f, 0x0d, 0x09, 0x48, 0x06, 0x91, 0x8a, 0x34, 0xc4, 0x02, 0x59, 0xc8, 0xcd, 0x1c, 0x90, 0x2d,
	0x87, 0x0c, 0xd7, 0x39, 0x18, 0x9a, 0x6b, 0x85, 0xca, 0x6d, 0x5b, 0x1b, 0xfd, 0xb4, 0xd1, 0xd2,
	0x46, 0x97, 0xc3, 0x4e, 0x37, 0x51, 0x2a, 0x59, 0x00, 0xb3, 0xc6, 0xb0, 0x78, 0x64, 0x28, 0x32,
	0x30, 0xc8, 0xb3, 0xbc, 0xcc, 0x76, 0xbc, 0xef, 0x86, 0xb8, 0xd0, 0x1c, 0x85, 0x92, 0xa5, 0xde,
	0xbf, 0x26, 0xf5, 0xb1, 0xa5, 0x4d, 0x2c, 0xda, 0xed, 0x92, 0x5a, 0xc8, 0x31, 0x4a, 0x83, 0x18,
	0xa4, 0xca, 0x5a, 0x4e, 0xcf, 0x19, 0xfc, 0xf2, 0x89, 0x5d, 0x4d, 0xf7, 0x1b, 0xf7, 0x2f, 0xa9,
	0xf2, 0x4c, 0x15, 0x12, 0x5b, 0x3f, 0xac, 0x76, 0x98, 0xfa, 0x2f, 0x0e, 0xa9, 0x4f, 0x39, 0xc2,
	0x44, 0x0b, 0x04, 0x2d, 0xb8, 0x7b, 0x49, 0x1a, 0x99, 0x90, 0x81, 0x41, 0xae, 0x31, 0x88, 0x39,
	0x82, 0x85, 0xd5, 0x46, 0x1d, 0x5a, 0x56, 0xa2, 0xc7, 0x4a, 0x74, 0x76, 0xec, 0xec, 0xd7, 0x33,
	0x21, 0xef, 0xf6, 0x81, 0x3d, 0xc9, 0xbd, 0x22, 0xcd, 0xaf, 0x74, 0xb0, 0x12, 0x32, 0x56, 0x2b,
	0xfb, 0x6a, 0x6d, 0xd4, 0x3e, 0x81, 0x4c, 0x0f, 0xff, 0xf2, 0xff, 0x98, 0x23, 0xe0, 0xde, 0x26,
	0xdc, 0xff, 0xa4, 0xb9, 0x06, 0xae, 0x4d, 0x20, 0x64, 0x80, 0x29, 0x04, 0x39, 0x37, 0xd8, 0xfa,
	0xd9, 0x73, 0x06, 0xbf, 0xfd, 0x86, 0x15, 0x6e, 0xe4, 0x2c, 0x85, 0x5b, 0x6e, 0x70, 0xec, 0xbf,
	0x6e, 0x3d, 0x67, 0xb3, 0xf5, 0x9c, 0xf7, 0xad, 0xe7, 0x3c, 0xef, 0xbc, 0xca, 0x66, 0xe7, 0x55,
	0xde, 0x76, 0x5e, 0xe5, 0xe1, 0x22, 0x11, 0x98, 0x16, 0x21, 0x8d, 0x54, 0xc6, 0xec, 0x39, 0xce,
	0x24, 0xe0, 0x4a, 0xe9, 0xf9, 0x61, 0x5a, 0x40, 0x9c, 0x80, 0x66, 0x4f, 0x27, 0xc7, 0x0c, 0xab,
	0xb6, 0xe2, 0xf9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xf9, 0xb4, 0xee, 0xed, 0x01, 0x00,
	0x00,
}

func (m *BasketCredit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasketCredit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasketCredit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DateCriteria) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DateCriteria) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DateCriteria) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.YearsInThePast != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.YearsInThePast))
		i--
		dAtA[i] = 0x18
	}
	if m.StartDateWindow != nil {
		{
			size, err := m.StartDateWindow.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MinStartDate != nil {
		{
			size, err := m.MinStartDate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasketCredit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *DateCriteria) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinStartDate != nil {
		l = m.MinStartDate.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.StartDateWindow != nil {
		l = m.StartDateWindow.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.YearsInThePast != 0 {
		n += 1 + sovTypes(uint64(m.YearsInThePast))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasketCredit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BasketCredit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasketCredit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *DateCriteria) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: DateCriteria: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DateCriteria: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStartDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinStartDate == nil {
				m.MinStartDate = &types.Timestamp{}
			}
			if err := m.MinStartDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDateWindow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartDateWindow == nil {
				m.StartDateWindow = &types.Duration{}
			}
			if err := m.StartDateWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field YearsInThePast", wireType)
			}
			m.YearsInThePast = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.YearsInThePast |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
