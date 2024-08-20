package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Hello, macOS!")

	// Create a label widget
	label := widget.NewLabel("Hello, World!")

	// Create a button widget with a click handler
	button := widget.NewButton("Click Me!", func() {
		label.SetText("Button Clicked!")
	})

	// Set the content of the window
	myWindow.SetContent(container.NewVBox(
		label,
		button,
	))

	// Set the window size
	myWindow.Resize(fyne.NewSize(300, 200))

	// Show the window and run the application
	myWindow.ShowAndRun()
}
