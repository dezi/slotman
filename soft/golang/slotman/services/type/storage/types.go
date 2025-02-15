package storage

import (
	"slotman/utils/simple"
	"time"
)

type Meta interface {
	GetUuid() (uuid *simple.UUIDHex)
	GetTime() (time *time.Time)
	SetTime(time *time.Time)
	GetDay() (want bool)
	GetSub() (sub string)
	GetTag() (tag string)
}

type File interface {
	GetUuid() (uuid *simple.UUIDHex)
	GetTime() (time *time.Time)
	SetTime(time *time.Time)
	GetDay() (want bool)
	GetSub() (sub string)
	GetTag() (tag string)
	GetFext() (fext string)
	GetData() (data []byte)
	SetData(data []byte)
}
