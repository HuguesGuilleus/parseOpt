// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"io"
	"os"
)

// Specification for an option or an flag
type Spec struct {
	// Name used in os.Args
	NameShort, NameLong string
	// Name used in os.Env or in file
	NameEnv string

	// Used in help, it describe this specification.
	Desc string
	// [Option only] It's the name of the option, used in help.
	OptionName string

	// True: -a|--aaa option -a|--aaa=option
	// False: -a\--aaaa=true|false
	NeedArg bool

	// [Flag only] Callback exectuted after the parsing.
	CBFlag func()
	// [Option only] Callback exectuted after the parsing.
	CBOption func([]string)
}

// The key is used for the maps in Option
func (s *Spec) key() string {
	if s == nil {
		return "="
	}
	if len(s.NameLong) != 0 {
		return s.NameLong
	}
	if len(s.NameShort) != 0 {
		return s.NameShort
	}
	if len(s.NameEnv) != 0 {
		return s.NameEnv
	}
	return ""
}

// A list of specification, it will be verified
// and used to parse Arg or/and Environment variable
type SpecList []*Spec

// Give the first item who NameShort equals key
func (list *SpecList) getShort(key string) (value *Spec) {
	for _, item := range *list {
		if item.NameShort == key {
			return item
		}
	}
	return
}

// Give the first item who NameLong equals key
func (list *SpecList) getLong(key string) (value *Spec) {
	for _, item := range *list {
		if item.NameLong == key {
			return item
		}
	}
	return
}

// Give the first item who NameEnv equals key
func (list *SpecList) getEnv(key string) (value *Spec) {
	for _, item := range *list {
		if item.NameEnv == key {
			return item
		}
	}
	return
}

// Get the spec by a key used in a Option struct
func (list *SpecList) get(key string) (spec *Spec) {
	if l := len(key); l == 0 {
		for _, spec := range *list {
			if spec.NameShort == "" && spec.NameLong == "" && spec.NameEnv == "" {
				return spec
			}
		}
		return nil
	} else if l == 1 {
		return list.getShort(key)
	} else {
		specLong := list.getLong(key)
		if specLong != nil {
			return specLong
		} else {
			return list.getEnv(key)
		}
	}
}

// It will be used for the option -h or --help
func (list *SpecList) Help(w io.Writer) {
	names := make([]string, len(*list))
	maxLong := 0
	for i, spec := range *list {
		names[i] = "\t\033[1m"
		if spec.NameShort != "" {
			names[i] += "-" + spec.NameShort + " "
		} else {
			names[i] += "   "
		}
		if spec.NameLong == "--" {
			names[i] += "--"
		} else if spec.NameLong != "" {
			names[i] += "--" + spec.NameLong
		}
		if l := len(names[i]); l > maxLong {
			maxLong = l
		}
	}
	spaces := ""
	for i := 0; i < maxLong; i++ {
		spaces += " "
	}
	for i, spec := range *list {
		io.WriteString(w, (names[i] + spaces)[:maxLong])
		if spec.NeedArg && spec.OptionName != "" {
			io.WriteString(w, " \033[0;4m"+spec.OptionName)
		}
		io.WriteString(w, "\033[0m "+spec.Desc+"\n")
	}
}

// Verify the spec, if possible add a help option
func (list *SpecList) toOption() (opt *Option) {
	if list.getShort("h") == nil && list.getLong("help") == nil {
		*list = append(*list, &Spec{
			NameShort: "h",
			NameLong:  "help",
			Desc:      "Print this help",
			NeedArg:   false,
			CBFlag: func() {
				list.Help(os.Stdout)
				os.Exit(0)
			},
		})
	}
	if dash := list.getLong("--"); dash == nil {
		*list = append(*list, &Spec{
			NameLong: "--",
			NeedArg:  false,
		})
	}
	return &Option{
		Flag:     make(map[string]bool),
		Option:   make(map[string][]string),
		spec:     *list,
		canRunCB: true,
	}
}
