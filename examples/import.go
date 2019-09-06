package main

import (
	"fmt"

	"github.com/pkg/errors" // <-- this is external package
)

func main() {
	err := errors.WithStack(fmt.Errorf("testfail"))
	fmt.Printf("%+v", err)
}
