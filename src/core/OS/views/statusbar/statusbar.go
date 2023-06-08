package statusbar

import (
	"godeck/src/core/OS/graphic"
	"image/color"
)

type StatusBar struct {
	*graphic.Screen
}

func (s *StatusBar) GetButtonColor(caller *graphic.Screen, x int, y int) color.RGBA {
	return color.RGBA{R: 0, G: 0, B: 255, A: 255}
}

func (s *StatusBar) ButtonPressed(caller *graphic.Screen, x int, y int) {
	println("Redraw SB")

	caller.AskForRedraw(x, y)
}

func CreateStatusBar() *StatusBar {
	statusBar := &StatusBar{
		Screen: &graphic.Screen{
			Width:  1,
			Height: 4,
		},
	}
	statusBar.Screen.Interactor = statusBar

	return statusBar
}