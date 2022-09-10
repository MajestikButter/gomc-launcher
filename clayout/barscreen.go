package clayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type BarScreen struct {
}

func (l *BarScreen) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w = fyne.Max(w, childSize.Width)
		h += childSize.Height
	}
	return fyne.NewSize(w, h+theme.Padding())
}

func (l *BarScreen) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	bar := objects[0]
	content := objects[1]

	h := bar.MinSize().Height
	bar.Resize(fyne.NewSize(containerSize.Width, h))
	content.Resize(fyne.NewSize(containerSize.Width, containerSize.Height-h-theme.Padding()))

	bar.Move(fyne.NewPos(0, 0))
	content.Move(fyne.NewPos(0, h+theme.Padding()))
}
