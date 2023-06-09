package OS

import (
	"godeck/src/core/OS/graphic"
	"godeck/src/core/OS/views/navigationview"
	"godeck/src/core/OS/views/statusbar"
)

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