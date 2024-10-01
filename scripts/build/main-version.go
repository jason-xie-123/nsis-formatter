package main

import (
	"fmt"
	packageVersion "nsis-formatter/version"
)

func main() {
	fmt.Printf("%v\n", packageVersion.Version)
}
