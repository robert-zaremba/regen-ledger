// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/marketplace/v1/types.proto

package marketplace

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

// Filter is used to create filtered buy orders which match credit batch
// sell orders based on selection criteria rather than matching individual
// sell orders
type Filter struct {
	// or is a list of criteria for matching credit batches. A credit which
	// matches this filter must match at least one of these criteria.
	Or []*Filter_Criteria `protobuf:"bytes,1,rep,name=or,proto3" json:"or,omitempty"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_eff2ca0c007ee426, []int{0}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return m.Size()
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetOr() []*Filter_Criteria {
	if m != nil {
		return m.Or
	}
	return nil
}

// Criteria is a simple filter criteria for matching a credit batch.
type Filter_Criteria struct {
	// Valid selector types are all
	// attributes which are assigned to credit batches by some authority such
	// as the credit issuer or a curator. Requiring some authority-based
	// selector ensures that buy orders cannot just match some randomly issued
	// credit based on location and dates.
	//
	// Types that are valid to be assigned to Selector:
	//	*Filter_Criteria_ClassSelector
	//	*Filter_Criteria_ProjectSelector
	//	*Filter_Criteria_BatchSelector
	Selector isFilter_Criteria_Selector `protobuf_oneof:"selector"`
}

func (m *Filter_Criteria) Reset()         { *m = Filter_Criteria{} }
func (m *Filter_Criteria) String() string { return proto.CompactTextString(m) }
func (*Filter_Criteria) ProtoMessage()    {}
func (*Filter_Criteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_eff2ca0c007ee426, []int{0, 0}
}
func (m *Filter_Criteria) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Filter_Criteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Filter_Criteria.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Filter_Criteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter_Criteria.Merge(m, src)
}
func (m *Filter_Criteria) XXX_Size() int {
	return m.Size()
}
func (m *Filter_Criteria) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter_Criteria.DiscardUnknown(m)
}

var xxx_messageInfo_Filter_Criteria proto.InternalMessageInfo

type isFilter_Criteria_Selector interface {
	isFilter_Criteria_Selector()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Filter_Criteria_ClassSelector struct {
	ClassSelector *ClassSelector `protobuf:"bytes,1,opt,name=class_selector,json=classSelector,proto3,oneof" json:"class_selector,omitempty"`
}
type Filter_Criteria_ProjectSelector struct {
	ProjectSelector *ProjectSelector `protobuf:"bytes,2,opt,name=project_selector,json=projectSelector,proto3,oneof" json:"project_selector,omitempty"`
}
type Filter_Criteria_BatchSelector struct {
	BatchSelector *BatchSelector `protobuf:"bytes,3,opt,name=batch_selector,json=batchSelector,proto3,oneof" json:"batch_selector,omitempty"`
}

func (*Filter_Criteria_ClassSelector) isFilter_Criteria_Selector()   {}
func (*Filter_Criteria_ProjectSelector) isFilter_Criteria_Selector() {}
func (*Filter_Criteria_BatchSelector) isFilter_Criteria_Selector()   {}

func (m *Filter_Criteria) GetSelector() isFilter_Criteria_Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *Filter_Criteria) GetClassSelector() *ClassSelector {
	if x, ok := m.GetSelector().(*Filter_Criteria_ClassSelector); ok {
		return x.ClassSelector
	}
	return nil
}

func (m *Filter_Criteria) GetProjectSelector() *ProjectSelector {
	if x, ok := m.GetSelector().(*Filter_Criteria_ProjectSelector); ok {
		return x.ProjectSelector
	}
	return nil
}

func (m *Filter_Criteria) GetBatchSelector() *BatchSelector {
	if x, ok := m.GetSelector().(*Filter_Criteria_BatchSelector); ok {
		return x.BatchSelector
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter_Criteria) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_Criteria_ClassSelector)(nil),
		(*Filter_Criteria_ProjectSelector)(nil),
		(*Filter_Criteria_BatchSelector)(nil),
	}
}

// ClassSelector is a selector for a credit class.
type ClassSelector struct {
	// class_id is the credit class ID.
	ClassId uint64 `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// project_location can be specified in three levels of granularity:
	// country, sub-national-code, or postal code. If just country is given,
	// for instance "US" then any credits in the "US" will be matched even
	// their project location is more specific, ex. "US-NY 12345". If
	// a country, sub-national-code and postal code are all provided then
	// only projects in that postal code will match.
	ProjectLocation string `protobuf:"bytes,2,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
	// start_date is the beginning of the period during which a credit batch
	// was quantified and verified. If it is empty then there is no start date
	// limit.
	MinStartDate *types.Timestamp `protobuf:"bytes,3,opt,name=min_start_date,json=minStartDate,proto3" json:"min_start_date,omitempty"`
	// max_end_date is the end of the period during which a credit batch was
	// quantified and verified. If it is empty then there is no end date
	// limit.
	MaxEndDate *types.Timestamp `protobuf:"bytes,4,opt,name=max_end_date,json=maxEndDate,proto3" json:"max_end_date,omitempty"`
}

func (m *ClassSelector) Reset()         { *m = ClassSelector{} }
func (m *ClassSelector) String() string { return proto.CompactTextString(m) }
func (*ClassSelector) ProtoMessage()    {}
func (*ClassSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_eff2ca0c007ee426, []int{1}
}
func (m *ClassSelector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassSelector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassSelector.Merge(m, src)
}
func (m *ClassSelector) XXX_Size() int {
	return m.Size()
}
func (m *ClassSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ClassSelector proto.InternalMessageInfo

func (m *ClassSelector) GetClassId() uint64 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *ClassSelector) GetProjectLocation() string {
	if m != nil {
		return m.ProjectLocation
	}
	return ""
}

func (m *ClassSelector) GetMinStartDate() *types.Timestamp {
	if m != nil {
		return m.MinStartDate
	}
	return nil
}

func (m *ClassSelector) GetMaxEndDate() *types.Timestamp {
	if m != nil {
		return m.MaxEndDate
	}
	return nil
}

// ProjectSelector is a selector for a project.
type ProjectSelector struct {
	// project_id is the project ID.
	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// start_date is the beginning of the period during which a credit batch
	// was quantified and verified. If it is empty then there is no start date
	// limit.
	MinStartDate *types.Timestamp `protobuf:"bytes,2,opt,name=min_start_date,json=minStartDate,proto3" json:"min_start_date,omitempty"`
	// max_end_date is the end of the period during which a credit batch was
	// quantified and verified. If it is empty then there is no end date
	// limit.
	MaxEndDate *types.Timestamp `protobuf:"bytes,3,opt,name=max_end_date,json=maxEndDate,proto3" json:"max_end_date,omitempty"`
}

func (m *ProjectSelector) Reset()         { *m = ProjectSelector{} }
func (m *ProjectSelector) String() string { return proto.CompactTextString(m) }
func (*ProjectSelector) ProtoMessage()    {}
func (*ProjectSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_eff2ca0c007ee426, []int{2}
}
func (m *ProjectSelector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProjectSelector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProjectSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectSelector.Merge(m, src)
}
func (m *ProjectSelector) XXX_Size() int {
	return m.Size()
}
func (m *ProjectSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectSelector proto.InternalMessageInfo

func (m *ProjectSelector) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *ProjectSelector) GetMinStartDate() *types.Timestamp {
	if m != nil {
		return m.MinStartDate
	}
	return nil
}

func (m *ProjectSelector) GetMaxEndDate() *types.Timestamp {
	if m != nil {
		return m.MaxEndDate
	}
	return nil
}

// BatchSelector is a selector for a credit batch.
type BatchSelector struct {
	// batch_id is the credit batch ID.
	BatchId uint64 `protobuf:"varint,1,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
}

