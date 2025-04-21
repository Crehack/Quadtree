package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au-dessus (0),
// à droite (1), au-dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	// Pré-calcul des positions adjacentes
	adjacent := [4][2]int{
		{relativeXPos, relativeYPos - 1}, // Haut
		{relativeXPos + 1, relativeYPos}, // Droite
		{relativeXPos, relativeYPos + 1}, // Bas
		{relativeXPos - 1, relativeYPos}, // Gauche
	}

	// Vérification des conditions bloquantes
	for i, pos := range adjacent {
		x, y := pos[0], pos[1]
		blocking[i] = !IsInBounds(f.Content, x, y) || GetTileValue(f.Content, x, y) == -1
	}

	// Si l'eau est bloquante, ajout de la vérification correspondante
	if configuration.Global.BlockingWater {
		for i, pos := range adjacent {
			x, y := pos[0], pos[1]
			if IsInBounds(f.Content, x, y) && GetTileValue(f.Content, x, y) == 4 {
				blocking[i] = true
			}
		}
	}

	return blocking
}

// Vérifie que les indices sont dans les limites du tableau
func IsInBounds(content [][]int, x, y int) bool {
	mapWidth := len(content[0])
	mapHeight := len(content)

	if configuration.Global.RoundWorld {
		// Gestion du wrap-around
		x = (x + mapWidth) % mapWidth
		y = (y + mapHeight) % mapHeight
	}

	return x >= 0 && y >= 0 && y < mapHeight && x < mapWidth
}

// GetTileValue retourne la valeur de la tuile en tenant compte du RoundWorld
func GetTileValue(content [][]int, x, y int) int {
	mapWidth := len(content[0])
	mapHeight := len(content)

	if configuration.Global.RoundWorld {
		x = (x + mapWidth) % mapWidth
		y = (y + mapHeight) % mapHeight
	}

	// Retourne -1 si hors limites (ne devrait pas arriver avec RoundWorld activé)
	if x < 0 || y < 0 || y >= mapHeight || x >= mapWidth {
		return -1
	}

	return content[y][x]
}
