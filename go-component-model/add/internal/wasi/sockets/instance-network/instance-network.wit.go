// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package instancenetwork represents the imported interface "wasi:sockets/instance-network@0.2.0".
package instancenetwork

import (
	"example.com/internal/wasi/sockets/network"
	"go.bytecodealliance.org/cm"
)

// Network represents the imported type alias "wasi:sockets/instance-network@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// InstanceNetwork represents the imported function "instance-network".
//
//	instance-network: func() -> network
//
//go:nosplit
func InstanceNetwork() (result Network) {
	result0 := wasmimport_InstanceNetwork()
	result = cm.Reinterpret[Network]((uint32)(result0))
	return
}
