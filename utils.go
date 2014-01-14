package mcpiapi

import (
	"math"
)

// PyramidHere draws a pyramid of the given height at the player's current location.
func PyramidHere(c Connection, height int) error {
	x, y, z, err := c.Player().GetTile()
	if err != nil {
		return err
	}

	err = Pyramid(c, x, y, z, height, SANDSTONE, 2, &PyramidSettings{Floor: true})
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
	Floor            bool // Draws a floor under the pyramid
	FloorBlockTypeId int  // Block ID to use for the floor
	FloorBlockData   int  // Block data to use for the floor
	ClearInside      bool // whether to clear the blocks from the inside of the pyramid
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
	// for stairs, automatically adjust the orientation
	d1, d2, d3, d4 := blockData, blockData, blockData, blockData
	if IsStairs(blockTypeId) {
		d1 = 0
		d2 = 1
		d3 = 2
		d4 = 3
	}
	for iy = height + y; iy >= y; iy-- {
		if dim == 0 {
			if IsStairs(blockTypeId) {
				err = c.World().SetBlocks(x, iy, z, x, iy, z, BaseMaterial(blockTypeId), 0)
				if err != nil {
					return err
				}
			} else {
				err = c.World().SetBlocks(x, iy, z, x, iy, z, blockTypeId, blockData)
				if err != nil {
					return err
				}
			}
		} else {
			err = c.World().SetBlocks(x-dim, iy, z-dim, x-dim, iy, z+dim, blockTypeId, d1)
			if err != nil {
				return err
			}
			err = c.World().SetBlocks(x+dim, iy, z-dim, x+dim, iy, z+dim, blockTypeId, d2)
			if err != nil {
				return err
			}
			err = c.World().SetBlocks(x-dim, iy, z-dim, x+dim, iy, z-dim, blockTypeId, d3)
			if err != nil {
				return err
			}
			err = c.World().SetBlocks(x-dim, iy, z+dim, x+dim, iy, z+dim, blockTypeId, d4)
			if err != nil {
				return err
			}
			if settings.ClearInside {
				err = c.World().SetBlocks(x-dim+1, iy, z-dim+1, x+dim-1, iy, z+dim-1, AIR, 0)
				if err != nil {
					return err
				}
			}
		}
		dim++
	}
	if settings.Floor {
		if settings.FloorBlockTypeId <= 0 {
			settings.FloorBlockTypeId = blockTypeId
			settings.FloorBlockData = blockData
			if IsStairs(settings.FloorBlockTypeId) {
				settings.FloorBlockTypeId = BaseMaterial(settings.FloorBlockTypeId)
				settings.FloorBlockData = 0
			}
		}
		err = c.World().SetBlocks(x-dim, iy, z-dim, x+dim, iy, z+dim, settings.FloorBlockTypeId, settings.FloorBlockData)
		if err != nil {
			return err
		}
	}
	return nil
}

// DegToRad converts degrees to radians.
func DegToRad(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// Arc draws an arc of radius r centered at (x, y, z) within the angles specified.
func Arc(c Connection, x, y, z, r int, xzStartRads, xyStartRads, xzRads, xyRads float64, blockTypeId, blockData int) error {

	// radians delta = 2 * pi / (2 * pi * r)
	deltaRad := 1.0 / float64(r)

	dxz := math.Min(xzStartRads, xzRads)
	dxy := math.Min(xyStartRads, xyRads)
	for dxz <= math.Max(xzStartRads, xzRads) || dxy <= math.Max(xyStartRads, xyRads) {
		dx := x + int(float64(r)*math.Cos(dxy)*math.Cos(dxz))
		dy := y + int(float64(r)*math.Sin(dxy))
		dz := z + int(float64(r)*math.Cos(dxy)*math.Sin(dxz))
		err := c.World().SetBlock(dx, dy, dz, blockTypeId, blockData)
		if err != nil {
			return err
		}
		if dxz <= math.Max(xzStartRads, xzRads) {
			dxz += deltaRad
		}
		if dxy <= math.Max(xyStartRads, xyRads) {
			dxy += deltaRad
		}
	}
	return nil
}

// Sphere draws a sphere of the given radius r around the coordinates (x, y, z).
func Sphere(c Connection, x, y, z, r, blockTypeId, blockData int) error {
	delta := 1.0 / float64(r)
	for d := 0.0; d <= 2.0*math.Pi; d += delta {
		err := Arc(c, x, y, z, r, 0.0, d, 2.0*math.Pi, d, blockTypeId, blockData)
		if err != nil {
			return err
		}
	}
	return nil
}
