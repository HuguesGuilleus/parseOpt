// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"bytes"
	"log"
	"testing"
)

func TestCerrorBool(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		e := Cerror{
			Reason: 0,
		}
		if e.Bool() {
			t.Error("Cerror with Reason==0, err.Bool() must return false")
		}
	})
	t.Run("error", func(t *testing.T) {
		e := Cerror{
			Reason: 1,
		}
		if e.Bool() != true {
			t.Error("Cerror with Reason!=0, err.Bool() must return true")
		}
	})
}

func TestCerrorPrint(t *testing.T) {
	oldLog := ErrLog
	defer func() { ErrLog = oldLog }()
	buff := &bytes.Buffer{}
	ErrLog = log.New(buff, "", 0)
	t.Run("Simple", func(t *testing.T) {
		err := Cerror{
			Reason: 1,
		}
		exp := err.Error() + "\n"
		err.print()
		if rec := buff.String(); rec != exp {
			t.Error("ErrLog:")
			t.Log("Expected:", exp)
			t.Log("Received:", rec)
		}
	})
	t.Run("No error", func(t *testing.T) {
		buff.Reset()
		err := Cerror{}
		err.print()
		if buff.Len() != 0 {
			t.Error("Unexpected ErrLog:", buff)
		}
	})
}
