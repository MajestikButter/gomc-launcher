package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/ccontainer"
	"github.com/MajestikButter/gomc-launcher/cwidget"
	"github.com/MajestikButter/gomc-launcher/game"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/logger"
	"github.com/MajestikButter/gomc-launcher/preset"
	"github.com/modfin/henry/mapz"
	"github.com/modfin/henry/slicez"
)

func (s *Screens) CreateProfiles(l *launcher.Launcher, game *game.Game, name string) *fyne.Container {
	defer logger.HandlePanic()

	le := len(game.Profiles) + 1
	content := make([]fyne.CanvasObject, le)
	sorted := slicez.SortFunc(mapz.Keys(game.Profiles), func(a string, b string) bool {
		return a < b
	})

	for i, pName := range sorted {
		prof := game.Profiles[pName]

		n := pName
		text := n

		if game.Selected() == prof {
			text = "[X] " + text
		}
		content[i] = preset.NewCardButton(text, "Select", prof.Icon(), func() {
			game.SelectedProfile = n
			s.SetContent(s.CreateProfiles(l, game, name))
		})
		i++
	}
	content[le-1] = widget.NewCard("", "", cwidget.NewIconButton(theme.ContentAddIcon(), func() {
		sel := game.Selected()
		p := ""
		if sel != nil {
			p = sel.Path
		}
		pName, prof := game.NewProfile("New Profile", p)
		s.SetContent(s.CreateProfiles(l, game, name))
		s.DialogEditProfile(l, game, prof, pName)
	}))

	scroll := container.NewVScroll(
		container.NewGridWrap(
			fyne.NewSize(170, 170),
			content...,
		),
	)
	return ccontainer.NewBarScreen(
		preset.NewHeaderBar("Profiles", func() {
			s.SetContent(s.CreateGames(l))
		}),
		ccontainer.NewProfiles(
			widget.NewRichTextFromMarkdown(fmt.Sprintf("# %s", name)),
			widget.NewButton("Edit", func() {
				defer logger.HandlePanic()

				sel := game.Selected()
				if sel == nil {
					return
				}

				// oldName := game.SelectedProfile
				name := game.SelectedProfile
				s.DialogEditProfile(l, game, sel, name)
				// d := dialog.NewCustom("Profile Settings", "Close", s.CreateEditProfile(l, game, game.Selected(), &name), s.Window)
				// d.Show()
				// d.SetOnClosed(func() {
				// 	if oldName != name {
				// 		game.RenameProfile(oldName, name)
				// 		s.SetContent(s.CreateProfiles(l, game, name))
				// 	}
				// })
			}),
			widget.NewButtonWithIcon("Play", theme.NavigateNextIcon(), func() {
				defer logger.HandlePanic()

				sel := game.Selected()
				if game.LaunchScript == "" || game.Destination == "" || sel == nil {
					return
				}

				game.LoadProfile(sel)

				launcher.RunScript(game.LaunchScript)
				if !l.KeepOpen {
					s.Window.Close()
				}
			}),
			widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
				defer logger.HandlePanic()

				// oldName := name

				s.DialogEditGame(l, game, name)
				// d = dialog.NewCustom("Game Settings", "Close", s.CreateEditGame(l, game, &name), s.Window)
				// d.Show()
				// d.SetOnClosed(func() {
				// 	if oldName != name {
				// 		l.RenameGame(oldName, name)
				// 		s.SetContent(s.CreateProfiles(l, game, name))
				// 	}
				// })
			}),
			widget.NewCard("", "", container.NewMax(scroll)),
		),
	)
}
