package speedi

func (sv *Service) GetSpeedControlsAttached() (tracks []bool) {

	tracks = make([]bool, 8)

	for track, attached := range sv.speedControlAttached {
		tracks[track] = attached
	}

	return
}
