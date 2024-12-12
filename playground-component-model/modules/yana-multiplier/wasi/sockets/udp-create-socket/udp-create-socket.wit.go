// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package udpcreatesocket represents the imported interface "wasi:sockets/udp-create-socket@0.2.0".
package udpcreatesocket

import (
	"go.bytecodealliance.org/cm"
	"yana-multiplier/wasi/sockets/network"
	"yana-multiplier/wasi/sockets/udp"
)

// Network represents the imported type alias "wasi:sockets/udp-create-socket@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// ErrorCode represents the type alias "wasi:sockets/udp-create-socket@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPAddressFamily represents the type alias "wasi:sockets/udp-create-socket@0.2.0#ip-address-family".
//
// See [network.IPAddressFamily] for more information.
type IPAddressFamily = network.IPAddressFamily

// UDPSocket represents the imported type alias "wasi:sockets/udp-create-socket@0.2.0#udp-socket".
//
// See [udp.UDPSocket] for more information.
type UDPSocket = udp.UDPSocket

// CreateUDPSocket represents the imported function "create-udp-socket".
//
//	create-udp-socket: func(address-family: ip-address-family) -> result<udp-socket,
//	error-code>
//
//go:nosplit
func CreateUDPSocket(addressFamily IPAddressFamily) (result cm.Result[UDPSocket, UDPSocket, ErrorCode]) {
	addressFamily0 := (uint32)(addressFamily)
	wasmimport_CreateUDPSocket((uint32)(addressFamily0), &result)
	return
}
