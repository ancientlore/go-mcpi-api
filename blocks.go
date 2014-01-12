package mcpiapi

// Block types
const (
	AIR                  = 0
	STONE                = 1
	GRASS                = 2
	DIRT                 = 3
	COBBLESTONE          = 4
	WOOD_PLANKS          = 5
	SAPLING              = 6
	BEDROCK              = 7
	WATER_FLOWING        = 8
	WATER                = WATER_FLOWING
	WATER_STATIONARY     = 9
	LAVA_FLOWING         = 10
	LAVA                 = LAVA_FLOWING
	LAVA_STATIONARY      = 11
	SAND                 = 12
	GRAVEL               = 13
	GOLD_ORE             = 14
	IRON_ORE             = 15
	COAL_ORE             = 16
	WOOD                 = 17
	LEAVES               = 18
	SPONGE               = 19 // added
	GLASS                = 20
	LAPIS_LAZULI_ORE     = 21
	LAPIS_LAZULI_BLOCK   = 22
	SANDSTONE            = 24
	BED                  = 26
	POWERED_RAIL         = 27 // added
	COBWEB               = 30
	GRASS_TALL           = 31
	DEAD_BUSH            = 32 // adeed
	WOOL                 = 35
	FLOWER_YELLOW        = 37
	FLOWER_CYAN          = 38
	MUSHROOM_BROWN       = 39
	MUSHROOM_RED         = 40
	GOLD_BLOCK           = 41
	IRON_BLOCK           = 42
	STONE_SLAB_DOUBLE    = 43
	STONE_SLAB           = 44
	BRICK_BLOCK          = 45
	TNT                  = 46
	BOOKSHELF            = 47
	MOSS_STONE           = 48
	OBSIDIAN             = 49
	TORCH                = 50
	FIRE                 = 51
	STAIRS_WOOD          = 53
	CHEST                = 54
	DIAMOND_ORE          = 56
	DIAMOND_BLOCK        = 57
	CRAFTING_TABLE       = 58
	SEEDS                = 59 // added
	FARMLAND             = 60
	FURNACE_INACTIVE     = 61
	FURNACE_ACTIVE       = 62
	SIGN_POST            = 63 // added
	DOOR_WOOD            = 64
	LADDER               = 65
	RAIL                 = 66 // added
	STAIRS_COBBLESTONE   = 67
	WALL_SIGN            = 68 // added
	DOOR_IRON            = 71
	REDSTONE_ORE         = 73
	GLOWING_REDSTONE_ORE = 74 // added
	SNOW                 = 78
	ICE                  = 79
	SNOW_BLOCK           = 80
	CACTUS               = 81
	CLAY                 = 82
	SUGAR_CANE           = 83
	FENCE                = 85
	PUMPKIN              = 86 // adeed
	NETHERRACK           = 87 // added
	GLOWSTONE_BLOCK      = 89
	JACK_O_LANTERN       = 91 // added
	CAKE_BLOCK           = 92 // added
	BEDROCK_INVISIBLE    = 95
	TRAP_DOOR            = 96 // added
	STONE_BRICK          = 98
	IRON_BARS            = 101
	GLASS_PANE           = 102
	MELON                = 103
	PUMPKIN_STEM         = 104 // added
	MELON_STEM           = 105 // added
	FENCE_GATE           = 107
	BRICK_STAIRS         = 108 // added
	STONE_BRICK_STAIRS   = 109 // added
	NETHER_BRICK         = 112 // added
	NETHER_BRICK_STAIRS  = 114 // added
	SANDSTONE_STAIRS     = 128 // added
	EMERALD_ORE          = 129 // added
	SPRUCE_WOOD_STAIRS   = 134 // added
	BIRCH_WOOD_STAIRS    = 135 // added
	JUNGLE_WOOD_STAIRS   = 136 // added
	COBBLESTONE_WALL     = 139 // added
	CARROTS              = 141 // added
	POTATO               = 142 // added
	QUARTZ_BLOCK         = 155 // added
	QUARTZ_STAIRS        = 156 // added
	WOODEN_DOUBLE_SLAB   = 157 // added
	WOODEN_SLAB          = 158 // added
	HAY_BLOCK            = 170 // added
	CARPET               = 171 // added
	BLOCK_OF_COAL        = 173 // added
	BEETROOT             = 244 // added
	STONE_CUTTER         = 245 // added
	GLOWING_OBSIDIAN     = 246
	NETHER_REACTOR_CORE  = 247
	UPDATE_GAME_BLOCK_1  = 248 // added
	UPDATE_GAME_BLOCK_2  = 249 // added
	LAST                 = 255 // added
)

// IsWater returns true if the given block type is any kind of water.
func IsWater(blockTypeId int) bool {
	switch blockTypeId {
	case WATER_FLOWING, WATER_STATIONARY:
		return true
	default:
		return false
	}
}

// IsLava returns true if the given block type is any kind of lava.
func IsLava(blockTypeId int) bool {
	switch blockTypeId {
	case LAVA_FLOWING, LAVA_STATIONARY:
		return true
	default:
		return false
	}
}

// IsFire returns true if the given block type is fire.
func IsFire(blockTypeId int) bool {
	switch blockTypeId {
	case FIRE:
		return true
	default:
		return false
	}
}

