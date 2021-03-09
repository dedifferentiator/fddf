package main

import (
	"fmt"
	"os"
	"runtime"

	fddf "github.com/dedifferentiator/fddf/internal"
)

func init() {
	if runtime.GOOS != "linux" {
		fmt.Fprint(os.Stderr, "Linux is the only supported os\n")
		os.Exit(1)
	}
}

func main() {
	pidd, err := fddf.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		if _, ok := err.(fddf.EArgNExist); ok {
			fddf.Usage()
		}
		os.Exit(1)
	}

	fddf.RunUI(pidd)
}
