package clayout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Expand struct {
	Vertical bool
}

func (l *Expand) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		if l.Vertical {
			w = fyne.Max(w, childSize.Width)
			h += childSize.Height
		} else {
			w += childSize.Width
			h = fyne.Max(h, childSize.Height)
		}
	}
	if l.Vertical {
		h += theme.Padding()
	} else {
		w += theme.Padding()
	}
	return fyne.NewSize(w, h)
}

func (l *Expand) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	first := objects[0]
	second := objects[1]

	w := second.MinSize().Width
	h := second.MinSize().Height
	s := second.MinSize()
	if l.Vertical {
		first.Resize(fyne.NewSize(containerSize.Width, containerSize.Height-theme.Padding()-h))
		second.Move(fyne.NewPos(0, containerSize.Height-h))
		s = fyne.NewSize(containerSize.Width, s.Height)
	} else {
		first.Resize(fyne.NewSize(containerSize.Width-theme.Padding()-w, containerSize.Height))
		second.Move(fyne.NewPos(containerSize.Width-w, 0))
		s = fyne.NewSize(s.Width, containerSize.Height)
	}
	second.Resize(s)
	first.Move(fyne.NewPos(0, 0))
}
