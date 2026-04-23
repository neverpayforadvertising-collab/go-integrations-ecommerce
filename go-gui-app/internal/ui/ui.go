package ui // Defines UI layer package

import (
	"fyne.io/fyne/v2" // Core Fyne interfaces and types
	"fyne.io/fyne/v2/app" // Provides app lifecycle management
	"fyne.io/fyne/v2/container" // Layout containers (VBox, HBox, etc.)
	"fyne.io/fyne/v2/widget" // UI components like buttons, labels, inputs

	"go-gui-app/internal/services" // Imports service layer for API calls
)

// StartApp initializes and runs the GUI application
func StartApp() {
	a := app.New() // Creates a new Fyne application instance (manages lifecycle)
	w := a.NewWindow("Go Integration GUI") // Creates a new window with a title

	output := widget.NewMultiLineEntry() // Creates a multi-line text input for displaying results
	output.SetPlaceHolder("API response will appear here...") // Sets placeholder text

	button := widget.NewButton("Fetch API", func() { // Creates a button with click handler
		output.SetText("Loading...") // Immediately updates UI to show loading state

		go func() { // Starts a goroutine to avoid blocking UI thread
			result := services.FetchData() // Calls service layer to fetch API data
			output.SetText(result) // Updates UI with result after API call completes
		}()
	})

	content := container.NewVBox( // Creates vertical layout container
		button, // Adds button at top
		output, // Adds output box below
	)

	w.SetContent(content) // Sets window content to defined layout
	w.Resize(fyne.NewSize(500, 400)) // Sets initial window size
	w.ShowAndRun() // Displays window and starts event loop
}