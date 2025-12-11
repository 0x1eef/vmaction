package main

import (
	"fmt"
	"os"

	"github.com/hardenedbsd/hardenedbsd-vm/internal/apt"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/curl"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/vm"
	"github.com/hardenedbsd/hardenedbsd-vm/internal/xz"
)

func main() {
	var (
		ip      string
		archive string
		image   string
		err     error
	)
	if err := apt.Run(); err != nil {
		abort("error: %s\n", err)
	}
	if archive, err = curl.Run(); err != nil {
		abort("error: %s\n", err)
	}
	if image, err = xz.Run(archive); err != nil {
		abort("error: %s\n", err)
	}
	if ip, err = vm.Run(image); err != nil {
		abort("error: %s\n", err)
	}
	fmt.Println("VM IP Address:", ip)
}

func abort(s string, v ...any) {
	fmt.Fprintf(os.Stderr, s, v...)
	os.Exit(1)
}