func (m *BatchSelector) Reset()         { *m = BatchSelector{} }
func (m *BatchSelector) String() string { return proto.CompactTextString(m) }
func (*BatchSelector) ProtoMessage()    {}
func (*BatchSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_eff2ca0c007ee426, []int{3}
}
func (m *BatchSelector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchSelector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchSelector.Merge(m, src)
}
func (m *BatchSelector) XXX_Size() int {
	return m.Size()
}
func (m *BatchSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchSelector.DiscardUnknown(m)
}

var xxx_messageInfo_BatchSelector proto.InternalMessageInfo

func (m *BatchSelector) GetBatchId() uint64 {
	if m != nil {
		return m.BatchId
	}
	return 0
}

func init() {
	proto.RegisterType((*Filter)(nil), "regen.ecocredit.marketplace.v1.Filter")
	proto.RegisterType((*Filter_Criteria)(nil), "regen.ecocredit.marketplace.v1.Filter.Criteria")
	proto.RegisterType((*ClassSelector)(nil), "regen.ecocredit.marketplace.v1.ClassSelector")
	proto.RegisterType((*ProjectSelector)(nil), "regen.ecocredit.marketplace.v1.ProjectSelector")
	proto.RegisterType((*BatchSelector)(nil), "regen.ecocredit.marketplace.v1.BatchSelector")
}

