// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package check

import (
	"github.com/HuguesGuilleus/parseOpt"
	"testing"
)

func Test(t *testing.T) {
	err := &parseOpt.Cerror{
		Reason: 2,
	}
	list := cerrorList{
		list: []*parseOpt.Cerror{
			&parseOpt.Cerror{
				Reason: 1,
			},
		},
		currentIndex: 5,
	}
	t.Run("NormalPush", func(t *testing.T) {
		list.push(err)
		if len(list.list) != 2 {
			t.Fail()
		}
	})
	t.Run("ErrorIndex", func(t *testing.T) {
		if err.Index != 5 {
			t.Error("err.index must be set (5):", err.Index)
		}
	})
	t.Run("NoPush", func(t *testing.T) {
		list.push(&parseOpt.Cerror{
			Reason: 0,
		})
		if len(list.list) != 2 {
			t.Fail()
		}
	})
}
