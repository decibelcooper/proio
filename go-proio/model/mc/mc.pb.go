// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proio/model/mc.proto

/*
	Package mc is a generated protocol buffer package.

	It is generated from these files:
		proio/model/mc.proto

	It has these top-level messages:
		MCParameters
		MapInt
		MapDouble
		Pythia8Parameters
*/
package mc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// this message is for general Monte Carlo generators
type MCParameters struct {
	Number    *uint64  `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Processid *int32   `protobuf:"varint,2,opt,name=processid" json:"processid,omitempty"`
	Weight    *float64 `protobuf:"fixed64,3,opt,name=weight" json:"weight,omitempty"`
	// keep metadata as key-value (int)
	Imap []*MapInt `protobuf:"bytes,4,rep,name=imap" json:"imap,omitempty"`
	// keep metadata as key-value (double)
	Dmap             []*MapDouble `protobuf:"bytes,5,rep,name=dmap" json:"dmap,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MCParameters) Reset()                    { *m = MCParameters{} }
func (m *MCParameters) String() string            { return proto.CompactTextString(m) }
func (*MCParameters) ProtoMessage()               {}
func (*MCParameters) Descriptor() ([]byte, []int) { return fileDescriptorMc, []int{0} }

func (m *MCParameters) GetNumber() uint64 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *MCParameters) GetProcessid() int32 {
	if m != nil && m.Processid != nil {
		return *m.Processid
	}
	return 0
}

func (m *MCParameters) GetWeight() float64 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return 0
}

func (m *MCParameters) GetImap() []*MapInt {
	if m != nil {
		return m.Imap
	}
	return nil
}

func (m *MCParameters) GetDmap() []*MapDouble {
	if m != nil {
		return m.Dmap
	}
	return nil
}

// map to store arbitrary data as key-int value
type MapInt struct {
	// key for integer value
	Key *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	// value
	Value            []int32 `protobuf:"zigzag32,2,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MapInt) Reset()                    { *m = MapInt{} }
func (m *MapInt) String() string            { return proto.CompactTextString(m) }
func (*MapInt) ProtoMessage()               {}
func (*MapInt) Descriptor() ([]byte, []int) { return fileDescriptorMc, []int{1} }

func (m *MapInt) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *MapInt) GetValue() []int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

// map to store arbitrary data as key-double value
type MapDouble struct {
	// key for double value
	Key *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	// value
	Value            []float64 `protobuf:"fixed64,2,rep,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MapDouble) Reset()                    { *m = MapDouble{} }
func (m *MapDouble) String() string            { return proto.CompactTextString(m) }
func (*MapDouble) ProtoMessage()               {}
func (*MapDouble) Descriptor() ([]byte, []int) { return fileDescriptorMc, []int{2} }

func (m *MapDouble) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *MapDouble) GetValue() []float64 {
	if m != nil {
		return m.Value
	}
	return nil
}

