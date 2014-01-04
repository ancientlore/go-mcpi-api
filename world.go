package mcpiapi

import (
	"fmt"
	"strconv"
)

type World object

func (obj World) Checkpoint() Checkpoint {
	return Checkpoint(obj)
}

func (obj World) GetBlock(x, y, z int) (blockTypeId int, err error) {
	s := fmt.Sprintf("world.getBlock(%d,%d,%d)", x, y, z)
	blockTypeId = 0
	var r string
	var i int64
	r, err = object(obj).sendReceive(s)
	if err != nil {
		return
	}
	i, err = strconv.ParseInt(r, 10, 32)
	if err != nil {
		return
	}
	blockTypeId = int(i)
	return
}

func (obj World) SetBlock(x, y, z, blockTypeId, blockData int) error {
	var s string
	if blockData < 0 {
		s = fmt.Sprintf("world.setBlock(%d,%d,%d,%d)", x, y, z, blockTypeId)
	} else {
		s = fmt.Sprintf("world.setBlock(%d,%d,%d,%d,%d)", x, y, z, blockTypeId, blockData)
	}
	return object(obj).send(s)
}

func (obj World) SetBlocks(x1, y1, z1, x2, y2, z2, blockTypeId, blockData int) error {
	var s string
	if blockData < 0 {
		s = fmt.Sprintf("world.setBlocks(%d,%d,%d,%d,%d,%d,%d)", x1, y1, z1, x2, y2, z2, blockTypeId)
	} else {
		s = fmt.Sprintf("world.setBlocks(%d,%d,%d,%d,%d,%d,%d,%d)", x1, y1, z1, x2, y2, z2, blockTypeId, blockData)
	}
	return object(obj).send(s)
}

func (obj World) GetHeight(x, z int) (y int, err error) {
	s := fmt.Sprintf("world.getHeight(%d,%d)", x, z)
	y = -1
	var r string
	var i int64
	r, err = object(obj).sendReceive(s)
	if err != nil {
		return
	}
	i, err = strconv.ParseInt(r, 10, 32)
	if err != nil {
		return
	}
	y = int(i)
	return
}

type Checkpoint object

func (obj Checkpoint) Save() error {
	s := "world.checkpoint.save()"
	return object(obj).send(s)
}

func (obj Checkpoint) Restore() error {
	s := "world.checkpoint.restore()"
	return object(obj).send(s)
}
