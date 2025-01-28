package gc9a01

import (
	"bytes"
	"errors"
	"golang.org/x/image/draw"
	"image"
	"slotman/utils/log"
)

func (se *GC9A01) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *GC9A01) BlipFullImage(img image.Image) (err error) {

	if img.Bounds().Size().X != ScreenWidth ||
		img.Bounds().Size().Y != ScreenHeight {

		//
		// Resize image.
		//

		log.Printf("########### resize....")

		rgb := image.NewRGBA(image.Rect(0, 0, ScreenWidth, ScreenHeight))
		draw.BiLinear.Scale(rgb, rgb.Bounds(), img, img.Bounds(), draw.Src, nil)
		img = rgb
	}

	if img.Bounds().Size().X != ScreenWidth ||
		img.Bounds().Size().Y != ScreenHeight {
		err = errors.New("invalid image size")
		return
	}

	rgba, ok := img.(*image.RGBA)
	if !ok {
		err = errors.New("image not rgba")
		return
	}

	wid := rgba.Bounds().Size().X
	hei := rgba.Bounds().Size().Y
	pix := rgba.Pix
	raw := make([]byte, wid*hei*3)

	src := 0
	dst := 0

	for x := 0; x < wid; x++ {

		stride := src

		for y := 0; y < hei; y++ {

			raw[dst] = pix[stride]
			stride++
			dst++

			raw[dst] = pix[stride]
			stride++
			dst++

			raw[dst] = pix[stride]
			stride++
			dst++

			stride++
		}

		src += rgba.Stride
	}

	err = se.BlipFullImageRaw(raw)
	return
}

func (se *GC9A01) BlipFullImageRaw(image []byte) (err error) {

	se.blipLock.Lock()
	defer se.blipLock.Unlock()

	if len(image) != ScreenWidth*ScreenHeight*3 {
		return errors.New("invalid image size")
	}

	minY := 0
	maxY := ScreenHeight - 1

	if se.blipLast != nil {

		for minY = 0; minY < ScreenHeight; minY++ {
			start := minY * ScreenWidth * 3
			end := start + ScreenWidth*3
			if !bytes.Equal(image[start:end], se.blipLast[start:end]) {
				break
			}
		}

		for maxY = ScreenHeight - 1; maxY > minY; maxY-- {
			start := maxY * ScreenWidth * 3
			end := start + ScreenWidth*3
			if !bytes.Equal(image[start:end], se.blipLast[start:end]) {
				break
			}
		}

		if minY == ScreenHeight {
			return nil
		}
	}

	err = se.SetFrame(Frame{X0: 0, Y0: uint16(minY), X1: ScreenWidth - 1, Y1: uint16(maxY)})
	if err != nil {
		return err
	}

	chunkSize := ScreenWidth * 3
	for y := minY; y <= maxY; y++ {
		chunkPos := y * chunkSize
		if y == minY {
			err = se.writeMem(image[chunkPos : chunkPos+chunkSize])
		} else {
			err = se.writeMemCont(image[chunkPos : chunkPos+chunkSize])
		}
		if err != nil {
			return err
		}
	}

	se.blipLast = make([]byte, len(image))
	copy(se.blipLast, image)

	return
}

func (se *GC9A01) BlipFullImageRawClaudeV2(image []byte) (err error) {
	se.blipLock.Lock()
	defer se.blipLock.Unlock()

	if len(image) != ScreenWidth*ScreenHeight*3 {
		return errors.New("invalid image size")
	}

	var startY, endY int

	// Find the first different scanline from the top
	for startY = 0; startY < ScreenHeight; startY++ {
		start := startY * ScreenWidth * 3
		end := start + ScreenWidth*3
		if !bytes.Equal(image[start:end], se.blipLast[start:end]) {
			break
		}
	}

	// Find the first different scanline from the bottom
	for endY = ScreenHeight - 1; endY > startY; endY-- {
		start := endY * ScreenWidth * 3
		end := start + ScreenWidth*3
		if !bytes.Equal(image[start:end], se.blipLast[start:end]) {
			break
		}
	}

	// If the image is identical to the last one, return without updating
	if startY == ScreenHeight {
		return nil
	}

	// Set the frame to update only the changed area
	err = se.SetFrame(Frame{X0: 0, Y0: uint16(startY), X1: ScreenWidth - 1, Y1: uint16(endY)})
	if err != nil {
		return err
	}

	// Update only the changed portion of the image
	chunkSize := ScreenWidth * 3
	for y := startY; y <= endY; y++ {
		chunkPos := y * chunkSize
		if y == startY {
			err = se.writeMem(image[chunkPos : chunkPos+chunkSize])
		} else {
			err = se.writeMemCont(image[chunkPos : chunkPos+chunkSize])
		}
		if err != nil {
			return err
		}
	}

	// Update the last image
	se.blipLast = make([]byte, len(image))
	copy(se.blipLast, image)

	return nil
}

