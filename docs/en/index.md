---
title: Index
---

[![GoDoc](https://godoc.org/github.com/HuguesGuilleus/parseOpt?status.svg)](https://godoc.org/github.com/HuguesGuilleus/parseOpt)

This module parse the environnement variables and arguments with a specification list, and return the result in a `Option` structure.


## Installation
```bash
go get github.com/HuguesGuilleus/parseOpt/
```


## Sommaire
{% include index_file.liquid %}


## Error log
The error display in `ErrLog` (type `*log.Logger`), who redirects to a intern `io.Writer` who set the line in red and write all in `os.Stderr`.
