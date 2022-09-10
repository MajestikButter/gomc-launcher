package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/launcher"
)

func (s *Screens) CreateMenu(l *launcher.Launcher) *fyne.Container {
	return container.NewVBox(
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
				d := dialog.NewCustom("Launcher Settings", "Close", s.CreateSettings(l), s.Window)
				d.Show()
			}),
		),
		layout.NewSpacer(),
		container.NewHBox(
			container.NewVBox(
				widget.NewButtonWithIcon("GitHub", theme.ColorPaletteIcon(), func() {
					s.OpenURL(GIT_URL)
				}),
				widget.NewButtonWithIcon("Twitter", theme.ColorPaletteIcon(), func() {
					s.OpenURL(TWT_URL)
				}),
				widget.NewButtonWithIcon("Youtube", theme.ColorPaletteIcon(), func() {
					s.OpenURL(YT_URL)
				}),
			),
			layout.NewSpacer(),
			widget.NewButton("Launch", func() {
				s.SetContent(s.CreateGames(l))
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
		container.NewHBox(
			widget.NewLabel("Created by MajestikButter"),
			layout.NewSpacer(),
			widget.NewLabel("V 0.0.1"),
		),
	)
}
