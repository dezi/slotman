package teams

import (
	"image"
	"slotman/services/impl/provider"
	"slotman/services/type/slotman"
)

const (
	Service provider.Service = "serviceTeams"
)

type Interface interface {
	GetName() (name provider.Service)

	GetAllTeams() (teams []*slotman.Team)
	GetTeam(name string) (team *slotman.Team, err error)
	GetScaledTeamLogo(team *slotman.Team, size int) (img *image.RGBA, err error)
	GetCircleTeamLogo(team *slotman.Team, size int) (img *image.RGBA, err error)

	UpdateTeam(team *slotman.Team)
}

func GetInstance() (iface Interface, err error) {

	baseProvider, err := provider.GetProvider(Service)
	if err != nil {
		return
	}

	iface = baseProvider.(Interface)
	if iface == nil {
		err = provider.ErrNotFound(Service)
		return
	}

	return
}
