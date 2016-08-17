package mcpiapi

import (
	"fmt"
	"strconv"
	"strings"
)

// Player provides methods to manipulate the player ("Steve").
type Player struct{
	obj object
	name string
}

// GetTile returns the world coordinates of the player's location.
func (obj Player) GetTile() (x, y, z int, err error) {
	s := obj.name + ".getTile()"
	fmt.Printf("%s\n",s)
	x = 0
	y = 0
	z = 0
	var r string
	var i int64
	r, err = object(obj.obj).sendReceive(s)
	if err != nil {
		return
	}
	arr := strings.Split(r, ",")
	arr2 := make([]*int, 3)
	arr2[0] = &x
	arr2[1] = &y
	arr2[2] = &z
	for index, rs := range arr {
		i, err = strconv.ParseInt(rs, 10, 32)
		if err != nil {
			return
		}
		*(arr2[index]) = int(i)
	}
	return
}

// SetTile moves the player the given world coordinates.
func (obj Player) SetTile(x, y, z int) error {
	s := fmt.Sprintf("%s.setTile(%d,%d,%d)",obj.name, x, y, z)
	return object(obj.obj).send(s)
}

// GetPos returns the player's position. Note the the player position
// is more granular than a world position.
func (obj Player) GetPos() (xf, yf, zf float64, err error) {
	s := obj.name + ".getPos()"
	xf = 0.0
	yf = 0.0
	zf = 0.0
	var r string
	r, err = object(obj.obj).sendReceive(s)
	if err != nil {
		return
	}
	arr := strings.Split(r, ",")
	arr2 := make([]*float64, 3)
	arr2[0] = &xf
	arr2[1] = &yf
	arr2[2] = &zf
	for index, rs := range arr {
		*(arr2[index]), err = strconv.ParseFloat(rs, 64)
		if err != nil {
			return
		}
	}
	return
}

// SetPos sets the player's position. Note that the player position
// is more granular than a world position.
func (obj Player) SetPos(xf, yf, zf float64) error {
	s := fmt.Sprintf("%s.setPos(%f,%f,%f)",obj.name, xf, yf, zf)
	return object(obj.obj).send(s)
}
