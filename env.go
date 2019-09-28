// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	reComment = regexp.MustCompile("\\s*#.*$|\\s*$")
	reKey     = regexp.MustCompile("\\s*([\\w\\d]+)\\s*\\=.*")
	reVal     = regexp.MustCompile(".*=\\s*(.+)\\s*")
)

func (opt *Option) ParseOsEnv() *Option {
	for _, spec := range opt.spec {
		if len(spec.NameEnv) == 0 {
			continue
		}
		value, exist := os.LookupEnv(spec.NameEnv)
		if !exist {
			continue
		}
		key := spec.key()
		if spec.NeedArg {
			opt.Option[key] = append(opt.Option[key], value)
		} else {
			opt.Flag[key] = toBool(value)
		}
	}
	opt.runCB()
	return opt
}

// Parse a key=value file and run CB
func (opt *Option) ParseFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		opt.includeLine(string(line))
	}
	// opt.runCb()
	return nil
}

// Include one line key=value to opt
func (opt *Option) includeLine(line string) {
	keyEnv, value, ok := parseLine(line)
	if ok {
		spec := opt.spec.getEnv(keyEnv)
		if spec == nil {
			ErrLog.Print("Unknown key: ",keyEnv)
			return
		}
		key := spec.key()
		if spec.NeedArg {
			opt.Option[key] = append(opt.Option[key], value)
		} else {
			opt.Flag[key] = toBool(value)
		}
	}
	return
}

// Parse key=value in two string
func parseLine(src string) (key, val string, ok bool) {
	line := reComment.ReplaceAllString(src, "")
	if line == "" {
		return
	}
	ok, _ = regexp.MatchString("\\s*[\\w\\d]+\\s*=\\s*[\\w\\d]+\\s*", line)
	if ok {
		key = reKey.ReplaceAllString(line, "$1")
		val = reVal.ReplaceAllString(line, "$1")
	} else {
		ErrLog.Printf("Bad syntax for environment: '%s'", src)
	}
	return
}
