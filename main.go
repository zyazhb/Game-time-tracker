package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(400, 300))

	output := widget.NewLabel("None")
	entry := widget.NewEntry()
	// textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			// log.Println("multiline:", textArea.Text)
			Pbool, Pname, _ := isProcessExist(entry.Text + ".exe")
			if Pbool {
				output.SetText(Pname + " is running")
			} else {
				output.SetText(Pname + " is not running")
			}

			// w.Close()
		},
	}

	// we can also append items
	// form.Append("Text", textArea)

	w.SetContent(container.NewVBox(
		form,
		output,
	))

	w.ShowAndRun()
}
