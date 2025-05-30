package teams

import (
	"errors"
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"slotman/services/type/slotman"
	"slotman/utils/imaging"
	"slotman/utils/log"
	"sort"
)

func (sv *Service) GetAllTeams() (teams []*slotman.Team) {

	sv.mapsLock.Lock()
	defer sv.mapsLock.Unlock()

	for _, team := range sv.teams {
		teams = append(teams, team)
	}

	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Name < teams[j].Name
	})

	return
}

func (sv *Service) GetTeam(name string) (team *slotman.Team, err error) {

	for _, mockupTeam := range mockupTeams {
		if mockupTeam.Name == name {
			team = mockupTeam
			return
		}
	}

	err = errors.New(fmt.Sprintf("unknown team <%s>", name))
	return
}

func (sv *Service) GetScaledTeamLogo(team *slotman.Team, size int) (img *image.RGBA, err error) {

	src, err := imaging.DecodeBase64Image(team.Logo)
	if err != nil {
		log.Cerror(err)
		return
	}

	img = image.NewRGBA(image.Rect(0, 0, size, size))
	draw.NearestNeighbor.Scale(img, img.Rect, src, src.Bounds(), draw.Over, nil)

	return
}

func (sv *Service) GetCircleTeamLogo(team *slotman.Team, size int) (img *image.RGBA, err error) {

	img, err = sv.GetScaledTeamLogo(team, size)
	if err != nil {
		return
	}

	img, err = imaging.ScaleToCircle(img, size, 8, "e0bf78")
	return
}

func (sv *Service) UpdateTeam(team *slotman.Team) {

	sv.mapsLock.Lock()
	defer sv.mapsLock.Unlock()

	sv.teams[team.Uuid] = team

	sv.teamLogoFull[team.Uuid] = nil
	sv.teamLogoSmall[team.Uuid] = nil

	sv.teamCarFull[team.Uuid] = nil
	sv.teamCarSmall[team.Uuid] = nil

	if team.Logo != "" {

		img, err := imaging.DecodeBase64Image(team.Logo)
		if err != nil {

			log.Cerror(err)

		} else {

			var full *image.RGBA
			full, err = imaging.ScaleToCircle(img, 240, 0, "")
			if err == nil {
				sv.teamLogoFull[team.Uuid] = full
			}

			var small *image.RGBA
			small, err = imaging.ScaleToCircle(img, 40, 2, "ff0000")
			if err == nil {
				sv.teamLogoSmall[team.Uuid] = small
			}
		}
	}

	if team.CarPic != "" {

		img, err := imaging.DecodeBase64Image(team.CarPic)
		if err != nil {

			log.Cerror(err)

		} else {

			var full *image.RGBA
			full, err = imaging.ScaleToCircle(img, 240, 0, "")
			if err == nil {
				sv.teamCarFull[team.Uuid] = full
			}

			var small *image.RGBA
			small, err = imaging.ScaleToCircle(img, 40, 2, "ff0000")
			if err == nil {
				sv.teamCarSmall[team.Uuid] = small
			}
		}
	}

	return
}
