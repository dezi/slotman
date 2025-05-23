package turner

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"os"
	"slotman/services/type/slotman"
	"slotman/things/galaxycore/gc9a01"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
	"time"
)

func (sv *Service) DoControlTask() {

	sv.checkDisplays()

	if sv.isProxyServer {
		return
	}

	state := sv.rce.GetRaceState()

	if state == slotman.RaceStateIdle {
		switch sv.loopCount % 2 {
		case 0:
			sv.displayState()
		case 1:
			sv.displayTeams()
		}
	}

	if state == slotman.RaceStateRaceWaiting ||
		state == slotman.RaceStateRaceRunning ||
		state == slotman.RaceStateRaceSuspended {
		sv.displayRace()
	}

	sv.loopCount++
}

func (sv *Service) displayRace() {

	tracksReady := sv.rce.GetTrackStates()

	haveTrack := false

	for inx := 0; inx < len(tracksReady); inx++ {

		if tracksReady[sv.raceIndex] > 0 {
			haveTrack = true
			break
		}

		sv.raceIndex = (sv.raceIndex + 1) % slotman.MaxTracks
	}

	defer func() { sv.raceIndex = (sv.raceIndex + 1) % slotman.MaxTracks }()

	if !haveTrack {
		return
	}

	raceInfo, err := sv.rce.GetRaceInfo(sv.raceIndex)
	if err != nil {
		log.Cerror(err)
		return
	}

	if raceInfo.PilotUuid == "" {
		err = errors.New("no pilot for track")
		log.Cerror(err)
		return
	}

	pilot, err := sv.plt.GetPilot(raceInfo.PilotUuid)
	if err != nil {
		log.Cerror(err)
		return
	}

	img, err := sv.getSlotmanLogo()
	if err != nil {
		log.Cerror(err)
		return
	}

	dc := gg.NewContextForRGBA(img)

	dc.DrawRectangle(0, 0, 240, 240)
	dc.SetHexColor("#00000080")
	dc.Fill()

	pilotImg, err := sv.plt.GetCircularPilotPic(pilot, 80)
	if err != nil {
		log.Cerror(err)
		return
	}

	dc.DrawImage(pilotImg, 120-80/2, 8)

	dc.SetHexColor(goldColor)
	dc.SetFontFace(sv.faceBoldNormal)

	var text string

	yPos := 120.0

	if raceInfo.Position == 0 {
		text = fmt.Sprintf("n.a. - %d", raceInfo.Rounds)
	} else {
		text = fmt.Sprintf("#%d - %d", raceInfo.Position, raceInfo.Rounds)
	}

	dc.DrawStringAnchored(text, 120, yPos, 0.5, 0)

	yPos += 30

	text = fmt.Sprintf("%0.1f - %0.1f", raceInfo.TopRound, raceInfo.ActRound)
	dc.DrawStringAnchored(text, 120, yPos, 0.5, 0)

	yPos += 30

	text = fmt.Sprintf("%0.0f - %0.0f", raceInfo.TopSpeed, raceInfo.ActSpeed)
	dc.DrawStringAnchored(text, 120, yPos, 0.5, 0)

	err = sv.blipFullImage(img)
	log.Cerror(err)
}

func (sv *Service) displayState() {

	img, err := sv.getSlotmanLogo()
	if err != nil {
		log.Cerror(err)
		return
	}

	raceState := sv.rce.GetRaceState()

	if raceState == slotman.RaceStateIdle {

		if sv.loopCount%4 == 0 {
			sv.displayHardware(img)
		}

		if sv.loopCount%4 == 2 {
			sv.displayControls(img)
		}
	}

	err = sv.blipFullImage(img)
	log.Cerror(err)
}

func (sv *Service) displayHardware(img *image.RGBA) {

	dc := gg.NewContextForRGBA(img)

	dc.DrawRectangle(0, 0, 240, 240)
	dc.SetHexColor("#00000080")
	dc.Fill()

	dc.SetHexColor(goldColor)
	dc.SetFontFace(sv.faceBoldLarge)
	dc.DrawStringAnchored("Host", 120, 44, 0.5, 0.0)
	dc.DrawStringAnchored("____", 120, 48, 0.5, 0.0)

	//
	// Hostname.
	//

	hostName, _ := os.Hostname()
	parts := strings.Split(hostName, ".")
	hostName = parts[0]

	dc.SetFontFace(sv.faceBoldLarge)
	dc.DrawStringAnchored(hostName, 120, float64(86+0*36), 0.5, 0.0)

	//
	// Ip-Address.
	//

	ipAddr := ""

	ip, err := simple.GetLocalIPV4()
	if err != nil {
		ipAddr = "No IP-Address"
	} else {
		ipAddr = ip.String()
	}

	dc.SetFontFace(sv.faceBoldNormal)
	dc.DrawStringAnchored(ipAddr, 120, float64(86+1*36), 0.5, 0.0)

	//
	// Operating system.
	//

	goos := simple.FirstUpper(simple.GOOS)

	dc.SetFontFace(sv.faceBoldNormal)
	dc.DrawStringAnchored(goos, 120, float64(86+2*36), 0.5, 0.0)

	//
	// Software version.
	//

	dc.SetFontFace(sv.faceBoldNormal)
	dc.DrawStringAnchored(slotman.Version, 120, float64(86+3*36), 0.5, 0.0)
}

