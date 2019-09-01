// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"fmt"
	"testing"
)

func ExampleOption() {
	opt := &Option{
		Flag: map[string]bool{
			"flag1": true,
			"flag2": false,
			"flag3": true,
		},
		Option: map[string][]string{
			"opt": []string{"aaa", "bbb"},
		},
	}
	fmt.Println(opt)
	// Output:
	// Flags:
	// 	flag1
	// 	flag3
	// Options:
	// 	opt: "aaa", "bbb"
}

func ExampleOptionEmpty() {
	opt := &Option{}
	fmt.Println(opt)
	// Output:
	// Flags:
	// 	<empty>
	// Options:
	// 	<empty>
}

func TestOptionRunCb(t *testing.T) {
	t.Run("NoRunCb", func(t *testing.T) {
		run := false
		opt := Option{
			spec: SpecList{
				&Spec{
					NameShort: "a",
					CBFlag: func() {
						run = true
					},
				},
			},
			Flag: map[string]bool{
				"a": true,
			},
			canRunCB: false,
		}
		opt.runCB()
		if run {
			t.Fail()
		}
	})
	t.Run("RunCB", func(t *testing.T) {
		runFlagA, runFlagB := false, false
		runOptionC, runOptionD := false, false
		opt := Option{
			spec: SpecList{
				&Spec{
					NameShort: "a",
					CBFlag:    func() { runFlagA = true },
				},
				&Spec{
					NameShort: "b",
					CBFlag:    func() { runFlagB = true },
				},
				&Spec{
					NameShort: "c",
					NeedArg:   true,
					CBOption:  func(_ []string) { runOptionC = true },
				},
				&Spec{
					NameShort: "d",
					NeedArg:   true,
					CBOption:  func(_ []string) { runOptionD = true },
				},
			},
			Flag: map[string]bool{
				"a": true,
				"b": false,
			},
			Option: map[string][]string{
				"c": []string{"yolo1", "yolo2", "yolo3"},
			},
			canRunCB: true,
		}
		opt.runCB()
		t.Run("FlagTrue", func(t *testing.T) {
			if runFlagA != true {
				t.Fail()
			}
		})
		t.Run("FlagFalse", func(t *testing.T) {
			if runFlagB != false {
				t.Fail()
			}
		})
		t.Run("OptionExist", func(t *testing.T) {
			if runOptionC != true {
				t.Fail()
			}
		})
		t.Run("OptionNoExist", func(t *testing.T) {
			if runOptionD != false {
				t.Fail()
			}
		})
	})
}
