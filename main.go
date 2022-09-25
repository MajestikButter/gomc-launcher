package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/MajestikButter/gomc-launcher/launcher"
	"github.com/MajestikButter/gomc-launcher/logger"
	"github.com/MajestikButter/gomc-launcher/screens"
)

var LAUNCHER *launcher.Launcher

func init() {
	defer logger.HandlePanic()

	l := launcher.New()
	l.Load()
	LAUNCHER = l
}

func main() {
	defer logger.HandlePanic()

	l := LAUNCHER

	a := app.New()

	w := a.NewWindow("GOMC Launcher")
	w.Resize(fyne.NewSize(0, 0))
	w.CenterOnScreen()

	aw := a.NewWindow("GO Launcher")
	aw.Resize(fyne.NewSize(500, 400))
	aw.CenterOnScreen()
	aw.RequestFocus()

	s := screens.Screens{App: a, Window: w}
	s.SetContent(s.CreateMenu(l))

	w.SetMaster()
	w.SetCloseIntercept(func() {
		if !w.FixedSize() {
			st := l.WindowSize
			s := w.Canvas().Size()

			st.X = s.Width
			st.Y = s.Height
		}

		w.Close()
	})

	ticker := time.NewTicker(150 * time.Second) // 2.5 Minutes
	go func() {
		for {
			<-ticker.C
			logger.Println("Autosaving launcher json files")
			l.Save()
		}
	}()
	logger.Println("Started save loop")

	w.ShowAndRun()

	l.Save()
	logger.WriteLog()
}
