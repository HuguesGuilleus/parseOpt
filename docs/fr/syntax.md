---
title: Syntaxe
lang: fr
---

# Option par la ligne de commande
Posons la spécification suivante sui est une option:
```go
	s := &Spec{
		NameShort:"a",
		NameLong:"aaa",
		NeedArg:true,
	}
```

Posons un drapeau:
```go
	s := &Spec{
		NameShort:"d",
	}
```

Les lignes suivantes sont équivalentes pour l'option a:
```
-a option
-a=option
-da option
-da=option
--aaa option
--aaa=option
```

Notons que comme il s'agit d'option, on peut faire passer plusieurs éléments pour une option. Exemple:
```
-a option1 -a option2 -a option3
```


# Drapeaux
## Ligne de commande
Posons le drapeau:
```go
	s := &parseOpt.Spec{
		NameShort: "d",
		NameShort: "ddd",
		NameEnv:   "DDD",
	}
```

### Vraie
Les lignes suivantes sont équivalentes
```
-d
-d=true
-d=True
-d=TRUE
-d=1
--ddd
--ddd=true
--ddd=True
--ddd=TRUE
--ddd=1
```

### Faux
Les lignes suivantes sont équivalentes
```
-d=false
-d=false
-d=FALSE
-d=0
--ddd=false
--ddd=false
--ddd=FALSE
--ddd=0
```

## Variable d'environment et fichier
On est obligé d'écrire un booléen.

### Vraie
```
DDD=true
DDD=True
DDD=TRUE
DDD=1
```

### Faux
```
DDD=false
DDD=False
DDD=FALSE
DDD=0
```
