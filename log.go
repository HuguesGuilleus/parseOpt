// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"log"
	"os"
)

// The logger of error in this module. The line is writing in red.
var ErrLog *log.Logger = log.New(&writerErr{}, "", 0)

// A custom writer for writing in red
type writerErr struct{}

func (w *writerErr) Write(text []byte) (n int, err error) {
	os.Stderr.Write([]byte("\033[31m"))
	n, err = os.Stderr.Write(text)
	os.Stderr.Write([]byte("\033[0m"))
	return
}
