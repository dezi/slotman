package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"math/rand"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

func (sv *Service) handleController(appId simple.UUIDHex, ws *websocket.Conn, jsonBytes []byte) (err error) {

	controller := &slotman.Controller{}
	err = json.Unmarshal(jsonBytes, controller)
	if err != nil {
		log.Cerror(err)
		return
	}

	if controller.Mode == "set" {

		log.Printf("Controller controller=%d isCalibrating=%v",
			controller.Controller, controller.IsCalibrating)

		go sv.fakeControllerCalibration(appId, ws, controller.Controller, controller.IsCalibrating)

		return
	}

	return
}

var isCalibratingNow bool

func (sv *Service) fakeControllerCalibration(appId simple.UUIDHex, ws *websocket.Conn, selected int, isCalibrating bool) {

	isCalibratingNow = isCalibrating

	controller := &slotman.Controller{
		What:          "controller",
		Mode:          "set",
		Controller:    selected,
		IsCalibrating: isCalibratingNow,
		MinValue:      0,
		MaxValue:      0,
	}

	startTime := time.Now().Unix()

	minValue := -1
	maxValue := -1

	for isCalibratingNow {

		time.Sleep(time.Millisecond * 20)

		value := int(rand.Int31() % 32000)

		if minValue < 0 || value < minValue {
			minValue = value
		}

		if maxValue < 0 || value > maxValue {
			maxValue = value
		}

		controller.MinValue = minValue
		controller.MaxValue = maxValue

		if time.Now().Unix()-startTime > 10 {
			controller.IsCalibrating = false
			isCalibratingNow = false
		}

		jsonBytes, err := simple.MarshalJsonClean(controller)
		if err != nil {
			log.Cerror(err)
			return
		}

		tryErr := ws.WriteMessage(websocket.TextMessage, jsonBytes)
		if tryErr != nil {
			break
		}
	}

	isCalibratingNow = false
}
