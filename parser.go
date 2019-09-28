// parseOpt
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parseOpt

import (
	"os"
)

// Parse the arguments from the os.Args
func (opt *Option) ParseArgOs() *Option {
	return opt.ParseArg(os.Args[1:])
}

// Parse the arguments from the os.Args
func (list *SpecList) ParseArgOs() (opt *Option) {
	return list.toOption().ParseArgOs()
}

// Parse Os env, the file and Os arguments, and run the callBack
func (opt *Option) ParseOsAndFile(file string) (*Option, error) {
	opt.canRunCB = false
	opt.ParseOsEnv()
	err := opt.ParseFile(file)
	opt.ParseArgOs()
	opt.canRunCB = true
	opt.runCB()
	return opt, err
}

// Parse the environment variable and argument from OS
func (opt *Option) ParseOs() *Option {
	opt.canRunCB = false
	opt.ParseOsEnv()
	opt.ParseArgOs()
	opt.canRunCB = true
	opt.runCB()
	return opt
}

// Parse the environment variable and argument from OS
func (list *SpecList) ParseOs() (opt *Option) {
	return list.toOption().ParseOs()
}
