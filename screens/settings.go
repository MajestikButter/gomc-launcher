package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/preset"
)

func (s *Screens) CreateSettings(l *launcher.Launcher) *fyne.Container {
	kCheck := widget.NewCheck("", func(b bool) {
		l.KeepOpen = b
	})
	kCheck.SetChecked(l.KeepOpen)

	return container.NewVBox(
		preset.NewSetting("Keep Launcher Open", kCheck),
		widget.NewSeparator(),
		widget.NewAccordion(
			widget.NewAccordionItem("Advanced", container.NewVBox(
				preset.NewFolderSetting(s.Window, "Logs Directory", l.LogsDirectory, false, func(path string) {
					l.LogsDirectory = path
				}, nil, nil),
			)),
		),
		widget.NewLabel("                                                                                                                      "),
	)
}