// this block is only for Pythia8
type Pythia8Parameters struct {
	WeightSum     *float64 `protobuf:"fixed64,1,opt,name=weight_sum,json=weightSum" json:"weight_sum,omitempty"`
	MergingWeight *float64 `protobuf:"fixed64,2,opt,name=merging_weight,json=mergingWeight" json:"merging_weight,omitempty"`
	// transverse momentum
	PtHat   *float64 `protobuf:"fixed64,3,opt,name=pt_hat,json=ptHat" json:"pt_hat,omitempty"`
	AlphaEm *float64 `protobuf:"fixed64,4,opt,name=alpha_em,json=alphaEm" json:"alpha_em,omitempty"`
	AlphaS  *float64 `protobuf:"fixed64,5,opt,name=alpha_s,json=alphaS" json:"alpha_s,omitempty"`
	// Q-scale used in evaluation of PDF’s (in GeV)
	ScaleQFac *float64 `protobuf:"fixed64,6,opt,name=scale_q_fac,json=scaleQFac" json:"scale_q_fac,omitempty"`
	// event weight
	Weight *float64 `protobuf:"fixed64,7,opt,name=weight" json:"weight,omitempty"`
	// fraction of beam momentum carried by first parton (”beam side”)
	X1 *float64 `protobuf:"fixed64,8,opt,name=x1" json:"x1,omitempty"`
	// fraction of beam momentum carried by second parton (”target side”)
	X2 *float64 `protobuf:"fixed64,9,opt,name=x2" json:"x2,omitempty"`
	// flavour code of first parton
	Id1 *uint64 `protobuf:"varint,10,opt,name=id1" json:"id1,omitempty"`
	// flavour code of second parton
	Id2              *uint64 `protobuf:"varint,11,opt,name=id2" json:"id2,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pythia8Parameters) Reset()                    { *m = Pythia8Parameters{} }
func (m *Pythia8Parameters) String() string            { return proto.CompactTextString(m) }
func (*Pythia8Parameters) ProtoMessage()               {}
func (*Pythia8Parameters) Descriptor() ([]byte, []int) { return fileDescriptorMc, []int{3} }

func (m *Pythia8Parameters) GetWeightSum() float64 {
	if m != nil && m.WeightSum != nil {
		return *m.WeightSum
	}
	return 0
}

func (m *Pythia8Parameters) GetMergingWeight() float64 {
	if m != nil && m.MergingWeight != nil {
		return *m.MergingWeight
	}
	return 0
}

func (m *Pythia8Parameters) GetPtHat() float64 {
	if m != nil && m.PtHat != nil {
		return *m.PtHat
	}
	return 0
}

func (m *Pythia8Parameters) GetAlphaEm() float64 {
	if m != nil && m.AlphaEm != nil {
		return *m.AlphaEm
	}
	return 0
}

func (m *Pythia8Parameters) GetAlphaS() float64 {
	if m != nil && m.AlphaS != nil {
		return *m.AlphaS
	}
	return 0
}

func (m *Pythia8Parameters) GetScaleQFac() float64 {
	if m != nil && m.ScaleQFac != nil {
		return *m.ScaleQFac
	}
	return 0
}

func (m *Pythia8Parameters) GetWeight() float64 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return 0
}

func (m *Pythia8Parameters) GetX1() float64 {
	if m != nil && m.X1 != nil {
		return *m.X1
	}
	return 0
}

func (m *Pythia8Parameters) GetX2() float64 {
	if m != nil && m.X2 != nil {
		return *m.X2
	}
	return 0
}

func (m *Pythia8Parameters) GetId1() uint64 {
	if m != nil && m.Id1 != nil {
		return *m.Id1
	}
	return 0
}

func (m *Pythia8Parameters) GetId2() uint64 {
	if m != nil && m.Id2 != nil {
		return *m.Id2
	}
	return 0
}

func init() {
	proto.RegisterType((*MCParameters)(nil), "proio.model.mc.MCParameters")
	proto.RegisterType((*MapInt)(nil), "proio.model.mc.MapInt")
	proto.RegisterType((*MapDouble)(nil), "proio.model.mc.MapDouble")
	proto.RegisterType((*Pythia8Parameters)(nil), "proio.model.mc.Pythia8Parameters")
}
func (m *MCParameters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MCParameters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Number != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Number))
	}
	if m.Processid != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Processid))
	}
	if m.Weight != nil {
		dAtA[i] = 0x19
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.Weight))))
		i += 8
	}
	if len(m.Imap) > 0 {
		for _, msg := range m.Imap {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Dmap) > 0 {
		for _, msg := range m.Dmap {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintMc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *MapInt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapInt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMc(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if len(m.Value) > 0 {
		for _, num := range m.Value {
			dAtA[i] = 0x10
			i++
			x1 := (uint32(num) << 1) ^ uint32((num >> 31))
			for x1 >= 1<<7 {
				dAtA[i] = uint8(uint64(x1)&0x7f | 0x80)
				x1 >>= 7
				i++
			}
			dAtA[i] = uint8(x1)
			i++
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *MapDouble) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapDouble) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key == nil {
		return 0, new(proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMc(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if len(m.Value) > 0 {
		for _, num := range m.Value {
			dAtA[i] = 0x11
			i++
			f2 := math.Float64bits(float64(num))
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(f2))
			i += 8
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Pythia8Parameters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pythia8Parameters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.WeightSum != nil {
		dAtA[i] = 0x9
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.WeightSum))))
		i += 8
	}
	if m.MergingWeight != nil {
		dAtA[i] = 0x11
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.MergingWeight))))
		i += 8
	}
	if m.PtHat != nil {
		dAtA[i] = 0x19
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.PtHat))))
		i += 8
	}
	if m.AlphaEm != nil {
		dAtA[i] = 0x21
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.AlphaEm))))
		i += 8
	}
	if m.AlphaS != nil {
		dAtA[i] = 0x29
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.AlphaS))))
		i += 8
	}
	if m.ScaleQFac != nil {
		dAtA[i] = 0x31
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.ScaleQFac))))
		i += 8
	}
	if m.Weight != nil {
		dAtA[i] = 0x39
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.Weight))))
		i += 8
	}
	if m.X1 != nil {
		dAtA[i] = 0x41
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.X1))))
		i += 8
	}
	if m.X2 != nil {
		dAtA[i] = 0x49
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.X2))))
		i += 8
	}
	if m.Id1 != nil {
		dAtA[i] = 0x50
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Id1))
	}
	if m.Id2 != nil {
		dAtA[i] = 0x58
		i++
		i = encodeVarintMc(dAtA, i, uint64(*m.Id2))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MCParameters) Size() (n int) {
	var l int
	_ = l
	if m.Number != nil {
		n += 1 + sovMc(uint64(*m.Number))
	}
	if m.Processid != nil {
		n += 1 + sovMc(uint64(*m.Processid))
	}
	if m.Weight != nil {
		n += 9
	}
	if len(m.Imap) > 0 {
		for _, e := range m.Imap {
			l = e.Size()
			n += 1 + l + sovMc(uint64(l))
		}
	}
	if len(m.Dmap) > 0 {
		for _, e := range m.Dmap {
			l = e.Size()
			n += 1 + l + sovMc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MapInt) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovMc(uint64(l))
	}
	if len(m.Value) > 0 {
		for _, e := range m.Value {
			n += 1 + sozMc(uint64(e))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MapDouble) Size() (n int) {
	var l int
	_ = l
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovMc(uint64(l))
	}
	if len(m.Value) > 0 {
		n += 9 * len(m.Value)
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Pythia8Parameters) Size() (n int) {
	var l int
	_ = l
	if m.WeightSum != nil {
		n += 9
	}
	if m.MergingWeight != nil {
		n += 9
	}
	if m.PtHat != nil {
		n += 9
	}
	if m.AlphaEm != nil {
		n += 9
	}
	if m.AlphaS != nil {
		n += 9
	}
	if m.ScaleQFac != nil {
		n += 9
	}
	if m.Weight != nil {
		n += 9
	}
	if m.X1 != nil {
		n += 9
	}
	if m.X2 != nil {
		n += 9
	}
	if m.Id1 != nil {
		n += 1 + sovMc(uint64(*m.Id1))
	}
	if m.Id2 != nil {
		n += 1 + sovMc(uint64(*m.Id2))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMc(x uint64) (n int) {
	return sovMc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MCParameters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMc
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
			return fmt.Errorf("proto: MCParameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MCParameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Number = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Processid", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Processid = &v
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.Weight = &v2
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Imap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Imap = append(m.Imap, &MapInt{})
			if err := m.Imap[len(m.Imap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dmap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dmap = append(m.Dmap, &MapDouble{})
			if err := m.Dmap[len(m.Dmap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMc
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
func (m *MapInt) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMc
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
			return fmt.Errorf("proto: MapInt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapInt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
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
				return ErrInvalidLengthMc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
				m.Value = append(m.Value, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMc
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
					m.Value = append(m.Value, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MapDouble) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMc
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
			return fmt.Errorf("proto: MapDouble: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapDouble: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
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
				return ErrInvalidLengthMc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.Value = append(m.Value, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMc
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.Value = append(m.Value, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Pythia8Parameters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMc
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
			return fmt.Errorf("proto: Pythia8Parameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pythia8Parameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field WeightSum", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.WeightSum = &v2
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field MergingWeight", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.MergingWeight = &v2
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field PtHat", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.PtHat = &v2
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlphaEm", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.AlphaEm = &v2
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlphaS", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.AlphaS = &v2
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScaleQFac", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.ScaleQFac = &v2
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.Weight = &v2
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X1", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.X1 = &v2
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X2", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.X2 = &v2
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id1", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Id1 = &v
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id2", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Id2 = &v
		default:
			iNdEx = preIndex
			skippy, err := skipMc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMc
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
func skipMc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMc
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
					return 0, ErrIntOverflowMc
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
					return 0, ErrIntOverflowMc
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
				return 0, ErrInvalidLengthMc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMc
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
				next, err := skipMc(dAtA[start:])
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
	ErrInvalidLengthMc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proio/model/mc.proto", fileDescriptorMc) }

var fileDescriptorMc = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x71, 0x9a, 0x74, 0xcb, 0x29, 0x4c, 0xcc, 0x1a, 0xc3, 0x93, 0x20, 0x8a, 0x2a, 0x21,
	0x45, 0x48, 0x4b, 0xd7, 0x70, 0xc3, 0xf5, 0xf8, 0x23, 0xb8, 0xa8, 0x34, 0xd2, 0x0b, 0x24, 0x6e,
	0x22, 0xd7, 0x31, 0x49, 0x44, 0x5c, 0x9b, 0xc4, 0x81, 0xed, 0x4d, 0xb8, 0xe4, 0x2d, 0x78, 0x05,
	0x2e, 0x79, 0x04, 0x54, 0x5e, 0x04, 0xc5, 0x4e, 0xa1, 0x95, 0x90, 0xb8, 0x3b, 0xdf, 0x2f, 0xdf,
	0x89, 0xed, 0xef, 0x1c, 0x38, 0x51, 0x8d, 0xac, 0xe4, 0x4c, 0xc8, 0x9c, 0xd7, 0x33, 0xc1, 0x62,
	0xd5, 0x48, 0x2d, 0xf1, 0x91, 0xa1, 0xb1, 0xa1, 0xb1, 0x60, 0xd3, 0x6f, 0x08, 0x6e, 0x2f, 0x9e,
	0x5d, 0xd1, 0x86, 0x0a, 0xae, 0x79, 0xd3, 0xe2, 0x53, 0x18, 0xaf, 0x3b, 0xb1, 0xe2, 0x0d, 0x41,
	0x21, 0x8a, 0xdc, 0x74, 0x50, 0xf8, 0x01, 0xf8, 0xaa, 0x91, 0x8c, 0xb7, 0x6d, 0x95, 0x13, 0x27,
	0x44, 0x91, 0x97, 0xfe, 0x05, 0x7d, 0xd7, 0x67, 0x5e, 0x15, 0xa5, 0x26, 0xa3, 0x10, 0x45, 0x28,
	0x1d, 0x14, 0x7e, 0x0c, 0x6e, 0x25, 0xa8, 0x22, 0x6e, 0x38, 0x8a, 0x26, 0xc9, 0x69, 0xbc, 0x7f,
	0x7a, 0xbc, 0xa0, 0xea, 0xf5, 0x5a, 0xa7, 0xc6, 0x83, 0xcf, 0xc1, 0xcd, 0x7b, 0xaf, 0x67, 0xbc,
	0x67, 0xff, 0xf0, 0x3e, 0x97, 0xdd, 0xaa, 0xe6, 0xa9, 0xb1, 0x4d, 0x2f, 0x60, 0x6c, 0xdb, 0xf1,
	0x5d, 0x18, 0x7d, 0xe0, 0x37, 0x04, 0x85, 0x4e, 0xe4, 0xa7, 0x7d, 0x89, 0x4f, 0xc0, 0xfb, 0x44,
	0xeb, 0x8e, 0x13, 0x27, 0x1c, 0x45, 0xc7, 0xa9, 0x15, 0xd3, 0x27, 0xe0, 0xff, 0xf9, 0xc9, 0xff,
	0x9a, 0xd0, 0xb6, 0xe9, 0xab, 0x03, 0xc7, 0x57, 0x37, 0xba, 0xac, 0xe8, 0xd3, 0x9d, 0x94, 0x1e,
	0x02, 0xd8, 0x17, 0x66, 0x6d, 0x27, 0x4c, 0x52, 0x28, 0xf5, 0x2d, 0x59, 0x76, 0x02, 0x3f, 0x82,
	0x23, 0xc1, 0x9b, 0xa2, 0x5a, 0x17, 0xd9, 0x10, 0x8b, 0x63, 0x2c, 0x77, 0x06, 0xfa, 0xd6, 0xa6,
	0x73, 0x0f, 0xc6, 0x4a, 0x67, 0x25, 0xdd, 0xa6, 0xe6, 0x29, 0xfd, 0x8a, 0x6a, 0x7c, 0x06, 0x87,
	0xb4, 0x56, 0x25, 0xcd, 0xb8, 0x20, 0xae, 0xf9, 0x70, 0x60, 0xf4, 0x0b, 0x81, 0xef, 0x83, 0x2d,
	0xb3, 0x96, 0x78, 0x36, 0x68, 0x23, 0x97, 0x38, 0x80, 0x49, 0xcb, 0x68, 0xcd, 0xb3, 0x8f, 0xd9,
	0x7b, 0xca, 0xc8, 0xd8, 0xde, 0xc8, 0xa0, 0x37, 0x2f, 0x29, 0xdb, 0x19, 0xd0, 0xc1, 0xde, 0x80,
	0x8e, 0xc0, 0xb9, 0x9e, 0x93, 0x43, 0xc3, 0x9c, 0xeb, 0xb9, 0xd1, 0x09, 0xf1, 0x07, 0x9d, 0xf4,
	0x31, 0x55, 0xf9, 0x9c, 0x80, 0xd9, 0x85, 0xbe, 0xb4, 0x24, 0x21, 0x93, 0x2d, 0x49, 0x2e, 0x97,
	0xdf, 0x37, 0x01, 0xfa, 0xb1, 0x09, 0xd0, 0xcf, 0x4d, 0x80, 0xbe, 0xfc, 0x0a, 0x6e, 0xc1, 0x64,
	0x67, 0x76, 0x97, 0xce, 0x82, 0xbd, 0xbb, 0x28, 0x2a, 0x5d, 0x76, 0xab, 0x98, 0x49, 0x31, 0xcb,
	0x39, 0xab, 0x56, 0xbc, 0x66, 0x52, 0x2a, 0xde, 0xcc, 0xec, 0x96, 0x16, 0xf2, 0x7c, 0x7f, 0x5d,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xa4, 0x4e, 0xee, 0xbf, 0x02, 0x00, 0x00,
}
