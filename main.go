package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"os"
)

const (
	CliName = "rvmbuilder"
	Version = "0.1.1"
)

const UTF8BOM = "\xef\xbb\xbf"

//go:embed sample.json
var SampleJson string

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s outputs a macro to export Aveva Review files and the Navisworks file from the json format options.\n", CliName)
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [-o file] {-s | json_file}\n", CliName)
		fmt.Fprintf(os.Stderr, "Example: %s sample.json > export.mac\n", CliName)
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\n")
		fmt.Fprint(os.Stderr, "When json_file is -, read standard input instead of a file.\n")
	}
	vflg := flag.Bool("v", false, "Display version")
	sflg := flag.Bool("s", false, "Output a sample JSON")
	oflg := flag.String("o", "", "Output file")
	flag.Parse()
	if *vflg {
		fmt.Fprintf(os.Stderr, "%s version %s\n", CliName, Version)
		return
	}
	out := os.Stdout
	if *oflg != "" {
		var err error
		out, err = os.Create(*oflg)
		chkErr(err)
	}
	if *sflg {
		fmt.Fprintln(out, SampleJson)
		return
	}
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}
	if flag.NArg() >= 2 {
		chkErr(errors.New("too many arguments"))
	}
	var f *os.File
	if flag.Arg(0) == "-" {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(flag.Arg(0))
		chkErr(err)
		defer f.Close()
	}
	rb, err := LoadRvmBuilder(f)
	chkErr(err)
	s, err := MakePml(rb)
	chkErr(err)
	fmt.Fprint(out, UTF8BOM+s)
}

func chkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
