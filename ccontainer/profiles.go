package ccontainer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/clayout"
)

func NewProfiles(title *widget.RichText, editProfile, play, editGame *widget.Button, scroll fyne.Widget) *fyne.Container {
	return container.New(&clayout.Profiles{}, title, editProfile, play, editGame, scroll)
}
