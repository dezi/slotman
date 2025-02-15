package race

import (
	"encoding/json"
	"slotman/services/type/server"
	"slotman/services/type/slotman"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) OnRequestFromClient(appId simple.UUIDHex, what string, reqBytes []byte) {
	switch what {
	case "init":
		err := sv.handleInit(appId, reqBytes)
		log.Cerror(err)
	case "race":
		err := sv.handleRace(appId, reqBytes)
		log.Cerror(err)
	case "pilot":
		err := sv.handlePilot(appId, reqBytes)
		log.Cerror(err)
	case "tracks":
		err := sv.handleTracks(appId, reqBytes)
		log.Cerror(err)
	}
}

func (sv *Service) handleTracks(appId simple.UUIDHex, reqBytes []byte) (err error) {

	tracks := slotman.Tracks{}
	err = json.Unmarshal(reqBytes, &tracks)
	if err != nil {
		log.Cerror(err)
		return
	}

	if tracks.Mode == "get" {

		tracks = sv.setup.Tracks

		reqBytes, err = simple.MarshalJsonClean(tracks)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = sv.srv.Transmit(appId, reqBytes)
		log.Cerror(err)

		return
	}

	if tracks.Mode == "set" {

		sv.setup.Tracks = tracks

		log.Printf("Tracks tracks=%d", sv.setup.Tracks.Tracks)

		err = sv.srv.Broadcast(appId, reqBytes)
		log.Cerror(err)

		return
	}

	return
}

func (sv *Service) handleRace(appId simple.UUIDHex, reqBytes []byte) (err error) {

	race := slotman.Race{}
	err = json.Unmarshal(reqBytes, &race)
	if err != nil {
		log.Cerror(err)
		return
	}

	if race.Mode == "get" {

		race = sv.setup.Race

		var resBytes []byte
		resBytes, err = simple.MarshalJsonClean(race)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = sv.srv.Transmit(appId, resBytes)
		log.Cerror(err)

		return
	}

	if race.Mode == "set" {

		sv.setup.Race = race

		log.Printf("Race title=%s tracks=%d rounds=%d",
			sv.setup.Race.Title, sv.setup.Race.Tracks, sv.setup.Race.Rounds)

		err = sv.srv.Broadcast(appId, reqBytes)
		log.Cerror(err)

		return
	}

	return
}

func (sv *Service) handlePilot(appId simple.UUIDHex, reqBytes []byte) (err error) {

	pilot := &slotman.Pilot{}
	err = json.Unmarshal(reqBytes, pilot)
	if err != nil {
		log.Cerror(err)
		return
	}

	if pilot.Mode == "set" {

		sv.setupLock.Lock()
		sv.setup.Pilots[pilot.Uuid] = pilot
		sv.setupLock.Unlock()

		log.Printf("Pilot first=%s last=%s car=%s",
			pilot.FirstName, pilot.LastName, pilot.Car)

		err = sv.srv.Broadcast(appId, reqBytes)
		log.Cerror(err)

		return
	}

	return
}

func (sv *Service) handleInit(appId simple.UUIDHex, reqBytes []byte) (err error) {

	init := server.Message{}
	err = json.Unmarshal(reqBytes, &init)
	if err != nil {
		log.Cerror(err)
		return
	}

	if init.Mode == "get" {

		tracks := sv.setup.Tracks

		var resBytes []byte
		resBytes, err = simple.MarshalJsonClean(tracks)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = sv.srv.Transmit(appId, resBytes)
		log.Cerror(err)

		race := sv.setup.Race

		resBytes, err = simple.MarshalJsonClean(race)
		if err != nil {
			log.Cerror(err)
			return
		}

		err = sv.srv.Transmit(appId, resBytes)
		log.Cerror(err)

		for _, pilot := range sv.setup.Pilots {

			pilot.What = "pilot"
			pilot.Mode = "set"

			reqBytes, err = simple.MarshalJsonClean(pilot)
			if err != nil {
				log.Cerror(err)
				return
			}

			err = sv.srv.Transmit(appId, resBytes)
			log.Cerror(err)
		}

		return
	}

	return
}
