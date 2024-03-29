package Framework

import (
	log "github.com/sirupsen/logrus"
	eventDomain "godeck/src/domain/events"
	"time"
)

func (p *ProdeckFramework) setupSleepEvent() {
	p.eventHub.Subscribe(eventDomain.DeviceWillSleepEvent, func(topic eventDomain.Event, data eventDomain.EventData) {
		log.Println("Device will sleep")

		p.isSleeping = true
		go p.setBrightness(0, true, false)
		(*p.currentOS).SleepEntered()

		p.eventHub.CancelPublishLater(eventDomain.DeviceWillSleepEvent)
	})
}

func (p *ProdeckFramework) prepareForSleep() {
	log.Println("Device will sleep in 5 minutes")
	p.eventHub.PublishLater(eventDomain.DeviceWillSleepEvent, nil, 10*time.Second)
}

func (p *ProdeckFramework) wakeUpFromSleep() bool {
	if p.isSleeping {
		p.isSleeping = false
		go p.setBrightness(p.currentBrightness, true, false)
		(*p.currentOS).SleepExited()
		return true
	}

	p.eventHub.CancelPublishLater(eventDomain.DeviceWillSleepEvent)
	return false
}