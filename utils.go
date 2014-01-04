package mcpiapi

import ()

func Pyramid(c Connection, height int) error {
	x, y, z, err := c.Player().GetTile()
	if err != nil {
		return err
	}

	dim := 0

	for iy := height + y; iy >= y; iy-- {
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
	err = c.Player().SetTile(x, y+height+1, z)
	if err != nil {
		return err
	}
	return nil
}
