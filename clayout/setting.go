package clayout

import (
	"fyne.io/fyne/v2"
)

type Setting struct {
}

func (l *Setting) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w = fyne.Max(w, childSize.Width)
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (l *Setting) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	label := objects[0]
	content := objects[1]

	h := label.MinSize().Height
	label.Resize(fyne.NewSize(containerSize.Width, h))
	content.Resize(fyne.NewSize(containerSize.Width, containerSize.Height-h))

	label.Move(fyne.NewPos(0, 0))
	content.Move(fyne.NewPos(0, h))
}