func (se *GC9A01) BlipFullImageRawClaudeV1(image []byte) (err error) {
	se.blipLock.Lock()
	defer se.blipLock.Unlock()

	if len(image) != ScreenWidth*ScreenHeight*3 {
		return errors.New("invalid image size")
	}

	// Find the top start index
	topStart := 0
	for i := 0; i < len(image); i += ScreenWidth * 3 {
		if !bytes.Equal(se.blipLast[i:i+ScreenWidth*3], image[i:i+ScreenWidth*3]) {
			topStart = i / (ScreenWidth * 3)
			break
		}
	}

	// Find the bottom end index
	bottomEnd := ScreenHeight - 1
	for i := len(image) - ScreenWidth*3; i >= 0; i -= ScreenWidth * 3 {
		if !bytes.Equal(se.blipLast[i:i+ScreenWidth*3], image[i:i+ScreenWidth*3]) {
			bottomEnd = i/(ScreenWidth*3) + 1
			break
		}
	}

	// If the image hasn't changed, return early
	if topStart >= bottomEnd {
		return nil
	}

	// Set the frame for the changed portion
	err = se.SetFrame(Frame{X0: 0, Y0: uint16(topStart), X1: ScreenWidth - 1, Y1: uint16(bottomEnd - 1)})
	if err != nil {
		return err
	}

	// Write the changed portion of the image
	startIndex := topStart * ScreenWidth * 3
	endIndex := bottomEnd * ScreenWidth * 3
	chunkSize := ScreenWidth * 4 * 3

	for chunkPos := startIndex; chunkPos < endIndex; chunkPos += chunkSize {
		endChunk := chunkPos + chunkSize
		if endChunk > endIndex {
			endChunk = endIndex
		}
		if chunkPos == startIndex {
			err = se.writeMem(image[chunkPos:endChunk])
		} else {
			err = se.writeMemCont(image[chunkPos:endChunk])
		}
		if err != nil {
			return err
		}
	}

	// Update blipLast with the new image
	se.blipLast = make([]byte, len(image))
	copy(se.blipLast, image)

	return nil
}

func (se *GC9A01) SetOrientation(orientation Orientation) (err error) {

	switch orientation {
	case 0:
		err = se.writeCommandBytes(CommandOrientation, byte(OrientMode0))
	case 1:
		err = se.writeCommandBytes(CommandOrientation, byte(OrientMode90))
	case 2:
		err = se.writeCommandBytes(CommandOrientation, byte(OrientMode180))
	case 3:
		err = se.writeCommandBytes(CommandOrientation, byte(OrientMode270))
	default:
		err = errors.New("wrong orientation")
	}

	return
}

func (se *GC9A01) SetFrame(frame Frame) (err error) {

	var data [4]byte

	data[0] = byte(frame.X0 >> 8)
	data[1] = byte(frame.X0)
	data[2] = byte(frame.X1 >> 8)
	data[3] = byte(frame.X1)

	_ = se.writeCommandBytes(CommandColAddrSet, data[:]...)

	data[0] = byte(frame.Y0 >> 8)
	data[1] = byte(frame.Y0)
	data[2] = byte(frame.Y1 >> 8)
	data[3] = byte(frame.Y1)

	_ = se.writeCommandBytes(CommandRowAddrSet, data[:]...)

	return
}

