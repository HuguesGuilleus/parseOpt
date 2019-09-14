// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"sort"
	"strings"
)

type Option struct {
	Flag     map[string]bool
	Option   map[string][]string
	spec     SpecList
	canRunCB bool
}

func (opt *Option) String() string {
	b := strings.Builder{}
	b.WriteString("Flags:\n")
	flags := []string{}
	for key, value := range opt.Flag {
		if value {
			flags = append(flags, key)
		}
	}
	if len(flags) > 0 {
		sort.Strings(flags)
		for _, flag := range flags {
			b.WriteRune('\t')
			b.WriteString(flag)
			b.WriteRune('\n')
		}
	} else {
		b.WriteString("\t<empty>\n")
	}
	b.WriteString("Options:\n")
	if len(opt.Option) != 0 {
		for key, list := range opt.Option {
			b.WriteString("\t" + key + ": ")
			for i, max := 0, len(list); i < max; i++ {
				b.WriteRune('"')
				b.WriteString(list[i])
				if i != max-1 {
					b.WriteString("\", ")
				} else {
					b.WriteString("\"\n")
				}
			}
		}
	} else {
		b.WriteString("\t<empty>\n")
	}
	return b.String()
}

// Run the Callbacks, it's used in functions of parsing
func (opt *Option) runCB() {
	if !opt.canRunCB {
		return
	}
	for _, spec := range opt.spec {
		key := spec.key()
		if spec.NeedArg {
			if opt.Option[key] != nil && spec.CBOption != nil {
				spec.CBOption(opt.Option[key])
			}
		} else {
			if opt.Flag[key] && spec.CBFlag != nil{
				spec.CBFlag()
			}
		}
	}
}
