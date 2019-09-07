// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package check

import (
	"github.com/HuguesGuilleus/parseOpt"
	"testing"
)

func TestCheck(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		test := &testing.T{}
		Check(test, &parseOpt.SpecList{
			&parseOpt.Spec{
				NameShort: "a",
				Desc:      "Description ...",
			},
		})
		if test.Failed() == true {
			t.Fail()
		}
	})
	t.Run("Fail", func(t *testing.T) {
		test := &testing.T{}
		Check(test, &parseOpt.SpecList{
			&parseOpt.Spec{
				NameShort: "a",
			},
		})
		if test.Failed() == false {
			t.Fail()
		}
	})
}
