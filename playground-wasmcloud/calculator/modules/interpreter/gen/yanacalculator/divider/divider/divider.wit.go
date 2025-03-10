// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package divider represents the imported interface "yanacalculator:divider/divider".
package divider

// Divide represents the imported function "divide".
//
//	divide: func(a: u32, b: u32) -> u32
//
//go:nosplit
func Divide(a uint32, b uint32) (result uint32) {
	a0 := (uint32)(a)
	b0 := (uint32)(b)
	result0 := wasmimport_Divide((uint32)(a0), (uint32)(b0))
	result = (uint32)((uint32)(result0))
	return
}
