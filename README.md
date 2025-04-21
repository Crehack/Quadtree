# project-quadtree

Code source initial pour le projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01).

Par GUIRADO Jessy [Kowaulsky] et TERRIEN Swan [Destycoon]

# Registres des modifications
### Première partie
15 Novembre 2024 : Les fonctions readFloorFromFile, updateFromFileFloor, ont été rendues compatibles avec la QUESTION 5.1 par GUIRADO Jessy. (commit 651ffbd6a42470b0b1f856bbb25472b109cb07bb)

21 Novembre 2024 : Les fonctions MakeNode, MakeFromArray, isHomogeneous ont été conçues pour répondre à la QUESTION 5.2 par TERRIEN Swan. (commit 57e390a5314e0a4a1ec5b4a2994eb685b7e400bd)

22 Novembre 2024 : La fonction getValueFromNode a été faite pour faire le reste de la QUESTION 5.2 par GUIRADO Jessy. (commit 567422d516fcdb2f174d363c7c65d213c5f3ed54)

### Deuxième partie
10 Décembre 2024 : Le prototype de la fonction RandomFloorGeneration a été fait pour QUESTION 6.1 par TERRIEN Swan. (commit cf37e58c4b03d0370d5cb1d7b95df2b5fc1bd10b)

10 Décembre 2024 : La fonction randomFloorGeneration (+ implémentation), ainsi que quelques autres modifications et optimisations ont été faites pour faire le reste de la QUESTION 6.1 par GUIRADO Jessy. (commit 600fd002bd601ae0be26448baf3128de3553bee5)

18 Décembre 2024 : La fonction Savefloor, ainsi que quelques autres modifications ont été faites pour faire le reste de la QUESTION 6.2 par GUIRADO Jessy. (commit be232e9deb724bbfcebc87d70912975c46b11516)

18 Décembre 2024 : La fonction blocking a été modifier pour ajouter des contrainte au personnage cependant le personnage peut apparaitre sur l'eau et la contrainte s'applique une fois hors de l'eau pour repondre a la question 6.5. Fait par TERRIEN Swan. (commit 8b3e3f10d48034202da55379d6c05ecbc49a8af1)

6 Janvier 2025 : Modification du code par défaut concernant Floor et ses fonctions Update et Draw pour introduire de l'eau animée pour répondre à la QUESTION 6.3 par GUIRADO Jessy. (commit 27e5f3ac0c2272af9d34c08249279107d4ec0139)

8 Janvier 2025 : Ajout de l'extension de la caméra qui bloque avec ajout des fonction size pour mesurer la taille du terrain et camerabloc permetant de bloquer la camera au bord du terrain pour repondre a la question 6.6 par TERRIEN Swan (commit 3cd02bff1b0d503bb7b10e95bbe1a227c7ed9a28)

8 Janvier 2025 : Ajout de l'extension Zoom (touche : PageUp, PageDown) dans game Update, Optimisation de la fonction floor Init pour repondre a la question 6.12 par TERRIEN Swan (commit : 72d8e96b7694d576dde05ed54004e190faf8967f)

10 Janvier 2025 : Ajout de la fonction updatedCinematic et changement des coordonnées de type int par float64 pour répondre à la QUESTION 6.8 par GUIRADO Jessy. (commit : f66eeb7a79b6ae50e7d48007a0c262649dda79aa)

12 Janvier 2025 : Ajout de l'extension des teleporteur (Touche T pour poser un teleporteur et Espace pour se teleporter une fois sur le tp) ajout des structure et correction de bug concernant la camera blocante (maintenant compatible avec toute les camera) pour repondre a la question 6.4 par TERRIEN Swan (commit : 133d8553888ec54a3b976dc362e2d316b6d5b40e)

13 Janvier 2025 : Ajout des fonction canMoveTo, RoundWorld, GetTileValue, IsInBounds pour repondre a la question 6.11 + correction de bug de compatibiliter entre les extension par TERRIEN Swan. (commit : de5c1cf1ca54b1dcbe19cb299cca8f20e209b672 )