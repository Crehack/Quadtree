package floor

import (
	"bufio"
	"log"
	"math/rand/v2"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.Content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.Content); y++ {
		f.Content[y] = make([]int, configuration.Global.NumTileX)
	}

	if len(f.FullContent) != 0 {
		return
	}

	// Ajout de f.FullContent pour qu'il fonctionne avec les nouvelles extensions ajoutées.
	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.FullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.FullContent = readFloorFromFile(configuration.Global.FloorFile)
		f.QuadtreeContent = quadtree.MakeFromArray(f.FullContent)
	}

	if configuration.Global.RandomFloor {
		f.Random = randomFloorGeneration()
		f.FullContent=f.Random
		f.QuadtreeContent = quadtree.MakeFromArray(f.Random)
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau

// TODO: add error checks for width etc
func readFloorFromFile(fileName string) (floorContent [][]int) {
	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		var new_line []int

		for _, char_floor_type := range fileScanner.Text() {
			int_floor_type, err := strconv.Atoi(string(char_floor_type))
			if err != nil {
				log.Fatal(err)
			}

			new_line = append(new_line, int_floor_type)
		}

		floorContent = append(floorContent, new_line)
	}

	readFile.Close()

	return floorContent
}

func randomFloorGeneration() (floorContent [][]int) {

	for y := 0; y < configuration.Global.FloorHeight; y++ {
		var new_line []int

		for x := 0; x < configuration.Global.FloorWidth; x++ {
			new_line = append(new_line, rand.IntN(5))
		}

		floorContent = append(floorContent, new_line)
	}

	return floorContent
}
