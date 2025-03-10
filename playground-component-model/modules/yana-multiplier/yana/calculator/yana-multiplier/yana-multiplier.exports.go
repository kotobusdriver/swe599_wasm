// Code generated by wit-bindgen-go. DO NOT EDIT.

package yanamultiplier

// Exports represents the caller-defined exports from "yana:calculator/yana-multiplier".
var Exports struct {
	// Multiply represents the caller-defined, exported function "multiply".
	//
	// TinyGo requires this import to be able to use the WASI CLI
	// export multiplier;
	//
	//	multiply: func(a: u32, b: u32) -> u32
	Multiply func(a uint32, b uint32) (result uint32)
}
