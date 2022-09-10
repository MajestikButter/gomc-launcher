package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/game"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/preset"
)

func (s *Screens) DialogEditGame(l *launcher.Launcher, g *game.Game, name string) {
	var d dialog.Dialog
	n := &name
	content := container.NewVBox(
		preset.NewInputSetting("Name", *n, func(text string) {
			*n = text
		}),
		widget.NewSeparator(),
		widget.NewAccordion(
			widget.NewAccordionItem("Advanced", container.NewVBox(
				preset.NewInputSetting("Launch Script", g.LaunchScript, func(text string) {
					g.LaunchScript = text
				}),
				preset.NewFileSetting(s.Window, "Profile Destination", g.Destination, func(path string) {
					g.Destination = path
				}),
				widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					msg := widget.NewLabel(
						fmt.Sprintf("This action is irreversible and cannot be undone. Be sure you would like to delete '%s' before clicking the 'Delete' button below", *n),
					)
					msg.Wrapping = fyne.TextWrapWord

					dialog.ShowCustomConfirm(
						"Are you sure you want to delete this game?",
						"Delete",
						"Cancel",
						msg,
						func(b bool) {
							if !b {
								return
							}
							delete(l.Games, *n)
							s.SetContent(s.CreateGames(l))
							d.Hide()
						},
						s.Window,
					)
				}),
			)),
		),
		widget.NewLabel("                                                                                                                      "),
	)

	d = dialog.NewCustom("Game Settings", "Close", content, s.Window)
	d.Show()

	oldName := name
	d.SetOnClosed(func() {
		if oldName != name {
			l.RenameGame(oldName, name)
			s.SetContent(s.CreateProfiles(l, g, name))
		}
	})
}
