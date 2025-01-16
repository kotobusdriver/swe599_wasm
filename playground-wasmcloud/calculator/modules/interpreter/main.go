//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world interpreter-component --out gen ./wit
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"go.wasmcloud.dev/component/net/wasihttp"
    "go.wasmcloud.dev/component/log/wasilog"
	"github.com/yana/calculator/interpreter-component/gen/yanacalculator/adder/adder"
	"github.com/yana/calculator/interpreter-component/gen/yanacalculator/subtractor/subtractor"
	"github.com/yana/calculator/interpreter-component/gen/yanacalculator/multiplier/multiplier"
	"github.com/yana/calculator/interpreter-component/gen/yanacalculator/divider/divider"
)

func init() {
	// Register the handleRequest function as the handler for all incoming requests.
	wasihttp.HandleFunc(handleRequest)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    logger := wasilog.ContextLogger("handle")
    op := r.FormValue("op")
    aa := r.FormValue("a")
    bb := r.FormValue("b")
    a64, _ := strconv.ParseUint(aa, 10, 32)
    b64, _ := strconv.ParseUint(bb, 10, 32)
    a := uint32(a64)
    b := uint32(b64)
    logger.Info("request diagnostics", "op", op, "a", a, "b", b)
    result := uint32(0)
    switch op {
        case "add":
            result = adder.Add(a, b)
    	case "sub":
    		result = subtractor.Subtract(a, b)
    	case "mul":
    		result = multiplier.Multiply(a, b)
    	case "div":
    		result = divider.Divide(a, b)
    	default:
    		logger.Info("Unknown operation is provided!")
    		fmt.Fprintf(w, "Unknown operation is provided!\n")
    		return
    }
	fmt.Fprintf(w, "Result = %d\n", result)
}

// Since we don't run this program like a CLI, the `main` function is empty. Instead,
// we call the `handleRequest` function when an HTTP request is received.
func main() {}
