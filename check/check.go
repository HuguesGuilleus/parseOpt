// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package check

import (
	"github.com/HuguesGuilleus/parseOpt"
	"testing"
)

// Check the list for test
func Check(t *testing.T, list *parseOpt.SpecList) {
	errList := &cerrorList{
		list: []*parseOpt.Cerror{},
	}
	verifySpecList(list, errList)
	for _, err := range errList.list {
		t.Error(err)
	}
}
