package cwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type IconButton struct {
	widget.Icon
	TappedFunc func()
}

func (t *IconButton) Tapped(_ *fyne.PointEvent) {
	if t.TappedFunc != nil {
		t.TappedFunc()
	}
}

func (t *IconButton) TappedSecondary(_ *fyne.PointEvent) {
}

func NewIconButton(res fyne.Resource, tapped func()) *IconButton {
	icon := &IconButton{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)
	icon.TappedFunc = tapped
	// canvas.NewImageFromFile("").

	return icon
}
