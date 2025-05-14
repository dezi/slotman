package things

// https://chatgpt.com/share/68245a4d-bc24-8012-8a0a-7176c8f44f70

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
