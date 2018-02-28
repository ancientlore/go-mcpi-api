package mcpiapi

import (
	"fmt"
	"strconv"
)

// World has methods for manipulating the Minecraft world.
type World object

// Checkpoint provides access to the checkpoint object for this world.
func (obj World) Checkpoint() Checkpoint {
	return Checkpoint(obj)
}

// GetBlock returns the block type at the given coordinates.
// Block types can be 0-108.
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

// SetBlock sets the block type and block data at the given coordinate.
// Block types are 0-255 and block data can be 0-15. Block data represents
// extra attributes like the wool color.
func (obj World) SetBlock(x, y, z, blockTypeId, blockData int) error {
	var s string
	if blockData < 0 {
		s = fmt.Sprintf("world.setBlock(%d,%d,%d,%d)", x, y, z, blockTypeId)
	} else {
		s = fmt.Sprintf("world.setBlock(%d,%d,%d,%d,%d)", x, y, z, blockTypeId, blockData)
	}
	return object(obj).send(s)
}

// SetBlocks sets a range of blocks to the block type and block data provided.
// Block types are 0-255 and block data can be 0-15. Block data represents
// extra attributes like the wool color.
// Note: Setting a huge number of blocks can cause lag in the Minecraft game.
func (obj World) SetBlocks(x1, y1, z1, x2, y2, z2, blockTypeId, blockData int) error {
	var s string
	if blockData < 0 {
		s = fmt.Sprintf("world.setBlocks(%d,%d,%d,%d,%d,%d,%d)", x1, y1, z1, x2, y2, z2, blockTypeId)
	} else {
		s = fmt.Sprintf("world.setBlocks(%d,%d,%d,%d,%d,%d,%d,%d)", x1, y1, z1, x2, y2, z2, blockTypeId, blockData)
	}
	return object(obj).send(s)
}

// GetHeight returns the height of the ground at the given coordinate.
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

// Setting is used to enable or disable various Minecraft world settings.
func (obj World) Setting(key string, enable bool) error {
	var s string
	val := 0
	if enable {
		val = 1
	}
	s = fmt.Sprintf("world.setting(\"%s\",%d)", key, val)
	return object(obj).send(s)
}

// Checkpoint has methods for managing saving and restoring the world's state.
type Checkpoint object

// Save saves the world's current state so that it can be restored later.
func (obj Checkpoint) Save() error {
	s := "world.checkpoint.save()"
	return object(obj).send(s)
}

// Restore restores the world's state to that of the last checkpoint.
func (obj Checkpoint) Restore() error {
	s := "world.checkpoint.restore()"
	return object(obj).send(s)
}
