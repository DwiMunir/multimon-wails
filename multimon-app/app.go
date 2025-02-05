package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Get all screens
	screens, err := runtime.ScreenGetAll(ctx)
	if err != nil {
		println("Error getting screens:", err.Error())
		return
	}

	// Calculate total width and maximum height
	var totalWidth int
	var maxHeight int

	// Calculate dimensions
	for _, screen := range screens {
		// Add screen width to total
		if screen.Size.Width > 0 {
			totalWidth += screen.Size.Width
		}

		// Keep track of maximum height
		if screen.Size.Height > maxHeight {
			maxHeight = screen.Size.Height
		}
	}

	// Set window size to cover all screens width
	runtime.WindowSetSize(ctx, totalWidth, maxHeight)

	// Set window position to the primary display (0,0)
	runtime.WindowSetPosition(ctx, 10, 10)

	// Ensure window is always on top
	runtime.WindowSetAlwaysOnTop(ctx, true)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return "Hello " + name + "! Welcome to MultiMon App"
}

// ShowMessage displays a message in the frontend
func (a *App) ShowMessage(message string) {
	runtime.EventsEmit(a.ctx, "show-message", message)
}
