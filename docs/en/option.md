---
title: The `Option` structure
---

```go
type Option struct {
	Flag map[string]bool
	Option map[string][]string
}
```

The `Option` structure save all parsed informations with two fiel:
- `Flag`: all flag meeting, `true`

- `Flag`: contient chaques drapeaux rencontrés, `true` si le drapeau est réglé vrai (par défaut), `false` si le drapeau est réglé à faux ou n'est pas présent.
- `Option`: Contient un tableau de chaînes de caractères de chaque options rencontrés.

## Key
The key to get a element in `Option.Flag` or `Option.Option` is the fisrt non empty string:
1. NameLong
2. NameShort
3. NameEnv
4. A empty string

## Notes

The `--` flag are saved as a flag with the `--` key.

The standard arguments (no *rattaché* to option) are saved in option and the key in emty string.


## Method
