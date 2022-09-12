package screens

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/game"
	"github.com/MajestikButter/gomc-launcher/preset"
)

func (s *Screens) DialogEditSubfolder(p *game.Profile, name string) {
	var d dialog.Dialog
	oldName := name
	n := &name
	pa := strings.ReplaceAll(path.Clean(p.Path), `\`, "/")
	content := container.NewHBox(
		container.NewVBox(
			preset.NewFolderSetting(s.Window, "Folder", *n,
				func(pathStr string) error {
					if !strings.HasPrefix(strings.ReplaceAll(path.Clean(pathStr), `\`, "/"), pa) {
						return errors.New("folder must be a subfolder of the profile path")
					}
					return nil
				},
				func(path string) {
					*n = path
				},
				func(text string) string {
					if !strings.HasPrefix(strings.ReplaceAll(path.Clean(text), `\`, "/"), pa) {
						return text
					}
					return text[len(pa):]
				},
				func(text string) string {
					return path.Join(pa, text)
				},
			),
			preset.NewFolderSetting(s.Window, "Destination", p.Subfolders[*n], nil, func(path string) {
				p.Subfolders[oldName] = path
			}, nil, nil),
			widget.NewSeparator(),
			widget.NewAccordion(
				widget.NewAccordionItem("Advanced", container.NewVBox(
					widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
						msg := widget.NewLabel(
							fmt.Sprintf("This action is irreversible and cannot be undone. Be sure you would like to delete '%s' before clicking the 'Delete' button below", *n),
						)
						msg.Wrapping = fyne.TextWrapWord

						dialog.ShowCustomConfirm(
							"Are you sure you want to delete this subfolder?",
							"Delete",
							"Cancel",
							msg,
							func(b bool) {
								if !b {
									return
								}
								delete(p.Subfolders, *n)
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
	d = dialog.NewCustom("Edit Subfolder", "Close", content, s.Window)
	d.Show()

	d.SetOnClosed(func() {
		if oldName != name {
			d := p.Subfolders[oldName]
			delete(p.Subfolders, oldName)
			p.Subfolders[name] = d
		}
	})
}
