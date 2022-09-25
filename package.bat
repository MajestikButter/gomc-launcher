go-winres make
fyne package -os windows -exe bin/gomc-launcher.exe -icon winres/icon.png
go run tools/zip.go