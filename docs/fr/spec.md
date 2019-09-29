---
title: Spécifications
---

Afin d'analyser les arguments, vous devez fournir des instructions pour l'analyse. Chaque option ou drapeau, est un élément de type `Spec` qui sont regourpés dans une liste `SpecList` qui servira à l'analyse. Ce module fait une distinction entre les différents éléments selon ce qu'il contiendront:
- Les *options*: ils contiendront une liste de chaîne de caractères.
- Les *drapeaux* ou *flags*: ils contiendront un booléen; vrai si présent, faux sinon.

```go
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

// A list of specification, it will be verified
// and used to parse Arg or/and Environment variable
type SpecList []*Spec
```



> ## Notes
Si tous ces noms sont vides, la structure correspond aux arguments standards passés au programme.
>
> **Specification d'aide:** Le programme ajoute par défaut une spécification dans la liste des spécifications d'aide si NameShort `h` et NameLong `help` n'existe pas.
> **Specification d'aide:** Si la liste de spécifications ne contient pas d'élément avec «h» pour `NameShort` et «help» pour `NameLong`, alos il ajoute une spécification d'aide conteant un *CallBack* qui liste les options présente et quitte le programme.


## Vérification
Vous pouvez vérifier la syntaxe de vos spécifications grâce au sous module `check` dans les testes.

```go
package main

import (
	"github.com/HuguesGuilleus/parseOpt/check"
	"testing"
)

func TestSpecList(t *testing.T) {
	// list est votre liste de spécifications
	check.Check(t, list)
}
```
