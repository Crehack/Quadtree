package configuration

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

// Configuration définit les élèments de la configuration
// du jeu. Pour ajouter un élèment de configuration il
// suffit d'ajouter un champs dans cette structure.
//
// Les champs directement lus dans le fichier de configuration sont :
//   - DebugMode : indique si on est en mode debug ou pas
//   - NumTileX, NumTileY : les nombres de cases affichées à l'écran
//     en largeur et hauteur.
//   - TileSize : la taille en pixels du côté d'une case.
//   - NumCharacterAnimImages : le nombre de d'images de l'animation du
//     personnage.
//   - NumFramePerCharacterAnimImage : le nombre d'appels à update ou
//     de 1/60 de seconde) qui ont lieu entre deux images de l'animation
//     du personnage.
//   - NumTileForDebug : le nombre de cases à ajouter à droite de l'écran
//     pour afficher les informations de debug
//   - CameraMode : le type de caméra à utiliser (0 pour une caméra fixe
//     et 1 pour une caméra qui suit le personnage).
//   - FloorKind : détermine la méthode à utiliser pour afficher le terrain
//     (quadrillage, lecture dans un fichier, quadtree, etc)
//   - FloorFile : le chemin d'un fichier où lire les informations sur le
//     terrain si nécessaire
//   - RandomFloor : détermine si le terrain est généré aléatoirement
//   - FloorWidth , FloorHeight : La taille du terrain aléatoire
//
// Les champs calculés à partir des précédents sont :
//   - ScreenWidth, ScreenHeight : la largeur et la hauteur de l'écran
//     en pixels (hors zone d'affichage pour le debug)
//   - ScreenCenterTileX, ScreenCenterTileY : les coordonnées de la case
//     au centre de l'écran, où sera placé la caméra.
type Configuration struct {
	DebugMode                     bool
	NumTileX, NumTileY            int
	TileSize                      int
	NumCharacterAnimImages        int
	NumFramePerCharacterAnimImage int
	NumTileForDebug               int
	CameraMode                    int
	FloorKind                     int
	FloorFile                     string

	RandomFloor bool
	FloorWidth  int
	FloorHeight int

	Save bool

	BlockingWater bool

	MovingWater bool

	CameraBlock bool

	Zoom bool

	Teleporteur bool

	RoundWorld bool

	CameraPrecise bool

	Infinie bool

	ScreenWidth, ScreenHeight            int `json:"-"`
	ScreenCenterTileX, ScreenCenterTileY int `json:"-"`
	MapWidth, MapHeight                  int `json:"-"`
}

// Global est la variable qui contient la configuration
// du jeu. Sa valeur est fixée à partir de la lecture d'un
// fichier de configuration par la fonction Load. C'est
// cette variable qu'il faut lire (configuration.Global)
// pour accéder à la configuration depuis d'autres paquets.
var Global Configuration

// Load se charge de lire un fichier de configuration, de
// remplir les champs obtenus par simple lecture, puis
// d'appeler la fonction qui remplit les champs calculés.
func Load(configurationFileName string) {
	content, err := os.ReadFile(configurationFileName)
	if err != nil {
		log.Fatal("Error while opening configuration file: ", err)
	}

	err = json.Unmarshal(content, &Global)
	if err != nil {
		log.Fatal("Error while reading configuration file: ", err)
	}

	SetComputedFields()
}

// setComputedFields se charge de remplir les champs calculés
// de la configuration à partir des autres champs.
func SetComputedFields() {
	Global.ScreenWidth = Global.NumTileX * Global.TileSize
	Global.ScreenHeight = Global.NumTileY * Global.TileSize
	Global.ScreenCenterTileX = Global.NumTileX / 2
	Global.ScreenCenterTileY = Global.NumTileY / 2
	Global.MapWidth, Global.MapHeight = size()
}

func size() (width, height int) {
	if Global.RandomFloor {
		width, height = Global.FloorWidth, Global.FloorHeight

	} else {
		var floorContent [][]int
		readFile, err := os.Open(Global.FloorFile)
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
		width, height = len(floorContent[0]), len(floorContent)
	}

	return width, height
}
