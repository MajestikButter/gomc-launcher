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
	ctext := widget.NewRichTextFromMarkdown(l.Changelog())
	ctext.Wrapping = fyne.TextWrapWord
	changelog := container.NewVScroll(ctext)
	changelog.SetMinSize(fyne.NewSize(700, 300))

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
				layout.NewSpacer(),
				widget.NewButtonWithIcon("GitHub", theme.ColorPaletteIcon(), func() {
					s.OpenURL(GIT_URL)
				}),
				widget.NewButtonWithIcon("Twitter", theme.ColorPaletteIcon(), func() {
					s.OpenURL(TWT_URL)
				}),
				widget.NewButtonWithIcon("Youtube", theme.ColorPaletteIcon(), func() {
					s.OpenURL(YT_URL)
				}),
				layout.NewSpacer(),
			),
			container.NewVBox(
				widget.NewRichTextFromMarkdown("# Changelogs"),
				changelog,
				widget.NewButton("Launch", func() {
					s.SetContent(s.CreateGames(l))
					m := l.WindowSize
					s.Resize(fyne.NewSize(m.X, m.Y))
				}),
			),
		),
		layout.NewSpacer(),
		container.NewHBox(
			widget.NewLabel("Created by MajestikButter"),
			layout.NewSpacer(),
			widget.NewLabel(l.Version()),
		),
	)
}
