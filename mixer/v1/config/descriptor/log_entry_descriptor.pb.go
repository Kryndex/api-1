// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/config/descriptor/log_entry_descriptor.proto

/*
	Package istio_mixer_v1_config_descriptor is a generated protocol buffer package.

	It is generated from these files:
		mixer/v1/config/descriptor/log_entry_descriptor.proto
		mixer/v1/config/descriptor/metric_descriptor.proto
		mixer/v1/config/descriptor/monitored_resource_descriptor.proto
		mixer/v1/config/descriptor/principal_descriptor.proto
		mixer/v1/config/descriptor/quota_descriptor.proto
		mixer/v1/config/descriptor/value_type.proto

	It has these top-level messages:
		LogEntryDescriptor
		MetricDescriptor
		MonitoredResourceDescriptor
		PrincipalDescriptor
		QuotaDescriptor
*/
package istio_mixer_v1_config_descriptor

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"
import sortkeys "github.com/gogo/protobuf/sortkeys"

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

// PayloadFormat details the currently supported logging payload formats.
// TEXT is the default payload format.
type LogEntryDescriptor_PayloadFormat int32

const (
	// Invalid, default value.
	PAYLOAD_FORMAT_UNSPECIFIED LogEntryDescriptor_PayloadFormat = 0
	// Indicates a payload format of raw text.
	TEXT LogEntryDescriptor_PayloadFormat = 1
	// Indicates that the payload is a serialized JSON object.
	JSON LogEntryDescriptor_PayloadFormat = 2
)

var LogEntryDescriptor_PayloadFormat_name = map[int32]string{
	0: "PAYLOAD_FORMAT_UNSPECIFIED",
	1: "TEXT",
	2: "JSON",
}
var LogEntryDescriptor_PayloadFormat_value = map[string]int32{
	"PAYLOAD_FORMAT_UNSPECIFIED": 0,
	"TEXT": 1,
	"JSON": 2,
}

func (LogEntryDescriptor_PayloadFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorLogEntryDescriptor, []int{0, 0}
}

// Defines the format of a single log entry.
type LogEntryDescriptor struct {
	// Required. The name of this descriptor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. A concise name for the log entry type, which can be displayed in
	// user interfaces. Use sentence case without an ending period, for example
	// "access log".
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. A description of the log entry type, which can be used in documentation.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Required. Format of the value of the payload attribute.
	PayloadFormat LogEntryDescriptor_PayloadFormat `protobuf:"varint,4,opt,name=payload_format,json=payloadFormat,proto3,enum=istio.mixer.v1.config.descriptor.LogEntryDescriptor_PayloadFormat" json:"payload_format,omitempty"`
	// Required. The template that will be populated with labels at runtime to
	// generate a log message; the labels describe the parameters for this template.
	//
	// The template strings must conform to go's text/template syntax.
	// TODO: unify templates to use our expression language when its finalized
	LogTemplate string `protobuf:"bytes,5,opt,name=log_template,json=logTemplate,proto3" json:"log_template,omitempty"`
	// Labels describe the parameters of this log's template string. The log
	// definition allows the user to map attribute expressions to actual values
	// for these labels at run time; the result of the evaluation must be of the
	// type described by the kind for each label.
	Labels map[string]ValueType `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=istio.mixer.v1.config.descriptor.ValueType"`
}

func (m *LogEntryDescriptor) Reset()      { *m = LogEntryDescriptor{} }
func (*LogEntryDescriptor) ProtoMessage() {}
func (*LogEntryDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptorLogEntryDescriptor, []int{0}
}

func (m *LogEntryDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogEntryDescriptor) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *LogEntryDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LogEntryDescriptor) GetPayloadFormat() LogEntryDescriptor_PayloadFormat {
	if m != nil {
		return m.PayloadFormat
	}
	return PAYLOAD_FORMAT_UNSPECIFIED
}

func (m *LogEntryDescriptor) GetLogTemplate() string {
	if m != nil {
		return m.LogTemplate
	}
	return ""
}

