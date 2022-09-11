package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

const FILE_MODE = 0777

var LOG_READER *os.File

var STDOUT *os.File

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

func WriteLog() {
	buf := make([]byte, 50000)
	n, err := LOG_READER.Read(buf)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// fmt.Println("Written to stdout:", string(buf[:n]))

	t := time.Now()
	l := path.Join(Path(), t.Format("LOG_2006-01-02_15-04-05.txt"))

	os.WriteFile(l, buf[:n], FILE_MODE)
}

func init() {
	lp := Path()
	if _, err := os.Stat(lp); os.IsNotExist(err) {
		os.MkdirAll(lp, FILE_MODE)
	}

	STDOUT = os.Stdout
	r, w, err := os.Pipe()
	LOG_READER = r
	if err != nil {
		panic(err)
	}
	os.Stdout = w
}

func timeString() string {
	t := time.Now()
	return t.Format("[01/02 15:04:05.00000]")
}

// Printf without time prefix
func RPrintf(format string, a ...any) {
	fmt.Printf(format, a...)
	STDOUT.WriteString(fmt.Sprintf(format, a...))
}

// similar to log.Print
func Print(a ...any) {
	a = append([]any{timeString()}, a...)
	fmt.Print(a...)
	STDOUT.WriteString(fmt.Sprint(a...))
}

// similar to log.Printf
func Printf(format string, a ...any) {
	format = timeString() + " " + format
	fmt.Printf(format, a...)
	STDOUT.WriteString(fmt.Sprintf(format, a...))
}

// similar to log.Println
func Println(a ...any) {
	a = append([]any{timeString()}, a...)
	fmt.Println(a...)
	STDOUT.WriteString(fmt.Sprintln(a...))
}
