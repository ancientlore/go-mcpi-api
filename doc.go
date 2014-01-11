/*
This package enables you to use the Minecraft Pi Edition API from Go. Each connection uses a single socket
and commands are sent over channels, making it safe to use the API from different goroutines.

Example:

	var c mcpiapi.Connection
	c.Open("192.168.1.115")
	defer c.Close()

	go func() { c.Chat().Post("Hello, World!") }()
	go func() { c.Chat().Post("Hello again!") }()

	err = c.World().SetBlock(0, 0, 0, mcpiapi.GOLD_BLOCK, 0)
	if err != nil {
		log.Print(err)
	}

Block types are 0-255. See the Micraft Wiki for information on block values.

Block data is 0-15, and is used to specify extra characteristics like wool color.

Regarding the coordinate system, (0,0,0) is where the world was spawned and is sea level. (X,Z) represents
the ground plane while Y points to the sky.

Be sure and read the Minecraft Pi Edition API speificication which is included in the installation at:

	mcpi/api/spec/mcpi_protocol_spec.txt

See also:

http://pi.minecraft.net/
http://www.minecraftwiki.net/wiki/Data_values_(Pocket_Edition)
http://golang.org/
*/
package mcpiapi
