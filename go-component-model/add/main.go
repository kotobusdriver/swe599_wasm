package main

import (
	"example.com/internal/docs/adder/adder"
)

func init() {
	adder.Exports.Add = func(x int32, y int32) int32 {
		return x + y
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