func (m *LogEntryDescriptor) GetLabels() map[string]ValueType {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*LogEntryDescriptor)(nil), "istio.mixer.v1.config.descriptor.LogEntryDescriptor")
	proto.RegisterEnum("istio.mixer.v1.config.descriptor.LogEntryDescriptor_PayloadFormat", LogEntryDescriptor_PayloadFormat_name, LogEntryDescriptor_PayloadFormat_value)
}
func (x LogEntryDescriptor_PayloadFormat) String() string {
	s, ok := LogEntryDescriptor_PayloadFormat_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *LogEntryDescriptor) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LogEntryDescriptor)
	if !ok {
		that2, ok := that.(LogEntryDescriptor)
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
	if this.Name != that1.Name {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.PayloadFormat != that1.PayloadFormat {
		return false
	}
	if this.LogTemplate != that1.LogTemplate {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	return true
}
func (this *LogEntryDescriptor) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&istio_mixer_v1_config_descriptor.LogEntryDescriptor{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "DisplayName: "+fmt.Sprintf("%#v", this.DisplayName)+",\n")
	s = append(s, "Description: "+fmt.Sprintf("%#v", this.Description)+",\n")
	s = append(s, "PayloadFormat: "+fmt.Sprintf("%#v", this.PayloadFormat)+",\n")
	s = append(s, "LogTemplate: "+fmt.Sprintf("%#v", this.LogTemplate)+",\n")
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]ValueType{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%#v: %#v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	if this.Labels != nil {
		s = append(s, "Labels: "+mapStringForLabels+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLogEntryDescriptor(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *LogEntryDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogEntryDescriptor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.DisplayName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(len(m.DisplayName)))
		i += copy(dAtA[i:], m.DisplayName)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if m.PayloadFormat != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(m.PayloadFormat))
	}
	if len(m.LogTemplate) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(len(m.LogTemplate)))
		i += copy(dAtA[i:], m.LogTemplate)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x32
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovLogEntryDescriptor(uint64(len(k))) + 1 + sovLogEntryDescriptor(uint64(v))
			i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintLogEntryDescriptor(dAtA, i, uint64(v))
		}
	}
	return i, nil
}

func encodeVarintLogEntryDescriptor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LogEntryDescriptor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovLogEntryDescriptor(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovLogEntryDescriptor(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovLogEntryDescriptor(uint64(l))
	}
	if m.PayloadFormat != 0 {
		n += 1 + sovLogEntryDescriptor(uint64(m.PayloadFormat))
	}
	l = len(m.LogTemplate)
	if l > 0 {
		n += 1 + l + sovLogEntryDescriptor(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLogEntryDescriptor(uint64(len(k))) + 1 + sovLogEntryDescriptor(uint64(v))
			n += mapEntrySize + 1 + sovLogEntryDescriptor(uint64(mapEntrySize))
		}
	}
	return n
}

