// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file describes the abstract model of services (and their instances) as
// represented in Istio. This model is independent of the underlying platform
// (Kubernetes, Mesos, etc.). Platform specific adapters found populate the
// model object with various fields, from the metadata found in the platform.
// The platform independent proxy code uses the representation in the model to
// generate the configuration files for the Layer 7 proxy sidecar. The proxy
// code is specific to individual proxy implementations

package protocol

import "strings"

// Instance defines network protocols for ports
type Instance string

const (
	// GRPC declares that the port carries gRPC traffic.
	GRPC Instance = "GRPC"
	// GRPCWeb declares that the port carries gRPC traffic.
	GRPCWeb Instance = "GRPC-Web"
	// HTTP declares that the port carries HTTP/1.1 traffic.
	// Note that HTTP/1.0 or earlier may not be supported by the proxy.
	HTTP Instance = "HTTP"
	// HTTP2 declares that the port carries HTTP/2 traffic.
	HTTP2 Instance = "HTTP2"
	// HTTPS declares that the port carries HTTPS traffic.
	HTTPS Instance = "HTTPS"
	// TCP declares the the port uses TCP.
	// This is the default protocol for a service port.
	TCP Instance = "TCP"
	// TLS declares that the port carries TLS traffic.
	// TLS traffic is assumed to contain SNI as part of the handshake.
	TLS Instance = "TLS"
	// UDP declares that the port uses UDP.
	// Note that UDP protocol is not currently supported by the proxy.
	UDP Instance = "UDP"
	// Mongo declares that the port carries MongoDB traffic.
	Mongo Instance = "Mongo"
	// Redis declares that the port carries Redis traffic.
	Redis Instance = "Redis"
	// MySQL declares that the port carries MySQL traffic.
	MySQL Instance = "MySQL"
	// Unsupported - value to signify that the protocol is unsupported.
	Unsupported Instance = "UnsupportedProtocol"
)

// Parse from string ignoring case
func Parse(s string) Instance {
	switch strings.ToLower(s) {
	case "tcp":
		return TCP
	case "udp":
		return UDP
	case "grpc":
		return GRPC
	case "grpc-web":
		return GRPCWeb
	case "http":
		return HTTP
	case "http2":
		return HTTP2
	case "https":
		return HTTPS
	case "tls":
		return TLS
	case "mongo":
		return Mongo
	case "redis":
		return Redis
	case "mysql":
		return MySQL
	}

	return Unsupported
}

// IsHTTP2 is true for protocols that use HTTP/2 as transport protocol
func (i Instance) IsHTTP2() bool {
	switch i {
	case HTTP2, GRPC, GRPCWeb:
		return true
	default:
		return false
	}
}

// IsHTTP is true for protocols that use HTTP as transport protocol
func (i Instance) IsHTTP() bool {
	switch i {
	case HTTP, HTTP2, GRPC, GRPCWeb:
		return true
	default:
		return false
	}
}

// IsTCP is true for protocols that use TCP as transport protocol
func (i Instance) IsTCP() bool {
	switch i {
	case TCP, HTTPS, TLS, Mongo, Redis, MySQL:
		return true
	default:
		return false
	}
}

// IsTLS is true for protocols on top of TLS (e.g. HTTPS)
func (i Instance) IsTLS() bool {
	switch i {
	case HTTPS, TLS:
		return true
	default:
		return false
	}
}

// IsGRPC is true for GRCP protocols.
func (i Instance) IsGRPC() bool {
	switch i {
	case GRPC, GRPCWeb:
		return true
	default:
		return false
	}
}