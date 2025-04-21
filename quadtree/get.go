package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du quadtree q.
func (q *Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {

	height := len(contentHolder)
	if height == 0 {
		return
	}
	width := len(contentHolder[0])

	for y := 0; y < height; y++ {
		rowY := topLeftY + y
		for x := 0; x < width; x++ {
			contentHolder[y][x] = getValueFromNode(q.root, topLeftX+x, rowY)
			if configuration.Global.Infinie {
				height := len(contentHolder)
				width := 0
				if height > 0 {
					width = len(contentHolder[0])
				}

				q.EnsureInfiniteGeneration(topLeftX, topLeftY, width, height)

			}
		}
	}
}

// Fonction récursive pour trouver la valeur correspondant aux coordonnées (x, y) dans le Quadtree
func getValueFromNode(n *node, x, y int) int {
	// Si le nœud est nul ou si les coordonnées (x, y) sont hors des limites du nœud, retourne -1

	if n == nil || x < n.topLeftX || x >= n.topLeftX+n.Width || y < n.topLeftY || y >= n.topLeftY+n.Height {
		return -1 // Valeur par défaut pour les zones hors limites
	}

	// Si le nœud est une feuille, retourne directement son contenu
	if n.isLeaf {
		return n.content
	}

	// Calcule les coordonnées centrales du nœud actuel
	midX := n.topLeftX + n.Width/2
	midY := n.topLeftY + n.Height/2

	// Utilise un cheminement basé sur les coordonnées pour réduire les comparaisons
	switch {
	case x < midX && y < midY: // En haut à gauche
		return getValueFromNode(n.topLeftNode, x, y)
	case x < midX && y >= midY: // En bas à gauche
		return getValueFromNode(n.bottomLeftNode, x, y)
	case x >= midX && y < midY: // En haut à droite
		return getValueFromNode(n.topRightNode, x, y)
	default: // En bas à droite
		return getValueFromNode(n.bottomRightNode, x, y)
	}
}
