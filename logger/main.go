package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

var log = []byte{}

var p string

func Path() string {
	if p == "" {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p = path.Join(cwd, "logs")
	}

	return p
}

var f string

func FilePath() string {
	if f == "" {
		t := time.Now()
		f = path.Join(Path(), t.Format("LOG_2006-01-02_15-04-05.txt"))
	}
	return f
}

func WriteLog() {
	os.WriteFile(FilePath(), log, 0777)
}

func writeString(data string) {
	log = append(log, []byte(data)...)
}

func timeString() string {
	t := time.Now()
	return t.Format("[01/02 15:04:05.00000]")
}

// Printf without time prefix
func RPrintf(format string, a ...any) {
	fmt.Printf(format, a...)
	writeString(fmt.Sprintf(format, a...))
}

// similar to log.Print
func Print(a ...any) {
	a = append([]any{timeString()}, a...)
	fmt.Print(a...)
	writeString(fmt.Sprint(a...))
}

// similar to log.Printf
func Printf(format string, a ...any) {
	format = timeString() + " " + format
	fmt.Printf(format, a...)
	writeString(fmt.Sprintf(format, a...))
}

// similar to log.Println
func Println(a ...any) {
	a = append([]any{timeString()}, a...)
	fmt.Println(a...)
	writeString(fmt.Sprintln(a...))
}
