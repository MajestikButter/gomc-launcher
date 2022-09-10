package preset

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/clayout"
	"github.com/MajestikButter/gomc-launcher/cwidget"
)

func NewCardButton(label, btnText string, image fyne.Resource, tapped func()) fyne.Widget {
	text := widget.NewLabel(label)
	text.TextStyle = fyne.TextStyle{Bold: true}
	text.Wrapping = fyne.TextWrapWord
	text.Alignment = fyne.TextAlignCenter
	btn := widget.NewButton(btnText, tapped)

	return widget.NewCard("", "", container.New(
		&clayout.Expand{Vertical: true},
		container.New(
			&clayout.Expand{Vertical: true},
			cwidget.NewIconButton(image, tapped),
			text,
		),
		btn,
	))
}
