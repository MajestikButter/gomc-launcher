package screens

import (
	"fmt"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/game"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/preset"
)

func (s *Screens) RefreshSubfolders(p *game.Profile, subVBox *fyne.Container) {
	subVBox.RemoveAll()
	for f := range p.Subfolders {
		n := f
		subVBox.Add(container.NewHBox(
			widget.NewLabel(f),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
				s.DialogEditSubfolder(p, n, subVBox)
			}),
		))
	}
	subVBox.Add(
		widget.NewButtonWithIcon("New Subfolder", theme.FolderNewIcon(), func() {
			p.Subfolders["new_subfolder"] = path.Join(p.Path, "new_subfolder")
			s.DialogEditSubfolder(p, "new_subfolder", subVBox)
		}),
	)
}

func (s *Screens) DialogEditProfile(l *launcher.Launcher, g *game.Game, p *game.Profile, name string) {
	var d dialog.Dialog
	oldName := name
	n := &name

	subVBox := container.NewVBox()
	s.RefreshSubfolders(p, subVBox)
	subScroll := container.NewVScroll(subVBox)
	subScroll.SetMinSize(fyne.NewSize(0, 80))

	content := container.NewHBox(
		container.NewVBox(
			preset.NewInputSetting("Name", *n, func(text string) {
				*n = text
			}),
			widget.NewSeparator(),
			widget.NewAccordion(
				widget.NewAccordionItem("Advanced", container.NewVBox(
					preset.NewFolderSetting(s.Window, "Path", p.Path, nil, func(path string) {
						p.Path = path
					}, nil, nil),

					widget.NewCard("", "",
						widget.NewAccordion(
							widget.NewAccordionItem("Subfolders", subScroll),
						),
					),

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
			widget.NewLabel("                                                                                                     "),
		),
		preset.NewIconSetting(s.Window, p),
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
