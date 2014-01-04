package mcpiapi

import (
	"fmt"
)

type Camera object

func (obj Camera) Mode() Mode {
	return Mode(obj)
}

type Mode object

func (obj Mode) SetNormal() error {
	s := "camera.mode.setNormal()"
	return object(obj).send(s)
}

func (obj Mode) SetThirdPerson() error {
	s := "camera.mode.setThirdPerson()"
	return object(obj).send(s)
}

func (obj Mode) SetFixed() error {
	s := "camera.mode.setFixed()"
	return object(obj).send(s)
}

func (obj Mode) SetPos(x, y, z int) error {
	s := fmt.Sprintf("camera.mode.setPos(%d,%d,%d)", x, y, z)
	return object(obj).send(s)
}
