package ccontainer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/MajestikButter/gomc-launcher/clayout"
)

func NewBarScreen(bar *fyne.Container, content *fyne.Container) *fyne.Container {
	return container.New(&clayout.BarScreen{}, bar, content)
}
