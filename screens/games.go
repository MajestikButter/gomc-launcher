package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/ccontainer"
	"github.com/MajestikButter/gomc-launcher/cwidget"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/logger"
	"github.com/MajestikButter/gomc-launcher/preset"
	"github.com/modfin/henry/mapz"
	"github.com/modfin/henry/slicez"
)

func (s *Screens) CreateGames(l *launcher.Launcher) *fyne.Container {
	defer logger.HandlePanic()

	le := len(l.Games) + 1
	content := make([]fyne.CanvasObject, le)
	sorted := slicez.SortFunc(mapz.Keys(l.Games), func(a string, b string) bool {
		return a < b
	})
	for i, name := range sorted {
		g := l.Games[name]
		n := name
		content[i] = preset.NewCardButton(name, "Open", g.Icon(), func() {
			s.SetContent(s.CreateProfiles(l, g, n))
		})
	}
	content[le-1] = widget.NewCard("", "", cwidget.NewIconButton(theme.ContentAddIcon(), func() {
		name, game := l.NewGame("New Game")
		s.SetContent(s.CreateGames(l))
		s.DialogEditGame(l, game, name)
	}))

	scroll := container.NewVScroll(
		container.NewGridWrap(
			fyne.NewSize(200, 200),
			content...,
		),
	)
	return ccontainer.NewBarScreen(
		preset.NewHeaderBar("Games", func() {
			s.SetContent(s.CreateMenu(l))
			m := s.Canvas().Size()
			l.WindowSize.X = m.Width
			l.WindowSize.Y = m.Height
			s.Resize(fyne.NewSize(0, 0))
		}),
		container.NewMax(scroll),
	)
}
