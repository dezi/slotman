package gps6mv2

import "errors"

func (se *GPS6MV2) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *GPS6MV2) GetCurrentPosition() (latitude, longitude, elevation float64, err error) {

	if se.Latitude == 0 && se.Longitude == 0 {
		err = errors.New("no data available")
		return
	}

	latitude = se.Latitude
	longitude = se.Longitude
	elevation = se.Elevation
	return
}
