---
title: Example
---

This programme displays random letters or digits. This exemple is in `_ex/` directory.

```go
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
```


Dans le main:
```go
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
```

## Flag
To display the letters:
```bash
./_ex
./_ex -n=false
./_ex --number=false
```

To display the digits:
```bash
./_ex -n
./_ex -n=true
./_ex --number=true
```

The value of the booleans are: `0`, `1`, `true`, `false`, `True`, `False`, `TRUE` and `FALSE`; an other value will generate a warning and will be set to true.

Les valeurs possibles des booléens sont: `0`, `1`, `true`, `false`, `True`, `False`, `TRUE`, `FALSE`. Une autre valeur donnera un avertissement et sera compté comme vrai.

## Option
To set the length:
```bash
./_ex -l=20
./_ex -l 20
./_ex --length=20
./_ex --length 20
```

## Flag and option
On peut combiner les drapeaux avec les options. Les syntaxes suivantes sont équivalentes:
We can associate the flag and the option, the foolowed syntax are equal:
```bash
./_ex -nl=20
./_ex -nl 20
./_ex -n -l 20
```

## Help
To display the options and flags:
```bash
./_ex -h
./_ex --help
```

## Check
The specification list can (must) be checked (if there are a description, ...) see the `_ex/main_test.go` file.
