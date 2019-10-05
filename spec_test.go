// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"bytes"
	"testing"
)

func TestSpecKey(t *testing.T) {
	t.Log("Order of priority: NameLong>NameShort>NameEnv>\"\"")
	t.Run("Long", func(t *testing.T) {
		spec := &Spec{
			NameLong: "aaa",
		}
		for _, short := range []string{"a", ""} {
			for _, env := range []string{"AAA", ""} {
				spec.NameShort = short
				spec.NameEnv = env
				if spec.key() != "aaa" {
					t.Fail()
				}
			}
		}
	})
	t.Run("Short", func(t *testing.T) {
		spec := &Spec{
			NameShort: "a",
		}
		for _, env := range []string{"AAA", ""} {
			spec.NameEnv = env
			if spec.key() != "a" {
				t.Error("The spec.key() must return NameShort if it's present")
			}
		}
	})
	t.Run("Env", func(t *testing.T) {
		spec := &Spec{
			NameEnv: "AAA",
		}
		if spec.key() != "AAA" {
			t.Fail()
		}
	})
	t.Run("Empty", func(t *testing.T) {
		spec := &Spec{}
		if spec.key() != "" {
			t.Fail()
		}
	})
}

func TestSpecListGetxxx(t *testing.T) {
	a := &Spec{
		NameShort: "a",
		NameLong:  "aaa",
		NameEnv:   "AAA",
	}
	list := SpecList{
		a,
		&Spec{
			NameShort: "b",
			NameLong:  "bbb",
			NameEnv:   "BBB",
		},
	}
	t.Run("Short exist", func(t *testing.T) {
		item := list.getShort("a")
		if item != a {
			t.Errorf("Expected: %+v\n", a)
			t.Logf("Receveid: %+v\n", item)
		}
	})
	t.Run("Short no exist", func(t *testing.T) {
		item := list.getShort("z")
		if item != nil {
			t.Errorf("Expected: %+v\n", nil)
			t.Logf("Receveid: %+v\n", item)
		}
	})
	t.Run("Long exist", func(t *testing.T) {
		item := list.getLong("aaa")
		if item != a {
			t.Errorf("Expected: %+v\n", a)
			t.Logf("Receveid: %+v\n", item)
		}
	})
	t.Run("Long no exist", func(t *testing.T) {
		item := list.getLong("zzz")
		if item != nil {
			t.Errorf("Expected: %+v\n", nil)
			t.Logf("Receveid: %+v\n", item)
		}
	})
	t.Run("Env exist", func(t *testing.T) {
		item := list.getEnv("AAA")
		if item != a {
			t.Errorf("Expected: %+v\n", a)
			t.Logf("Receveid: %+v\n", item)
		}
	})
	t.Run("Env no exist", func(t *testing.T) {
		item := list.getEnv("ZZZ")
		if item != nil {
			t.Errorf("Expected: %+v\n", nil)
			t.Logf("Receveid: %+v\n", item)
		}
	})
}
func TestSpecListGet(t *testing.T) {
	onlyShort := &Spec{NameShort: "a"}
	onlyLong := &Spec{NameLong: "aaa"}
	booth := &Spec{NameShort: "b", NameLong: "bbb"}
	tree := &Spec{NameShort: "c", NameLong: "ccc", NameEnv: "CCC"}
	empty := &Spec{}
	list := SpecList{
		onlyShort,
		onlyLong,
		booth,
		empty,
		tree,
	}
	t.Run("Short", func(t *testing.T) {
		if p := list.get("a"); p != onlyShort {
			t.Errorf("(%p) %p\n", onlyShort, p)
		}
	})
	t.Run("Long", func(t *testing.T) {
		if p := list.get("aaa"); p != onlyLong {
			t.Errorf("(%p) %p\n", onlyLong, p)
		}
	})
	t.Run("Both", func(t *testing.T) {
		if p := list.get("b"); p != booth {
			t.Errorf("By short: (%p) %p\n", booth, p)
		}
		if p := list.get("bbb"); p != booth {
			t.Errorf("By long: (%p) %p\n", booth, p)
		}
	})
	t.Run("Tree", func(t *testing.T) {
		if p := list.get("c"); p != tree {
			t.Errorf("By short: (%p) %p\n", tree, p)
		}
		if p := list.get("ccc"); p != tree {
			t.Errorf("By long:  (%p) %p\n", tree, p)
		}
		if p := list.get("CCC"); p != tree {
			t.Errorf("By env:   (%p) %p\n", tree, p)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		if p := list.get(""); p != empty {
			t.Errorf("(%p) %p\n", empty, p)
		}
	})
	t.Run("Nothing", func(t *testing.T) {
		if p := list.get("z"); p != nil {
			t.Errorf("(%v) %p\n", nil, p)
		}
	})
}

func TestSpecListHelp(t *testing.T) {
	expected := []string{
		"\t\033[1m-a       \033[0m Description ...",
		"\t\033[1m   --aaa \033[0m Description ...",
		"\t\033[1m-b --bbbb\033[0m Description ...",
		"\t\033[1m-c        \033[0;4moption\033[0m Description ...",
		"\t\033[1m   --    \033[0m Yolo",
	}
	list := SpecList{
		&Spec{
			NameShort: "a",
			Desc:      "Description ...",
		},
		&Spec{
			NameLong: "aaa",
			Desc:     "Description ...",
		},
		&Spec{
			NameShort: "b",
			NameLong:  "bbbb",
			Desc:      "Description ...",
		},
		&Spec{
			NameShort:  "c",
			NeedArg:    true,
			Desc:       "Description ...",
			OptionName: "option",
		},
		&Spec{
			NameLong: "--",
			NeedArg:  false,
			Desc:     "Yolo",
		},
	}
	buff := &bytes.Buffer{}
	list.Help(buff)
	for line, expectedLine := range expected {
		returnedLine, _ := buff.ReadBytes(byte('\n'))
		returnedLine = returnedLine[:len(returnedLine)-1]
		if expectedLine != string(returnedLine) {
			t.Errorf("Line %d:\n", line)
			t.Log("Exp:", expectedLine)
			t.Log("Ret:", string(returnedLine))
			t.Log("Exp:", []byte(expectedLine))
			t.Log("Ret:", returnedLine)
		}
	}
}

func TestSpecListToOption(t *testing.T) {
	t.Run("help no add (short)", func(t *testing.T) {
		list := SpecList{
			&Spec{
				NameShort: "h",
				Desc:      "Description",
			},
		}
		opt := list.toOption()
		if len(opt.spec) != 2 {
			t.Error("More item was add to the list")
			t.Logf("%+v\n", list)
		}
	})
	t.Run("help no add (long)", func(t *testing.T) {
		list := SpecList{
			&Spec{
				NameLong: "help",
				Desc:     "Description",
			},
		}
		opt := list.toOption()
		if len(opt.spec) != 2 {
			t.Error("More item was add to the list")
			t.Logf("%+v\n", list)
		}
	})
	t.Run("help add", func(t *testing.T) {
		list := SpecList{
			&Spec{
				NameLong: "yolo",
				Desc:     "Description",
			},
		}
		opt := list.toOption()
		h := opt.spec.getShort("h")
		longHelp := opt.spec.getLong("help")
		if h == nil || h != longHelp || h.CBFlag == nil || h.NeedArg {
			t.Error("The help option was not added to the list")
		}
	})
	t.Run("AddDoubleDash", func(t *testing.T) {
		list := SpecList{
			&Spec{
				NameLong: "yolo",
				Desc:     "Description",
			},
		}
		dash := list.toOption().spec.getLong("--")
		if dash == nil {
			t.Error("The dash option was not insterted to the list")
			return
		}
		if dash.NeedArg {
			t.Error("In the new double dash, NeedArg must be false")
		}
	})
}
