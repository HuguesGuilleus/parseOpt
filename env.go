// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
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
	return opt
}

// func (opt *Option) ParseFile() *Option

func (opt *Option) ParseLine(line string) *Option {
	keyEnv, value, ok := parseLine(line)
	if ok {
		spec := opt.spec.getEnv(keyEnv)
		if spec == nil {
			ErrLog.Print("Unknown key: ",keyEnv)
			return opt
		}
		key := spec.key()
		if spec.NeedArg {
			opt.Option[key] = append(opt.Option[key], value)
			if spec.CBOption != nil {
				spec.CBOption(opt.Option[key])
			}
		} else {
			b := toBool(value)
			opt.Flag[key] = b
			if b && spec.CBFlag != nil {
				spec.CBFlag()
			}
		}
	}
	return opt
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
