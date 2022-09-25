package preset

import (
	"errors"
	"os"
	"path"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MajestikButter/gomc-launcher/ccontainer"
	"github.com/MajestikButter/gomc-launcher/clayout"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/logger"
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
	"github.com/modfin/henry/slicez"
)

func NewSetting(label string, content fyne.CanvasObject) fyne.Widget {
	defer logger.HandlePanic()

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
	window fyne.Window, label, pathStr, filter string, validator func(path string) error,
	changed func(path string), setText func(path string) string, valPath func(text string) string,
) fyne.Widget {
	sPath := pathStr

	pathw := widget.NewEntry()
	if setText != nil {
		pathw.SetText(setText(sPath))
	} else {
		pathw.SetText(sPath)
	}

	filterSplit := slicez.Map(strings.Split(filter, ";"), func(a string) []string { return strings.Split(a, ".") })

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

		_, f := path.Split(strings.ReplaceAll(s, `\`, "/"))
		if len(filterSplit) > 0 && !slicez.SomeFunc(filterSplit, func(a []string) bool {
			if len(a) < 2 {
				a = append(a, "*")
			}
			if a[0] == "*" && a[1] == "*" {
				return true
			}
			if a[0] == "*" && strings.HasSuffix(f, a[1]) {
				return true
			}
			return a[1] == "*" && strings.HasPrefix(f, a[0])
		}) {
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
			defer logger.HandlePanic()

			// d := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			// 	if uc != nil {
			// 		sPath = uc.URI().Path()
			// 		uc.Close()

			// 		if setText != nil {
			// 			pathw.SetText(setText(sPath))
			// 		} else {
			// 			pathw.SetText(sPath)
			// 		}
			// 	}
			// }, window)
			// d.SetFilter(filter)

			// u, err := storage.ParseURI("file://" + b)
			// if err != nil {
			// 	dialog.NewError(err, window).Show()
			// 	return
			// }

			// ul, err := storage.ListerForURI(u)
			// if err != nil {
			// 	dialog.NewError(err, window).Show()
			// 	return
			// }
			// d.SetLocation(ul)
			// d.Show()

			s := sPath
			if _, err := os.Stat(s); os.IsNotExist(err) {
				s = launcher.DATA_PATH
			}
			b, f := path.Split(strings.ReplaceAll(s, `\`, "/"))

			res, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
				Title:       "Select Folder",
				Role:        "SelectFolder",
				Folder:      strings.ReplaceAll(b, `/`, `\`),
				FileName:    f,
				FileFilters: []cfd.FileFilter{{Pattern: filter}},
			})

			if err == cfd.ErrorCancelled {
				return
			} else if err != nil {
				dialog.NewError(err, window).Show()
				return
			}
			if setText != nil {
				pathw.SetText(setText(res))
			} else {
				pathw.SetText(res)
			}
		}),
	))
}

func NewFolderSetting(
	window fyne.Window, label, pathStr string, validator func(path string) error,
	changed func(path string), setText, unsetText func(string) string,
) fyne.Widget {
	defer logger.HandlePanic()

	sPath := pathStr

	pathw := widget.NewEntry()
	if setText != nil {
		pathw.SetText(setText(sPath))
	} else {
		pathw.SetText(sPath)
	}
	pathw.Validator = func(s string) error {
		defer logger.HandlePanic()

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
		defer logger.HandlePanic()

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
			defer logger.HandlePanic()

			// d := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
			// 	if lu != nil {
			// 		sPath = lu.Path()
			// 		if setText != nil {
			// 			pathw.SetText(setText(sPath))
			// 		} else {
			// 			pathw.SetText(sPath)
			// 		}
			// 	}
			// }, window)

			s := sPath
			if unsetText != nil {
				s = unsetText(s)
			}

			// u, err := storage.ParseURI("file://" + s)
			// if err != nil {
			// 	dialog.NewError(err, window).Show()
			// 	return
			// }

			// ul, err := storage.ListerForURI(u)
			// if err != nil {
			// 	dialog.NewError(err, window).Show()
			// 	return
			// }
			// d.SetLocation(ul)
			// d.Show()

			if _, err := os.Stat(s); os.IsNotExist(err) {
				s = launcher.DATA_PATH
			}
			b, f := path.Split(strings.ReplaceAll(s, `\`, "/"))
			res, err := cfdutil.ShowPickFolderDialog(cfd.DialogConfig{
				Title:    "Select Folder",
				Role:     "SelectFolder",
				Folder:   strings.ReplaceAll(b, `/`, `\`),
				FileName: f,
			})

			if err == cfd.ErrorCancelled {
				return
			} else if err != nil {
				dialog.NewError(err, window).Show()
				return
			}
			if setText != nil {
				pathw.SetText(setText(res))
			} else {
				pathw.SetText(res)
			}
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
		NewFileSetting(window, "Icon", provider.IconPath(), "*.png;*.jpg", func(path string) error {
			return nil
		}, func(path string) {
			provider.SetIconPath(path)
			icon.SetResource(provider.Icon())
		}, nil, nil),
	)
}
