package quadtree

import "testing"

type requestedWindow struct {
	topLeftX, topLeftY int
	Width, Height      int
	expectedResult     [][]int
}

// Vérife la fonction GetContent avec des quadtrees où les
// réductions dans le cas où toute la partie du terrain
// représentée par un nœud n'a qu'un seul type de sol n'ont
// pas été effectuées.
func TestGetContentNoReduction(t *testing.T) {

	trees := []Quadtree{
		{Width: 2, Height: 2,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2,
				topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
				topRightNode:    &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
				bottomRightNode: &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 4},
			},
		},
		{Width: 6, Height: 5,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 6, Height: 5,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 3, Height: 2,
					topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
					topRightNode: &node{topLeftX: 1, topLeftY: 0, Width: 2, Height: 1,
						topLeftNode:     &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 2, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
					bottomLeftNode: &node{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 1, Width: 2, Height: 1,
						topLeftNode:     &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomRightNode: &node{topLeftX: 2, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 4},
					},
				},
				topRightNode: &node{topLeftX: 3, topLeftY: 0, Width: 3, Height: 2,
					topLeftNode: &node{topLeftX: 3, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 4},
					topRightNode: &node{topLeftX: 4, topLeftY: 0, Width: 2, Height: 1,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 5, topLeftY: 0, Width: 1, Height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 0, Width: 1, Height: 1, isLeaf: true, content: 2},
					},
					bottomLeftNode: &node{topLeftX: 3, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 1},
					bottomRightNode: &node{topLeftX: 4, topLeftY: 1, Width: 2, Height: 1,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 1, Width: 1, Height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 5, topLeftY: 1, Width: 1, Height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 1, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
				},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 3, Height: 3,
					topLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 3},
					topRightNode: &node{topLeftX: 1, topLeftY: 2, Width: 2, Height: 1,
						topLeftNode:     &node{topLeftX: 1, topLeftY: 2, Width: 1, Height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 1, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 4},
						bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					bottomLeftNode: &node{topLeftX: 0, topLeftY: 3, Width: 1, Height: 2,
						topLeftNode:     &node{topLeftX: 0, topLeftY: 3, Width: 0, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 0, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 4},
						bottomLeftNode:  &node{topLeftX: 0, topLeftY: 4, Width: 0, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 0, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 1},
					},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 3, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 1, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
						topRightNode:    &node{topLeftX: 2, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomLeftNode:  &node{topLeftX: 1, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 2},
						bottomRightNode: &node{topLeftX: 2, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 3},
					},
				},
				bottomRightNode: &node{topLeftX: 3, topLeftY: 2, Width: 3, Height: 3,
					topLeftNode: &node{topLeftX: 3, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 2},
					topRightNode: &node{topLeftX: 4, topLeftY: 2, Width: 2, Height: 1,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 2, Width: 1, Height: 0, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 5, topLeftY: 2, Width: 1, Height: 0, isLeaf: true, content: 0},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 4},
					},
					bottomLeftNode: &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 2,
						topLeftNode:     &node{topLeftX: 3, topLeftY: 3, Width: 0, Height: 1, isLeaf: true, content: 0},
						topRightNode:    &node{topLeftX: 3, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 3},
						bottomLeftNode:  &node{topLeftX: 3, topLeftY: 4, Width: 0, Height: 1, isLeaf: true, content: 0},
						bottomRightNode: &node{topLeftX: 3, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 4},
					},
					bottomRightNode: &node{topLeftX: 4, topLeftY: 3, Width: 2, Height: 2,
						topLeftNode:     &node{topLeftX: 4, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 4},
						topRightNode:    &node{topLeftX: 5, topLeftY: 3, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomLeftNode:  &node{topLeftX: 4, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 1},
						bottomRightNode: &node{topLeftX: 5, topLeftY: 4, Width: 1, Height: 1, isLeaf: true, content: 2},
					},
				},
			},
		},
	}

	requests := [][]requestedWindow{
		{
			{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2, expectedResult: [][]int{{1, 2}, {3, 4}}},
			{topLeftX: 0, topLeftY: 0, Width: 1, Height: 1, expectedResult: [][]int{{1}}},
			{topLeftX: 1, topLeftY: 0, Width: 1, Height: 1, expectedResult: [][]int{{2}}},
			{topLeftX: 0, topLeftY: 1, Width: 1, Height: 1, expectedResult: [][]int{{3}}},
			{topLeftX: 1, topLeftY: 1, Width: 1, Height: 1, expectedResult: [][]int{{4}}},
			{topLeftX: 0, topLeftY: 0, Width: 2, Height: 1, expectedResult: [][]int{{1, 2}}},
			{topLeftX: 0, topLeftY: 1, Width: 2, Height: 1, expectedResult: [][]int{{3, 4}}},
			{topLeftX: 0, topLeftY: 0, Width: 1, Height: 2, expectedResult: [][]int{{1}, {3}}},
			{topLeftX: 1, topLeftY: 0, Width: 1, Height: 2, expectedResult: [][]int{{2}, {4}}},
			{topLeftX: 1, topLeftY: 0, Width: 2, Height: 2, expectedResult: [][]int{{2, -1}, {4, -1}}},
			{topLeftX: 0, topLeftY: 1, Width: 2, Height: 2, expectedResult: [][]int{{3, 4}, {-1, -1}}},
			{topLeftX: 1, topLeftY: 1, Width: 2, Height: 2, expectedResult: [][]int{{4, -1}, {-1, -1}}},
			{topLeftX: 2, topLeftY: 2, Width: 2, Height: 2, expectedResult: [][]int{{-1, -1}, {-1, -1}}},
			{topLeftX: -1, topLeftY: 0, Width: 2, Height: 2, expectedResult: [][]int{{-1, 1}, {-1, 3}}},
			{topLeftX: 0, topLeftY: -1, Width: 2, Height: 2, expectedResult: [][]int{{-1, -1}, {1, 2}}},
			{topLeftX: -1, topLeftY: -1, Width: 2, Height: 2, expectedResult: [][]int{{-1, -1}, {-1, 1}}},
			{topLeftX: -2, topLeftY: -2, Width: 2, Height: 2, expectedResult: [][]int{{-1, -1}, {-1, -1}}},
		},
		{
			{topLeftX: 0, topLeftY: 0, Width: 6, Height: 5, expectedResult: [][]int{
				{1, 2, 3, 4, 1, 2},
				{2, 3, 4, 1, 2, 3},
				{3, 4, 1, 2, 3, 4},
				{4, 1, 2, 3, 4, 1},
				{1, 2, 3, 4, 1, 2},
			}},
			{topLeftX: 0, topLeftY: 0, Width: 4, Height: 4, expectedResult: [][]int{
				{1, 2, 3, 4},
				{2, 3, 4, 1},
				{3, 4, 1, 2},
				{4, 1, 2, 3},
			}},
			{topLeftX: 1, topLeftY: 0, Width: 4, Height: 4, expectedResult: [][]int{
				{2, 3, 4, 1},
				{3, 4, 1, 2},
				{4, 1, 2, 3},
				{1, 2, 3, 4},
			}},
			{topLeftX: 2, topLeftY: 0, Width: 4, Height: 4, expectedResult: [][]int{
				{3, 4, 1, 2},
				{4, 1, 2, 3},
				{1, 2, 3, 4},
				{2, 3, 4, 1},
			}},
			{topLeftX: 0, topLeftY: 1, Width: 4, Height: 4, expectedResult: [][]int{
				{2, 3, 4, 1},
				{3, 4, 1, 2},
				{4, 1, 2, 3},
				{1, 2, 3, 4},
			}},
			{topLeftX: 1, topLeftY: 1, Width: 4, Height: 4, expectedResult: [][]int{
				{3, 4, 1, 2},
				{4, 1, 2, 3},
				{1, 2, 3, 4},
				{2, 3, 4, 1},
			}},
			{topLeftX: 2, topLeftY: 1, Width: 4, Height: 4, expectedResult: [][]int{
				{4, 1, 2, 3},
				{1, 2, 3, 4},
				{2, 3, 4, 1},
				{3, 4, 1, 2},
			}},
			{topLeftX: 4, topLeftY: 2, Width: 4, Height: 4, expectedResult: [][]int{
				{3, 4, -1, -1},
				{4, 1, -1, -1},
				{1, 2, -1, -1},
				{-1, -1, -1, -1},
			}},
			{topLeftX: -2, topLeftY: -1, Width: 4, Height: 4, expectedResult: [][]int{
				{-1, -1, -1, -1},
				{-1, -1, 1, 2},
				{-1, -1, 2, 3},
				{-1, -1, 3, 4},
			}},
			{topLeftX: -1, topLeftY: 3, Width: 8, Height: 3, expectedResult: [][]int{
				{-1, 4, 1, 2, 3, 4, 1, -1},
				{-1, 1, 2, 3, 4, 1, 2, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1},
			}},
		},
	}

	for i, rSet := range requests {

		for _, r := range rSet {

			res := make([][]int, r.Height)
			for y := range res {
				res[y] = make([]int, r.Width)
			}

			trees[i].GetContent(r.topLeftX, r.topLeftY, res)

			if !equalContent(r.expectedResult, res) {
				treeAsText := trees[i].GetAsText("fullContent")
				t.Fatalf("Le résultat attendu pour le quadtree \n%s\n avec (topLeftX, topLeftY) = (%d, %d) était le tableau %v mais le résultat retourné par GetContent est %v", treeAsText, r.topLeftX, r.topLeftY, r.expectedResult, res)
			}
		}

	}

}

