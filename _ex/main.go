package main

import (
	"crypto/rand"
	"fmt"
	"github.com/HuguesGuilleus/parseOpt"
)

// A list of spec used for parsing
var spec = parseOpt.SpecList{
	// A spec of an flag
	&parseOpt.Spec{
		NameShort: "n",
		NameLong:  "number",
		Desc:      "Display digit",
	},
	// A spec of an option
	&parseOpt.Spec{
		NameShort:  "l",
		NameLong:   "length",
		NeedArg:    true,
		OptionName: "length",
		Desc:       "The number of characters",
	},
}

func main() {
	// We parse the arguments
	opt := spec.ParseOs()

	// We get the number of character
	length := 10 // default value
	if opt.Option["length"] != nil {
		fmt.Sscanf(opt.Option["length"][0], "%d", &length)
	}

	// Generate random number
	nb := make([]byte, length)
	if _, err := rand.Read(nb); err != nil {
		fmt.Println(err)
		return
	}

	// Display the characters either letter (by default)
	// or digit (if the flag exist)
	modeNumber := opt.Flag["number"]
	for _, char := range nb {
		if modeNumber {
			fmt.Printf("%c", '0'+char%10)
		} else {
			fmt.Printf("%c", 'A'+char%26)
		}
	}
	fmt.Println()
}
