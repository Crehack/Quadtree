package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY

	if configuration.Global.RandomFloor {
		f.updateQuadtreeFloor(topLeftX, topLeftY)

	}

	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}

	// Incrémente le compteur de frames d'animation
	f.animationFrameCount++

	// Si le compteur atteint le nombre de frames nécessaires pour changer d'image
	if f.animationFrameCount >= numFramesPerWaterAnimImage {
		// Réinitialise le compteur à 0 pour le prochain cycle
		f.animationFrameCount = 0

		// Passe à l'image suivante de l'animation
		// Le compteur d'animation "f.animationFrame" est incrémenté
		// et cyclé grâce à l'opérateur modulo (%) pour revenir à 0
		// après avoir atteint la dernière frame (totalWaterFrames - 1)
		f.animationFrame = (f.animationFrame + 1) % totalWaterFrames
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.Content); y++ {
		for x := 0; x < len(f.Content[y]); x++ {
			absX := topLeftX
			if absX < 0 {
				absX = -absX
			}
			absY := topLeftY
			if absY < 0 {
				absY = -absY
			}
			f.Content[y][x] = ((x + absX%2) + (y + absY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
//
// la version actuelle recopie FullContent dans Content, ce qui n'est pas
// le comportement attendu dans le rendu du projet
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
	// for y := 0; y < len(f.Content); y++ {
	// 	for x := 0; x < len(f.Content[y]); x++ {
	// 		if y < len(f.FullContent) && x < len(f.FullContent[y]) {
	// 			f.Content[y][x] = f.FullContent[y][x]
	// 		} else {
	// 			f.Content[y][x] = -1
	// 		}
	// 	}
	// }

	for y := range f.Content {
		if 0 > y {
			continue
		}

		for x := range f.Content[y] {
			if 0 > x {
				continue
			}

			var real_y int = y + topLeftY
			var real_x int = x + topLeftX

			if (real_y >= 0 && real_y < len(f.FullContent)) && (real_x >= 0 && real_x < len(f.FullContent[real_y])) {
				f.Content[y][x] = f.FullContent[real_y][real_x]
			} else {
				f.Content[y][x] = -1
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {

	f.QuadtreeContent.GetContent(topLeftX, topLeftY, f.Content)
}
