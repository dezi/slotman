package things

import "slotman/utils/simple"

var (
	ThingSystemUuid = simple.ZeroUuidHex()
)

func SetGlobalThingSystemUuid(uuid simple.UUIDHex) {
	ThingSystemUuid = uuid
}

func GetGlobalThingSystemUuid() (uuid simple.UUIDHex) {
	uuid = ThingSystemUuid
	return
}
