//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world multiplier-component --out gen ./wit
package main

import (
	"github.com/yana/calculator/multiplier-component/gen/yanacalculator/multiplier/multiplier"
)

func init() {
	multiplier.Exports.Multiply = func(a uint32, b uint32) uint32 {
		return a * b
	}
}


// Since we don't run this program like a CLI, the `main` function is empty. Instead,
// we call the `handleRequest` function when an HTTP request is received.
func main() {}
