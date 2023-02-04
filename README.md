go-mcpi-api
===========

[![GoDoc](https://godoc.org/github.com/ancientlore/go-mcpi-api?status.svg)](https://godoc.org/github.com/ancientlore/go-mcpi-api)

This library enables you to use the [Minecraft Pi Edition](http://pi.minecraft.net/) API from [Go](http://golang.org/). Each connection uses a single socket and commands are sent over channels, making it safe to use the API from different goroutines.

中文版本README: [用go語言翻轉Minecraft - go-mcpi-api](http://holacode.blogspot.tw/2016/08/go-gominecraft-go-mcpi-api.html)

Also see [FrozenKP's version](https://github.com/FrozenKP/go-mcpi-api).

## FAQ

**What is MCPI ?** 

> MCPI is an API which can let you control Minecraft with code.

> It's original edition was written with python; you can learn it [here](https://github.com/teachthenet/TeachCraft-Challenges).

**Can I using it with a normal edition Minecraft ?**

> Yes, you can. You can refer to [this project](https://github.com/teachthenet/TeachCraft-Server).

## Example

	var c mcpiapi.Connection
	c.Open("192.168.1.115", "player1")
	defer c.Close()

	go func() { c.Chat().Post("Hello, World!") }()
	go func() { c.Chat().Post("Hello again!") }()

	err = c.World().SetBlock(0, 0, 0, mcpiapi.GOLD_BLOCK, 0)
	if err != nil {
		log.Print(err)
	}

## Pocket Edition Information

Block types are 0-108. See the [Minecraft Wiki](http://www.minecraftwiki.net/wiki/Data_values_(Pocket_Edition)) for information on block values.

Block data is 0-15, and is used to specify extra characteristics like wool color.

Regarding the coordinate system, (0,0,0) is where the world was spawned and is sea level. (X,Z) represents the ground plane while Y points to the sky.

Be sure and read the [Minecraft Pi Edition](http://pi.minecraft.net/) API speificication which is included in the installation at `mcpi/api/spec/mcpi_protocol_spec.txt`.

The library includes some utility functions for shapes - more may be added over time.

NOTE: Events have not been tested.

## Installation

    go get github.com/ancientlore/go-mcpi-api
