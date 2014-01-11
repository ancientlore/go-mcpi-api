package mcpiapi

import ()

// PyramidHere draws a pyramid of the given height at the player's current location.
func PyramidHere(c Connection, height int) error {
	x, y, z, err := c.Player().GetTile()
	if err != nil {
		return err
	}

	err = Pyramid(c, x, y, z, height, SANDSTONE, 0, &PyramidSettings{Floor: true})
	if err != nil {
		return err
	}

	err = c.Player().SetTile(x, y+height+1, z)
	if err != nil {
		return err
	}

	return nil
}

// PyramidSettings is used to specify additional information controlling rendering a pyramid.
type PyramidSettings struct {
	Floor bool // Draws a floor under the pyramid
}

// Pyramid draws a pyramid of the given height at the specified location using the block type and
// settings provided.
func Pyramid(c Connection, x, y, z, height, blockTypeId, blockData int, settings *PyramidSettings) error {
	var err error
	dim := 0
	var iy int
	if settings == nil {
		settings = &PyramidSettings{}
	}
	for iy = height + y; iy >= y; iy-- {
		err = c.World().SetBlocks(x-dim, iy, z-dim, x-dim, iy, z+dim, SANDSTONE, 0)
		if err != nil {
			return err
		}
		err = c.World().SetBlocks(x+dim, iy, z-dim, x+dim, iy, z+dim, SANDSTONE, 0)
		if err != nil {
			return err
		}
		err = c.World().SetBlocks(x-dim, iy, z-dim, x+dim, iy, z-dim, SANDSTONE, 0)
		if err != nil {
			return err
		}
		err = c.World().SetBlocks(x-dim, iy, z+dim, x+dim, iy, z+dim, SANDSTONE, 0)
		if err != nil {
			return err
		}
		dim++
	}
	if settings.Floor {
		err = c.World().SetBlocks(x-dim, iy, z-dim, x+dim, iy, z+dim, SANDSTONE, 0)
		if err != nil {
			return err
		}
	}
	return nil
}
