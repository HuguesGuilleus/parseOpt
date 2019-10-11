---
title: Sommaire
---

[(en) ![GoDoc](https://godoc.org/github.com/HuguesGuilleus/parseOpt?status.svg)](https://godoc.org/github.com/HuguesGuilleus/parseOpt)

Ce module analyse les variables d'environnement et les arguments grâce à une liste `SpecList` de spécification `Spec`; le résultat est retourné dans une structure `Option`.


## Installation
```bash
go get github.com/HuguesGuilleus/parseOpt/
```


## Sommaire
{% include index_file.liquid %}



## Journaux d'erreurs
Les erreur s'affiche dans `ErrLog` (type `*log.Logger`), qui redirige vers un `io.Writer` interne qui règle la ligne en rouge et écrit dans `os.Stderr`.
