package connector

type OSConnector interface {
	ButtonDown(button int)
}

type FrameworkConnector interface {
	SetBrightness(value int)
}

type OSBuilder func() *OSConnector