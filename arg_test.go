// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"bytes"
	"log"
	"testing"
)

func TestParseArg(t *testing.T) {
	cbOpt := 0
	cbFlagA := true
	cbFlagB := false
	opt := &Option{
		Flag:   make(map[string]bool),
		Option: make(map[string][]string),
		spec: SpecList{
			&Spec{
				NameLong: "flagA",
				CBFlag:   func() { cbFlagA = false },
				Desc:     "Desciption",
			},
			&Spec{
				NameLong: "flagB",
				CBFlag:   func() { cbFlagB = true },
				Desc:     "Desciption",
			},
			&Spec{
				NameShort: "o",
				NameLong:  "opt",
				NeedArg:   true,
				CBOption: func(_ []string) {
					cbOpt++
				},
				Desc:       "Desciption",
				OptionName: "opt",
			},
		},
		canRunCB: true,
	}
	input := []string{"yolo", "--",
		"--flagA", "--flagB=false",
		"--opt", "value1", "-o", "value2", "--opt=value3"}
	returned := opt.ParseArg(input)
	t.Run("Same Option", func(t *testing.T) {
		if opt != returned {
			t.Error("Returned value:", returned)
		}
	})
	t.Run("Argument standart", func(t *testing.T) {
		returned := opt.Option[""]
		expected := []string{"yolo"}
		if test2SliceString(returned, expected) {
			t.Fail()
			t.Log("Expected:", expected)
			t.Log("Returned:", returned)
		}
	})
	t.Run("Flag Value", func(t *testing.T) {
		if opt.Flag["flagA"] != true {
			t.Fail()
		}
		if opt.Flag["flagB"] != false {
			t.Fail()
		}
	})
	t.Run("Flag CallBack", func(t *testing.T) {
		if cbFlagA {
			t.Error("The callback was not exectuted")
		}
		if cbFlagB {
			t.Error("The callBack of false flag must not be executed")
		}
	})
	t.Run("Option Value", func(t *testing.T) {
		returned := opt.Option["opt"]
		expected := []string{"value1", "value2", "value3"}
		if test2SliceString(returned, expected) {
			t.Fail()
			t.Log("Returned:", returned)
			t.Log("Expected:", expected)
		}
	})
	t.Run("Option CallBack", func(t *testing.T) {
		if cbOpt != 1 {
			t.Error("cbOpt (1):", cbOpt)
		}
	})
	t.Run("Double Dash", func(t *testing.T) {
		if opt.Flag["--"] != true {
			t.Fail()
		}
	})
}
func TestGetNext(t *testing.T) {
	oldLog := ErrLog
	defer func() { ErrLog = oldLog }()
	buff := &bytes.Buffer{}
	ErrLog = log.New(buff, "", 0)
	args := []arg{
		// For supplement
		arg{
			Type: argType_opt,
			Supp: "yolo",
		},
		// For get the next item
		arg{
			Type: argType_opt,
		},
		arg{
			Type:    argType_arg,
			Content: "swag",
		},
		// The next item is no a argument
		arg{
			Type:    argType_opt,
			Content: "stupidité",
		},
		arg{
			Type: argType_opt,
		},
		// There are no next item
		arg{
			Type:    argType_opt,
			Content: "feux",
		},
	}
	t.Run("Supplement", func(t *testing.T) {
		buff.Reset()
		pass := false
		value, ok := getNext(args, 0, &pass)
		if value != "yolo" {
			t.Error("The function must get the supplement")
			t.Log("Supplement (yolo): ", value)
		}
		if ok != true {
			t.Error("In this case there are no error")
		}
		if pass != false {
			t.Error("pass must be false")
		}
		if buff.Len() != 0 {
			t.Error("ErrLog must be empty")
			t.Log(buff.String())
		}
	})
	t.Run("Next item", func(t *testing.T) {
		buff.Reset()
		pass := false
		value, ok := getNext(args, 1, &pass)
		if value != "swag" {
			t.Error("The function must get the next item")
			t.Log("Supplement (swag): ", value)
		}
		if ok != true {
			t.Error("In this case there are no error")
		}
		if pass != true {
			t.Error("pass must be true")
		}
		if buff.Len() != 0 {
			t.Error("ErrLog must be empty")
			t.Log(buff.String())
		}
	})
	t.Run("Next item is an option", func(t *testing.T) {
		buff.Reset()
		pass := false
		_, ok := getNext(args, 3, &pass)
		if ok != false {
			t.Error("ok must be false")
		}
		if pass != false {
			t.Error("pass must be false")
		}
		err := buff.String()
		exp := "We need an argument to the option 'stupidité'\n"
		if err != exp {
			t.Error("ErrLog not same:")
			t.Log("Expected:", exp, ";;;")
			t.Log("Returned:", err, ";;;")
		}
	})
	t.Run("Last element", func(t *testing.T) {
		buff.Reset()
		pass := false
		_, ok := getNext(args, 5, &pass)
		if ok != false {
			t.Error("ok must be false")
		}
		if pass != false {
			t.Error("pass must be false")
		}
		err := buff.String()
		exp := "The last option 'feux' need an argument\n"
		if err != exp {
			t.Error("ErrLog not same:")
			t.Log("Expected:", exp, ";;;")
			t.Log("Returned:", err, ";;;")
		}
	})
}
func TestToBool(t *testing.T) {
	oldLog := ErrLog
	defer func() { ErrLog = oldLog }()
	buff := &bytes.Buffer{}
	ErrLog = log.New(buff, "", 0)
	t.Run("True", func(t *testing.T) {
		for _, input := range []string{"true", "True", "TRUE", "1", ""} {
			buff.Reset()
			if toBool(input) != true {
				t.Errorf("From '%s' bad return (true): false", input)
			}
			if buff.Len() != 0 {
				t.Errorf("From '%s' errLog must be empty: %s", input, buff.String())
			}
		}
	})
	t.Run("False", func(t *testing.T) {
		for _, input := range []string{"false", "False", "FALSE", "0"} {
			buff.Reset()
			if toBool(input) != false {
				t.Errorf("From '%s' bad return (false): true", input)
			}
			if buff.Len() != 0 {
				t.Errorf("From '%s' errLog must be empty: %s", input, buff.String())
			}
		}
	})
	t.Run("Error", func(t *testing.T) {
		for _, input := range []string{"fALSe", "egtbdsmfgrf", "46853"} {
			buff.Reset()
			if toBool(input) != true {
				t.Error("Return of bad string is <true>")
			}
			if buff.Len() == 0 {
				t.Error("It must be log an error")
			}
		}
	})
}

