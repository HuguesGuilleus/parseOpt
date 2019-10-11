---
title: La structure `Option`
---

```go
// Option stocke toutes les informations qui ont été traitées.
type Option struct {
	// Flag contient chaque drapeaux rencontrés.
	// true si le drapeau est réglé vrai (par défaut);
	// false si le drapeau est réglé à faux ou n’est pas présent.
	Flag map[string]bool
	// Option contient un tableau de chaînes de caractères
	// de chaque options rencontrés.
	Option map[string][]string
}
```

## Clé
La clé permettant d'accéder à l'élément souhaité (drapeau ou option) est le premier élément non vide de la liste suivante:
1. NameLong
2. NameShort
3. NameEnv
4. Sinon chaîne vide.

**Note:** L'argument `--` est enregistré dans les drapeaux avec comme clé `--`.

**Note:** Les arguments standards (non rattachés à une option) sont enregistrés dans les options une chaîne vide comme clé.
