package quadtree

import (
	"testing"
)

// Vérifie la fonction MakeFromArray dans le cas où le
// terrain est un carré dont le côté a un nombre de cases
// qui est une puissance de 2. C'est la cas le plus simple
// pour créer le quadtree.
//
// Les réductions dans le cas où toute la partie du terrain
// représentée par un nœud n'a qu'un seul type de sol n'ont
// pas besoin d'être effectuées pour passer ces tests.
func TestMakeFromArrayPowerOfTwo(t *testing.T) {

	arrays := [][][]int{
		{{1}},
		{{0, 1}, {2, 3}},
		{{0, 1, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}},
		{
			{0, 1, 2, 3, 0, 1, 2, 3},
			{0, 1, 2, 3, 0, 1, 2, 3},
			{0, 1, 2, 3, 0, 1, 2, 3},
			{0, 1, 2, 3, 0, 1, 2, 3},
			{0, 1, 2, 3, 0, 1, 2, 3},
			{0, 1, 2, 3, 0, 1, 2, 3},
			{0, 1, 2, 3, 0, 1, 2, 3},
			{0, 1, 2, 3, 0, 1, 2, 3},
		},
		{{0, 0, 1, 1}, {0, 0, 0, 1}, {1, 1, 0, 0}, {1, 1, 1, 0}},
	}

	trees := []Quadtree{
		{Width: 1, Height: 1,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
		},
		{Width: 2, Height: 2,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2,
				topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
				bottomRightNode: &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
			},
		},
		{Width: 4, Height: 4,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 4, Height: 4,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
				},
				topRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
					topRightNode:    &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 3},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 1, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
				},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 2},
					topRightNode:    &node{topLeftX: 3, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 3},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 3},
				},
			},
		},
		{Width: 8, Height: 8,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 8, Height: 8,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 4, Height: 4,
					topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					topRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 3, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
					bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 0, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 1, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 0, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 1, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 3, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 2, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
				},
				topRightNode: &node{topLeftX: 4, topLeftY: 0, Width: 4, Height: 4,
					topLeftNode: &node{topLeftX: 4, topLeftY: 0, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 5, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					topRightNode: &node{topLeftX: 6, topLeftY: 0, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 6, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 7, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 6, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 7, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
					bottomLeftNode: &node{topLeftX: 4, topLeftY: 2, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 5, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					bottomRightNode: &node{topLeftX: 6, topLeftY: 2, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 6, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 7, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 6, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 7, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 4, Width: 4, Height: 4,
					topLeftNode: &node{topLeftX: 0, topLeftY: 4, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 0, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 1, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 0, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 1, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					topRightNode: &node{topLeftX: 2, topLeftY: 4, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 2, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 3, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 2, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 3, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
					bottomLeftNode: &node{topLeftX: 0, topLeftY: 6, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 0, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 1, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 0, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 1, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					bottomRightNode: &node{topLeftX: 2, topLeftY: 6, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 2, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 3, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 2, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 3, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
				},
				bottomRightNode: &node{topLeftX: 4, topLeftY: 4, Width: 4, Height: 4,
					topLeftNode: &node{topLeftX: 4, topLeftY: 4, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 5, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					topRightNode: &node{topLeftX: 6, topLeftY: 4, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 6, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 7, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 6, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 7, topLeftY: 5, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
					bottomLeftNode: &node{topLeftX: 4, topLeftY: 6, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 5, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					bottomRightNode: &node{topLeftX: 6, topLeftY: 6, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 6, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 2},
						topRightNode:    &node{topLeftX: 7, topLeftY: 6, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 6, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 7, topLeftY: 7, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
				},
			},
		},
		{Width: 4, Height: 4,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 4, Height: 4,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2, isLeaf: true, content: 0},
				topRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					topRightNode:    &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 2, Height: 2, isLeaf: true, content: 1},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 3, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 0},
				},
			},
		},
	}

	for i, a := range arrays {
		fullContentArray := a
		fullContentQuadtree := trees[i]

		madeQuadtree := MakeFromArray(fullContentArray)

		if !equal(fullContentQuadtree, madeQuadtree, true) {
			fullContentAsText := fullContentQuadtree.GetAsText("attendu")
			madeAsText := madeQuadtree.GetAsText("obtenu par MakeFromArray")
			t.Fatalf("Le résultat attendu pour le tableau %v était ce quadtree :\n%s\net le résultat obtenu par MakeFromArray est :\n%s\nMême en ignorant les réductions ces quadtrees ne sont pas similaires.", fullContentArray, fullContentAsText, madeAsText)
		}
	}
}

// Vérifie la fonction MakeFromArray dans le cas général où
// le terrain est un rectangle quelconque.
//
// Les réductions dans le cas où toute la partie du terrain
// représentée par un nœud n'a qu'un seul type de sol n'ont
// pas besoin d'être effectuées pour passer ces tests.
func TestMakeFromArrayAnyRectangle(t *testing.T) {

	arrays := [][][]int{
		{{0, 1, 2}, {0, 1, 2}, {0, 1, 2}},
		{{0, 1}},
		{{0}, {1}},
		{{0, 1, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}},
		{
			{0, 1, 2, 3},
			{0, 1, 2, 3},
			{0, 1, 2, 3},
			{0, 1, 2, 3},
			{0, 1, 2, 3},
		},
	}

	trees := []Quadtree{
		{Width: 3, Height: 3,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 3, Height: 3,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
				topRightNode: &node{topLeftX: 1, topLeftY: 0, Width: 2, Height: 1,
					topLeftNode:     &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
					bottomLeftNode:  &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 2, isLeaf: true, content: 0},
				bottomRightNode: &node{topLeftX: 1, topLeftY: 1, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
					topRightNode:    &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
					bottomLeftNode:  &node{topLeftX: 1, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 2},
				},
			},
		},
		{Width: 2, Height: 1,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 1,
				topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
				bottomRightNode: &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
			},
		},
		{Width: 1, Height: 2,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 2,
				topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 0, Height: 1, isLeaf: true, content: 0},
				topRightNode:    &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, Width: 0, Height: 1, isLeaf: true, content: 0},
				bottomRightNode: &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
			},
		},
		{Width: 4, Height: 3,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 4, Height: 3,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 1,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
				},
				topRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 2, Height: 1,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 3},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 1, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
				},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 1, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
					topRightNode:    &node{topLeftX: 3, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 3},
				},
			},
		},
		{Width: 4, Height: 5,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 4, Height: 5,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
				},
				topRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
					topRightNode:    &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 3},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 2, Height: 3,
					topLeftNode:    &node{topLeftX: 0, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
					topRightNode:   &node{topLeftX: 1, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode: &node{topLeftX: 0, topLeftY: 3, Width: 1, Height: 2, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 3, Width: 1, Height: 2,
						topLeftNode:     &node{topLeftX: 1, topLeftY: 3, Width: 0, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 1, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 1, topLeftY: 4, Width: 0, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 1, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
				},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 2, Height: 3,
					topLeftNode:  &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 2},
					topRightNode: &node{topLeftX: 3, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 3},
					bottomLeftNode: &node{topLeftX: 2, topLeftY: 3, Width: 1, Height: 2,
						topLeftNode:     &node{topLeftX: 2, topLeftY: 3, Width: 0, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 2, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomLeftNode:  &node{topLeftX: 2, topLeftY: 4, Width: 0, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 2, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 2},
					},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 2,
						topLeftNode:     &node{topLeftX: 3, topLeftY: 3, Width: 0, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 3, topLeftY: 4, Width: 0, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 3, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
				},
			},
		},
	}

	for i, a := range arrays {
		fullContentArray := a
		fullContentQuadtree := trees[i]

		madeQuadtree := MakeFromArray(fullContentArray)

		if !equal(fullContentQuadtree, madeQuadtree, true) {
			fullContentAsText := fullContentQuadtree.GetAsText("attendu")
			madeAsText := madeQuadtree.GetAsText("obtenu par MakeFromArray")
			t.Fatalf("Le résultat attendu pour le tableau %v était ce quadtree :\n%s\net le résultat obtenu par MakeFromArray est :\n%s\nMême en ignorant les réductions ces quadtrees ne sont pas similaires.", fullContentArray, fullContentAsText, madeAsText)
		}
	}

}

