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
}

func (p *ProdeckOS) ButtonLongPressed(button int) {

}

func (p *ProdeckOS) SleepEntered() {

}

func (p *ProdeckOS) SleepExited() {

}

func (p *ProdeckOS) ButtonDown(button int) {
	p.screenManager.screen.ButtonPressed(button)
}

func (p *ProdeckOS) ButtonUp(button int) {

}

func initOS(connector connector.FrameworkConnector) *ProdeckOS {
	serial := connector.GetSerialNumber()
	log.WithFields(log.Fields{
		"serial": serial,
	}).Println("Init OS")

	config := config.NewConfig()
	deviceConfig := config.GetDeviceConfig(serial)
	_ = deviceConfig
	proos := ProdeckOS{
		connector: connector,
		serial:    serial,
	}
	proos.createScreenManager(connector.GetHeight(), connector.GetWidth(), deviceConfig.Brightness)

	connector.SetBrightness(proos.screenManager.brightness)

	return &proos
}

func StartOS(connector connector.FrameworkConnector) *ProdeckOS {
	proos := initOS(connector)

	return proos
}