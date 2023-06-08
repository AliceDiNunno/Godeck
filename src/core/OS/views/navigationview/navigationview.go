package navigationview

import (
	"godeck/src/core/OS/graphic"
	"image/color"
)

type NavigationView struct {
	*graphic.Screen
}

func (n *NavigationView) Name() string {
	return "NavigationView"
}

func (n *NavigationView) ButtonPressed(caller *graphic.Screen, x int, y int) {
	println("Redraw NV")
	caller.AskForRedraw(x, y)
}

func (n *NavigationView) GetButtonColor(caller *graphic.Screen, x int, y int) color.RGBA {
	return color.RGBA{R: 255, G: 0, B: 0, A: 255}
}

func CreateNavigationView(width int, height int) *NavigationView {
	navigationView := &NavigationView{
		Screen: &graphic.Screen{
			Width:  width,
			Height: height,
		},
	}
	navigationView.Screen.Interactor = navigationView

	return navigationView
}