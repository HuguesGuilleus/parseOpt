# parseOpt

It's a module for parse and analyse environments variables and argument and give it' in a `Option` structure.

```go
type Option struct {
	Flag     map[string]bool
	Option   map[string][]string
}
```
