package main

import (
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/creack/pty"
)

func main() {
	a := app.New()
	w := a.NewWindow("germ")

	ui := widget.NewTextGrid()       // Create a new TextGrid
	ui.SetText("I'm on a terminal!") // Set text to display

	c := exec.Command("/bin/bash")
	p, err := pty.Start(c)

	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer c.Process.Kill()

	p.Write([]byte("ls\r"))
	time.Sleep(1 * time.Second)
	b := make([]byte, 1024)
	_, err = p.Read(b)
	if err != nil {
		fyne.LogError("Failed to read pty", err)
	}
	//s := fmt.Sprintf("read bytes from pty.\nContent:%s", string(b))
	//fmt.Println(s)
	ui.SetText(string(b))

	ui_layout := layout.NewGridWrapLayout(fyne.NewSize(700, 400))
	container := container.New(ui_layout, ui)
	w.SetContent(container)

	w.ShowAndRun()

}
