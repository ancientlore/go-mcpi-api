package mcpiapi

import (
	"fmt"
)

// Camera has methods for manipulating the scene displayed to the user.
type Camera object

// Mode returns the mode object for the camera
func (obj Camera) Mode() Mode {
	return Mode(obj)
}

// Mode has methods for manipulating the camera mode.
type Mode object

// SetNormal sets the camera mode to the normal first-person view.
func (obj Mode) SetNormal() error {
	s := "camera.mode.setNormal()"
	return object(obj).send(s)
}

// SetThirdPerson sets the camera to the third-person view.
func (obj Mode) SetThirdPerson() error {
	s := "camera.mode.setThirdPerson()"
	return object(obj).send(s)
}

// SetFixed sets the camera to a fixed view. Use SetPos to move it around.
func (obj Mode) SetFixed() error {
	s := "camera.mode.setFixed()"
	return object(obj).send(s)
}

// SetPos moves the camera to the given coordinate.
func (obj Mode) SetPos(x, y, z int) error {
	s := fmt.Sprintf("camera.mode.setPos(%d,%d,%d)", x, y, z)
	return object(obj).send(s)
}