// Vérifie la fonction MakeFromArray dans le cas général où
// le terrain est un rectangle quelconque.
//
// Les réductions dans le cas où toute la partie du terrain
// représentée par un nœud n'a qu'un seul type de sol doivent
// être effectuées pour passer ces tests.
func TestMakeFromArrayReduction(t *testing.T) {

	arrays := [][][]int{
		{{0, 0, 1, 1}, {0, 0, 0, 1}, {1, 1, 0, 0}, {1, 1, 1, 0}},
		{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
		},
		{
			{1, 1, 1, 2, 2, 2, 2},
			{1, 1, 1, 2, 2, 2, 2},
			{1, 1, 1, 2, 2, 2, 2},
			{3, 3, 3, 1, 1, 1, 1},
			{3, 3, 3, 1, 1, 1, 1},
			{3, 3, 3, 1, 1, 1, 1},
			{3, 3, 3, 1, 1, 1, 1},
		},
		{
			{1, 1, 1, 3, 3, 2, 2},
			{1, 1, 1, 2, 2, 4, 4},
			{1, 1, 1, 2, 2, 4, 4},
			{3, 3, 3, 1, 1, 1, 1},
			{3, 3, 3, 1, 1, 1, 1},
			{3, 3, 3, 1, 1, 1, 1},
			{3, 3, 3, 1, 1, 1, 1},
		},
	}

	trees := []Quadtree{
		{Width: 4, Height: 4,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 4, Height: 4,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2, isLeaf: true, content: 0},
				topRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					topRightNode:    &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 2, Height: 2, isLeaf: true, content: 1},
				bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 2, Height: 2,
					topLeftNode:     &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
					topRightNode:    &node{topLeftX: 3, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 0},
					bottomLeftNode:  &node{topLeftX: 2, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomRightNode: &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 0},
				},
			},
		},
		{Width: 8, Height: 8,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 8, Height: 8, isLeaf: true, content: 0},
		},
		{Width: 8, Height: 8,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 8, Height: 8, isLeaf: true, content: 1},
		},
		{Width: 7, Height: 7,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 7, Height: 7, isLeaf: true, content: 1},
		},
		{Width: 7, Height: 7,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 7, Height: 7,
				topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 3, Height: 3, isLeaf: true, content: 1},
				topRightNode:    &node{topLeftX: 3, topLeftY: 0, Width: 4, Height: 3, isLeaf: true, content: 2},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 3, Width: 3, Height: 4, isLeaf: true, content: 3},
				bottomRightNode: &node{topLeftX: 3, topLeftY: 3, Width: 4, Height: 4, isLeaf: true, content: 1},
			},
		},
		{Width: 7, Height: 7,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 7, Height: 7,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 3, Height: 3, isLeaf: true, content: 1},
				topRightNode: &node{topLeftX: 3, topLeftY: 0, Width: 4, Height: 3,
					topLeftNode:     &node{topLeftX: 3, topLeftY: 0, Width: 2, Height: 1, isLeaf: true, content: 3},
					topRightNode:    &node{topLeftX: 5, topLeftY: 0, Width: 2, Height: 1, isLeaf: true, content: 2},
					bottomLeftNode:  &node{topLeftX: 3, topLeftY: 1, Width: 2, Height: 2, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 5, topLeftY: 1, Width: 2, Height: 2, isLeaf: true, content: 4},
				},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 3, Width: 3, Height: 4, isLeaf: true, content: 3},
				bottomRightNode: &node{topLeftX: 3, topLeftY: 3, Width: 4, Height: 4, isLeaf: true, content: 1},
			},
		},
	}

	for i, a := range arrays {
		fullContentArray := a
		fullContentQuadtree := trees[i]

		madeQuadtree := MakeFromArray(fullContentArray)

		if !equal(fullContentQuadtree, madeQuadtree, false) {
			fullContentAsText := fullContentQuadtree.GetAsText("attendu")
			madeAsText := madeQuadtree.GetAsText("obtenu par MakeFromArray")
			t.Fatalf("Le résultat attendu pour le tableau %v était ce quadtree :\n%s\net le résultat obtenu par MakeFromArray est :\n%s\nMême en ignorant les réductions ces quadtrees ne sont pas similaires.", fullContentArray, fullContentAsText, madeAsText)
		}
	}
}

