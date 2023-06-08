package OS

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/graphic"
	"godeck/src/core/OS/views/navigationview"
	"godeck/src/core/OS/views/statusbar"
	"time"
)

type ScreenManager struct {
	screen     *graphic.Screen
	brightness int
}

type HomeScreen struct {
	Screen *graphic.Screen
}

func (p *ProdeckOS) createHomeScreen() {
	home := &HomeScreen{
		Screen: p.screenManager.screen,
	}

	statusBar := statusbar.CreateStatusBar()

	navigationView := navigationview.CreateNavigationView(
		p.screenManager.screen.Width-statusBar.Width,
		p.screenManager.screen.Height)

	// Create layout
	layout := &graphic.Layout{
		Parent: home.Screen,
	}

	layout.AddEntry(*statusBar.Screen)
	layout.AddEntry(*navigationView.Screen)

	home.Screen.SetLayout(layout)
}

func (p *ProdeckOS) createScreenManager(height int, width int, brightness int) {
	p.screenManager = &ScreenManager{
		screen: &graphic.Screen{
			Width:  width,
			Height: height,
		},
		brightness: brightness,
	}

	p.createHomeScreen()

	p.startRefreshLoop()
}

func (p *ProdeckOS) startRefreshLoop() {
	screen := p.screenManager.screen

	for x := 0; x < screen.Width; x++ {
		for y := 0; y < p.screenManager.screen.Height; y++ {
			entryColor := screen.Layout.GetEntryColor(x, y)
			p.connector.SetButtonColor(screen.ButtonNumber(x, y), entryColor)
		}
	}

	go func() {
		for _ = range time.Tick(time.Second) {
			waitingForRedrew := screen.ButtonsWaitingForRedraw()
			log.WithFields(log.Fields{
				"waitingForRedrew": len(waitingForRedrew),
			}).Println("Refresh screen ")

			/*for x := 0; x < screen.Width; x++ {
				for y := 0; y < p.screenManager.screen.Height; y++ {
					entryColor := screen.Layout.GetEntryColor(x, y)
					p.connector.SetButtonColor(screen.ButtonNumber(x, y), entryColor)
				}
			}*/
		}
	}()
}