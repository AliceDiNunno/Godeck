package Framework

import (
	log "github.com/sirupsen/logrus"
	eventDomain "godeck/src/domain/events"
	"time"
)

func (p *ProdeckFramework) setupButtonEvent() {
	p.eventHub.Subscribe(eventDomain.ButtonStateChangedEvent, func(topic eventDomain.Event, data eventDomain.EventData) {
		if p.wakeUpFromSleep() {
			p.lastInteractionTime = time.Now().Unix()
			p.prepareForSleep()
			return
		}

		button := data["button"].(int)
		pressed := data["state"].(bool)

		log.WithFields(log.Fields(data)).Info(topic)

		//setting last interaction time only if the button is released

		if !pressed {
			p.lastInteractionTime = time.Now().Unix()

			p.buttonPressedUp(button)

			p.prepareForSleep()
		} else {
			p.buttonPressedDown(button)
		}
	})
}

func (p *ProdeckFramework) buttonPressedUp(button int) {
	image, ok := p.imageCache.images[button]
	if ok {
		p.SetButtonImage(button, image, false)
	}
}

func (p *ProdeckFramework) buttonPressedDown(button int) {
	p.outlineCurrentButton(button)
	if p.currentOS != nil {
		(*p.currentOS).ButtonDown(button)
	}
}