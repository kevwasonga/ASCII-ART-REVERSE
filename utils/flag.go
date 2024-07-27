package utils

import "flag"

var ReversePtr = flag.String("reverse", "", "reverse the ascii art")

func ParseFlag() {
	flag.Parse()
}
