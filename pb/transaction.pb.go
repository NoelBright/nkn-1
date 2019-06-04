// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/transaction.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UnsignedTx struct {
	Payload    *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	Nonce      uint64   `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Fee        int64    `protobuf:"varint,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Attributes []byte   `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (m *UnsignedTx) Reset()      { *m = UnsignedTx{} }
func (*UnsignedTx) ProtoMessage() {}
func (*UnsignedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_4881c9c550dd9620, []int{0}
}
func (m *UnsignedTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnsignedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnsignedTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UnsignedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsignedTx.Merge(dst, src)
}
func (m *UnsignedTx) XXX_Size() int {
	return m.Size()
}
func (m *UnsignedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsignedTx.DiscardUnknown(m)
}

var xxx_messageInfo_UnsignedTx proto.InternalMessageInfo

func (m *UnsignedTx) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UnsignedTx) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *UnsignedTx) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *UnsignedTx) GetAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Transaction struct {
	UnsignedTx *UnsignedTx `protobuf:"bytes,1,opt,name=unsigned_tx,json=unsignedTx" json:"unsigned_tx,omitempty"`
	Programs   []*Program  `protobuf:"bytes,2,rep,name=programs" json:"programs,omitempty"`
}

func (m *Transaction) Reset()      { *m = Transaction{} }
func (*Transaction) ProtoMessage() {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_4881c9c550dd9620, []int{1}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetUnsignedTx() *UnsignedTx {
	if m != nil {
		return m.UnsignedTx
	}
	return nil
}

func (m *Transaction) GetPrograms() []*Program {
	if m != nil {
		return m.Programs
	}
	return nil
}

func init() {
	proto.RegisterType((*UnsignedTx)(nil), "pb.UnsignedTx")
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
}
func (this *UnsignedTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnsignedTx)
	if !ok {
		that2, ok := that.(UnsignedTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Payload.Equal(that1.Payload) {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	if this.Fee != that1.Fee {
		return false
	}
	if !bytes.Equal(this.Attributes, that1.Attributes) {
		return false
	}
	return true
}
func (this *Transaction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transaction)
	if !ok {
		that2, ok := that.(Transaction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UnsignedTx.Equal(that1.UnsignedTx) {
		return false
	}
	if len(this.Programs) != len(that1.Programs) {
		return false
	}
	for i := range this.Programs {
		if !this.Programs[i].Equal(that1.Programs[i]) {
			return false
		}
	}
	return true
}
func (this *UnsignedTx) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.UnsignedTx{")
	if this.Payload != nil {
		s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	}
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Fee: "+fmt.Sprintf("%#v", this.Fee)+",\n")
	s = append(s, "Attributes: "+fmt.Sprintf("%#v", this.Attributes)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Transaction) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.Transaction{")
	if this.UnsignedTx != nil {
		s = append(s, "UnsignedTx: "+fmt.Sprintf("%#v", this.UnsignedTx)+",\n")
	}
	if this.Programs != nil {
		s = append(s, "Programs: "+fmt.Sprintf("%#v", this.Programs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTransaction(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *UnsignedTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnsignedTx) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.Payload.Size()))
		n1, err := m.Payload.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Nonce != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.Nonce))
	}
	if m.Fee != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.Fee))
	}
	if len(m.Attributes) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Attributes)))
		i += copy(dAtA[i:], m.Attributes)
	}
	return i, nil
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UnsignedTx != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTransaction(dAtA, i, uint64(m.UnsignedTx.Size()))
		n2, err := m.UnsignedTx.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Programs) > 0 {
		for _, msg := range m.Programs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintTransaction(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedUnsignedTx(r randyTransaction, easy bool) *UnsignedTx {
	this := &UnsignedTx{}
	if r.Intn(10) != 0 {
		this.Payload = NewPopulatedPayload(r, easy)
	}
	this.Nonce = uint64(uint64(r.Uint32()))
	this.Fee = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Fee *= -1
	}
	v1 := r.Intn(100)
	this.Attributes = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Attributes[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedTransaction(r randyTransaction, easy bool) *Transaction {
	this := &Transaction{}
	if r.Intn(10) != 0 {
		this.UnsignedTx = NewPopulatedUnsignedTx(r, easy)
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Programs = make([]*Program, v2)
		for i := 0; i < v2; i++ {
			this.Programs[i] = NewPopulatedProgram(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTransaction interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTransaction(r randyTransaction) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTransaction(r randyTransaction) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneTransaction(r)
	}
	return string(tmps)
}
func randUnrecognizedTransaction(r randyTransaction, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTransaction(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTransaction(dAtA []byte, r randyTransaction, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTransaction(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateTransaction(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateTransaction(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTransaction(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTransaction(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTransaction(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTransaction(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *UnsignedTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTransaction(uint64(m.Nonce))
	}
	if m.Fee != 0 {
		n += 1 + sovTransaction(uint64(m.Fee))
	}
	l = len(m.Attributes)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	return n
}

func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnsignedTx != nil {
		l = m.UnsignedTx.Size()
		n += 1 + l + sovTransaction(uint64(l))
	}
	if len(m.Programs) > 0 {
		for _, e := range m.Programs {
			l = e.Size()
			n += 1 + l + sovTransaction(uint64(l))
		}
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *UnsignedTx) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnsignedTx{`,
		`Payload:` + strings.Replace(fmt.Sprintf("%v", this.Payload), "Payload", "Payload", 1) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Fee:` + fmt.Sprintf("%v", this.Fee) + `,`,
		`Attributes:` + fmt.Sprintf("%v", this.Attributes) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Transaction) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Transaction{`,
		`UnsignedTx:` + strings.Replace(fmt.Sprintf("%v", this.UnsignedTx), "UnsignedTx", "UnsignedTx", 1) + `,`,
		`Programs:` + strings.Replace(fmt.Sprintf("%v", this.Programs), "Program", "Program", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTransaction(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UnsignedTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: UnsignedTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnsignedTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &Payload{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes[:0], dAtA[iNdEx:postIndex]...)
			if m.Attributes == nil {
				m.Attributes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
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
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnsignedTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UnsignedTx == nil {
				m.UnsignedTx = &UnsignedTx{}
			}
			if err := m.UnsignedTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Programs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Programs = append(m.Programs, &Program{})
			if err := m.Programs[len(m.Programs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTransaction
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
				next, err := skipTransaction(dAtA[start:])
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
	ErrInvalidLengthTransaction = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/transaction.proto", fileDescriptor_transaction_4881c9c550dd9620) }

var fileDescriptor_transaction_4881c9c550dd9620 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4a, 0xc3, 0x50,
	0x14, 0xc6, 0x73, 0x9a, 0xfa, 0x87, 0x13, 0x91, 0x12, 0x3a, 0x84, 0x0e, 0x87, 0x50, 0x10, 0xb3,
	0x98, 0x80, 0xae, 0x4e, 0x3e, 0x81, 0x84, 0x3a, 0x4b, 0x6e, 0x9b, 0x5e, 0x03, 0xf6, 0xde, 0x4b,
	0x72, 0x03, 0x15, 0x17, 0x1f, 0xc1, 0xc7, 0xf0, 0x11, 0x7c, 0x04, 0xc7, 0x8e, 0x1d, 0xcd, 0xed,
	0xe2, 0xd8, 0xd1, 0x51, 0x9a, 0xc4, 0xa6, 0xdb, 0xf9, 0x7e, 0x1f, 0x07, 0x7e, 0xe7, 0xe0, 0x50,
	0xb1, 0x48, 0xe7, 0x89, 0x28, 0x92, 0xa9, 0xce, 0xa4, 0x08, 0x55, 0x2e, 0xb5, 0x74, 0x7b, 0x8a,
	0x8d, 0xae, 0x78, 0xa6, 0x9f, 0x4a, 0x16, 0x4e, 0xe5, 0x22, 0xe2, 0x92, 0xcb, 0xa8, 0xae, 0x58,
	0x39, 0xaf, 0x53, 0x1d, 0xea, 0xa9, 0x59, 0x19, 0x0d, 0x14, 0x8b, 0x54, 0xf2, 0xf2, 0x2c, 0x93,
	0xd9, 0x21, 0xc9, 0x25, 0xcf, 0x93, 0x45, 0x43, 0xc6, 0xaf, 0x88, 0x0f, 0xa2, 0xc8, 0xb8, 0x48,
	0x67, 0x93, 0xa5, 0x7b, 0x81, 0x27, 0xed, 0x82, 0x07, 0x3e, 0x04, 0xce, 0xb5, 0x13, 0x2a, 0x16,
	0xde, 0x37, 0x28, 0xfe, 0xef, 0xdc, 0x21, 0x1e, 0x09, 0x29, 0xa6, 0xa9, 0xd7, 0xf3, 0x21, 0xe8,
	0xc7, 0x4d, 0x70, 0x07, 0x68, 0xcf, 0xd3, 0xd4, 0xb3, 0x7d, 0x08, 0xec, 0x78, 0x37, 0xba, 0x84,
	0x98, 0x68, 0x9d, 0x67, 0xac, 0xd4, 0x69, 0xe1, 0xf5, 0x7d, 0x08, 0xce, 0xe2, 0x03, 0x32, 0xe6,
	0xe8, 0x4c, 0xba, 0x43, 0xdd, 0x08, 0x9d, 0xb2, 0x75, 0x79, 0xd4, 0xcb, 0xd6, 0xe0, 0x7c, 0x67,
	0xd0, 0x29, 0xc6, 0x58, 0x76, 0xba, 0x97, 0x78, 0xda, 0x5e, 0x53, 0x78, 0x3d, 0xdf, 0xde, 0xfb,
	0x36, 0x2c, 0xde, 0x97, 0x77, 0xb7, 0xab, 0x8a, 0xac, 0x75, 0x45, 0xd6, 0xb6, 0x22, 0xf8, 0xad,
	0x08, 0xde, 0x0c, 0xc1, 0x87, 0x21, 0xf8, 0x34, 0x04, 0x5f, 0x86, 0x60, 0x65, 0x08, 0xbe, 0x0d,
	0xc1, 0x8f, 0x21, 0x6b, 0x6b, 0x08, 0xde, 0x37, 0x64, 0xad, 0x36, 0x64, 0xad, 0x37, 0x64, 0xb1,
	0xe3, 0xfa, 0x55, 0x37, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0x5c, 0xc6, 0xcd, 0x99, 0x01,
	0x00, 0x00,
}