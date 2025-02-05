package main

import (
	"context"
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Base options that apply to all platforms
	appOptions := &options.App{
		Title:  "MultiMon App",
		Width:  1280, // This will be overridden by startup
		Height: 800,  // This will be overridden by startup
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
		},
		Bind: []interface{}{
			app,
		},
		EnableDefaultContextMenu: true,
		Frameless:                true,
		AlwaysOnTop:              true,
		MinWidth:                 1280,
		MinHeight:                720,
		Fullscreen:               false,
	}

	// Platform specific options
	if runtime.GOOS == "darwin" {
		appOptions.Mac = &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		}
	} else if runtime.GOOS == "windows" {
		appOptions.Windows = &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
			Theme:                windows.SystemDefault,
		}
	}

	// Run the application
	err := wails.Run(appOptions)
	if err != nil {
		println("Error:", err.Error())
	}
}
