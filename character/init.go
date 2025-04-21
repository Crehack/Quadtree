package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init(floorWidth, floorHeight int , f floor.Floor) {
	c.animationStep = 1

	c.X = floorWidth /2
	c.Y = floorHeight /2

	if configuration.Global.BlockingWater {
		c.X , c.Y = floor.Terre(f,  c.X , c.Y)
	}
	

}

