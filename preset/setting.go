package preset

import (
	"errors"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/ccontainer"
	"github.com/MajestikButter/gomc-launcher/clayout"
)

func NewSetting(label string, content fyne.CanvasObject) fyne.Widget {
	text := widget.NewRichTextFromMarkdown("### " + label)
	content.Move(fyne.NewPos(0, text.MinSize().Height))
	return widget.NewCard("", "",
		ccontainer.NewSetting(
			text,
			content,
		),
	)
}

func NewInputSetting(label, text string, changed func(text string)) fyne.Widget {
	input := widget.NewEntry()
	input.SetText(text)
	input.OnChanged = changed

	return NewSetting(label, input)
}

func NewFileSetting(window fyne.Window, label, pathStr string, changed func(path string)) fyne.Widget {
	sPath := pathStr
	// if len(path) > max_path_chars {
	// 	sPath = sPath[:max_path_chars-3] + "..."
	// }

	pathw := widget.NewEntry()
	pathw.SetText(sPath)
	pathw.Validator = func(s string) error {
		if len(strings.Split(strings.ReplaceAll(path.Clean(s), `\`, "/"), "/")) < 3 {
			return errors.New("path is too short")
		}
		stat, err := os.Stat(s)
		if err != nil || !stat.IsDir() {
			return errors.New("invalid path")
		}
		return nil
	}
	pathw.OnChanged = func(s string) {
		if pathw.Validate() == nil {
			sPath = s
			changed(sPath)
		}
	}

	return NewSetting(label, container.New(
		&clayout.Expand{},
		pathw,
		widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
			d := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
				if lu != nil {
					sPath = lu.Path()
					// if len(path) > max_path_chars {
					// 	sPath = sPath[:max_path_chars-3] + "..."
					// }
					pathw.SetText(sPath)
					changed(sPath)
				}
			}, window)

			u, err := storage.ParseURI("file://" + pathStr)
			if err != nil {
				dialog.NewError(err, window).Show()
				return
			}

			ul, err := storage.ListerForURI(u)
			if err != nil {
				dialog.NewError(err, window).Show()
				return
			}
			d.SetLocation(ul)
			d.Show()
		}),
	))
}
