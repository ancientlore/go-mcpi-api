package mcpiapi

import (
	"fmt"
	"strconv"
	"strings"
)

type Player object

func (obj Player) GetTile() (x, y, z int, err error) {
	s := "player.getTile()"
	x = 0
	y = 0
	z = 0
	var r string
	var i int64
	r, err = object(obj).sendReceive(s)
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

func (obj Player) SetTile(x, y, z int) error {
	s := fmt.Sprintf("player.setTile(%d,%d,%d)", x, y, z)
	return object(obj).send(s)
}

func (obj Player) GetPos() (xf, yf, zf float64, err error) {
	s := "player.getPos()"
	xf = 0.0
	yf = 0.0
	zf = 0.0
	var r string
	r, err = object(obj).sendReceive(s)
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

func (obj Player) SetPos(xf, yf, zf float64) error {
	s := fmt.Sprintf("player.setPos(%f,%f,%f)", xf, yf, zf)
	return object(obj).send(s)
}
