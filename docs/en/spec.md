---
title: Specifications
---

To parse arguments, you must give some information. Every option or flag is type of `Spec` who was regrouped in a `SpecList` who need to parse. This module make a difference between:
- The *options*: contain a list of string.
- The *flags*: contain a booolean, true if presnet, else false.

```go
// Specification for an option or an flag
type Spec struct {
	// Name used in os.Args
	NameShort, NameLong string
	// Name used in os.Env or in file
	NameEnv string

	// Used in help, it describe this specification.
	Desc string
	// [Option only] It's the name of the option, used in help.
	OptionName string

	// True: -a|--aaa option -a|--aaa=option
	// False: -a\--aaaa=true|false
	NeedArg bool

	// [Flag only] Callback exectuted after the parsing.
	CBFlag func()
	// [Option only] Callback exectuted after the parsing.
	CBOption func([]string)
}

// A list of specification, it will be verified
// and used to parse Arg or/and Environment variable
type SpecList []*Spec
```


## Notes
In a specification item, if all Name are empty, the structure *correspond* to standard argument.

**Help specification**: if the specification list *don't contain* element with “h” for `NameShort` or “help” for `nameLong`, *so* it will add a help specification with a callback who list all flag and all option and exit the program.

The specification with `NameLong` *correspondant à* “`--`” will be a flag.

A specification with `NameShort` and `NameLong` empty *correspond aux* stadart arguments.

## Check
You can check in the test the syntax of your specification with the sub module `check`.

```go
package main

import (
	"github.com/HuguesGuilleus/parseOpt/check"
	"testing"
)

func TestSpecList(t *testing.T) {
	// list is your specification list
	check.Check(t, list)
}
```
