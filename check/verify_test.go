// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package check

import (
	"github.com/HuguesGuilleus/parseOpt"
	"testing"
)

func TestVerifySpecNameShort(t *testing.T) {
	for _, name := range []string{"", "a", "5"} {
		testSpecNoError(t, &parseOpt.Spec{
			NameShort: name,
			Desc:      "Desc ...",
		})
	}
	testSpecError("Too Long", t, parseOpt.E_NAME_TOOLONG, &parseOpt.Spec{
		NameShort: "aaa",
		Desc:      "description...",
	})
	for _, char := range "-= \t" {
		testSpecError("Unautorised char", t, parseOpt.E_NAME_CHAR, &parseOpt.Spec{
			NameShort: string(char),
			Desc:      "description...",
		})
	}
}

func TestVerifySpecNameLong(t *testing.T) {
	for _, name := range []string{"aaa", "", "--"} {
		testSpecNoError(t, &parseOpt.Spec{
			NameLong: name,
			Desc:     "desc ...",
		})
	}
	testSpecError("Too short", t, parseOpt.E_NAME_TOOSHORT, &parseOpt.Spec{
		NameLong: "a",
		Desc:     "description...",
	})
	for _, char := range "= \t" {
		testSpecError("Unautorized character", t, parseOpt.E_NAME_CHAR, &parseOpt.Spec{
			NameLong: string(char) + "yolo",
			Desc:     "Yolo ...",
		})
	}
	for _, input := range []string{"-yolo", "--yolo"} {
		testSpecError("dash", t, parseOpt.E_NAME_BEGIN_DASH, &parseOpt.Spec{
			NameLong: input,
			Desc:     "desc ...",
		})
	}
}

func TestVerifSpecNameEnv(t *testing.T) {
	for _, name := range []string{"aaa", ""} {
		testSpecNoError(t, &parseOpt.Spec{
			NameEnv: name,
			Desc:    "desc ...",
		})
	}
	for _, char := range "= \t" {
		testSpecError("Unautorised char", t, parseOpt.E_NAME_CHAR, &parseOpt.Spec{
			NameEnv: string(char),
			Desc:    "desc ...",
		})
	}
}

func TestVerifSpecCallBack(t *testing.T) {
	testSpecNoError(t, &parseOpt.Spec{
		NameLong:   "aaa",
		Desc:       "desc ...",
		OptionName: "opt",
		NeedArg:    true,
		CBOption:   func(_ []string) {},
	})
	testSpecNoError(t, &parseOpt.Spec{
		NameLong: "aaa",
		Desc:     "desc ...",
		CBFlag:   func() {},
	})
	testSpecError("CbFlagOnOption", t, parseOpt.E_CBFLAG_EXIST, &parseOpt.Spec{
		NameLong:   "aaa",
		Desc:       "desc ...",
		OptionName: "opt",
		NeedArg:    true,
		CBFlag:     func() {},
	})
	testSpecError("CbOptiononFlag", t, parseOpt.E_CBOPT_EXIST, &parseOpt.Spec{
		NameLong: "aaa",
		Desc:     "desc ...",
		CBOption: func(_ []string) {},
	})
}

func TestVerifSpecDesc(t *testing.T) {
	testSpecNoError(t, &parseOpt.Spec{
		NameShort: "a",
		Desc:      "A great flag!",
	})
	testSpecNoError(t, &parseOpt.Spec{
		NameShort:  "a",
		NeedArg:    true,
		Desc:       "A great option!",
		OptionName: "prog.go",
	})
	testSpecError("noDesc", t, parseOpt.E_NODESC, &parseOpt.Spec{
		NameShort: "a",
	})
	testSpecError("noOptionName", t, parseOpt.E_NOOPTIONNAME, &parseOpt.Spec{
		NameShort: "a",
		Desc:      "Une supper option!",
		NeedArg:   true,
	})
}

func TestVerify2Dash(t *testing.T) {
	testSpecNoError(t, &parseOpt.Spec{
		NameShort: "a",
		NameEnv:   "AAA",
		NameLong:  "--",
		Desc:      "The two dash flag",
		CBFlag:    func() {},
	})
	testSpecError("NeedArg", t, parseOpt.E_DOUBLEDASH_NEEDARG, &parseOpt.Spec{
		NameLong:   "--",
		Desc:       "The two dash flag",
		NeedArg:    true,
		OptionName: "yolo",
	})
}

