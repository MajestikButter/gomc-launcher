package clayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Profiles struct {
}

func (l *Profiles) MinSize(objects []fyne.CanvasObject) fyne.Size {
	title := objects[0]
	editProfile := objects[1]
	play := objects[2]
	editGame := objects[3]
	scroll := objects[4]
	return scroll.MinSize().Add(editGame.MinSize()).Max(title.MinSize().Add(play.MinSize()).Add(editProfile.MinSize()))
}

func (l *Profiles) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	title := objects[0]
	editProfile := objects[1]
	play := objects[2]
	editGame := objects[3]
	scroll := objects[4]

	bMinSize := title.MinSize().Max(play.MinSize().Max(editProfile.MinSize()))

	w, h := containerSize.Width, containerSize.Height

	egs := editGame.MinSize()
	scroll.Move(fyne.NewPos(0, 0))
	scroll.Resize(
		fyne.NewSize(
			w-egs.Width-theme.Padding(),
			h-bMinSize.Height-theme.Padding(),
		),
	)

	editGame.Move(fyne.NewPos(w-egs.Width, 0))
	editGame.Resize(egs)

	ts := title.MinSize()
	title.Move(fyne.NewPos(0, h-ts.Height))
	title.Resize(ts)

	ps := play.MinSize()
	pw := ps.Width
	play.Move(fyne.NewPos(w-pw, h-ps.Height))
	play.Resize(ps)

	es := editGame.MinSize()
	editProfile.Move(fyne.NewPos(w-pw-theme.Padding()-es.Width, h-es.Height))
	editProfile.Resize(es)

	// h := bar.MinSize().Height
	// bar.Resize(fyne.NewSize(containerSize.Width, h))
	// content.Resize(fyne.NewSize(containerSize.Width, containerSize.Height-h-theme.Padding()))

	// bar.Move(fyne.NewPos(0, 0))
	// content.Move(fyne.NewPos(0, h+theme.Padding()))
}
