package main

import (
	"fmt"
	"os"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/virt"
)

func main() {
	image, err := curl.Run()
	if err != nil {
		abort("error: %s\n", err)
	}
	if err := virt.Run(image); err != nil {
		abort("error: %s\n", err)
	}
}

func abort(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s, v...)
	os.Exit(1)
}
