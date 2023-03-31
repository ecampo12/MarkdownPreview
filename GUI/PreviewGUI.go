package main

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/russross/blackfriday/v2"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Markdown Preview")

	mdEntry := widget.NewMultiLineEntry()
	mdEntry.Wrapping = fyne.TextWrapWord
	mdEntry.Resize(fyne.NewSize(400, 400))

	htmlLabel := widget.NewLabel("")
	htmlLabel.Wrapping = fyne.TextWrapWord
	htmlLabel.Resize(fyne.NewSize(400, 400))

	// Toolbar items
	openMenuItem := fyne.NewMenuItem("Open", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				bytes, err := ioutil.ReadAll(reader)
				if err == nil {
					mdEntry.SetText(string(bytes))
				}
				reader.Close()
			}
		}, myWindow)
	})

	saveMenuItem := fyne.NewMenuItem("Save", func() {
		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err == nil && writer != nil {
				_, err := writer.Write([]byte(mdEntry.Text))
				if err != nil {
					dialog.ShowError(err, myWindow)
				}
				writer.Close()
			}
		}, myWindow)
	})

	mainMenu := fyne.NewMenu(
		"File",
		openMenuItem,
		saveMenuItem,
	)

	mdEntry.OnChanged = func(text string) {
		html := blackfriday.Run([]byte(text))
		htmlLabel.SetText(string(html))
	}

	content := container.New(layout.NewGridLayoutWithColumns(2), mdEntry, htmlLabel)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.SetMainMenu(fyne.NewMainMenu(mainMenu))
	myWindow.ShowAndRun()
}
