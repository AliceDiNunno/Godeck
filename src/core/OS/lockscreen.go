package OS

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"image/color"
)

type LockScreen struct {
	isLocked bool

	*graphic.Screen
}

func (n *LockScreen) ButtonPressed(caller *graphic.Screen, x int, y int) {
	println("Redraw LS")
	caller.AskForRedraw(x, y)
}

func (n *LockScreen) GetButtonColor(caller *graphic.Screen, x int, y int) color.RGBA {
	log.Println("LockScreen GetButtonColor")
	return color.RGBA{R: 255, G: 255, B: 0, A: 255}
}

func CreateLockScreen(width int, height int) *LockScreen {
	navigationView := &LockScreen{
		Screen: &graphic.Screen{
			Width:  width,
			Height: height,
		},
		isLocked: false,
	}
	navigationView.Screen.Interactor = navigationView
	return navigationView
}