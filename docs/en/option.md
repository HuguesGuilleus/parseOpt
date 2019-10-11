---
title: The `Option` structure
---

```go
// Option save all parsed informations with two fiel:
type Option struct {
	// Flag save all matched flags.
	// true if the flmag is set to true (by default);
	// false if the flag is set to false or do not exist.
	Flag map[string]bool
	// Option contain a slice of string of every matched option.
	Option map[string][]string
}
```

## Key
The key to get a element in `Option.Flag` or `Option.Option` is the first non empty string:
1. NameLong
2. NameShort
3. NameEnv
4. A empty string

**Note:** The `--` flag are saved as a flag with the `--` key.

**Note:** The standard arguments (no *rattaché* to option) are saved in option and the key in emty string.