func TestVerifySpecListOk(t *testing.T) {
	list := &parseOpt.SpecList{
		&parseOpt.Spec{
			NameShort: "a",
			NameLong:  "aaa",
			NameEnv:   "AAA",
			Desc:      "Description ...",
			CBFlag:    func() {},
		},
		&parseOpt.Spec{
			NameShort:  "b",
			NameLong:   "bbb",
			NameEnv:    "BBB",
			NeedArg:    true,
			OptionName: "optionName",
			CBOption:   func(_ []string) {},
			Desc:       "Description ...",
		},
	}
	errList := &cerrorList{
		list: []*parseOpt.Cerror{},
	}
	verifySpecList(list, errList)
	if len(errList.list) != 0 {
		t.Error("Error Unexpected:", errList.list)
		t.Log("Input:", list)
	}
}

func TestVerifySpecListItemError(t *testing.T) {
	testSpecListError(t, parseOpt.E_NODESC, &parseOpt.SpecList{
		&parseOpt.Spec{
			NameShort: "a",
			Desc:      "Description ...",
		},
		&parseOpt.Spec{
			NameShort: "b",
		},
	})
}

func TestVerifySpecListSameName(t *testing.T) {
	t.Run("Short", func(t *testing.T) {
		testSpecListError(t, parseOpt.E_SAME_NAMESHORT, &parseOpt.SpecList{
			&parseOpt.Spec{
				NameShort: "a",
				Desc:      "Description ...",
			},
			&parseOpt.Spec{
				NameShort: "a",
				Desc:      "Description ...",
			},
		})
	})
	t.Run("Long", func(t *testing.T) {
		testSpecListError(t, parseOpt.E_SAME_NAMELONG, &parseOpt.SpecList{
			&parseOpt.Spec{
				NameLong: "aaa",
				Desc:     "Description ...",
			},
			&parseOpt.Spec{
				NameLong: "aaa",
				Desc:     "Description ...",
			},
		})
	})
	t.Run("Env", func(t *testing.T) {
		testSpecListError(t, parseOpt.E_SAME_NAMEENV, &parseOpt.SpecList{
			&parseOpt.Spec{
				NameEnv: "AAA",
				Desc:    "Description ...",
			},
			&parseOpt.Spec{
				NameEnv: "AAA",
				Desc:    "Description ...",
			},
		})
	})
	t.Run("Empty", func(t *testing.T) {
		testSpecListError(t, parseOpt.E_SAME_EMPTY, &parseOpt.SpecList{
			&parseOpt.Spec{
				Desc: "Description ...",
			},
			&parseOpt.Spec{
				Desc: "Description ...",
			},
		})
	})
}

func testSpecListError(t *testing.T, err int, list *parseOpt.SpecList) {
	errList := &cerrorList{
		list: []*parseOpt.Cerror{},
	}
	verifySpecList(list, errList)
	if len(errList.list) != 1 {
		t.Error("We expected one error:", errList.list)
		return
	}
	if errList.list[0].Reason != err {
		t.Errorf("Expected err %d, %v", err, parseOpt.Cerror{
			Reason: err,
		})
		t.Logf("Received error %d, %v:", errList.list[0].Reason, errList.list[0])
	}
	if errList.list[0].Index != 1 {
		t.Error("The index must be 1; received:", errList.list[0].Index)
	}
}

// Tool for case sin error
func testSpecNoError(t *testing.T, spec *parseOpt.Spec) {
	t.Run("Normal", func(t *testing.T) {
		errList := &cerrorList{
			list: []*parseOpt.Cerror{},
		}
		verifySpec(spec, errList)
		if len(errList.list) != 0 {
			t.Error("Error Unexpected:", errList.list)
			t.Log("Input:", spec)
		}
	})
}

// Tool for test with an error
func testSpecError(testName string, t *testing.T, err int, spec *parseOpt.Spec) {
	t.Run(testName, func(t *testing.T) {
		errList := cerrorList{
			list: []*parseOpt.Cerror{},
		}
		verifySpec(spec, &errList)
		if len(errList.list) != 1 {
			t.Errorf("One error expected, received %d error(s).", len(errList.list))
			t.Log("Input:", spec)
			t.Log("Error expected: ", err, parseOpt.Cerror{
				Reason: err,
			})
			t.Log("Error received: ", errList.list)
			return
		}
		if errList.list[0].Reason != err {
			t.Error("It's not the good error")
			t.Log("Input:", spec)
			t.Log("Error expected: ", err, parseOpt.Cerror{
				Reason: err,
			})
			t.Log("Error received: ", errList.list[0].Reason, errList.list[0])
		}
	})
}
