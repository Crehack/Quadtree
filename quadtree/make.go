package quadtree

// Fonction pour créer un Quadtree à partir d'un tableau 2D représentant un terrain
func MakeFromArray(floorContent [][]int) Quadtree {
	// Validation rapide du tableau d'entrée
	if len(floorContent) == 0 || len(floorContent[0]) == 0 {
		return Quadtree{}
	}

	// Construction du Quadtree
	return Quadtree{
		Width:  len(floorContent[0]), // Largeur du tableau
		Height: len(floorContent),    // Hauteur du tableau
		root:   makeNode(0, 0, len(floorContent[0]), len(floorContent), floorContent),
	}
}

// Fonction récursive pour construire un nœud du Quadtree
func makeNode(x, y, width, height int, floorContent [][]int) *node {
	// Si la région est homogène, retourne un nœud feuille
	if isHomogeneous(x, y, width, height, floorContent) {
		return &node{
			topLeftX:        x,
			topLeftY:        y,
			Width:           width,
			Height:          height,
			content:         floorContent[y][x],
			isLeaf:          true,
			topLeftNode:     nil,
			topRightNode:    nil,
			bottomLeftNode:  nil,
			bottomRightNode: nil,
		}
	}

	// Si la région est trop petite pour être divisée, retourne un nœud feuille
	if width == 1 && height == 1 {
		return &node{
			topLeftX:        x,
			topLeftY:        y,
			Width:           width,
			Height:          height,
			content:         floorContent[y][x],
			isLeaf:          true,
			topLeftNode:     nil,
			topRightNode:    nil,
			bottomLeftNode:  nil,
			bottomRightNode: nil,
		}
	}

	// Calcul des dimensions pour diviser la région en 4 sous-régions
	midX := x + width/2
	midY := y + height/2

	// Création du nœud parent avec ses sous-régions
	return &node{
		topLeftX:        x,
		topLeftY:        y,
		Width:           width,
		Height:          height,
		content:         -1,
		isLeaf:          false,
		topLeftNode:     makeNode(x, y, midX-x, midY-y, floorContent),                        // Haut gauche
		topRightNode:    makeNode(midX, y, width-(midX-x), midY-y, floorContent),             // Haut droit
		bottomLeftNode:  makeNode(x, midY, midX-x, height-(midY-y), floorContent),            // Bas gauche
		bottomRightNode: makeNode(midX, midY, width-(midX-x), height-(midY-y), floorContent), // Bas droit
	}
}

// Vérifie si toutes les cases dans une région ont la même valeur
func isHomogeneous(x, y, width, height int, floorContent [][]int) bool {
	// Utilisation d'une variable unique pour comparer les valeurs
	value := floorContent[y][x]

	// Boucle optimisée pour parcourir les cases
	for i := y; i < y+height; i++ {
		for j := x; j < x+width; j++ {
			if floorContent[i][j] != value {
				return false
			}
		}
	}
	return true
}
