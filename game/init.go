package game

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	g.character.Init(g.floor.GetWidth(), g.floor.GetHeight(), g.floor)
	g.camera.Init(float64(g.character.X), float64(g.character.Y))
}
