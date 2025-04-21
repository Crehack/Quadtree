package teleporteur

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (tp Tp) DrawTp(screen *ebiten.Image, cX, cY int) {
	// Téléporteur 1
	if tp.Tp1_set {
		x1pos := (tp.X1 - cX + configuration.Global.NumTileX/2) * configuration.Global.TileSize
		y1pos := (tp.Y1 - cY + configuration.Global.NumTileY/2) * configuration.Global.TileSize


		frame := tp.FrameCounter / 15 % 2 
		shiftX := frame * configuration.Global.TileSize
		shiftY := 13 * configuration.Global.TileSize

		

		op1 := &ebiten.DrawImageOptions{}
		op1.GeoM.Translate(float64(x1pos), float64(y1pos))
		screen.DrawImage(assets.WaterImage.SubImage(
			image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
		).(*ebiten.Image), op1)
	}

	// Téléporteur 2
	if tp.Tp2_set {
		x2pos := (tp.X2 - cX + configuration.Global.NumTileX/2) * configuration.Global.TileSize
		y2pos := (tp.Y2 - cY + configuration.Global.NumTileY/2) * configuration.Global.TileSize

		frame := tp.FrameCounter / 15 % 2 
		shiftX := frame * configuration.Global.TileSize
		shiftY := 13 * configuration.Global.TileSize

		

		op1 := &ebiten.DrawImageOptions{}
		op1.GeoM.Translate(float64(x2pos), float64(y2pos))
		screen.DrawImage(assets.WaterImage.SubImage(
			image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
		).(*ebiten.Image), op1)
	}
}


