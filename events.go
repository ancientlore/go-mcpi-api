package mcpiapi

import (
	"strconv"
	"strings"
)

type Events object

type Block object

type Hit struct {
	PositionX   int
	PositionY   int
	PositionZ   int
	SurfaceX    int
	SurfaceY    int
	SurfaceZ    int
	BlockTypeId int
}

func (obj Events) Block() Block {
	return Block(obj)
}

func (obj Events) Clear() error {
	s := "events.clear()"
	return object(obj).send(s)
}

func (obj Block) Hits() (hits []Hit, err error) {
	s := "events.block.hits()"
	hits = make([]Hit, 0)
	var r string
	var i int64
	r, err = object(obj).sendReceive(s)
	if err != nil {
		return
	}
	if r == "" {
		return
	}
	harr := strings.Split(r, "|")
	for _, h := range harr {
		var hit Hit
		arr := strings.Split(h, ",")
		arr2 := make([]*int, 7)
		arr2[0] = &hit.PositionX
		arr2[1] = &hit.PositionY
		arr2[2] = &hit.PositionZ
		arr2[3] = &hit.SurfaceX
		arr2[4] = &hit.SurfaceY
		arr2[5] = &hit.SurfaceZ
		arr2[6] = &hit.BlockTypeId
		for index, rs := range arr {
			i, err = strconv.ParseInt(rs, 10, 32)
			if err != nil {
				return
			}
			*(arr2[index]) = int(i)
		}
		hits = append(hits, hit)
	}
	return
}
