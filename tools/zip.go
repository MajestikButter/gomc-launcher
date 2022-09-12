package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
)

func ZipFile(archive *zip.Writer, sFile, dFile string) {
	fmt.Println("Zipping File \n  '" + sFile + "'\n  '" + dFile + "'")
	dw, err := archive.Create(dFile)
	if err != nil {
		panic(err)
	}

	sr, err := os.Open(sFile)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(dw, sr); err != nil {
		panic(err)
	}
	sr.Close()
}

func ZipDir(archive *zip.Writer, sDir, dDir string) {
	fmt.Println("Zipping Directory \n  '" + sDir + "'\n  '" + dDir + "'")
	files, err := os.ReadDir(sDir)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		s, d := path.Join(sDir, f.Name()), path.Join(dDir, f.Name())
		if f.IsDir() {
			ZipDir(archive, s, d)
		} else {
			ZipFile(archive, s, d)
		}
	}
}

func Zip(paths [][]string) {
	res, err := os.Create("bin/gomc-launcher.zip")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	archive := zip.NewWriter(res)

	for _, p := range paths {
		i, err := os.Stat(p[0])
		if err != nil {
			panic(err)
		}
		if i.IsDir() {
			ZipDir(archive, p[0], p[1])
		} else {
			ZipFile(archive, p[0], p[1])
		}
	}

	archive.Close()
}

func main() {
	Zip([][]string{
		{"bin/gomc-launcher.exe", "gomc-launcher.exe"},
		{"assets/", "assets/"},
		{"changelog.md", "changelog.md"},
		{"LICENSE", "LICENSE"},
	})
}
