package character

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw permet d'afficher le personnage dans une *ebiten.Image
// (en pratique, celle qui représente la fenêtre de jeu) en
// fonction des caractéristiques du personnage (position, orientation,
// étape d'animation, etc.) et de la position de la caméra (le personnage
// est affiché relativement à la caméra).
func (c Character) Draw(screen *ebiten.Image, camX, camY int) {
	// Calcul des décalages en fonction de l'orientation et de la progression du mouvement
	xShift, yShift := 0, 0
	switch c.orientation {
	case orientedDown:
		yShift = int(c.shift)
	case orientedUp:
		yShift = int(-c.shift)
	case orientedLeft:
		xShift = int(-c.shift)
	case orientedRight:
		xShift = int(c.shift)
	}

	// Position absolue pour l'affichage, relative à la caméra
	xTileForDisplay := c.X - camX + configuration.Global.ScreenCenterTileX
	yTileForDisplay := c.Y - camY + configuration.Global.ScreenCenterTileY
	xPos := xTileForDisplay*configuration.Global.TileSize + xShift
	yPos := yTileForDisplay*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift

	// Configuration de l'animation
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	// Déplacement dans la sprite sheet en fonction de l'animation et de l'orientation
	shiftX := configuration.Global.TileSize
	if c.moving {
		shiftX += c.animationStep * configuration.Global.TileSize
	}
	shiftY := c.orientation * configuration.Global.TileSize

	// Extraction de la bonne image et affichage
	characterSprite := assets.CharacterImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image)

	screen.DrawImage(characterSprite, op)
}
