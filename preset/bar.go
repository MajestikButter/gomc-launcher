package preset

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewHeaderBar(header string, backFunc func()) *fyne.Container {
	content := []fyne.CanvasObject{}

	if backFunc != nil {
		content = append(content, widget.NewButtonWithIcon("Back", theme.NavigateBackIcon(), backFunc))
	}

	content = append(content,
		layout.NewSpacer(),
		widget.NewRichTextFromMarkdown("# "+header),
		layout.NewSpacer(),
	)

	return container.NewHBox(content...)
}
