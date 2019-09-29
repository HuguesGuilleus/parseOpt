---
title: Spécifications
---

Afin d'analyser les arguments, vous devez fournir des instructions pour l'analyse. Chaque option ou drapeau, est un élément de type `Spec` qui sont regroupées dans une liste `SpecList` qui servira à l'analyse. Ce module fait une distinction entre les différents éléments selon ce qu'il contiendront:
- Les *options*: ils contiendront une liste de chaîne de caractères.
- Les *drapeaux* ou *flags*: ils contiendront un booléen; vrai si présent, faux sinon.

```go
// Specification for an option or an flag
// Spécification pour une option ou un drapeau
type Spec struct {
	// Non utilisé dans os.Args
	NameShort, NameLong string
	// Nom utilisé dans os.Env ou dans un fichier
	NameEnv string

	// Utilisé dans l'aide pour décrire cette spécification
	Desc string
	// [Uniquement les options] Le nom de l'option dans l'aide
	OptionName string

	// Vrai pour une option, faux sinon
	NeedArg bool

	// [Uniquement les drapeaux] CallBack exécuté après l'analyse.
	CBFlag func()
	// [Uniquemen pour les option] CallBack exécuté après l'analyse
	CBOption func([]string)
}

// Une liste de spécification utilisé pour l'analyse
type SpecList []*Spec
```


## Notes
Dans un élément de spécification, si tous ces noms sont vides, la structure correspond aux arguments standards passés au programme.

**Specification d'aide:** Si la liste de spécifications ne contient pas d'élément avec «h» pour `NameShort` ou «help» pour `NameLong`, alors il ajoute une spécification d'aide contenant un *CallBack* qui liste les drapeaux et les options présentent puis quitte le programme.

La spécification avec `NameLong` correspondant à «`--`» sera forcément un drapeau.

Une spécification avec `NameLong` et `NameShort` vide correspond aux arguments standards.


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