func (sv *Service) displayControls(img *image.RGBA) {

	tachosAttached := sv.tco.GetTachosAttached()
	motoronsAttached := sv.sdo.GetMotoronsAttached()
	speedControlsAttached := sv.sdi.GetSpeedControlsAttached()
	tracksVoltage := sv.rce.GetTrackVoltages()

	maxControls := 0
	for _, attached := range speedControlsAttached {
		if attached {
			maxControls++
		}
	}

	dc := gg.NewContextForRGBA(img)

	dc.DrawRectangle(0, 0, 240, 240)
	dc.SetHexColor("#00000080")
	dc.Fill()

	dc.SetHexColor(goldColor)
	dc.SetFontFace(sv.faceBoldLarge)
	dc.DrawStringAnchored("Controls", 120, 44, 0.5, 0.0)
	dc.DrawStringAnchored("________", 120, 48, 0.5, 0.0)

	showMin := 0
	showMax := 4

	if maxControls > 4 && sv.loopCount%8 == 6 {
		showMin = 4
		showMax = 8
	}

	for track := showMin; track < showMax; track++ {

		dc.SetHexColor(goldColor)
		dc.SetFontFace(sv.faceBoldLarge)
		text := fmt.Sprintf("%d:", track+1)
		dc.DrawString(text, 30, float64(86+track*36))

		if motoronsAttached[track] {
			dc.SetHexColor("00ff00")
		} else {
			dc.SetHexColor("ff0000")
		}
		dc.DrawString("M", 60, float64(86+track*36))

		if speedControlsAttached[track] {
			dc.SetHexColor("00ff00")
		} else {
			dc.SetHexColor("ff0000")
		}
		dc.DrawString("C", 90, float64(86+track*36))

		if tachosAttached[track] {
			dc.SetHexColor("00ff00")
		} else {
			dc.SetHexColor("ff0000")
		}
		dc.DrawString("T", 120, float64(86+track*36))

		if motoronsAttached[track] {
			text = fmt.Sprintf("%0.1fV", float64(tracksVoltage[track])/1000)
			dc.SetHexColor("00ff00")
			dc.SetFontFace(sv.faceRegularNormal)
			dc.DrawString(text, 145, float64(86+track*36))
		}
	}
}

func (sv *Service) displayPilots() {

	pilots := sv.plt.GetAllPilots()

	sv.pilotIndex = (sv.pilotIndex + 1) % len(pilots)

	pilot := pilots[sv.pilotIndex]

	img, err := sv.plt.GetScaledPilotPic(pilot, 240)
	if err != nil {
		log.Cerror(err)
		return
	}

	//dc := gg.NewContextForRGBA(img)
	//
	//teams := sv.tms.GetAllTeams()
	//
	//for _, team := range teams {
	//
	//	if pilot.Team != team.Name {
	//		continue
	//	}
	//
	//	var teamImg *image.RGBA
	//	teamImg, err = sv.tms.GetCircleTeamLogo(team, circleSize*2)
	//	if err != nil {
	//		log.Cerror(err)
	//		continue
	//	}
	//
	//	dc.DrawImage(teamImg, 120-circleSize, 240-circleSize*2)
	//
	//	break
	//}

	err = sv.blipFullImage(img)
	log.Cerror(err)
}

func (sv *Service) displayTeams() {

	teams := sv.tms.GetAllTeams()

	sv.teamIndex = (sv.teamIndex + 1) % len(teams)

	team := teams[sv.teamIndex]

	img, err := sv.tms.GetScaledTeamLogo(team, 240)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = sv.blipFullImage(img)
	log.Cerror(err)

	dc := gg.NewContextForRGBA(img)

	pilots := sv.plt.GetAllPilots()

	members := 0
	for _, pilot := range pilots {
		if pilot.Team == team.Name {
			members++
		}
	}

	positions := evenPositions
	if members%2 == 1 {
		positions = oddPositions
	}

	count := 0

	for _, pilot := range pilots {

		if pilot.Team != team.Name {
			continue
		}

		time.Sleep(time.Millisecond * 1000)

		var pilotImg *image.RGBA
		pilotImg, err = sv.plt.GetCircularPilotPic(pilot, circleSize)
		if err != nil {
			log.Cerror(err)
			continue
		}

		if (count+1)<<1 > len(positions) {
			break
		}

		dc.DrawImage(pilotImg, int(positions[0+count<<1]), int(positions[1+count<<1]))
		count++

		err = sv.blipFullImage(img)
		log.Cerror(err)
	}

	time.Sleep(time.Millisecond * 2000)
}

func (sv *Service) checkDisplays() {

	if sv.isProxyClient {
		return
	}

	if sv.turnDisplay1 == nil {

		turnDisplay1 := gc9a01.NewGC9A01("/dev/spidev0.0", 25)
		turnDisplay1.SetHandler(sv)

		tryErr := turnDisplay1.Open()
		if tryErr == nil {
			sv.turnDisplay1 = turnDisplay1
		} else {
			log.Cerror(tryErr)
		}
	}

	if sv.turnDisplay2 == nil {

		turnDisplay2 := gc9a01.NewGC9A01("/dev/spidev0.1", 25)
		turnDisplay2.SetHandler(sv)

		tryErr := turnDisplay2.Open()
		if tryErr == nil {
			sv.turnDisplay2 = turnDisplay2
		} else {
			log.Cerror(tryErr)
		}
	}
}