func init() {
	proto.RegisterFile("regen/ecocredit/marketplace/v1/types.proto", fileDescriptor_eff2ca0c007ee426)
}

var fileDescriptor_eff2ca0c007ee426 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0x6a, 0x13, 0x41,
	0x14, 0xc6, 0xb3, 0x9b, 0x52, 0xd3, 0x69, 0xd3, 0xca, 0x5c, 0x69, 0xc0, 0x55, 0x72, 0x55, 0x0b,
	0x9d, 0xa1, 0xf5, 0x56, 0x51, 0x52, 0x15, 0x03, 0x5e, 0x48, 0x2a, 0x0a, 0x22, 0x84, 0xd9, 0x99,
	0xe3, 0x76, 0xec, 0xee, 0xce, 0x32, 0x7b, 0x5a, 0xe3, 0x5b, 0xf8, 0x10, 0x3e, 0x80, 0x8f, 0xe1,
	0x65, 0x6f, 0x04, 0x2f, 0x25, 0x79, 0x01, 0x1f, 0x41, 0x76, 0x76, 0x37, 0xd9, 0x48, 0x30, 0x8a,
	0x97, 0xf3, 0x71, 0xbe, 0x8f, 0xdf, 0xf9, 0x33, 0xe4, 0xc0, 0x42, 0x04, 0x29, 0x07, 0x69, 0xa4,
	0x05, 0xa5, 0x91, 0x27, 0xc2, 0x9e, 0x03, 0x66, 0xb1, 0x90, 0xc0, 0x2f, 0x8f, 0x38, 0x7e, 0xcc,
	0x20, 0x67, 0x99, 0x35, 0x68, 0x68, 0xe0, 0x6a, 0xd9, 0xbc, 0x96, 0x35, 0x6a, 0xd9, 0xe5, 0x51,
	0xef, 0x76, 0x64, 0x4c, 0x14, 0x03, 0x77, 0xd5, 0xe1, 0xc5, 0x3b, 0x8e, 0x3a, 0x81, 0x1c, 0x45,
	0x92, 0x95, 0x01, 0xfd, 0x9f, 0x3e, 0xd9, 0x7c, 0xaa, 0x63, 0x04, 0x4b, 0x1f, 0x12, 0xdf, 0xd8,
	0x1b, 0xde, 0x9d, 0xf6, 0xfe, 0xf6, 0x31, 0x67, 0x7f, 0x0e, 0x66, 0xa5, 0x87, 0x9d, 0x58, 0x8d,
	0x60, 0xb5, 0x18, 0xf9, 0xc6, 0xf6, 0x3e, 0xfb, 0xa4, 0x53, 0x0b, 0xf4, 0x15, 0xd9, 0x95, 0xb1,
	0xc8, 0xf3, 0x71, 0x0e, 0x31, 0x48, 0x74, 0xc9, 0xde, 0xfe, 0xf6, 0xf1, 0xe1, 0xba, 0xe4, 0x93,
	0xc2, 0x75, 0x5a, 0x99, 0x9e, 0xb5, 0x46, 0x5d, 0xd9, 0x14, 0xe8, 0x5b, 0x72, 0x3d, 0xb3, 0xe6,
	0x3d, 0x48, 0x5c, 0x24, 0xfb, 0x2e, 0x79, 0x2d, 0xf3, 0x8b, 0xd2, 0xd7, 0xc8, 0xde, 0xcb, 0x96,
	0xa5, 0x82, 0x3a, 0x14, 0x28, 0xcf, 0x16, 0xd9, 0xed, 0xbf, 0xa3, 0x1e, 0x14, 0xae, 0x26, 0x75,
	0xd8, 0x14, 0x06, 0x84, 0x74, 0xea, 0xc4, 0xfe, 0x37, 0x8f, 0x74, 0x97, 0x9a, 0xa4, 0x37, 0x49,
	0xa7, 0x9c, 0x95, 0x56, 0x6e, 0x4a, 0x1b, 0xa3, 0x6b, 0xee, 0x3d, 0x54, 0xf4, 0xee, 0xa2, 0xdd,
	0xd8, 0x48, 0x81, 0xda, 0xa4, 0xae, 0xdd, 0xad, 0x39, 0xfb, 0xf3, 0x4a, 0xa6, 0x8f, 0xc8, 0x6e,
	0xa2, 0xd3, 0x71, 0x8e, 0xc2, 0xe2, 0x58, 0x09, 0x84, 0x8a, 0xbd, 0xc7, 0xca, 0x23, 0x60, 0xf5,
	0x11, 0xb0, 0x97, 0xf5, 0x11, 0x8c, 0x76, 0x12, 0x9d, 0x9e, 0x16, 0x86, 0xc7, 0x02, 0x81, 0xde,
	0x27, 0x3b, 0x89, 0x98, 0x8c, 0x21, 0x55, 0xa5, 0x7f, 0x63, 0xad, 0x9f, 0x24, 0x62, 0xf2, 0x24,
	0x55, 0x85, 0xbb, 0xff, 0xc5, 0x23, 0x7b, 0xbf, 0x8d, 0x98, 0xde, 0x22, 0xa4, 0xc6, 0x9f, 0xf7,
	0xb6, 0x55, 0x29, 0x43, 0xb5, 0x02, 0xd9, 0xff, 0x4f, 0xe4, 0xf6, 0x3f, 0x21, 0x1f, 0x90, 0xee,
	0xd2, 0xe2, 0x8a, 0x4d, 0x94, 0xfb, 0x5f, 0x6c, 0xc2, 0xbd, 0x87, 0x6a, 0xf0, 0xfa, 0xeb, 0x34,
	0xf0, 0xae, 0xa6, 0x81, 0xf7, 0x63, 0x1a, 0x78, 0x9f, 0x66, 0x41, 0xeb, 0x6a, 0x16, 0xb4, 0xbe,
	0xcf, 0x82, 0xd6, 0x9b, 0x07, 0x91, 0xc6, 0xb3, 0x8b, 0x90, 0x49, 0x93, 0x70, 0x77, 0x26, 0x87,
	0x29, 0xe0, 0x07, 0x63, 0xcf, 0xab, 0x57, 0x0c, 0x2a, 0x02, 0xcb, 0x27, 0xab, 0xbf, 0x74, 0xb8,
	0xe9, 0x20, 0xef, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x3d, 0x56, 0xea, 0xf8, 0x03, 0x00,
	0x00,
}

