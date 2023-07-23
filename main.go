package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	typeName    = flag.String("type", "", "name of the enum type. It will be alias to int")
	valuesList  = flag.String("values", "", "comma-separated list of values names")
	packageName = flag.String("package", "", "package name")
	output      = flag.String("output", "", "output filename (default enum_<type name>.go)")
	noPrefix    = flag.Bool("noprefix", false, "do not add type name as prefix to values")
)

const usagePrefix = "Usage of enum:\n" +
	"\tAdd following line to your golang code:\n" +
	"\t//enum -package=<name of the package> -type=<name of the enum type> -values=<coma separated list of the values>\n" +
	"\tThen run following command:\n" +
	"\tgo generate\n" +
	"\tGenerated code will be placed into the file enum_<type name>.go\n" +
	"\tFor more details: https://github.com/mpkondrashin/enum\n" +
	"Flags:\n"

func Usage() {
	fmt.Fprintf(os.Stderr, usagePrefix)
	flag.PrintDefaults()
}

func main() {
	//	log.SetFlags(0)
	//	log.SetPrefix("enum: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*packageName)+len(*typeName)+len(*valuesList) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	values := strings.Split(*valuesList, ",")
	fileName := fmt.Sprintf("enum_%s.go", strings.ToLower(*typeName))
	if *output != "" {
		fileName = *output
	}
	if err := Run(fileName, *packageName, *typeName, *noPrefix, values...); err != nil {
		fmt.Println(err)
	}
}
