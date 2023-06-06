package OS

import (
	"godeck/src/core/connector"
)

type ProdeckOS struct {
	connector connector.FrameworkConnector

	brightness int
}

func mapval(x int, in_min int, in_max int, out_min int, out_max int) int {
	return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min
}

func (p *ProdeckOS) ButtonDown(button int) {
	newBrightness := mapval(button, 0, 32, 0, 100)
	p.connector.SetBrightness(newBrightness)
}

func StartOS(connector connector.FrameworkConnector) connector.OSConnector {
	proos := ProdeckOS{
		connector: connector,

		brightness: 1,
	}

	println("Starting OS")

	connector.SetBrightness(proos.brightness)

	return &proos
}