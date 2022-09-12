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

func NewFileSetting(
	window fyne.Window, label, pathStr string, filter storage.FileFilter, validator func(path string) error,
	changed func(path string), setText func(path string) string, valPath func(text string) string,
) fyne.Widget {
	sPath := pathStr

	pathw := widget.NewEntry()
	if setText != nil {
		pathw.SetText(setText(sPath))
	} else {
		pathw.SetText(sPath)
	}
	pathw.Validator = func(s string) error {
		if validator != nil {
			err := validator(s)
			if err != nil {
				return err
			}
		} else if s == "" {
			return errors.New("path cannot be empty")
		}

		if valPath != nil {
			s = valPath(s)
		}

		u, err := storage.ParseURI("file://" + s)
		if err != nil {
			return err
		}

		if !filter.Matches(u) {
			return errors.New("invalid file type")
		}

		stat, err := os.Stat(s)
		if err != nil || stat.IsDir() {
			return errors.New("file does not exist")
		}
		return nil
	}
	pathw.OnChanged = func(s string) {
		if pathw.Validate() == nil {
			sPath = s
			if changed != nil {
				changed(sPath)
			}
		}
	}

	return NewSetting(label, container.New(
		&clayout.Expand{},
		pathw,
		widget.NewButtonWithIcon("", theme.FileImageIcon(), func() {
			d := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
				if uc != nil {
					sPath = uc.URI().Path()
					uc.Close()

					if setText != nil {
						pathw.SetText(setText(sPath))
					} else {
						pathw.SetText(sPath)
					}
				}
			}, window)
			d.SetFilter(filter)

			b, _ := path.Split(strings.ReplaceAll(path.Clean(sPath), `\`, "/"))
			u, err := storage.ParseURI("file://" + b)
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

func NewFolderSetting(
	window fyne.Window, label, pathStr string, validator func(path string) error,
	changed func(path string), setText func(path string) string, unsetText func(text string) string,
) fyne.Widget {
	sPath := pathStr

	pathw := widget.NewEntry()
	if setText != nil {
		pathw.SetText(setText(sPath))
	} else {
		pathw.SetText(sPath)
	}
	pathw.Validator = func(s string) error {
		if unsetText != nil {
			s = unsetText(s)
		}

		if validator != nil {
			err := validator(s)
			if err != nil {
				return err
			}
		} else if s == "" {
			return errors.New("path cannot be empty")
		}

		if len(strings.Split(strings.ReplaceAll(path.Clean(s), `\`, "/"), "/")) < 3 {
			return errors.New("path is too short")
		}
		stat, err := os.Stat(s)
		if err != nil || !stat.IsDir() {
			return errors.New("path does not exist")
		}
		return nil
	}
	pathw.OnChanged = func(s string) {
		if pathw.Validate() == nil {
			sPath = s
			if changed != nil {
				changed(sPath)
			}
		}
	}

	return NewSetting(label, container.New(
		&clayout.Expand{},
		pathw,
		widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
			d := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
				if lu != nil {
					sPath = lu.Path()
					if setText != nil {
						pathw.SetText(setText(sPath))
					} else {
						pathw.SetText(sPath)
					}
				}
			}, window)

			s := sPath
			if unsetText != nil {
				s = unsetText(s)
			}
			u, err := storage.ParseURI("file://" + s)
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

type IconProvider interface {
	Icon() fyne.Resource
	IconPath() string
	SetIconPath(path string)
}

func NewIconSetting(window fyne.Window, provider IconProvider) *fyne.Container {
	icon := widget.NewIcon(provider.Icon())
	return container.NewVBox(
		container.NewGridWrap(
			fyne.NewSize(270, 200),
			icon,
		),
		NewFileSetting(window, "Icon", provider.IconPath(), storage.NewExtensionFileFilter([]string{".png", ".jpg"}), func(path string) error {
			return nil
		}, func(path string) {
			provider.SetIconPath(path)
			icon.SetResource(provider.Icon())
		}, nil, nil),
	)
}
