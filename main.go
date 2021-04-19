package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	DbInit()
	a := app.New()
	w := a.NewWindow("Game time tracker BY-ZYA")
	w.Resize(fyne.NewSize(400, 300))

	output := widget.NewLabel("No output")
	mygametime := widget.NewLabel("No time")
	entry := widget.NewEntry()
	// textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			Pbool, Pname, _ := isProcessExist(entry.Text + ".exe")
			if Pbool {
				output.SetText(Pname + " is running")
			} else {
				output.SetText(Pname + " is not running")
			}
			StartTime, EndTime := AddNewGame(Pname)
			mygametime.SetText("Start Time:" + StartTime + "\nEnd Time:" + EndTime + "\nTotal Run:")
		},
	}

	w.SetContent(container.NewVBox(
		form,
		output,
		mygametime,
	))

	w.ShowAndRun()
}
