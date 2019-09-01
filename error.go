// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"fmt"
)

const (
	E_CBFLAG_EXIST       = 1
	E_CBOPT_EXIST        = 2
	E_NAME_TOOLONG       = 3
	E_NAME_TOOSHORT      = 4
	E_NAME_CHAR          = 5
	E_NAME_BEGIN_DASH    = 6
	E_NODESC             = 7
	E_NOOPTIONNAME       = 8
	E_DOUBLEDASH_NEEDARG = 9

	E_SAME_NAMESHORT = 10
	E_SAME_NAMELONG  = 11
	E_SAME_NAMEENV   = 12
	E_SAME_EMPTY     = 13

	E_ARG_EMPTY       = 20
	E_ARG_DASH        = 21
	E_ARG_EQUAL       = 22
	E_ARG_2DASH_EQUAL = 23
	E_SUPBTOBOOL      = 30
)

// Custom error for this module
type Cerror struct {
	Reason int
	Index  int
	Str    string
}

func (e Cerror) Error() string {
	switch e.Reason {

	case E_CBFLAG_EXIST:
		return "A specification item with NeedArg==false don't need Flag CallBack."
	case E_CBOPT_EXIST:
		return "A specification item with NeedArg==true don't need Option CallBack."
	case E_NAME_TOOLONG:
		return fmt.Sprintf("A Spec.NameShort must have one character (index %d, '%s')", e.Index, e.Str)
	case E_NAME_TOOSHORT:
		return fmt.Sprintf("A Spec.NameLong should not have one character (index %d, '%s')", e.Index, e.Str)
	case E_NAME_CHAR:
		return fmt.Sprintf("Illegal character (=,space,dash for NameShort) in Spec.NameXxx (index %d, '%s')", e.Index, e.Str)
	case E_NAME_BEGIN_DASH:
		return fmt.Sprintf("A Spec.NameLong should not begin by dash (index %d, '%s')", e.Index, e.Str)
	case E_NODESC:
		return fmt.Sprintf("Write a Description (index %d)", e.Index)
	case E_NOOPTIONNAME:
		return fmt.Sprintf("Write the name of this option (index %d)", e.Index)
	case E_DOUBLEDASH_NEEDARG:
		return fmt.Sprintf("The \"--\" item is a flag by definition. (index %d)", e.Index)

	case E_SAME_NAMESHORT:
		return fmt.Sprintf("There are an other specification item with the same NameShort (index: %d, '%s')", e.Index, e.Str)
	case E_SAME_NAMELONG:
		return fmt.Sprintf("There are an other specification item with the same NameLong (index: %d, '%s')", e.Index, e.Str)
	case E_SAME_NAMEENV:
		return fmt.Sprintf("There are an other specification item with the same NameEnv (index: %d, '%s')", e.Index, e.Str)
	case E_SAME_EMPTY:
		return fmt.Sprintf("There are an other empty specification item (index: %d)", e.Index)

	case E_ARG_EMPTY:
		return fmt.Sprintf("The %d argument is empty.", e.Index)
	case E_ARG_DASH:
		return fmt.Sprintf("The %d argument is a simple dash.", e.Index)
	case E_ARG_EQUAL:
		return fmt.Sprintf("The %d argument is an option who ended by a equals sign.", e.Index)
	case E_SUPBTOBOOL:
		return fmt.Sprintf("Not a good bolean: '%s'.", e.Str)
	}
	return fmt.Sprintf("Unknow error: %d\n", e.Reason)
}
func (err Cerror) Bool() bool {
	if err.Reason != 0 {
		return true
	} else {
		return false
	}
}

func (e Cerror) print() {
	if e.Bool() {
		ErrLog.Print(e.Error())
	}
}
