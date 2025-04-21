package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/character"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/teleporteur"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).

func (g *Game) Update() error {
	
	
	if configuration.Global.Teleporteur {
		if inpututil.IsKeyJustPressed(ebiten.KeyT) {
			setTeleporter(&g.tp, g.character.X, g.character.Y)
		}
		
		
		if g.tp.Tp1_set && g.tp.Tp2_set && inpututil.IsKeyJustPressed(ebiten.KeySpace){
			teleportCharacter(&g.tp, &g.character)
		}
		if g.tp.Tp1_set || g.tp.Tp2_set {
			g.tp.FrameCounter++
		}
		
	}


	if configuration.Global.Zoom {

		if inpututil.IsKeyJustPressed(ebiten.KeyPageUp) {
			if configuration.Global.NumTileX+2 <= configuration.Global.MapWidth && configuration.Global.NumTileY+2 <= configuration.Global.MapHeight {
				configuration.Global.NumTileX += 2
				configuration.Global.NumTileY += 2

				configuration.Global.ScreenWidth += 2 * configuration.Global.TileSize
				configuration.Global.ScreenHeight += 2 * configuration.Global.TileSize

				configuration.Global.ScreenCenterTileX += 1
				configuration.Global.ScreenCenterTileY += 1

				g.floor.Init()
			}

		}
		if inpututil.IsKeyJustPressed(ebiten.KeyPageDown) {
			if configuration.Global.NumTileX-2 > 1 && configuration.Global.NumTileY-2 > 1 {
				configuration.Global.NumTileX -= 2
				configuration.Global.NumTileY -= 2

				configuration.Global.ScreenWidth -= 2 * configuration.Global.TileSize
				configuration.Global.ScreenHeight -= 2 * configuration.Global.TileSize

				configuration.Global.ScreenCenterTileX -= 1
				configuration.Global.ScreenCenterTileY -= 1

				g.floor.Init()
			}

		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) && configuration.Global.Save {
		fmt.Println("file saved")
		g.floor.Savefloor()
	}



	
	g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, int(g.camera.X), int(g.camera.Y)) , g.floor)
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(int(g.camera.X), int(g.camera.Y))

	
	

	return nil
}


func setTeleporter(tp *teleporteur.Tp, x, y int) {
	if !tp.Tp1_set && !tp.Tp2_set {
		
		tp.Tp1_set = true
		tp.X1, tp.Y1 = x, y
		
	} else if tp.Tp1_set && !tp.Tp2_set {
		
		tp.Tp2_set = true
		tp.X2, tp.Y2 = x, y
		
	} else if tp.Tp1_set && tp.Tp2_set {
		
		tp.X1, tp.Y1 = tp.X2, tp.Y2
		tp.X2, tp.Y2 = x, y
		
	}
}

func teleportCharacter(tp *teleporteur.Tp, character *character.Character ) {
	if character.X == tp.X1 && character.Y == tp.Y1  {
		
		character.X = tp.X2
		character.Y = tp.Y2
		
	} else if character.X == tp.X2 && character.Y == tp.Y2 {
		
		character.X = tp.X1
		character.Y = tp.Y1
		
	}
}