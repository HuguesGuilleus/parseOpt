package main

import (
	"github.com/HuguesGuilleus/parseOpt/check"
	"testing"
)

func TestSpecList(t *testing.T) {
	check.Check(t, &spec)
}
