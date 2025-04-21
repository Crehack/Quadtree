package floor

import (
	"log"
	"os"
	"strconv"
	"strings"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - Content : partie du terrain qui doit être affichée à l'écran
//   - FullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	Content         [][]int
	FullContent     [][]int
	QuadtreeContent quadtree.Quadtree

	Random [][]int

	animationFrameCount int // Compteur du nombre de cycles écoulés depuis la dernière transition d'animation.
	// Ce compteur permet de mesurer le temps avant de passer à la frame d'animation suivante.
	// Il est incrémenté à chaque cycle de jeu (appel à Update) et réinitialisé une fois qu'il atteint
	// le seuil défini par numFramesPerWaterAnimImage.

	animationFrame int // Index de la frame d'animation actuellement affichée.
	// Cet index est utilisé pour sélectionner la frame correspondante dans le sprite sheet.
	// L'index est cyclique et revient à 0 lorsque toutes les frames disponibles (totalWaterFrames) ont été affichées.
}

// types d'affichage du terrain disponibles
const (
	GridFloor int = iota
	FromFileFloor
	QuadTreeFloor
)

const (
	numFramesPerWaterAnimImage = 8 // Nombre de cycles nécessaires avant de passer à la prochaine frame
	totalWaterFrames           = 3 // Nombre total de frames d'animation pour l'eau
)

// GetHeight retourne la hauteur (en cases) du terrain
// à partir du tableau FullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeight() (height int) {
	return len(f.FullContent)
}

// GetWidth retourne la largeur (en cases) du terrain
// à partir du tableau FullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidth() (width int) {
	if len(f.FullContent) > 0 {
		width = len(f.FullContent[0])
	}
	return
}

func (f Floor) Savefloor() {
	floorContent := f.FullContent

	// os.O_CREATE : Crée le fichier s'il n'existe pas
	// os.O_WRONLY : Ouvre le fichier en mode écriture uniquement
	// os.O_TRUNC  : Vide le contenu du fichier s'il existe déjà

	// 4 => lecture
	// 2 => écriture
	// 1 => exécution
	// 0  6    6   6
	// |  |    |   |
	// |  |    |   └── Permissions des "autres" (utilisateurs non liés) --- 4 + 2
	// |  |    └────── Permissions du "groupe" (même groupe que le propriétaire) --- 4 + 2
	// |  └─────────── Permissions du "propriétaire" (créateur du fichier) --- 4 + 2
	// └────────────── Type de fichier (généralement 0 pour les fichiers normaux) --- 0

	file, err := os.OpenFile("../floor-files/saved", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, y := range floorContent {

		// Convertit chaque ligne en une chaîne de caractères avec des espaces
		line := make([]string, len(y))
		for x, val := range y {
			line[x] = strconv.Itoa(val)
		}

		// Écrit la ligne dans le fichier suivie d'un retour à la ligne
		file.WriteString(strings.Join(line, "") + "\n")
	}
}

func Terre(f Floor, b int, a int) (x int, y int) {
	floorContent := f.FullContent

	if floorContent[a][b] == 4 {
		for i := 0; i < len(floorContent); i++ {

			for j := 0; j < len(floorContent[i]); j++ {

				if f.FullContent[i][j] != 4 {

					return j, i
				}
			}
		}
	}
	return a, b

}
