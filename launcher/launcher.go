package launcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/MajestikButter/gomc-launcher/game"
)

type Launcher struct {
	*State
	*Settings
	Games map[string]*game.Game
}

func RunScript(script string) error {
	cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", script)
	return cmd.Start()
}

// func (l *Launcher) Launch(family string) {
// 	l.runWin(fmt.Sprintf("shell:appsFolder\\%s!App", family))
// }

func (l *Launcher) GetGame(name string) (*game.Game, bool) {
	game, ok := l.Games[name]
	return game, ok
}

func (l *Launcher) NewGame(name string) (string, *game.Game) {
	g := game.NewGame("", "", "")
	n := l.ValidifyName(name)
	l.Games[n] = g
	return n, g
}

func (l *Launcher) RenameGame(oldName, newName string) error {
	game, exists := l.GetGame(oldName)
	if !exists {
		return errors.New("game does not exist in the launcher")
	}

	delete(l.Games, oldName)
	l.Games[l.ValidifyName(newName)] = game
	return nil
}

var NAME_REGEX, _ = regexp.Compile(`\(\d+\)$`)

func (l *Launcher) ValidifyName(name string) string {
	_, exists := l.GetGame(name)
	if exists {
		numStr := NAME_REGEX.FindString(name)
		if numStr != "" {
			num, _ := strconv.Atoi(numStr[1 : len(numStr)-1])
			name = NAME_REGEX.ReplaceAllString(name, fmt.Sprintf("(%v)", num+1))
		} else {
			name = name + "(1)"
		}
		return l.ValidifyName(name)
	} else {
		return name
	}
}

// var lastLoad = time.Now()

func (l *Launcher) Load() error {
	err := l.Settings.Load()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(GAMES_FILE)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &l.Games)
	if err != nil {
		return err
	}

	return l.State.Load()
}

func (l *Launcher) Save() error {
	err := l.Settings.Save()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(l.Games, "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile(GAMES_FILE, data, FILE_MODE)

	return l.State.Save()
}
