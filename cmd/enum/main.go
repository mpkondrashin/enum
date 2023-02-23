package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mpkondrashin/enum"
)

var (
	typeName    = flag.String("type", "", "output file name; default srcdir/<type>_string.go; must be set")
	valuesList  = flag.String("values", "", "comma-separated list of values names; must be set")
	packageName = flag.String("package", "", "package name")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of enum:\n")
	fmt.Fprintf(os.Stderr, "\tenum [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tenum [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttps://;ekfnq;efjnq;efpkg.go.dev/golang.org/x/tools/cmd/stringer\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("enum: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*packageName) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	if len(*valuesList) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	values := strings.Split(*valuesList, ",")

	err := enum.Run(*packageName, *typeName, values...)
	if err != nil {
		panic(err)
	}
}
