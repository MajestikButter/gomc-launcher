package clayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type EditScreen struct {
}

func (l *EditScreen) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h+theme.Padding())
}

func (l *EditScreen) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	bar := objects[0]
	content := objects[1]

	h := bar.MinSize().Height
	bar.Resize(fyne.NewSize(containerSize.Width, h))
	content.Resize(fyne.NewSize(containerSize.Width, containerSize.Height-h-theme.Padding()))

	bar.Move(fyne.NewPos(0, 0))
	content.Move(fyne.NewPos(0, h+theme.Padding()))
}
