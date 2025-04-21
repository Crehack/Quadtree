package quadtree

import (
	"fmt"
	"math/rand"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// EnsureInfiniteGeneration s'assure que le Quadtree peut gérer la génération infinie avec des valeurs aléatoires.
func (q *Quadtree) EnsureInfiniteGeneration(x, y, width, height int) {

	if !configuration.Global.Infinie {
		return
	}

	for x <= q.root.topLeftX || y <= q.root.topLeftY || x+width >= q.root.topLeftX+q.root.Width || y+height >= q.root.topLeftY+q.root.Height {
		q.expand()
	
	}
}

// expand agrandit le quadtree en doublant sa taille et en ajoutant des valeurs aléatoires.
func (q *Quadtree) expand() {
	newWidth := q.Width * 2
	newHeight := q.Height * 2

	newRoot := &node{
		topLeftX: q.root.topLeftX - q.Width/2,
		topLeftY: q.root.topLeftY - q.Height/2,
		Width:    newWidth,
		Height:   newHeight,
		isLeaf:   false,
	}

	if q.root.topLeftX < newRoot.topLeftX+newWidth/2 {
		if q.root.topLeftY < newRoot.topLeftY+newHeight/2 {
			newRoot.topLeftNode = q.root
		} else {
			newRoot.bottomLeftNode = q.root
		}
	} else {
		if q.root.topLeftY < newRoot.topLeftY+newHeight/2 {
			newRoot.topRightNode = q.root
		} else {
			newRoot.bottomRightNode = q.root
		}
	}

	if newRoot.topLeftNode == nil {
		newRoot.topLeftNode = createRandomNode(newRoot.topLeftX, newRoot.topLeftY, q.Width, q.Height)
	}
	if newRoot.topRightNode == nil {
		newRoot.topRightNode = createRandomNode(newRoot.topLeftX+q.Width, newRoot.topLeftY, q.Width, q.Height)
	}
	if newRoot.bottomLeftNode == nil {
		newRoot.bottomLeftNode = createRandomNode(newRoot.topLeftX, newRoot.topLeftY+q.Height, q.Width, q.Height)
	}
	if newRoot.bottomRightNode == nil {
		newRoot.bottomRightNode = createRandomNode(newRoot.topLeftX+q.Width, newRoot.topLeftY+q.Height, q.Width, q.Height)
	}

	q.root = newRoot
	q.Width = newWidth
	q.Height = newHeight
	fmt.Println(q.Width, q.Height)
}

// createRandomNode crée un nœud feuille avec des valeurs aléatoires.
func createRandomNode(x, y, width, height int) *node {

	randomValue := rand.Intn(5)
	return &node{
		topLeftX:        x,
		topLeftY:        y,
		Width:           width,
		Height:          height,
		content:         randomValue,
		isLeaf:          true,
		topLeftNode:     nil,
		topRightNode:    nil,
		bottomLeftNode:  nil,
		bottomRightNode: nil,
	}
}
