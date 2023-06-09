package graphic

import (
	"image/color"
)

// Even though the ScreenInteractor may be the same object as the Screen, it may not be the same object memory-wise so we need to pass the pointer of the Screen
// This may be sub-optimal
type ScreenInteractor interface {
	GetButtonColor(caller *Screen, x int, y int) color.RGBA
	ButtonPressed(caller *Screen, x int, y int)
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
}

func (s *Screen) ButtonsWaitingForRedraw() []Point {
	selfButtons := []Point{}

	if s.buttonsWaitingForRedraw != nil {
		selfButtons = s.buttonsWaitingForRedraw
	}

	if s.Layout != nil {
		layoutButtons := s.Layout.ButtonsWaitingForRedraw()

		selfButtons = append(selfButtons, layoutButtons...)
	}

	return selfButtons
}

func (s *Screen) SetLayout(layout *Layout) {
	s.Layout = layout
}

func (s *Screen) GetEntryColor(x int, y int) color.Color {
	if s.Layout != nil {
		return s.Layout.GetEntryColor(x, y)
	}

	//remove from the list of buttons waiting for redraw
	for i, point := range s.buttonsWaitingForRedraw {
		if point.X == x && point.Y == y {
			s.buttonsWaitingForRedraw = append(s.buttonsWaitingForRedraw[:i], s.buttonsWaitingForRedraw[i+1:]...)
			break
		}
	}
	return s.Interactor.GetButtonColor(s, x, y)
}