package game

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/MajestikButter/gomc-launcher/logger"
	"github.com/hectane/go-acl"
	"golang.org/x/sys/windows"
)

const FILE_MODE = 0777

type IconStruct struct {
	IconPathR string `json:"iconPath"`
	icon      fyne.Resource
	iconPath  string
}

func (s *IconStruct) Icon() fyne.Resource {
	if s.icon == nil || s.iconPath != s.IconPathR {
		i, err := fyne.LoadResourceFromPath(s.IconPathR)
		if err != nil {
			return theme.AccountIcon()
		}
		s.icon = i
		s.iconPath = s.IconPathR
	}
	return s.icon
}

func (s *IconStruct) IconPath() string {
	return s.IconPathR
}

func (s *IconStruct) SetIconPath(path string) {
	s.IconPathR = path
}

type Profile struct {
	IconStruct
	Path       string            `json:"path"`
	Subfolders map[string]string `json:"subfolders"`
}

type Game struct {
	IconStruct
	Profiles        map[string]*Profile `json:"profiles"`
	LaunchScript    string              `json:"launchScript"`
	Destination     string              `json:"destination"`
	SecurityID      string              `json:"securityID"`
	SelectedProfile string              `json:"selectedProfile"`
}

var lastLoad = time.Now()

func (g *Game) LoadProfile(profile *Profile) {
	if time.Since(lastLoad).Milliseconds() < 300 {
		return
	}
	lastLoad = time.Now()

	sid := g.SecurityID

	logger.Println("Creating main symbolic link")
	g.CreateSymlink(g.Destination, profile.Path, sid)
	logger.Println("Creating subfolder symbolic links (if any)")
	for rel, abs := range profile.Subfolders {
		logger.Println("Subfolder", rel, abs)
		p := path.Join(profile.Path, rel)
		g.CreateSymlink(p, abs, sid)
	}
}

func (g *Game) TryMoveFolder(from string, count int) error {
	if count > 100 {
		return errors.New("attempted to move folder more than 100 times")
	}
	if count > 0 {
		n := fmt.Sprintf("%s.copy_%v", from, count)
		if _, err := os.Stat(n); os.IsNotExist(err) {
			err := os.Rename(from, n)
			if err != nil {
				return err
			}
			logger.Print("Moved '%s' to '%s'\n", from, path.Base(n))
		} else {
			return g.TryMoveFolder(from, count+1)
		}
	} else {
		n := fmt.Sprintf("%s.copy", from)
		if _, err := os.Stat(n); os.IsNotExist(err) {
			err := os.Rename(from, n)
			if err != nil {
				return err
			}
		} else {
			return g.TryMoveFolder(from, count+1)
		}
	}
	return nil
}

func (g *Game) CreateSymlink(from, to, secId string) {
	logger.Printf("Creating symbolic link {from: '%s' to: '%s'}\n", from, to)

	if _, err := os.Stat(from); err == nil {
		logger.Println("Existing directory found")
		if _, err := os.Readlink(from); err == nil {
			logger.Println("Removing existing symbolic link")

			err := os.Remove(from)
			if err != nil {
				panic(err)
			}
		} else {
			logger.Println("Existing directory is not a symbolic link, attempting to move")
			err := g.TryMoveFolder(from, 0)
			if err != nil {
				panic(err)
			}
		}
	}
	if _, err := os.Stat(to); os.IsNotExist(err) {
		os.MkdirAll(to, FILE_MODE)
	}

	logger.Println("Symlinking directories")
	err := os.Symlink(to, from)
	if err != nil {
		panic(err)
	}

	if secId != "" {
		sid, err := windows.StringToSid(secId)
		if err != nil {
			panic(err)
		}

		logger.Println("Applying required minecraft permissions to symbolic link target", to, sid)
		err = acl.Apply(to, true, true, acl.GrantSid(windows.GENERIC_ALL, sid))
		if err != nil {
			panic(err)
		}
	}
	logger.Println("Created symbolic link")
}

func (g *Game) NewProfile(name string, path string) (string, *Profile) {
	profile := &Profile{
		IconStruct{},
		path,
		map[string]string{},
	}
	n := g.ValidifyName(name)
	g.Profiles[n] = profile
	return n, profile
}

func (g *Game) GetProfile(name string) (*Profile, bool) {
	v, ok := g.Profiles[name]
	return v, ok
}

var NAME_REGEX, _ = regexp.Compile(`\(\d+\)$`)

func (g *Game) ValidifyName(name string) string {
	_, exists := g.GetProfile(name)
	if exists {
		numStr := NAME_REGEX.FindString(name)
		if numStr != "" {
			num, _ := strconv.Atoi(numStr[1 : len(numStr)-1])
			name = NAME_REGEX.ReplaceAllString(name, fmt.Sprintf("(%v)", num+1))
		} else {
			name = name + "(1)"
		}
		return g.ValidifyName(name)
	} else {
		return name
	}
}

func (g *Game) RenameProfile(oldName, newName string) error {
	_, exists := g.GetProfile(oldName)
	if !exists {
		return errors.New("profile does not exist in the current game")
	}

	if oldName == g.SelectedProfile {
		g.SelectedProfile = newName
	}

	profile := g.Profiles[oldName]
	delete(g.Profiles, oldName)
	g.Profiles[g.ValidifyName(newName)] = profile
	return nil
}

func (g *Game) Selected() *Profile {
	v, _ := g.GetProfile(g.SelectedProfile)
	return v
}

// Template games

const REL_FAMILY, PRE_FAMILY = "Microsoft.MinecraftUWP_8wekyb3d8bbwe", "Microsoft.MinecraftWindowsBeta_8wekyb3d8bbwe"
const REL_SID, PRE_SID = "S-1-15-2-1958404141-86561845-1752920682-3514627264-368642714-62675701-733520436", "S-1-15-3-424268864-5579737-879501358-346833251-474568803-887069379-4040235476"

var REL_COM_MOJANG = path.Join(os.Getenv("LOCALAPPDATA"), "Packages", REL_FAMILY, "LocalState", "games", "com.mojang")
var PRE_COM_MOJANG = path.Join(os.Getenv("LOCALAPPDATA"), "Packages", PRE_FAMILY, "LocalState", "games", "com.mojang")

func cwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

func NewGame(launchScript, destination, secId string) *Game {
	g := &Game{
		IconStruct{},
		map[string]*Profile{},
		launchScript,
		destination,
		secId,
		"Default",
	}
	g.NewProfile("Default", cwd())
	return g
}

func newMCProf(family, secId, name string) *Game {
	c := cwd()
	pPath := path.Join(c, "data", "profiles", name, "default")
	g := NewGame(
		fmt.Sprintf("shell:appsFolder\\%s!App", family),
		path.Join(os.Getenv("LOCALAPPDATA"), "Packages", family, "LocalState", "games", "com.mojang"),
		secId,
	)
	g.SetIconPath(path.Join(c, "assets", strings.ToLower(name)+".png"))
	g.NewProfile("Default", pPath)
	return g
}

func NewMCRelease() *Game {
	return newMCProf(REL_FAMILY, REL_SID, "release")
}

func NewMCPreview() *Game {
	return newMCProf(PRE_FAMILY, PRE_SID, "preview")
}
