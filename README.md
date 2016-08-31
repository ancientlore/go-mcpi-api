[![GoDoc](https://godoc.org/github.com/FrozenKP/go-mcpi-api?status.svg)](https://godoc.org/github.com/FrozenKP/go-mcpi-api)

中文版本README: [用go語言翻轉Minecraft - go-mcpi-api](http://holacode.blogspot.tw/2016/08/go-gominecraft-go-mcpi-api.html)

This is a library forked from [ancientlore](https://github.com/ancientlore/go-mcpi-api).

It can enable you to use the [Minecraft PI](http://pi.minecraft.net/) Edition API with [Go](https://golang.org/).

----------
**What is MCPI ?** 

> MCPI is a API which can let you control Minecraft with code.

> It's original edition was written with python, you can learn it [here](https://github.com/teachthenet/TeachCraft-Challenges).

**Can I using it at a normal edition Minecraft ?**

> Yes, you can. You can refer to [this project](https://github.com/teachthenet/TeachCraft-Server).

**What's difference between this project with the original one ?**

> The original project can only be used in single-player mode. 

> However, my edition can be used in both single-player mode and multiple-players mode.

**What's difference when using this project ?**

> The only difference is Open().

> You have to add the player name when using Open().

> Example : `c.Open("140.113.195.200","Frozen")`
	


----------
**Installation**

    go get github.com/FrozenKP/go-mcpi-api