func TestCorrectArg(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		input := []string{"-a", "--abc", "yolo", "-a=486", "--abc=ABCDE", "--"}
		returned := correctArg(input)
		if test2SliceString(input, returned) {
			t.Error(input)
			t.Log(returned)
		}
	})
	t.Run("Problem", func(t *testing.T) {
		oldLog := ErrLog
		defer func() { ErrLog = oldLog }()
		buffer := &bytes.Buffer{}
		ErrLog = log.New(buffer, "", 0)
		// Test the return value
		input := []string{"swag", "-", "yolo", ""}
		expected := []string{"swag", "yolo"}
		received := correctArg(input)
		if test2SliceString(expected, received) {
			t.Error("Error return value:")
			t.Logf("input: %+q\n", input)
			t.Logf("expected: %+q\n", expected)
			t.Logf("returned: %+q\n", received)
		}
		// First line
		strRec, err := buffer.ReadString(byte('\n'))
		strExp := Cerror{
			Reason: E_ARG_DASH,
			Index:  2,
		}.Error() + "\n"
		if err != nil || strRec != strExp {
			t.Error("read error: ", err)
			t.Logf("str exp: '%s'\n", strExp)
			t.Logf("str rec: '%s'\n", strRec)
		}
		// Second line
		strRec, err = buffer.ReadString(byte('\n'))
		strExp = Cerror{
			Reason: E_ARG_EMPTY,
			Index:  4,
		}.Error() + "\n"
		if err != nil || strRec != strExp {
			t.Error("read error: ", err)
			t.Logf("str exp: '%s'\n", strExp)
			t.Logf("str rec: '%s'\n", strRec)
		}
	})
}
func TestVerifyArg(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		inputs := []string{"-a", "--abc", "yolo", "-a=486", "--abc=ABCDE", "--"}
		for _, input := range inputs {
			err := verifyArg(input)
			if err.Bool() {
				t.Error(input, err)
			}
		}
	})
	t.Run("Empty Arg", func(t *testing.T) {
		err := verifyArg("")
		if err.Reason != E_ARG_EMPTY {
			t.Error(err)
		}
	})
	t.Run("Just dash", func(t *testing.T) {
		err := verifyArg("-")
		if err.Reason != E_ARG_DASH {
			t.Error(err)
		}
	})
	t.Run("Option (short) ended by equal", func(t *testing.T) {
		err := verifyArg("-zedf=")
		if err.Reason != E_ARG_EQUAL {
			t.Error(err)
		}
	})
	t.Run("Option (long) ended by equal", func(t *testing.T) {
		err := verifyArg("--zedf=")
		if err.Reason != E_ARG_EQUAL {
			t.Error(err)
		}
	})
	t.Run("2dashequal", func(t *testing.T) {
		err := verifyArg("--=48653")
		if err.Reason != E_ARG_2DASH_EQUAL {
			t.Error("Err Reason:", err.Reason)
		}
	})
}

