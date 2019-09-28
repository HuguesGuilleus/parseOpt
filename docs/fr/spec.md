---
title: Spéfifications
lang: fr
---

# Spécifications
Chaque élément de spécification est une structure Spec.
`// type SpecList []*Spec`

Ce module fait une distinction entre les différents éléments selon ce qu'il contiendront:
- Les *options*: ils contiendront une liste de chaîne de caractères.
- Les *drapeaux* ou *flags*: ils contiendront un booléen. Par défaut si le drapeau est présent sans autres valeur, la valeur sera vrai.

La structure `Spec` contient `NameShort` pour les arguments courts, `NameLong` pour les arguments longs et `NameEnv` pour les variables d'environnement ou les fichier de configurations. Si tous ces noms sont vides, la structure correspond aux arguments passés au programme. Exemple:
```bash
# Notons que -y est une option
$ ./prog -y yolo swag1 swag2

```

```go
package main
import "fmt"

func main()  {
	fmt.Println("Hello World")
}
```

## CallBack
Vous pouvez ajouter une fonction qui sera appelée après l'analyse des arguments ou environnement. Si la spécification correspond à un drapeau, utiliser `CBFlag` sinon utilisez `CBOpt`.

Suivant si la spécification est une option ou un drapeau


++:
help "" --

## Arguments

## Environnement
