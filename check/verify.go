// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package check

import (
	".."
)

// Verify the names syntax, the existence of a description and
// OptionName for Option and the callback
func verifySpec(s *parseOpt.Spec, errList *cerrorList) {
	verifySpecNameShort(s, errList)
	verifySpecNameLong(s, errList)
	verifySpecNameEnv(s, errList)
	verifySpecDesc(s, errList)
	verifySpecCallBack(s, errList)
}

func verifySpecNameShort(s *parseOpt.Spec, errList *cerrorList) {
	if l := len(s.NameShort); l == 1 {
		name := s.NameShort[0]
		if name <= ' ' || name == '-' || name == '=' {
			errList.push(&parseOpt.Cerror{
				Reason: parseOpt.E_NAME_CHAR,
			})
		}
	} else if l > 1 {
		errList.push(&parseOpt.Cerror{
			Reason: parseOpt.E_NAME_TOOLONG,
		})
	}
}

func verifySpecNameLong(s *parseOpt.Spec, errList *cerrorList) {
	if l := len(s.NameLong); l == 1 {
		errList.push(&parseOpt.Cerror{
			Reason: parseOpt.E_NAME_TOOSHORT,
		})
	} else if l > 1 {
		if s.NameLong == "--" {
			if s.NeedArg {
				errList.push(&parseOpt.Cerror{
					Reason: parseOpt.E_DOUBLEDASH_NEEDARG,
				})
			}
		} else {
			if s.NameLong != "--" && s.NameLong[0] == '-' {
				errList.push(&parseOpt.Cerror{
					Reason: parseOpt.E_NAME_BEGIN_DASH,
				})
			}
			for _, char := range s.NameLong {
				if char <= ' ' || char == '=' {
					errList.push(&parseOpt.Cerror{
						Reason: parseOpt.E_NAME_CHAR,
					})
				}
			}
		}
	}
}

func verifySpecNameEnv(s *parseOpt.Spec, errList *cerrorList) {
	for _, char := range s.NameEnv {
		if char <= ' ' || char == '=' {
			errList.push(&parseOpt.Cerror{
				Reason: parseOpt.E_NAME_CHAR,
				Str:    string(char),
			})
		}
	}
}

func verifySpecCallBack(s *parseOpt.Spec, errList *cerrorList) {
	if s.NeedArg == true && s.CBFlag != nil {
		s.CBFlag = nil
		errList.push(&parseOpt.Cerror{
			Reason: parseOpt.E_CBFLAG_EXIST,
		})
	} else if s.NeedArg == false && s.CBOption != nil {
		s.CBOption = nil
		errList.push(&parseOpt.Cerror{
			Reason: parseOpt.E_CBOPT_EXIST,
		})
	}
}

func verifySpecDesc(s *parseOpt.Spec, errList *cerrorList) {
	if len(s.Desc) == 0 {
		errList.push(&parseOpt.Cerror{
			Reason: parseOpt.E_NODESC,
		})
	}
	if s.NeedArg && len(s.OptionName) == 0 {
		errList.push(&parseOpt.Cerror{
			Reason: parseOpt.E_NOOPTIONNAME,
		})
	}
}

// Verify the syntaxe of all item and the no item duplication
func verifySpecList(list *parseOpt.SpecList, errList *cerrorList) {
	for i, spec := range *list {
		errList.currentIndex = i
		verifySpec(spec, errList)
		if spec.NameShort == "" && spec.NameLong == "" && spec.NameEnv == "" {
			for _, prev := range (*list)[:i] {
				if prev.NameShort == "" && prev.NameLong == "" && prev.NameEnv == "" {
					errList.push(&parseOpt.Cerror{
						Reason: parseOpt.E_SAME_EMPTY,
					})
				}
			}
		} else {
			for _, prev := range (*list)[:i] {
				if prev.NameShort == spec.NameShort && len(spec.NameShort) != 0 {
					errList.push(&parseOpt.Cerror{
						Reason: parseOpt.E_SAME_NAMESHORT,
						Str:    spec.NameShort,
					})
				}
				if prev.NameLong == spec.NameLong && len(spec.NameLong) != 0 {
					errList.push(&parseOpt.Cerror{
						Reason: parseOpt.E_SAME_NAMELONG,
						Str:    spec.NameLong,
					})
				}
				if prev.NameEnv == spec.NameEnv && len(spec.NameEnv) != 0 {
					errList.push(&parseOpt.Cerror{
						Reason: parseOpt.E_SAME_NAMEENV,
						Str:    spec.NameEnv,
					})
				}
			}
		}
	}
}