func (se *GC9A01) Initialize() (err error) {

	_ = se.writeCommandBytes(0xEF)
	_ = se.writeCommandBytes(0xEB, 0x14)
	_ = se.writeCommandBytes(0xFE)
	_ = se.writeCommandBytes(0xEF)
	_ = se.writeCommandBytes(0xEB, 0x14)
	_ = se.writeCommandBytes(0x84, 0x40)
	_ = se.writeCommandBytes(0x85, 0xFF)
	_ = se.writeCommandBytes(0x86, 0xFF)
	_ = se.writeCommandBytes(0x87, 0xFF)
	_ = se.writeCommandBytes(0x88, 0x0A)
	_ = se.writeCommandBytes(0x89, 0x21)
	_ = se.writeCommandBytes(0x8A, 0x00)
	_ = se.writeCommandBytes(0x8B, 0x80)
	_ = se.writeCommandBytes(0x8C, 0x01)
	_ = se.writeCommandBytes(0x8D, 0x01)
	_ = se.writeCommandBytes(0x8E, 0xFF)
	_ = se.writeCommandBytes(0x8F, 0xFF)
	_ = se.writeCommandBytes(0xB6, 0x00, 0x00)
	_ = se.writeCommandBytes(CommandOrientation, byte(OrientMode180))
	_ = se.writeCommandBytes(CommandColorMode, byte(ColorMode18Bit))
	_ = se.writeCommandBytes(0x90, 0x08, 0x08, 0x08, 0x08)
	_ = se.writeCommandBytes(0xBD, 0x06)
	_ = se.writeCommandBytes(0xBC, 0x00)
	_ = se.writeCommandBytes(0xFF, 0x60, 0x01, 0x04)
	_ = se.writeCommandBytes(0xC3, 0x13)
	_ = se.writeCommandBytes(0xC4, 0x13)
	_ = se.writeCommandBytes(0xC9, 0x22)
	_ = se.writeCommandBytes(0xBE, 0x11)
	_ = se.writeCommandBytes(0xE1, 0x10, 0x0E)
	_ = se.writeCommandBytes(0xDF, 0x21, 0x0c, 0x02)
	_ = se.writeCommandBytes(0xF0, 0x45, 0x09, 0x08, 0x08, 0x26, 0x2A)
	_ = se.writeCommandBytes(0xF1, 0x43, 0x70, 0x72, 0x36, 0x37, 0x6F)
	_ = se.writeCommandBytes(0xF2, 0x45, 0x09, 0x08, 0x08, 0x26, 0x2A)
	_ = se.writeCommandBytes(0xF3, 0x43, 0x70, 0x72, 0x36, 0x37, 0x6F)
	_ = se.writeCommandBytes(0xED, 0x1B, 0x0B)
	_ = se.writeCommandBytes(0xAE, 0x77)
	_ = se.writeCommandBytes(0xCD, 0x63)
	_ = se.writeCommandBytes(0x70, 0x07, 0x07, 0x04, 0x0E, 0x0F, 0x09, 0x07, 0x08, 0x03)
	_ = se.writeCommandBytes(0xE8, 0x34)
	_ = se.writeCommandBytes(0x62, 0x18, 0x0D, 0x71, 0xED, 0x70, 0x70, 0x18, 0x0F, 0x71, 0xEF, 0x70, 0x70)
	_ = se.writeCommandBytes(0x63, 0x18, 0x11, 0x71, 0xF1, 0x70, 0x70, 0x18, 0x13, 0x71, 0xF3, 0x70, 0x70)
	_ = se.writeCommandBytes(0x64, 0x28, 0x29, 0xF1, 0x01, 0xF1, 0x00, 0x07)
	_ = se.writeCommandBytes(0x66, 0x3C, 0x00, 0xCD, 0x67, 0x45, 0x45, 0x10, 0x00, 0x00, 0x00)
	_ = se.writeCommandBytes(0x67, 0x00, 0x3C, 0x00, 0x00, 0x00, 0x01, 0x54, 0x10, 0x32, 0x98)
	_ = se.writeCommandBytes(0x74, 0x10, 0x85, 0x80, 0x00, 0x00, 0x4E, 0x00)
	_ = se.writeCommandBytes(0x98, 0x3e, 0x07)
	_ = se.writeCommandBytes(0x35)
	_ = se.writeCommandBytes(0x21)
	_ = se.writeCommandBytes(0x11)
	_ = se.writeCommandBytes(0x29)

	err = se.SetFrame(Frame{X0: 0, Y0: 0, X1: ScreenWidth - 1, Y1: ScreenHeight - 1})
	if err != nil {
		_ = se.spiDev.Close()
		se.spiDev = nil
		return
	}

	err = se.SetOrientation(Orientation180Degree)
	if err != nil {
		_ = se.spiDev.Close()
		se.spiDev = nil
		return
	}

	return
}
