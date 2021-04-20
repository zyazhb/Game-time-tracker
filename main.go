package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "log"
)

func main() {
	DbInit()
	a := app.New()
	w := a.NewWindow("Game time tracker BY-ZYA")
	w.Resize(fyne.NewSize(400, 300))

	output := widget.NewLabel("No output")
	Mygametime := widget.NewLabel("No time")
	Totalgametime := widget.NewLabel("No Totaltime")
	entry := widget.NewEntry()
	// textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			//log.Println("Form submitted:", entry.Text)
			Pbool, Pname, _ := isProcessExist(entry.Text + ".exe")
			if Pbool {
				output.SetText(Pname + " is running")
			} else {
				output.SetText(Pname + " is not running")
			}
			StartTime, EndTime, Totaltime := AddNewGame(Pname)
			// Totaltime := AddTotalTime(entry.Text)
			Mygametime.SetText("Start Time:" + StartTime.Format("2006-01-02 15:04:05") + "\nEnd Time:" + EndTime.Format("2006-01-02 15:04:05"))
			Totalgametime.SetText("\nTotal Run:" + Totaltime.String())
		},
		OnCancel: func() {
		},
	}

	w.SetContent(container.NewVBox(
		form,
		output,
		Mygametime,
		Totalgametime,
	))

	w.ShowAndRun()
}