func (m *Filter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Filter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Filter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Or) > 0 {
		for iNdEx := len(m.Or) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Or[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Filter_Criteria) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Filter_Criteria) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Filter_Criteria) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Selector != nil {
		{
			size := m.Selector.Size()
			i -= size
			if _, err := m.Selector.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Filter_Criteria_ClassSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Filter_Criteria_ClassSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ClassSelector != nil {
		{
			size, err := m.ClassSelector.MarshalToSizedBuffer(dAtA[:i])
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
func (m *Filter_Criteria_ProjectSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Filter_Criteria_ProjectSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProjectSelector != nil {
		{
			size, err := m.ProjectSelector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Filter_Criteria_BatchSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Filter_Criteria_BatchSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BatchSelector != nil {
		{
			size, err := m.BatchSelector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *ClassSelector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxEndDate != nil {
		{
			size, err := m.MaxEndDate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
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
		dAtA[i] = 0x1a
	}
	if len(m.ProjectLocation) > 0 {
		i -= len(m.ProjectLocation)
		copy(dAtA[i:], m.ProjectLocation)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ProjectLocation)))
		i--
		dAtA[i] = 0x12
	}
	if m.ClassId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ClassId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProjectSelector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxEndDate != nil {
		{
			size, err := m.MaxEndDate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
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
		dAtA[i] = 0x12
	}
	if m.ProjectId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ProjectId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BatchSelector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BatchId))
		i--
		dAtA[i] = 0x8
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
func (m *Filter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Or) > 0 {
		for _, e := range m.Or {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Filter_Criteria) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Selector != nil {
		n += m.Selector.Size()
	}
	return n
}

func (m *Filter_Criteria_ClassSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClassSelector != nil {
		l = m.ClassSelector.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Filter_Criteria_ProjectSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProjectSelector != nil {
		l = m.ProjectSelector.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Filter_Criteria_BatchSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchSelector != nil {
		l = m.BatchSelector.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *ClassSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClassId != 0 {
		n += 1 + sovTypes(uint64(m.ClassId))
	}
	l = len(m.ProjectLocation)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.MinStartDate != nil {
		l = m.MinStartDate.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.MaxEndDate != nil {
		l = m.MaxEndDate.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ProjectSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProjectId != 0 {
		n += 1 + sovTypes(uint64(m.ProjectId))
	}
	if m.MinStartDate != nil {
		l = m.MinStartDate.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.MaxEndDate != nil {
		l = m.MaxEndDate.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *BatchSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchId != 0 {
		n += 1 + sovTypes(uint64(m.BatchId))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Filter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Filter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Filter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Or", wireType)
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
			m.Or = append(m.Or, &Filter_Criteria{})
			if err := m.Or[len(m.Or)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *Filter_Criteria) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Criteria: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Criteria: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassSelector", wireType)
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
			v := &ClassSelector{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Selector = &Filter_Criteria_ClassSelector{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectSelector", wireType)
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
			v := &ProjectSelector{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Selector = &Filter_Criteria_ProjectSelector{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSelector", wireType)
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
			v := &BatchSelector{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Selector = &Filter_Criteria_BatchSelector{v}
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
func (m *ClassSelector) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ClassSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			m.ClassId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClassId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectLocation", wireType)
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
			m.ProjectLocation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEndDate", wireType)
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
			if m.MaxEndDate == nil {
				m.MaxEndDate = &types.Timestamp{}
			}
			if err := m.MaxEndDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *ProjectSelector) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProjectSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			m.ProjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEndDate", wireType)
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
			if m.MaxEndDate == nil {
				m.MaxEndDate = &types.Timestamp{}
			}
			if err := m.MaxEndDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *BatchSelector) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BatchSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
			}
			m.BatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchId |= uint64(b&0x7F) << shift
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
