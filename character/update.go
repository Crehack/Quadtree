package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool, f floor.Floor) {

	if !c.moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.orientation = orientedRight
			if (!blocking[1] && (!configuration.Global.RoundWorld || configuration.Global.CameraBlock)) ||
				(configuration.Global.RoundWorld && !configuration.Global.CameraBlock && c.canMoveTo(c.X+1, c.Y, f)) {
				c.xInc = 1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.orientation = orientedLeft
			if (!blocking[3] && (!configuration.Global.RoundWorld || configuration.Global.CameraBlock)) ||
				(configuration.Global.RoundWorld && !configuration.Global.CameraBlock && c.canMoveTo(c.X-1, c.Y, f)) {
				c.xInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.orientation = orientedUp
			if (!blocking[0] && (!configuration.Global.RoundWorld || configuration.Global.CameraBlock)) ||
				(configuration.Global.RoundWorld && !configuration.Global.CameraBlock && c.canMoveTo(c.X, c.Y-1, f)) {
				c.yInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.orientation = orientedDown
			if (!blocking[2] && (!configuration.Global.RoundWorld || configuration.Global.CameraBlock)) ||
				(configuration.Global.RoundWorld && !configuration.Global.CameraBlock && c.canMoveTo(c.X, c.Y+1, f)) {
				c.yInc = 1
				c.moving = true
			}
		}
	} else {
		c.animationFrameCount++
		if configuration.Global.CameraPrecise {
			c.shift += 0.75
		}

		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			if !configuration.Global.CameraPrecise {
				c.shift += float64(shiftStep)
			}

			c.animationStep = -c.animationStep
			if c.shift > float64(configuration.Global.TileSize-shiftStep) {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0

				if configuration.Global.RoundWorld && !configuration.Global.CameraBlock {
					c.RoundWorld()
				}
			}
		}
	}
}

func (c *Character) RoundWorld() {
	c.X = (c.X + configuration.Global.MapWidth) % configuration.Global.MapWidth
	c.Y = (c.Y + configuration.Global.MapHeight) % configuration.Global.MapHeight
}

func (c *Character) canMoveTo(newX, newY int, f floor.Floor) bool {
	mapWidth := configuration.Global.MapWidth
	mapHeight := configuration.Global.MapHeight

	X := (newX + mapWidth) % mapWidth
	Y := (newY + mapHeight) % mapHeight

	// Vérification si la case cible est bloquante
	tile := f.FullContent[Y][X]
	return tile != -1 && (!configuration.Global.BlockingWater || tile != 4)
}
