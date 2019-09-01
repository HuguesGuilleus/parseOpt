// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"bytes"
	"log"
	"testing"
)

func TestOptionParseLine(t *testing.T) {
	// Init log
	oldLog := ErrLog
	defer func() { ErrLog = oldLog }()
	buff := &bytes.Buffer{}
	ErrLog = log.New(buff, "", 0)
	// Init callBack and Option
	cbFlag, cbOption := true, true
	opt := &Option{
		Flag:   make(map[string]bool),
		Option: make(map[string][]string),
		spec: SpecList{
			&Spec{
				NameEnv: "YOLO",
				CBFlag: func() {
					cbFlag = false
				},
			},
			&Spec{
				NameEnv: "SWAG",
				NeedArg: true,
				CBOption: func(_ []string) {
					cbOption = false
				},
			},
		},
	}
	// Test
	t.Run("Flag", func(t *testing.T) {
		buff.Reset()
		opt.ParseLine("YOLO=True")
		t.Run("Value", func(t *testing.T) {
			if opt.Flag["YOLO"] != true {
				t.Fail()
			}
		})
		t.Run("CallBack", func(t *testing.T) {
			if cbFlag != false {
				t.Fail()
			}
		})
		t.Run("NoErrLog", func(t *testing.T) {
			if buff.Len() != 0 {
				t.Error("No ErrLog:", buff)
			}
		})
	})
	t.Run("Option", func(t *testing.T) {
		buff.Reset()
		opt.ParseLine("SWAG = lapin")
		t.Run("Value", func(t *testing.T) {
			if test2SliceString(opt.Option["SWAG"], []string{"lapin"}) {
				t.Error("Expected:", []string{"lapin"})
				t.Log("Received:", opt.Option["SWAG"])
			}
		})
		t.Run("CallBack", func(t *testing.T) {
			if cbOption {
				t.Fail()
			}
		})
		t.Run("NoErrLog", func(t *testing.T) {
			if buff.Len() != 0 {
				t.Error("No ErrLog:", buff)
			}
		})
	})
	t.Run("No exist", func(t *testing.T) {
		buff.Reset()
		opt.ParseLine("ZZZ=yyy")
		expected := "Unknown key: ZZZ\n"
		if err := buff.String(); err != expected {
			t.Error("Bad ErrLog:")
			t.Log("Expected:", expected)
			t.Log("Received:", err)
			t.Log("Expected:", []byte(expected))
			t.Log("Received:", []byte(err))
		}
	})
}

func TestParseLine(t *testing.T) {
	oldLog := ErrLog
	defer func() { ErrLog = oldLog }()
	buff := &bytes.Buffer{}
	ErrLog = log.New(buff, "", 0)
	t.Run("Normal", func(t *testing.T) {
		names := []string{"Normal", "Commentary", "Space and commentary"}
		inputs := []string{"yolo=swag", "yolo=swag#Carpe Diem", "  yolo =\tswag    #Carpe Diem"}
		for i, input := range inputs {
			t.Run(names[i], func(t *testing.T) {
				buff.Reset()
				key, value, ok := parseLine(input)
				if key != "yolo" || value != "swag" || ok != true {
					t.Error("Input: ", input)
					t.Log("Expected: yolo; swag; true")
					t.Logf("Returned: %s; %s; %t", key, value, ok)
				}
				if buff.Len() != 0 {
					t.Error("It's a normal line, unxepected ErrLog:", buff)
				}
			})
		}
	})
	t.Run("Comment line", func(t *testing.T) {
		buff.Reset()
		_, _, ok := parseLine(" \t#dwshdyhj")
		if ok != false {
			t.Error("Comment line must return ok=false")
		}
		if buff.Len() != 0 {
			t.Error("Comment line is not a error, unxepected ErrLog:", buff)
		}
	})
	t.Run("Space line", func(t *testing.T) {
		buff.Reset()
		_, _, ok := parseLine(" \t   ")
		if ok != false {
			t.Error("Space line must return ok=false")
		}
		if buff.Len() != 0 {
			t.Error("Space line is not a error, unxepected ErrLog:", buff)
		}
	})
	t.Run("ErrorSyntax", func(t *testing.T) {
		for _, input := range []string{"yolo", "yolo=", "=yolo"} {
			buff.Reset()
			_, _, ok := parseLine(input)
			if ok != false {
				t.Log("Input:", input)
				t.Error("Bad syntax must return ok=false")
			}
			exp := "Bad syntax for environment: '" + input + "'\n"
			if err := buff.String(); err != exp {
				t.Log("Input:", input)
				t.Error("Bap ErrLog:")
				t.Log("Expected:", exp)
				t.Log("Returned:", err)
			}
		}
	})
}
