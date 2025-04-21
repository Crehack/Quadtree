package floor

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche la partie visible du sol sur l'écran en fonction de la position de la caméra.
func (f Floor) Draw(screen *ebiten.Image, cameraX, cameraY int) {
	mapWidth := configuration.Global.MapWidth   // Largeur de la carte
	mapHeight := configuration.Global.MapHeight // Hauteur de la carte
	tileSize := configuration.Global.TileSize

	// Parcours des lignes et colonnes de la section visible du sol.
	for y := 0; y < len(f.Content); y++ {
		for x := 0; x < len(f.Content[y]); x++ {
			mapX := (cameraX + x) - configuration.Global.NumTileX/2
			mapY := (cameraY + y) - configuration.Global.NumTileY/2

			// Gestion des coordonnées pour un monde circulaire.
			if configuration.Global.RoundWorld && !configuration.Global.CameraBlock {
				mapX = (mapX + mapWidth) % mapWidth
				mapY = (mapY + mapHeight) % mapHeight
			}

			// Si la position est en dehors des limites de la carte, on passe à la prochaine tuile.
			if mapX < 0 || mapX >= mapWidth || mapY < 0 || mapY >= mapHeight {
				continue
			}

			// Récupération de la tuile correspondante.
			tile := f.FullContent[mapY][mapX]
			if tile == -1 { 
				continue
			}

			// Préparation de l'affichage.
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))

			// Gestion des différents types de tuiles.
			if !configuration.Global.MovingWater { // Mode sans eau en mouvement.
				shiftX := tile * tileSize
				screen.DrawImage(
					assets.FloorImage.SubImage(image.Rect(shiftX, 0, shiftX+tileSize, tileSize)).(*ebiten.Image), op,
				)
			} else if tile != 4 { // Tuiles non aquatiques en mode eau en mouvement.
				shiftX := tile * tileSize
				screen.DrawImage(
					assets.FloorImage.SubImage(image.Rect(shiftX, 0, shiftX+tileSize, tileSize)).(*ebiten.Image), op,
				)
			} else { // Tuiles aquatiques avec animation.
				shiftX := (2 + f.animationFrame) * tileSize
				shiftY := 2 * tileSize
				screen.DrawImage(
					assets.WaterImage.SubImage(image.Rect(shiftX, shiftY, shiftX+tileSize, shiftY+tileSize)).(*ebiten.Image), op,
				)
			}
		}
	}
}
