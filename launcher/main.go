package launcher

import (
	"os"
	"path"

	"github.com/MajestikButter/gomc-launcher/game"
	"github.com/MajestikButter/gomc-launcher/logger"
)

const FILE_MODE = 0777

var ASSETS_PATH, DATA_PATH, PROFILES_PATH, GAMES_FILE, SETTINGS_FILE, STATE_FILE string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ASSETS_PATH = path.Join(cwd, "assets")
	DATA_PATH = path.Join(cwd, "data")
	PROFILES_PATH = path.Join(DATA_PATH, "profiles")
	GAMES_FILE = path.Join(DATA_PATH, "games.json")
	SETTINGS_FILE = path.Join(DATA_PATH, "settings.json")
	STATE_FILE = path.Join(DATA_PATH, "state.json")
}

func New() *Launcher {
	return &Launcher{
		&State{
			"Default Release",
			"Default Preview",
			&Vec2{600, 400},
		},
		&Settings{false, logger.Path()},
		map[string]*game.Game{
			"Release": game.NewMCRelease(),
			"Preview": game.NewMCPreview(),
		},
	}
}
