// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy/v1alpha1/config/egress_rule.proto

package istio_proxy_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Egress rules describe the properties of a service outside Istio. When transparent proxying
// is used, egress rules signify a white listed set of domains that microserves in the mesh
// are allowed to access. A subset of routing rules and all destination policies can be applied
// on the service targeted by an egress rule. The destination of an egress rule is allowed to
// contain wildcards (e.g., *.foo.com). Currently, only HTTP-based services can be expressed
// through the egress rule. If TLS origination from the sidecar is desired, the protocol
// associated with the service port must be marked as HTTPS, and the service is expected to
// be accessed over HTTP (e.g., http://gmail.com:443). The sidecar will automatically upgrade
// the connection to TLS when initiating a connection with the external service.
//
// For example, the following egress rule describes the set of services hosted under the *.foo.com domain
//
//     kind: EgressRule
//     metadata:
//       name: foo-egress-rule
//     spec:
//       destination:
//         service: *.foo.com
//       ports:
//         - port: 80
//           protocol: http
//         - port: 443
//           protocol: https
//
type EgressRule struct {
	// REQUIRED: Hostname or a wildcard domain name associated with the external service.
	// ONLY the "service" field of destination will be taken into consideration. Name,
	// namespace, domain and labels are ignored. Routing rules and destination policies that
	// refer to these external services must have identical specification for the destination
	// as the corresponding egress rule. Wildcard domain specifications must conform to format
	// allowed by Envoy's Virtual Host specification, such as “*.foo.com” or “*-bar.foo.com”.
	// The character '*' in a domain specification indicates a non-empty string. Hence, a wildcard
	// domain of form “*-bar.foo.com” will match “baz-bar.foo.com” but not “-bar.foo.com”.
	Destination *IstioService `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
	// REQUIRED: list of ports on which the external service is available.
	Ports []*EgressRule_Port `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty"`
	// Forward all the external traffic through a dedicated egress proxy. It is used in some scenarios
	// where there is a requirement that all the external traffic goes through special dedicated nodes/pods.
	// These dedicated egress nodes could then be more closely monitored for security vulnerabilities.
	//
	// The default is false, i.e. the sidecar forwards external traffic directly to the external service.
	UseEgressProxy bool `protobuf:"varint,3,opt,name=use_egress_proxy,json=useEgressProxy" json:"use_egress_proxy,omitempty"`
}

func (m *EgressRule) Reset()                    { *m = EgressRule{} }
func (m *EgressRule) String() string            { return proto.CompactTextString(m) }
func (*EgressRule) ProtoMessage()               {}
func (*EgressRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EgressRule) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *EgressRule) GetPorts() []*EgressRule_Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *EgressRule) GetUseEgressProxy() bool {
	if m != nil {
		return m.UseEgressProxy
	}
	return false
}

// Port describes the properties of a specific TCP port of an external service.
type EgressRule_Port struct {
	// A valid non-negative integer port number.
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	// The protocol to communicate with the external services.
	// MUST BE one of HTTP|HTTPS|GRPC|HTTP2.
	Protocol string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *EgressRule_Port) Reset()                    { *m = EgressRule_Port{} }
func (m *EgressRule_Port) String() string            { return proto.CompactTextString(m) }
func (*EgressRule_Port) ProtoMessage()               {}
func (*EgressRule_Port) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *EgressRule_Port) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EgressRule_Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*EgressRule)(nil), "istio.proxy.v1.config.EgressRule")
	proto.RegisterType((*EgressRule_Port)(nil), "istio.proxy.v1.config.EgressRule.Port")
}

func init() { proto.RegisterFile("proxy/v1alpha1/config/egress_rule.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0x07, 0x70, 0xb2, 0x1f, 0xb2, 0x4e, 0x41, 0x24, 0x20, 0x94, 0x9e, 0x8a, 0xc2, 0x9a, 0x53,
	0x4a, 0x57, 0xf0, 0xe4, 0x75, 0x0f, 0xde, 0x96, 0xf8, 0x00, 0x4b, 0xad, 0xe3, 0x1a, 0x28, 0x9d,
	0x32, 0x49, 0x8a, 0x3e, 0xb7, 0x2f, 0x20, 0x49, 0x45, 0xf7, 0xb0, 0x7b, 0x9b, 0x0c, 0xf3, 0x9b,
	0xc9, 0x1f, 0xee, 0x07, 0xa6, 0xcf, 0xaf, 0x6a, 0xac, 0x9b, 0x6e, 0xf8, 0x68, 0xea, 0xaa, 0xa5,
	0xfe, 0xdd, 0x1e, 0x2a, 0x3c, 0x30, 0x3a, 0xb7, 0xe7, 0xd0, 0xa1, 0x1e, 0x98, 0x3c, 0xc9, 0x1b,
	0xeb, 0xbc, 0x25, 0x9d, 0xc6, 0xf5, 0x58, 0xeb, 0x69, 0xb0, 0x58, 0x9f, 0xf6, 0x4c, 0xc1, 0xe3,
	0x11, 0xbf, 0xfd, 0x16, 0x00, 0xdb, 0xb4, 0xd4, 0x84, 0x0e, 0xe5, 0x16, 0xb2, 0x37, 0x74, 0xde,
	0xf6, 0x8d, 0xb7, 0xd4, 0xe7, 0xa2, 0x14, 0x2a, 0xdb, 0xdc, 0xe9, 0x93, 0x37, 0xf4, 0x73, 0xec,
	0xbe, 0x20, 0x8f, 0xb6, 0x45, 0x73, 0xec, 0xe4, 0x13, 0x2c, 0x07, 0x62, 0xef, 0xf2, 0x59, 0x39,
	0x57, 0xd9, 0x66, 0x7d, 0x66, 0xc1, 0xff, 0x61, 0xbd, 0x23, 0xf6, 0x66, 0x42, 0x52, 0xc1, 0x75,
	0x70, 0xb8, 0xff, 0xcd, 0x9a, 0x50, 0x3e, 0x2f, 0x85, 0x5a, 0x99, 0xab, 0xe0, 0x70, 0x42, 0xbb,
	0xd8, 0x2d, 0x1e, 0x61, 0x11, 0xa1, 0x94, 0xb0, 0x88, 0x34, 0xfd, 0x77, 0x69, 0x52, 0x2d, 0x0b,
	0x58, 0xa5, 0x88, 0x2d, 0x75, 0xf9, 0xac, 0x14, 0xea, 0xd2, 0xfc, 0xbd, 0x5f, 0x2f, 0x52, 0xf5,
	0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x94, 0x72, 0x6c, 0x8f, 0x66, 0x01, 0x00, 0x00,
}