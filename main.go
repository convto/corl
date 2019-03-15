package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

var (
	help   bool
	encode bool
	decode bool
)

func init() {
	flag.BoolVar(&help, "h", false, "Print help.")
	flag.BoolVar(&encode, "e", false, "Encode URL. eg, corl -e `$URL`")
	flag.BoolVar(&decode, "d", false, "Decode URL. eg, corl -d `$URL`")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if (encode || decode) && len(flag.Args()) == 0 {
		fmt.Println("URL parameter is required when using -e or -d")
		os.Exit(1)
	}
	if encode && decode {
		fmt.Println("Only one of encode and decode can be used")
		os.Exit(1)
	}
	rawURL := flag.Arg(0)
	if encode {
		encURL := url.PathEscape(rawURL)
		fmt.Println(encURL)
		os.Exit(0)
	}
	if decode {
		decURL, err := url.PathUnescape(rawURL)
		if err != nil {
			fmt.Printf("URL decode err.\n%v\n", err)
			os.Exit(1)
		}
		fmt.Println(decURL)
		os.Exit(0)
	}
	fmt.Printf("This is cmd for url convert. It supports decode and encode.\nIf you want to know how to use, please type\n\n  corl -h\n\n")
}
