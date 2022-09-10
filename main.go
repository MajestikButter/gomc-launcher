package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/logger"
	"github.com/MajestikButter/gomc-launcher/screens"
	"github.com/go-errors/errors"
)

func HandlePanic() {
	if e := recover(); e != nil {
		err := errors.Wrap(e, 2)
		logger.RPrintf("\n\n====================================[ ERROR ]====================================\n\nMessage: %s\n\nStack: %s\n\n=================================================================================\n\n", err.Error(), err.Stack())
		logger.WriteLog()
		os.Exit(1)
	}
}

var LAUNCHER *launcher.Launcher

func init() {
	defer HandlePanic()

	l := launcher.New()
	l.Load()
	LAUNCHER = l
}

func main() {
	defer HandlePanic()

	l := LAUNCHER

	a := app.New()

	w := a.NewWindow("GO Launcher")
	ws := l.WindowSize
	w.Resize(fyne.NewSize(ws.X, ws.Y))
	w.CenterOnScreen()

	aw := a.NewWindow("GO Launcher")
	aw.Resize(fyne.NewSize(500, 400))
	aw.CenterOnScreen()
	aw.RequestFocus()

	s := screens.Screens{App: a, Window: w}
	s.SetContent(s.CreateMenu(l))

	w.SetMaster()
	w.SetCloseIntercept(func() {
		st := l.WindowSize
		s := w.Content().Size()

		st.X = s.Width
		st.Y = s.Height

		w.Close()
	})
	w.ShowAndRun()

	l.Save()
	// defer HandlePanic()

	// fltk.InitStyles()
	// fltk.SetScheme("gtk+")

	// WinState := LAUNCHER.State.Window
	// Size := WinState.Size
	// Pos := WinState.Position

	// ERR_WIN = CreateErrWindow()
	// EDIT_WIN = CreateEditWindow()
	// MAIN_WIN = CreateMainWindow(Size, Pos)

	// ticker := time.NewTicker(150 * time.Second) // 2.5 Minutes
	// go func() {
	// 	for {
	// 		<-ticker.C
	// 		logger.Println("Autosaving launcher json files")
	// 		LAUNCHER.Save()
	// 	}
	// }()
	// logger.Println("Started save loop")

	// logger.Println("Showing launcher window")

	// LoadScreen(SCREENS.Home)
	// MAIN_WIN.Show()

	// fltk.Run()

	// logger.Println("Saving launcher files and log")
	// LAUNCHER.Save()
	// logger.WriteLog()
}