func sovLogEntryDescriptor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLogEntryDescriptor(x uint64) (n int) {
	return sovLogEntryDescriptor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LogEntryDescriptor) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]ValueType{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&LogEntryDescriptor{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", this.DisplayName) + `,`,
		`Description:` + fmt.Sprintf("%v", this.Description) + `,`,
		`PayloadFormat:` + fmt.Sprintf("%v", this.PayloadFormat) + `,`,
		`LogTemplate:` + fmt.Sprintf("%v", this.LogTemplate) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLogEntryDescriptor(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LogEntryDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogEntryDescriptor
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
			return fmt.Errorf("proto: LogEntryDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogEntryDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryDescriptor
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
				return ErrInvalidLengthLogEntryDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryDescriptor
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
				return ErrInvalidLengthLogEntryDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryDescriptor
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
				return ErrInvalidLengthLogEntryDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadFormat", wireType)
			}
			m.PayloadFormat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PayloadFormat |= (LogEntryDescriptor_PayloadFormat(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogTemplate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryDescriptor
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
				return ErrInvalidLengthLogEntryDescriptor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogTemplate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryDescriptor
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
				return ErrInvalidLengthLogEntryDescriptor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]ValueType)
			}
			var mapkey string
			var mapvalue ValueType
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLogEntryDescriptor
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLogEntryDescriptor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthLogEntryDescriptor
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLogEntryDescriptor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (ValueType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLogEntryDescriptor(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthLogEntryDescriptor
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogEntryDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogEntryDescriptor
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
func skipLogEntryDescriptor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogEntryDescriptor
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
					return 0, ErrIntOverflowLogEntryDescriptor
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
					return 0, ErrIntOverflowLogEntryDescriptor
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
				return 0, ErrInvalidLengthLogEntryDescriptor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLogEntryDescriptor
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
				next, err := skipLogEntryDescriptor(dAtA[start:])
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
	ErrInvalidLengthLogEntryDescriptor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogEntryDescriptor   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/v1/config/descriptor/log_entry_descriptor.proto", fileDescriptorLogEntryDescriptor)
}

var fileDescriptorLogEntryDescriptor = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x8d, 0xdb, 0xae, 0x02, 0x97, 0x55, 0x95, 0x4f, 0x51, 0x0f, 0x56, 0xd8, 0xa9, 0xd2, 0x90,
	0xa3, 0x15, 0x21, 0x21, 0x4e, 0x64, 0x6b, 0x2a, 0x0d, 0x95, 0xb6, 0xca, 0x02, 0x1a, 0xa7, 0xc8,
	0x5b, 0xdd, 0xc8, 0xc2, 0x89, 0xad, 0xc4, 0x54, 0xe4, 0x86, 0xf8, 0x05, 0xfc, 0x0c, 0x7e, 0x0a,
	0xc7, 0x1d, 0x39, 0xd2, 0x70, 0xe1, 0xb8, 0x9f, 0x80, 0xe2, 0x04, 0x3a, 0x84, 0x44, 0x25, 0x6e,
	0x9f, 0xdf, 0xf7, 0xbe, 0xf7, 0x3d, 0x3f, 0x1b, 0x3e, 0x49, 0xf8, 0x7b, 0x96, 0xb9, 0x9b, 0x13,
	0xf7, 0x5a, 0xa6, 0x6b, 0x1e, 0xbb, 0x2b, 0x96, 0x5f, 0x67, 0x5c, 0x69, 0x99, 0xb9, 0x42, 0xc6,
	0x11, 0x4b, 0x75, 0x56, 0x44, 0x3b, 0x90, 0xa8, 0x4c, 0x6a, 0x89, 0x1c, 0x9e, 0x6b, 0x2e, 0x89,
	0x19, 0x26, 0x9b, 0x13, 0x52, 0x0f, 0x93, 0x1d, 0x6f, 0x78, 0xfc, 0x0f, 0xe1, 0x0d, 0x15, 0xef,
	0x58, 0xa4, 0x0b, 0xc5, 0x6a, 0xb9, 0xa3, 0x8f, 0x1d, 0x88, 0x66, 0x32, 0xf6, 0xab, 0x65, 0x93,
	0xdf, 0x3c, 0x84, 0x60, 0x27, 0xa5, 0x09, 0xb3, 0x81, 0x03, 0x46, 0xf7, 0x03, 0x53, 0xa3, 0x87,
	0xf0, 0xc1, 0x8a, 0xe7, 0x4a, 0xd0, 0x22, 0x32, 0xbd, 0x96, 0xe9, 0xf5, 0x1a, 0x6c, 0x5e, 0x51,
	0x1c, 0xd8, 0xfb, 0xb5, 0x8c, 0xcb, 0xd4, 0x6e, 0x37, 0x8c, 0x1d, 0x84, 0x38, 0xec, 0x2b, 0x5a,
	0x08, 0x49, 0x57, 0xd1, 0x5a, 0x66, 0x09, 0xd5, 0x76, 0xc7, 0x01, 0xa3, 0xfe, 0xf8, 0x94, 0xec,
	0xbb, 0x17, 0xf9, 0xdb, 0x26, 0x59, 0xd6, 0x52, 0x53, 0xa3, 0x14, 0x1c, 0xaa, 0xbb, 0xc7, 0xca,
	0x6f, 0x95, 0xa3, 0x66, 0x89, 0x12, 0x54, 0x33, 0xfb, 0xa0, 0x76, 0x23, 0x64, 0x1c, 0x36, 0x10,
	0xba, 0x84, 0x5d, 0x41, 0xaf, 0x98, 0xc8, 0xed, 0xae, 0xd3, 0x1e, 0xf5, 0xc6, 0xcf, 0xff, 0xcb,
	0xc5, 0xcc, 0x48, 0x18, 0x34, 0x68, 0xf4, 0x86, 0x6b, 0xd8, 0xbb, 0x03, 0xa3, 0x01, 0x6c, 0xbf,
	0x65, 0x45, 0x13, 0x67, 0x55, 0x22, 0x0f, 0x1e, 0x98, 0xc7, 0x30, 0x31, 0xf6, 0xc7, 0xc7, 0xfb,
	0x37, 0xbf, 0xae, 0xe8, 0x61, 0xa1, 0x58, 0x50, 0x4f, 0x3e, 0x6b, 0x3d, 0x05, 0x47, 0x67, 0xf0,
	0xf0, 0x8f, 0x10, 0x10, 0x86, 0xc3, 0xa5, 0xf7, 0x66, 0xb6, 0xf0, 0x26, 0xd1, 0x74, 0x11, 0xbc,
	0xf4, 0xc2, 0xe8, 0xd5, 0xfc, 0x62, 0xe9, 0x9f, 0x9d, 0x4f, 0xcf, 0xfd, 0xc9, 0xc0, 0x42, 0xf7,
	0x60, 0x27, 0xf4, 0x2f, 0xc3, 0x01, 0xa8, 0xaa, 0x17, 0x17, 0x8b, 0xf9, 0xa0, 0x75, 0xfa, 0xe8,
	0x66, 0x8b, 0xad, 0xaf, 0x5b, 0x6c, 0xdd, 0x6e, 0x31, 0xf8, 0x50, 0x62, 0xf0, 0xb9, 0xc4, 0xe0,
	0x4b, 0x89, 0xc1, 0x4d, 0x89, 0xc1, 0xb7, 0x12, 0x83, 0x1f, 0x25, 0xb6, 0x6e, 0x4b, 0x0c, 0x3e,
	0x7d, 0xc7, 0xd6, 0x55, 0xd7, 0xfc, 0x9c, 0xc7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x0f,
	0x6e, 0x1e, 0xc1, 0x02, 0x00, 0x00,
}