func TestParseRawArg(t *testing.T) {
	input := []string{"-ab", "-b=yolo", "--abc", "--abc=yolo", "swag", "--"}
	expected := []arg{
		arg{Type: argType_opt, Content: "a"},
		arg{Type: argType_opt, Content: "b"},
		arg{Type: argType_opt, Content: "b", Supp: "yolo"},
		arg{Type: argType_opt, Content: "abc"},
		arg{Type: argType_opt, Content: "abc", Supp: "yolo"},
		arg{Type: argType_arg, Content: "swag"},
		arg{Type: argType_2dash},
	}
	returned := parseRawArg(input)
	if test2SiliceArg(expected, returned) {
		t.Error("Input: ", input)
		t.Logf("Expected: %+v\n", expected)
		t.Logf("Returned: %+v\n", returned)
	}
}
func TestParseShortOpt(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		input := "-ab"
		expected := []arg{
			arg{Type: argType_opt, Content: "a"},
			arg{Type: argType_opt, Content: "b"},
		}
		returned := parseShortOpt(input)
		if test2SiliceArg(expected, returned) {
			t.Error("Input: ", input)
			t.Logf("Expected: %+v\n", expected)
			t.Logf("Returned: %+v\n", returned)
		}
	})
	t.Run("Supplement", func(t *testing.T) {
		input := "-ab=yolo"
		expected := []arg{
			arg{Type: argType_opt, Content: "a"},
			arg{
				Type:    argType_opt,
				Content: "b",
				Supp:    "yolo",
			},
		}
		returned := parseShortOpt(input)
		if test2SiliceArg(expected, returned) {
			t.Error("Input: ", input)
			t.Logf("Expected: %+v\n", expected)
			t.Logf("Returned: %+v\n", returned)
		}
	})
}
func TestParseLongOpt(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		input := "--abc"
		expected := arg{
			Type:    argType_opt,
			Content: "abc",
		}
		returned := parseLongOpt(input)
		if test2Arg(expected, returned) {
			t.Error("Input: ", input)
			t.Logf("Expected: %+v\n", expected)
			t.Logf("Returned: %+v\n", returned)
		}
	})
	t.Run("Supplement", func(t *testing.T) {
		input := "--abc=yolo"
		expected := arg{
			Type:    argType_opt,
			Content: "abc",
			Supp:    "yolo",
		}
		returned := parseLongOpt(input)
		if test2Arg(expected, returned) {
			t.Error("Input: ", input)
			t.Logf("Expected: %+v\n", expected)
			t.Logf("Returned: %+v\n", returned)
		}
	})
}

func test2SliceString(a, b []string) (err bool) {
	if len(a) != len(b) {
		return true
	}
	for i := range a {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}
func test2SiliceArg(a, b []arg) (err bool) {
	if len(a) != len(b) {
		return true
	}
	for i := range a {
		if test2Arg(a[i], b[i]) {
			return true
		}
	}
	return false
}
func test2Arg(a, b arg) (err bool) {
	if a.Type != b.Type {
		return true
	}
	if a.Content != b.Content {
		return true
	}
	if a.Supp != b.Supp {
		return true
	}
	return false
}
