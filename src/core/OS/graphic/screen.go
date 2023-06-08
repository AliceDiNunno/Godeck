package graphic

import (
	"github.com/davecgh/go-spew/spew"
	"image/color"
)

type ScreenInteractor interface {
	GetButtonColor(caller *Screen, x int, y int) color.RGBA
	ButtonPressed(caller *Screen, x int, y int)
	Name() string
}

type Screen struct {
	Width  int
	Height int

	Interactor ScreenInteractor
	Layout     *Layout

	buttonsWaitingForRedraw []Point
}

func (s *Screen) ButtonPressed(number int) {
	x, y := s.ButtonPosition(number)

	//Check if the button is in the layout
	if s.Layout != nil {
		if !s.Layout.isSpaceAvailable(Point{x, y}) {
			s.Layout.ButtonPressed(x, y)
			return
		}
	}

	if s.Interactor != nil {
		s.Interactor.ButtonPressed(s, x, y)
	}
}

func (s *Screen) NumberOfButtons() int {
	return s.Width * s.Height
}

func (s *Screen) ButtonPosition(button int) (int, int) {
	return button % s.Width, button / s.Width
}

func (s *Screen) ButtonNumber(x int, y int) int {
	return y*s.Width + x
}

func (s *Screen) AskForRedraw(x int, y int) {
	spew.Dump(s)
	if s.buttonsWaitingForRedraw == nil {
		s.buttonsWaitingForRedraw = []Point{}
	}
	//check if the button is already waiting for redraw
	for _, point := range s.buttonsWaitingForRedraw {
		if point.X == x && point.Y == y {
			return
		}
	}
	s.buttonsWaitingForRedraw = append(s.buttonsWaitingForRedraw, Point{X: x, Y: y})
	spew.Dump(s.buttonsWaitingForRedraw)
}

func (s *Screen) ButtonsWaitingForRedraw() []Point {
	spew.Dump(s.buttonsWaitingForRedraw)

	var name = "root"

	selfButtons := []Point{}

	if s.Interactor != nil {
		name = s.Interactor.Name()
	}
	if s.buttonsWaitingForRedraw != nil {
		selfButtons = s.buttonsWaitingForRedraw
	}

	spew.Dump(name, selfButtons)

	if s.Layout != nil {
		layoutButtons := s.Layout.ButtonsWaitingForRedraw()
		spew.Dump("layout", layoutButtons)

		selfButtons = append(selfButtons, layoutButtons...)
	}

	return selfButtons
}

func (s *Screen) SetLayout(layout *Layout) {
	s.Layout = layout
}