// Indique si deux quadtrees sont similaires, éventuellement
// en rentrand dans les parties réduites.
func equal(q1, q2 Quadtree, expendReductions bool) bool {

	if q1.Width != q2.Width || q1.Height != q2.Height {
		return false
	}

	return equalNode(q1.root, q2.root, expendReductions)
}

// Indique si deux nœuds sont similaires, éventuellement
// en rentrant dans les parties réduites.
func equalNode(n1, n2 *node, expendReductions bool) bool {

	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.topLeftX != n2.topLeftX ||
		n1.topLeftY != n2.topLeftY ||
		n1.Width != n2.Width ||
		n1.Height != n2.Height {
		return false
	}

	if n1.isLeaf && n2.isLeaf {
		return n1.content == n2.content || n1.Width == 0 || n1.Height == 0
	}

	if n1.isLeaf || n2.isLeaf {
		if !expendReductions {
			return false
		}

		if n1.isLeaf {
			return equalToLeaf(n2, n1.content)
		}

		return equalToLeaf(n1, n2.content)
	}

	return equalNode(n1.topLeftNode, n2.topLeftNode, expendReductions) &&
		equalNode(n1.topRightNode, n2.topRightNode, expendReductions) &&
		equalNode(n1.bottomLeftNode, n2.bottomLeftNode, expendReductions) &&
		equalNode(n1.bottomRightNode, n2.bottomRightNode, expendReductions)
}

// Étant donné un type de sol on vérifie si un nœud pourrait être
// réduit à une feuille ayant ce type de sol
func equalToLeaf(n *node, content int) bool {

	if n == nil {
		return false
	}

	if n.isLeaf {
		return n.content == content || n.Width == 0 || n.Height == 0
	}

	return equalToLeaf(n.topLeftNode, content) &&
		equalToLeaf(n.topRightNode, content) &&
		equalToLeaf(n.bottomLeftNode, content) &&
		equalToLeaf(n.bottomRightNode, content)

}
