package simple

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type UUIDHex string

const hexFormat = "%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x"

func IsValidUUID(uuid string) (valid bool) {

	//
	// cd9b978c-0d8c-4448-b5c7-88a7f0287d7d
	// 012345678901234567890123456789012345
	//

	if len(uuid) != 36 {
		return
	}

	_, tryErr := UuidFromHexString(UUIDHex(uuid))
	valid = tryErr == nil

	return
}

func NewUuidHex() (uuid UUIDHex) {
	uuid = UuidToHexString(NewUuid())
	return
}

func NewUuid() (uuid []byte) {

	uuid = make([]byte, 16)
	_, _ = rand.Read(uuid)

	// variant bits; see section 4.1.1

	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3

	uuid[6] = uuid[6]&^0xf0 | 0x40

	return
}

func UuidToBase64(uuid []byte) (uuid64 string) {

	uuid64 = base64.StdEncoding.EncodeToString(uuid)
	return
}

func UuidFromBase64(uuid64 string) (uuid []byte, err error) {

	uuid, err = base64.StdEncoding.DecodeString(uuid64)
	return
}

func ZeroUuidHex() (uuidHex UUIDHex) {

	uuidHex = UuidToHexString(make([]byte, 16))
	return
}

func UuidToHexString(uuid []byte) (uuidHex UUIDHex) {

	//
	// 123e4567-e89b-12d3-a456-426655440000
	//

	uuidHex = UUIDHex(fmt.Sprintf(hexFormat,
		uuid[0], uuid[1], uuid[2], uuid[3],
		uuid[4], uuid[5],
		uuid[6], uuid[7],
		uuid[8], uuid[9],
		uuid[10], uuid[11], uuid[12], uuid[13], uuid[14], uuid[15]))

	return
}

func UuidFromHexString(uuidHex UUIDHex) (uuid []byte, err error) {

	uuid = make([]byte, 16)

	_, err = fmt.Sscanf(string(uuidHex), hexFormat,
		&uuid[0], &uuid[1], &uuid[2], &uuid[3],
		&uuid[4], &uuid[5],
		&uuid[6], &uuid[7],
		&uuid[8], &uuid[9],
		&uuid[10], &uuid[11], &uuid[12], &uuid[13], &uuid[14], &uuid[15])

	return
}

func UuidHexFromSha256(data []byte) (uuid UUIDHex) {

	hasher := sha256.New()
	hasher.Write(data)
	sum := hasher.Sum(nil)

	uuid = UuidToHexString(sum[:16])
	return
}
