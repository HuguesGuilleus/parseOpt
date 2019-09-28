---
title: Journaux d'évènements
---

Les erreur s'affiche dans `ErrLog` (type `*log.Logger`), qui redirige vers un `io.Writer` interne qui encadre la ligne par du rouge dans `os.Stderr`.
