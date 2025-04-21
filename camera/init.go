package camera

// Init met en place une caméra.
func (c *Camera) Init(characterPosX, characterPosY float64) {
	c.X = characterPosX
	c.Y = characterPosY
}
