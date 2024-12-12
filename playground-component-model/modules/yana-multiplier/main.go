package main

import (
	yanamultiplier "yana-multiplier/yana/calculator/yana-multiplier"
)

func init() {
	yanamultiplier.Exports.Multiply = func(a uint32, b uint32) uint32 {
		return a * b
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
