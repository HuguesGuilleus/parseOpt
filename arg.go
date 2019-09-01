// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import ()

const (
	argType_arg   = 1
	argType_opt   = 2
	argType_2dash = 3
)

type arg struct {
	Type    int
	Content string
	Supp    string
}

// Correct, parse and add a array of string to opt
func (opt *Option) ParseArg(rawArg []string) *Option {
	args := parseRawArg(correctArg(rawArg))
	pass := false
	for i, a := range args {
		if pass {
			pass = false
			continue
		}
		switch a.Type {
		case argType_arg:
			opt.Option[""] = append(opt.Option[""], a.Content)
		case argType_2dash:
			opt.Flag["--"] = true
		case argType_opt:
			spec := opt.spec.get(a.Content)
			if spec != nil {
				if spec.NeedArg {
					if value, ok := getNext(args, i, &pass); ok {
						key := spec.key()
						opt.Option[key] = append(opt.Option[key], value)
					}
				} else {
					opt.Flag[spec.key()] = toBool(a.Supp)
				}
			} else {
				ErrLog.Print("This option is unknown: ", a.Content)
			}
		}
	}
	opt.runCB()
	return opt
}

// Get the argument of an option, with the supplement or the next element.
func getNext(args []arg, i int, pass *bool) (value string, ok bool) {
	if supp := args[i].Supp; len(supp) != 0 {
		return supp, true
	} else if len(args) <= i+1 {
		ErrLog.Printf("The last option '%s' need an argument", args[i].Content)
		return "", false
	} else if next := args[i+1]; next.Type == argType_arg {
		*pass = true
		return next.Content, true
	} else {
		ErrLog.Printf("We need an argument to the option '%s'", args[i].Content)
		return "", false
	}
}

// It use for parsing the boolean of flag or environnement
func toBool(value string) bool {
	switch value {
	case "true", "True", "TRUE", "1", "":
		return true
	case "false", "False", "FALSE", "0":
		return false
	default:
		ErrLog.Println(Cerror{
			Reason: E_SUPBTOBOOL,
			Str:    value,
		})
		return true
	}
}

// Delete argument who have not a good syntax.
// The error are writed in ErrLog.
func correctArg(rawArg []string) (newArg []string) {
	for i, a := range rawArg {
		err := verifyArg(a)
		if err.Bool() == false {
			newArg = append(newArg, a)
		} else {
			err.Index = i + 1
			ErrLog.Print(err)
		}
	}
	return
}

// Cheak the suntax of a simgle arg
func verifyArg(a string) (err Cerror) {
	if a == "" {
		return Cerror{
			Reason: E_ARG_EMPTY,
		}
	} else if a == "-" {
		return Cerror{
			Reason: E_ARG_DASH,
		}
	} else if a[0] == '-' && a[len(a)-1] == '=' {
		return Cerror{
			Reason: E_ARG_EQUAL,
		}
	}
	if len(a) > 3 && a[:3] == "--=" {
		return Cerror{
			Reason: E_ARG_2DASH_EQUAL,
		}
	} else {
		return
	}
}

// Parse a slice of arg to get the type of each element
func parseRawArg(rawArg []string) (args []arg) {
	for _, a := range rawArg {
		if a[0] == '-' {
			if a == "--" {
				args = append(args, arg{
					Type: argType_2dash,
				})
			} else if a[1] == '-' {
				args = append(args, parseLongOpt(a))
			} else {
				args = append(args, parseShortOpt(a)...)
			}
		} else {
			args = append(args, arg{
				Type:    argType_arg,
				Content: a,
			})
		}
	}
	return
}

// Parse a simple arg of type: -abc or -abc=value
func parseShortOpt(raw string) (args []arg) {
	for i, a := range raw[1:] {
		if a != '=' {
			args = append(args, arg{
				Type:    argType_opt,
				Content: string(a),
			})
		} else {
			args[i-1].Supp = raw[i+2:]
			return
		}
	}
	return
}

// parse a long arg of type: --aaa or --abc=value
func parseLongOpt(raw string) (a arg) {
	a.Type = argType_opt
	raw = raw[2:]
	for i, char := range raw {
		if char == '=' {
			a.Content = raw[:i]
			a.Supp = raw[i+1:]
			return
		}
	}
	a.Content = raw
	return
}