// IsAir returns true if the given block type is air.
func IsAir(blockTypeId int) bool {
	switch blockTypeId {
	case AIR:
		return true
	default:
		return false
	}
}

// IsStairs returns true if the given block type is any kind of stairs.
func IsStairs(blockTypeId int) bool {
	switch blockTypeId {
	case STAIRS_WOOD, STAIRS_COBBLESTONE, BRICK_STAIRS, STONE_BRICK_STAIRS, NETHER_BRICK_STAIRS, SANDSTONE_STAIRS, SPRUCE_WOOD_STAIRS, BIRCH_WOOD_STAIRS, JUNGLE_WOOD_STAIRS, QUARTZ_STAIRS:
		return true
	default:
		return false
	}
}

// IsSlab returns true if the given block type is any kind of slab, which is half-height.
func IsSlab(blockTypeId int) bool {
	switch blockTypeId {
	case STONE_SLAB, WOODEN_SLAB:
		return true
	default:
		return false
	}
}

// BaseMaterial returns a block ID of the base material for the given type.
// The intent of the function is to determine the base type of stairs and
// other manufactured objects.
func BaseMaterial(blockTypeId int) int {
	switch blockTypeId {
	case AIR:
	case STONE:
	case GRASS:
		return DIRT
	case DIRT:
	case COBBLESTONE:
	case WOOD_PLANKS:
		return WOOD
	case SAPLING:
	case BEDROCK:
	case WATER_FLOWING:
	case WATER_STATIONARY:
	case LAVA_FLOWING:
	case LAVA_STATIONARY:
	case SAND:
	case GRAVEL:
	case GOLD_ORE:
	case IRON_ORE:
	case COAL_ORE:
	case WOOD:
	case LEAVES:
	case SPONGE:
	case GLASS:
	case LAPIS_LAZULI_ORE:
	case LAPIS_LAZULI_BLOCK:
	case SANDSTONE:
	case BED:
	case POWERED_RAIL:
	case COBWEB:
	case GRASS_TALL:
	case DEAD_BUSH:
	case WOOL:
	case FLOWER_YELLOW:
	case FLOWER_CYAN:
	case MUSHROOM_BROWN:
	case MUSHROOM_RED:
	case GOLD_BLOCK:
	case IRON_BLOCK:
	case STONE_SLAB_DOUBLE:
		return STONE
	case STONE_SLAB:
		return STONE
	case BRICK_BLOCK:
	case TNT:
	case BOOKSHELF:
	case MOSS_STONE:
		return STONE
	case OBSIDIAN:
	case TORCH:
	case FIRE:
	case STAIRS_WOOD:
		return WOOD
	case CHEST:
	case DIAMOND_ORE:
	case DIAMOND_BLOCK:
	case CRAFTING_TABLE:
	case SEEDS:
	case FARMLAND:
	case FURNACE_INACTIVE:
	case FURNACE_ACTIVE:
	case SIGN_POST:
		return WOOD
	case DOOR_WOOD:
		return WOOD
	case LADDER:
		return WOOD
	case RAIL:
	case STAIRS_COBBLESTONE:
		return COBBLESTONE
	case WALL_SIGN:
		return WOOD
	case DOOR_IRON:
		return IRON_BLOCK
	case REDSTONE_ORE:
	case GLOWING_REDSTONE_ORE:
	case SNOW:
	case ICE:
	case SNOW_BLOCK:
	case CACTUS:
	case CLAY:
	case SUGAR_CANE:
	case FENCE:
		return WOOD
	case PUMPKIN:
	case NETHERRACK:
	case GLOWSTONE_BLOCK:
	case JACK_O_LANTERN:
	case CAKE_BLOCK:
	case BEDROCK_INVISIBLE:
	case TRAP_DOOR:
		return WOOD
	case STONE_BRICK:
	case IRON_BARS:
	case GLASS_PANE:
		return GLASS
	case MELON:
	case PUMPKIN_STEM:
	case MELON_STEM:
	case FENCE_GATE:
		return WOOD
	case BRICK_STAIRS:
		return BRICK_BLOCK
	case STONE_BRICK_STAIRS:
		return STONE_BRICK
	case NETHER_BRICK:
	case NETHER_BRICK_STAIRS:
		return NETHER_BRICK
	case SANDSTONE_STAIRS:
		return SANDSTONE
	case EMERALD_ORE:
	case SPRUCE_WOOD_STAIRS:
		return WOOD
	case BIRCH_WOOD_STAIRS:
		return WOOD
	case JUNGLE_WOOD_STAIRS:
		return WOOD
	case COBBLESTONE_WALL:
		return COBBLESTONE
	case CARROTS:
	case POTATO:
	case QUARTZ_BLOCK:
	case QUARTZ_STAIRS:
		return QUARTZ_BLOCK
	case WOODEN_DOUBLE_SLAB:
		return WOOD
	case WOODEN_SLAB:
		return WOOD
	case HAY_BLOCK:
	case CARPET:
	case BLOCK_OF_COAL:
	case BEETROOT:
	case STONE_CUTTER:
	case GLOWING_OBSIDIAN:
	case NETHER_REACTOR_CORE:
	case UPDATE_GAME_BLOCK_1:
	case UPDATE_GAME_BLOCK_2:
	case LAST:
	}
	return blockTypeId
}
