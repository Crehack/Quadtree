package camera

import (
	"math"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY int) {

	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(characterPosX, characterPosY)

	case Cinematic:
		c.updateCinematic(float64(characterPosX), float64(characterPosY))

	}
	if configuration.Global.CameraBlock {
		c.Camerabloc()
	}

}

// updateStatic est la mise-à-jour d'une caméra qui reste
// toujours à la position (0,0). Cette fonction ne fait donc
// rien.
func (c *Camera) updateStatic() {}

// updateFollowCharacter est la mise-à-jour d'une caméra qui
// suit toujours le personnage. Elle prend en paramètres deux
// entiers qui indiquent les coordonnées du personnage et place
// la caméra au même endroit.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int) {
	c.X = float64(characterPosX)
	c.Y = float64(characterPosY)
}

// updateCinematic ajuste la position de la caméra pour suivre le personnage de manière fluide.
// La caméra ne bouge que si le personnage s'éloigne d'une certaine distance,
// et le mouvement est interpolé avec un facteur de lissage.
func (c *Camera) updateCinematic(characterX, characterY float64) {
	// Facteur de lissage : contrôle la fluidité du mouvement de la caméra.
	// Une valeur plus basse rend le mouvement plus rapide et direct,
	// tandis qu'une valeur plus élevée le rend plus progressif.
	smoothingFactor := 0.1

	// Seuil de sensibilité pour déclencher un mouvement de la caméra.
	// Si le personnage ne s'éloigne pas suffisamment de la caméra (en fonction de ce seuil),
	// la caméra ne bougera pas. Cela évite des mouvements inutiles pour des petits déplacements.
	moveThreshold := 2.0

	// Calcul de la distance entre la position actuelle de la caméra
	// et la position du personnage sur l'axe X.
	deltaX := characterX - c.X

	// Calcul de la distance entre la position actuelle de la caméra
	// et la position du personnage sur l'axe Y.
	deltaY := characterY - c.Y

	// Si la distance sur l'axe X dépasse le seuil, on met à jour la position de la caméra.
	// Le mouvement est lissé grâce au facteur de lissage, pour éviter des sauts brusques.
	if math.Abs(deltaX) > moveThreshold {
		// Mise à jour de la position X de la caméra avec interpolation.
		c.X += deltaX * smoothingFactor
	}

	// Si la distance sur l'axe Y dépasse le seuil, on met à jour la position de la caméra.
	// Le mouvement est également lissé de manière similaire à l'axe X.
	if math.Abs(deltaY) > moveThreshold {
		// Mise à jour de la position Y de la caméra avec interpolation.
		c.Y += deltaY * smoothingFactor
	}
}

func (c *Camera) Camerabloc() {

	halfScreenWidth := float64(configuration.Global.NumTileX) / 2
	halfScreenHeight := float64(configuration.Global.NumTileY) / 2

	// Limites horizontales
	if c.X-halfScreenWidth < 0 {
		c.X = halfScreenWidth
	} else if c.X+halfScreenWidth > float64(configuration.Global.MapWidth) {
		c.X = float64(configuration.Global.MapWidth) - halfScreenWidth
	}

	// Limites verticales
	if c.Y-halfScreenHeight < 0 {
		c.Y = halfScreenHeight
	} else if c.Y+halfScreenHeight > float64(configuration.Global.MapHeight) {
		c.Y = float64(configuration.Global.MapHeight) - halfScreenHeight
	}
}