// Vérife la fonction GetContent avec des quadtrees où les
// réductions dans le cas où toute la partie du terrain
// représentée par un nœud n'a qu'un seul type de sol ont
// été effectuées.
func TestGetContentReduction(t *testing.T) {

	trees := []Quadtree{
		{Width: 8, Height: 8,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 8, Height: 8,
				topLeftNode: &node{topLeftX: 0, topLeftY: 0, Width: 4, Height: 4,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 0, Width: 2, Height: 2, isLeaf: true, content: 1},
					topRightNode:    &node{topLeftX: 2, topLeftY: 0, Width: 2, Height: 2, isLeaf: true, content: 2},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 2, Width: 2, Height: 2, isLeaf: true, content: 2},
					bottomRightNode: &node{topLeftX: 2, topLeftY: 2, Width: 2, Height: 2, isLeaf: true, content: 3},
				},
				topRightNode: &node{topLeftX: 4, topLeftY: 0, Width: 4, Height: 4,
					topLeftNode:     &node{topLeftX: 4, topLeftY: 0, Width: 2, Height: 2, isLeaf: true, content: 3},
					topRightNode:    &node{topLeftX: 6, topLeftY: 0, Width: 2, Height: 2, isLeaf: true, content: 4},
					bottomLeftNode:  &node{topLeftX: 4, topLeftY: 2, Width: 2, Height: 2, isLeaf: true, content: 4},
					bottomRightNode: &node{topLeftX: 6, topLeftY: 2, Width: 2, Height: 2, isLeaf: true, content: 1},
				},
				bottomLeftNode:  &node{topLeftX: 0, topLeftY: 4, Width: 4, Height: 4, isLeaf: true, content: 1},
				bottomRightNode: &node{topLeftX: 4, topLeftY: 4, Width: 4, Height: 4, isLeaf: true, content: 2},
			},
		},
		{Width: 6, Height: 5,
			root: &node{topLeftX: 0, topLeftY: 0, Width: 6, Height: 5,
				topLeftNode:  &node{topLeftX: 0, topLeftY: 0, Width: 3, Height: 2, isLeaf: true, content: 1},
				topRightNode: &node{topLeftX: 3, topLeftY: 0, Width: 3, Height: 2, isLeaf: true, content: 2},
				bottomLeftNode: &node{topLeftX: 0, topLeftY: 2, Width: 3, Height: 3,
					topLeftNode:     &node{topLeftX: 0, topLeftY: 2, Width: 1, Height: 1, isLeaf: true, content: 1},
					topRightNode:    &node{topLeftX: 1, topLeftY: 2, Width: 2, Height: 1, isLeaf: true, content: 1},
					bottomLeftNode:  &node{topLeftX: 0, topLeftY: 3, Width: 1, Height: 2, isLeaf: true, content: 3},
					bottomRightNode: &node{topLeftX: 1, topLeftY: 3, Width: 2, Height: 2, isLeaf: true, content: 3},
				},
				bottomRightNode: &node{topLeftX: 3, topLeftY: 2, Width: 3, Height: 3, isLeaf: true, content: 4},
			},
		},
	}

	requests := [][]requestedWindow{
		{
			{topLeftX: 0, topLeftY: 0, Width: 8, Height: 8, expectedResult: [][]int{
				{1, 1, 2, 2, 3, 3, 4, 4},
				{1, 1, 2, 2, 3, 3, 4, 4},
				{2, 2, 3, 3, 4, 4, 1, 1},
				{2, 2, 3, 3, 4, 4, 1, 1},
				{1, 1, 1, 1, 2, 2, 2, 2},
				{1, 1, 1, 1, 2, 2, 2, 2},
				{1, 1, 1, 1, 2, 2, 2, 2},
				{1, 1, 1, 1, 2, 2, 2, 2},
			}},
			{topLeftX: 0, topLeftY: 0, Width: 5, Height: 5, expectedResult: [][]int{
				{1, 1, 2, 2, 3},
				{1, 1, 2, 2, 3},
				{2, 2, 3, 3, 4},
				{2, 2, 3, 3, 4},
				{1, 1, 1, 1, 2},
			}}, {topLeftX: 3, topLeftY: 3, Width: 5, Height: 5, expectedResult: [][]int{
				{3, 4, 4, 1, 1},
				{1, 2, 2, 2, 2},
				{1, 2, 2, 2, 2},
				{1, 2, 2, 2, 2},
				{1, 2, 2, 2, 2},
			}},
			{topLeftX: 1, topLeftY: 1, Width: 6, Height: 6, expectedResult: [][]int{
				{1, 2, 2, 3, 3, 4},
				{2, 3, 3, 4, 4, 1},
				{2, 3, 3, 4, 4, 1},
				{1, 1, 1, 2, 2, 2},
				{1, 1, 1, 2, 2, 2},
				{1, 1, 1, 2, 2, 2},
			}},
			{topLeftX: 5, topLeftY: 2, Width: 7, Height: 7, expectedResult: [][]int{
				{4, 1, 1, -1, -1, -1, -1},
				{4, 1, 1, -1, -1, -1, -1},
				{2, 2, 2, -1, -1, -1, -1},
				{2, 2, 2, -1, -1, -1, -1},
				{2, 2, 2, -1, -1, -1, -1},
				{2, 2, 2, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1},
			}},
		},
		{
			{topLeftX: 0, topLeftY: 0, Width: 6, Height: 5, expectedResult: [][]int{
				{1, 1, 1, 2, 2, 2},
				{1, 1, 1, 2, 2, 2},
				{1, 1, 1, 4, 4, 4},
				{3, 3, 3, 4, 4, 4},
				{3, 3, 3, 4, 4, 4},
			}},
			{topLeftX: 0, topLeftY: 0, Width: 4, Height: 4, expectedResult: [][]int{
				{1, 1, 1, 2},
				{1, 1, 1, 2},
				{1, 1, 1, 4},
				{3, 3, 3, 4},
			}},
			{topLeftX: 1, topLeftY: 0, Width: 4, Height: 4, expectedResult: [][]int{
				{1, 1, 2, 2},
				{1, 1, 2, 2},
				{1, 1, 4, 4},
				{3, 3, 4, 4},
			}},
			{topLeftX: 2, topLeftY: 0, Width: 4, Height: 4, expectedResult: [][]int{
				{1, 2, 2, 2},
				{1, 2, 2, 2},
				{1, 4, 4, 4},
				{3, 4, 4, 4},
			}},
			{topLeftX: 0, topLeftY: 1, Width: 4, Height: 4, expectedResult: [][]int{
				{1, 1, 1, 2},
				{1, 1, 1, 4},
				{3, 3, 3, 4},
				{3, 3, 3, 4},
			}},
			{topLeftX: 1, topLeftY: 1, Width: 4, Height: 4, expectedResult: [][]int{
				{1, 1, 2, 2},
				{1, 1, 4, 4},
				{3, 3, 4, 4},
				{3, 3, 4, 4},
			}},
			{topLeftX: 2, topLeftY: 1, Width: 4, Height: 4, expectedResult: [][]int{
				{1, 2, 2, 2},
				{1, 4, 4, 4},
				{3, 4, 4, 4},
				{3, 4, 4, 4},
			}},
			{topLeftX: 4, topLeftY: 2, Width: 4, Height: 4, expectedResult: [][]int{
				{4, 4, -1, -1},
				{4, 4, -1, -1},
				{4, 4, -1, -1},
				{-1, -1, -1, -1},
			}},
			{topLeftX: -2, topLeftY: -1, Width: 4, Height: 4, expectedResult: [][]int{
				{-1, -1, -1, -1},
				{-1, -1, 1, 1},
				{-1, -1, 1, 1},
				{-1, -1, 1, 1},
			}},
			{topLeftX: -1, topLeftY: 3, Width: 8, Height: 3, expectedResult: [][]int{
				{-1, 3, 3, 3, 4, 4, 4, -1},
				{-1, 3, 3, 3, 4, 4, 4, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1},
			}},
		},
	}

	for i, rSet := range requests {

		for _, r := range rSet {

			res := make([][]int, r.Height)
			for y := range res {
				res[y] = make([]int, r.Width)
			}

			trees[i].GetContent(r.topLeftX, r.topLeftY, res)

			if !equalContent(r.expectedResult, res) {
				treeAsText := trees[i].GetAsText("fullContent")
				t.Fatalf("Le résultat attendu pour le quadtree \n%s\n avec (topLeftX, topLeftY) = (%d, %d) était le tableau %v mais le résultat retourné par GetContent est %v", treeAsText, r.topLeftX, r.topLeftY, r.expectedResult, res)
			}
		}

	}

}

// Vérifie si deux tableaux de tableaux d'entiers sont
// égaux ou pas.
func equalContent(c1, c2 [][]int) bool {

	if len(c1) != len(c2) {
		return false
	}

	for y := 0; y < len(c1); y++ {

		if len(c1[y]) != len(c2[y]) {
			return false
		}

		for x := 0; x < len(c1[y]); x++ {
			if c1[y][x] != c2[y][x] {
				return false
			}
		}

	}

	return true

}
