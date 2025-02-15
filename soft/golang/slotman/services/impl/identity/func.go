package identity

import "slotman/utils/simple"

func (sv *Service) GetBoxTag() (boxTag string) {
	return sv.globalBoxTag
}

func (sv *Service) GetBoxIdentity() (boxIdentity simple.UUIDHex) {
	boxIdentity = sv.globalBoxIdentity
	return
}

func (sv *Service) GetStoragePath() (storagePath string) {
	storagePath = sv.globalStoragePath + "/" + string(sv.globalBoxIdentity)
	return
}
