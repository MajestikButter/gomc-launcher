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

func (s *Screens) DialogEditProfile(l *launcher.Launcher, g *game.Game, p *game.Profile, name string) {
	var d dialog.Dialog
	oldName := name
	n := &name
	content := container.NewVBox(
		preset.NewInputSetting("Name", *n, func(text string) {
			*n = text
		}),
		widget.NewSeparator(),
		widget.NewAccordion(
			widget.NewAccordionItem("Advanced", container.NewVBox(
				preset.NewFileSetting(s.Window, "Path", p.Path, func(path string) {
					p.Path = path
				}),
				widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					msg := widget.NewLabel(
						fmt.Sprintf("This action is irreversible and cannot be undone. Be sure you would like to delete '%s' before clicking the 'Delete' button below", *n),
					)
					msg.Wrapping = fyne.TextWrapWord

					dialog.ShowCustomConfirm(
						"Are you sure you want to delete this profile?",
						"Delete",
						"Cancel",
						msg,
						func(b bool) {
							if !b {
								return
							}
							delete(g.Profiles, *n)
							s.SetContent(s.CreateProfiles(l, g, *n))
							d.Hide()
						},
						s.Window,
					)
				}),
			)),
		),
		widget.NewLabel("                                                                                                                      "),
	)
	d = dialog.NewCustom("Edit Profile", "Close", content, s.Window)
	d.Show()

	d.SetOnClosed(func() {
		if oldName != name {
			g.RenameProfile(oldName, name)
			s.SetContent(s.CreateProfiles(l, g, name))
		}
	})
}
