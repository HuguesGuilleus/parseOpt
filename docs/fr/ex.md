---
title: Exemple
---

Le programme suivant affiche des lettres ou des chiffres générés aléatoirement. L'exemple est présent dans le répertoire `_ex/` (l'exemple est en anglais).


```go
// Une liste de spécification utilisé pour l'analyse
var spec = parseOpt.SpecList{
	// Une spécification de dapreau
	&parseOpt.Spec{
		NameShort: "n",
		NameLong:  "number",
		Desc:      "Affiche des chiffres",
	},
	// Une spécification d'option
	&parseOpt.Spec{
		NameShort:  "l",
		NameLong:   "length",
		NeedArg:    true,
		OptionName: "taille",
		Desc:       "Le nombre de caractères",
	},
}
```


Dans le main:
```go
func main() {
	// On analyse les arguments
	opt := spec.ParseOs()

	// Nous récupèrons le nombre de caractère à générer
	length := 10 // valeur par défaut
	if opt.Option["length"] != nil {
		fmt.Sscanf(opt.Option["length"][0], "%d", &length)
	}

	// Nous générons les nombres
	nb := make([]byte, length)
	if _, err := rand.Read(nb); err != nil {
		fmt.Println(err)
		return
	}

	// Nous affichons les caractères générés soit
	// sous forme de nombre, soit sous forme de lettre.
	modeNumber := opt.Flag["number"]
	for _, char := range nb {
		if modeNumber {
			fmt.Printf("%c", '0'+char%10)
		} else {
			fmt.Printf("%c", 'A'+char%26)
		}
	}
	fmt.Println()
}
```

## Drapeau
Pour afficher des lettres
```bash
./_ex
./_ex -n=false
./_ex --number=false
```

Pour afficher des chiffres:
```bash
./_ex -n
./_ex -n=true
./_ex --number=true
```

Les valeurs possibles des booléens sont: `0`, `1`, `true`, `false`, `True`, `False`, `TRUE`, `FALSE`. Une autre valeur donnera un avertissement et sera compté comme vrai.

## Option
On peut modifier la taille:
```bash
./_ex -l=20
./_ex -l 20
./_ex --length=20
./_ex --length 20
```

## Drapeau et option
On peut combiner les drapeaux avec les options. Les syntaxes suivantes sont équivalentes:
```bash
./_ex -nl=20
./_ex -nl 20
./_ex -n -l 20
```

## Aide
On peut afficher les options possibles:
```bash
./_ex -h
./_ex --help
```

## Vérification
La liste de spécification peut-être vérifiée (la présence d'un description, ...) voir le fichier `_ex/main_test.go`.

[Vérification des spécifications](./spec#vérification)
