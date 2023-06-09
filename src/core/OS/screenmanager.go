package OS

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/imagebuilder"
	"time"
)

type ScreenManager struct {
	lockScreen *LockScreen
	screen     *graphic.Screen
	brightness int
}

func (m ScreenManager) ButtonPressed(button int) {

}

func (p *ProdeckOS) createScreenManager(height int, width int, brightness int) {
	p.screenManager = &ScreenManager{
		screen: &graphic.Screen{
			Width:  width,
			Height: height,
		},
		lockScreen: CreateLockScreen(width, height),
		brightness: brightness,
	}

	p.createHomeScreen()

	p.startRefreshLoop()
}

func (p *ProdeckOS) currentScreen() *graphic.Screen {
	print("lock")
	if p.screenManager.lockScreen.isLocked {
		println("isLocked")
		return p.screenManager.lockScreen.Screen
	} else {
		println("isNotLocked")
		return p.screenManager.screen
	}
}

func (p *ProdeckOS) refreshAllScreen() {
	screen := p.currentScreen()

	log.Println("isLocked: ", p.screenManager.lockScreen.isLocked)

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < screen.Height; y++ {
			/*entryColor := screen.GetEntryColor(x, y)
			p.connector.SetButtonColor(screen.ButtonNumber(x, y), entryColor)*/

			img := imagebuilder.CreateImage(96, 96)
			png := imagebuilder.LoadPng("./resources/icon.png")

			if img == nil {
				log.Errorln("img is nil")
				continue
			}

			if png == nil {
				log.Errorln("png is nil")
				continue
			}

			newimg := imagebuilder.AddPngToBaseImage(img, png, 0, 0)

			p.connector.SetButtonImage(screen.ButtonNumber(x, y), newimg)
		}
	}
}

func (p *ProdeckOS) startRefreshLoop() {
	screen := p.currentScreen()

	p.refreshAllScreen()

	go func() {
		wasLocked := false
		for _ = range time.Tick(time.Second) {
			waitingForRedrew := screen.ButtonsWaitingForRedraw()

			if wasLocked && !p.screenManager.lockScreen.isLocked ||
				!wasLocked && p.screenManager.lockScreen.isLocked {
				wasLocked = p.screenManager.lockScreen.isLocked

				p.refreshAllScreen()
			}

			for _, button := range waitingForRedrew {
				//If button is pressed, don't redraw it now
				if p.buttonsState[screen.ButtonNumber(button.X, button.Y)] {
					continue
				}
				entryColor := screen.GetEntryColor(button.X, button.Y)
				p.connector.SetButtonColor(screen.ButtonNumber(button.X, button.Y), entryColor)
			}

			log.WithFields(log.Fields{
				"waitingForRedrew": len(waitingForRedrew),
			}).Println("Refresh screen ")
		}
	}()
}