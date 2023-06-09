package OS

import (
	log "github.com/sirupsen/logrus"
	"godeck/src/core/OS/config"
	"godeck/src/core/connector"
)

type ProdeckOS struct {
	connector connector.FrameworkConnector

	serial        string
	screenManager *ScreenManager

	buttonsState map[int]bool
}

func (p *ProdeckOS) ButtonLongPressed(button int) {
	p.buttonsState[button] = false
	log.Println("Button long pressed")
}

func (p *ProdeckOS) SleepEntered() {
	println("lock screen")
	p.screenManager.lockScreen.isLocked = true
}

func (p *ProdeckOS) SleepExited() {

}

func (p *ProdeckOS) ButtonDown(button int) {
	p.buttonsState[button] = true
	p.screenManager.ButtonPressed(button)
}

func (p *ProdeckOS) ButtonUp(button int) {
	p.buttonsState[button] = false
}

func initOS(connector connector.FrameworkConnector) *ProdeckOS {
	serial := connector.GetSerialNumber()
	log.WithFields(log.Fields{
		"serial": serial,
	}).Println("Init OS")

	cfg := config.NewConfig()
	deviceConfig := cfg.GetDeviceConfig(serial)
	_ = deviceConfig
	proos := ProdeckOS{
		connector:    connector,
		serial:       serial,
		buttonsState: map[int]bool{},
	}

	for i := 0; i < connector.GetHeight()*connector.GetWidth(); i++ {
		proos.buttonsState[i] = false
	}

	proos.createScreenManager(connector.GetHeight(), connector.GetWidth(), deviceConfig.Brightness)

	connector.SetBrightness(proos.screenManager.brightness)

	return &proos
}

func StartOS(connector connector.FrameworkConnector) *ProdeckOS {
	proos := initOS(connector)

	return proos
}