package ccontainer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/MajestikButter/gomc-launcher/clayout"
)

func NewSetting(label fyne.CanvasObject, content fyne.CanvasObject) *fyne.Container {
	return container.New(&clayout.Setting{}, label, content)
}
