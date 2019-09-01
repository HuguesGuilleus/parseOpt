// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package check

import (
	".."
)

type cerrorList struct {
	// The list of error
	list []*parseOpt.Cerror
	// The current index of a parseOpt.SpecList
	currentIndex int
}

func (errList *cerrorList) push(err *parseOpt.Cerror) {
	if err.Bool() {
		err.Index = errList.currentIndex
		errList.list = append(errList.list, err)
	}
}
