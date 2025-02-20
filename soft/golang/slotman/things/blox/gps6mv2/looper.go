package gps6mv2

import (
	"errors"
	"fmt"
	"slotman/things"
	"slotman/utils/log"
	"strings"
	"time"
)

func (se *GPS6MV2) evalLoop() {

	if se.isProbe {
		return
	}

	log.Printf("GPS6MV2 evalLoop started...")
	defer log.Printf("GPS6MV2 evalLoop done.")

	var err error

	for se.IsOpen {

		line := ""
		select {
		case <-time.After(time.Millisecond * 100):
		case line = <-se.results:
		}

		if line == "" {
			continue
		}

		crcHave := 0x00

		for inx := 1; inx < len(line); inx++ {

			if line[inx] == '*' {
				break
			}

			crcHave = crcHave ^ int(line[inx])
		}

		crcNeed := -1

		parts := strings.Split(line, ",")

		if len(parts) > 0 {

			crcParts := strings.Split(parts[len(parts)-1], "*")

			if len(crcParts) == 2 {
				_, _ = fmt.Sscanf(crcParts[1], "%x", &crcNeed)
			}

			//
			// Store last part cleaned from checksum.
			//

			parts[len(parts)-1] = crcParts[0]
		}

		if crcNeed != crcHave {
			err := errors.New("checksum mismatch")
			log.Cerror(err)
			continue
		}

		switch parts[0] {
		case "$GPRMC":
			//
			// Recommended minimum specific GPS/Transit data.
			//
			// $GPRMC,150338.00,A,5334.08742,N,01008.40794,E,0.494, ,180225, ,  ,A*76
			//    0      1      2     3      4      5      6    7  8   9   10 11 12
			//
			//  1 = UTC of position fix
			//  2 = Data status (V=navigation receiver warning)
			//  3 = Latitude of fix
			//  4 = N or S
			//  5 = Longitude of fix
			//  6 = E or W
			//  7 = Speed over ground in knots
			//  8 = Track made good in degrees True
			//  9 = UT date
			// 10 = Magnetic variation degrees (Easterly var. subtracts from true course)
			// 11 = E or W
			// 12 = Mode indicator:
			//		A: Autonomous mode
			//		D: Differential mode
			//		E: Estimated (dead reckoning) mode
			//		M: Manual input mode
			//		S: Simulator mode
			//		N: Data not valid
			//

			if len(parts) != 13 {
				continue
			}

			if parts[2] != "A" {
				continue
			}

			var latitude float64

			_, err = fmt.Sscanf(parts[3], "%f", &latitude)
			if err != nil {
				continue
			}

			latitude /= 100
			if parts[4] == "S" {
				latitude = -latitude
			}

			var longitude float64
			_, err = fmt.Sscanf(parts[5], "%f", &longitude)
			if err != nil {
				continue
			}

			longitude /= 100
			if parts[6] == "W" {
				longitude = -longitude
			}

			se.Latitude = latitude
			se.Longitude = longitude

			//log.Printf("GPS position latitude=%f longitude=%f", longitude, latitude)

			handler := se.handler
			if handler != nil {
				handler.OnGPSPosition(se, se.Latitude, se.Longitude, se.Elevation)
			}

		case "$GPGGA":
			//
			// GPS fix data and undulation
			//
			// $GPGGA,172655.00,5334.07581,N,01008.39612,E,1,06,1.40,41.6,M,45.0,M,   ,*63
			//    0      1          2      3      4      5 6 7  8    9   10  11 12 13  14
			//
			//  1 = UTC of Position
			//  2 = Latitude
			//  3 = N or S
			//  4 = Longitude
			//  5 = E or W
			//  6 = GPS quality indicator (0=invalid; 1=GPS fix; 2=Diff. GPS fix)
			//  7 = Number of satellites in use [not those in view]
			//  8 = Horizontal dilution of position
			//  9 = Antenna altitude above/below mean sea level (geoid)
			// 10 = Meters  (Antenna height unit)
			// 11 = Geoidal separation (Diff. between WGS-84 earth ellipsoid and
			//      mean sea level.  -=geoid is below WGS-84 ellipsoid)
			// 12 = Meters  (Units of geoidal separation)
			// 13 = Age in seconds since last update from diff. reference station
			// 14 = Diff. reference station ID#
			//

			if len(parts) != 15 {
				continue
			}

			var elevation float64
			_, err = fmt.Sscanf(parts[9], "%f", &elevation)
			if err != nil {
				continue
			}

			var geoidal float64
			_, err = fmt.Sscanf(parts[11], "%f", &geoidal)
			if err != nil {
				continue
			}

			se.Elevation = elevation

			//log.Printf("GPS position elevation=%f geoidal=%f", elevation, geoidal)

		case "$GPGSV":
			//
			// GPS satellites in view
			//
			// $GPGSV,3,1,11,01,01,360, ,06,10,100,  ,10,13,279,23,12,64,237,31*73
			//    0   1 2  3  4  5   6 7  8  9  10 11 12 13  14 15 16 17  18 19
			//
			// 0 = Message ID $GPGSV
			// 1 = Total number of messages of this type in this cycle
			// 2 = Message number
			// 3 = Total number of SVs visible
			// 4 = SV PRN number
			// 5 = Elevation, in degrees, 90° maximum
			// 6 = Azimuth, degrees from True North, 000° through 359°
			// 7 = SNR, 00 through 99 dB (null when not tracking)
			// 8–11  = Information about second SV, same format as fields 4 through 7
			// 12–15 = Information about third  SV, same format as fields 4 through 7
			// 16–19 = Information about fourth SV, same format as fields 4 through 7
			//

		case "$GPGSA":
			//
			// GPS DOP and active satellites
			//
			// $GPGSA,A,3,12,25,24,15,19,17,,,,,,,2.33,1.59,1.70*03
			//   0    1 2           3-14          15   16   17
			//
			// 0 = Message ID $GPGSA
			// 1 = Mode:
			//       M=Manual, forced to operate in 2D or 3D
			//       A=Automatic, 3D/2D
			// 2 = Mode:
			//       1=Fix not available
			//       2=2D
			//       3=3D
			// 3-14 = IDs of SVs used in position fix (null for unused fields)
			// 15 = PDOP
			// 16 = HDOP
			// 17 = VDOP
			//

		case "$GPGLL":
			//
			// Geographic Position, Latitude / Longitude and time.
			//
			// $GPGLL,5109.0262317,N,11401.8407304,W,202725.00,A,D*79
			// $GPGLL,  5334.08050,N,  01008.40075,E,080946.00,A,A*69
			//   0          1      2       3       4      5    6 7
			//
			// 0 = Message ID $GPGLL
			// 1 = Latitude in dd mm,mmmm format (0-7 decimal places)
			// 2 = Direction of latitude N: North S: South
			// 3 = Longitude in ddd mm,mmmm format (0-7 decimal places)
			// 4 = Direction of longitude E: East W: West
			// 5 = UTC of position in hhmmss.ss format
			// 6 = Status indicator:
			//		A: Data valid
			//		V: Data not valid
			// 7 = Mode indicator:
			//		A: Autonomous mode
			//		D: Differential mode
			//		E: Estimated (dead reckoning) mode
			//		M: Manual input mode
			//		S: Simulator mode
			//		N: Data not valid
			//
		case "$GPVTG":
			//
			// Track made good and ground speed.
			//
			// $GPVTG,224.592,T,224.592,M,0.003,N,0.005,K,D*20
			// $GPVTG,       ,T,       ,M,0.222,N,0.411,K,A*25
			//    0       1   2     3   4    5  6   7   8 9
			//
			// 1 = True track made good
			// 2 = Fixed text 'T' indicates that track made good is relative to true north
			// 3 = Magnetic track made good
			// 4 = Fixed text 'M' indicates that track made good is relative to magnetic north
			// 5 = Speed over ground in knots
			// 6 = Fixed text 'N' indicates that speed over ground in knots
			// 7 = Speed over ground in kilometers/hour
			// 8 = Fixed text 'K' indicates that speed over ground is in kilometers/hour
			// 9 = Indicator
			// 		A Autonomous
			//		D Differential
			//		E Estimated (dead reckoning) mode
			//		M Manual input
			//		N Data not valid
			//
		default:
			log.Printf("Line <%s> crcNeed=%02x crcHave=%02x", line, crcNeed, crcHave)
		}
	}
}

func (se *GPS6MV2) readLoop() {

	if !se.isProbe {
		log.Printf("GPS6MV2 readLoop started...")
		defer log.Printf("GPS6MV2 readLoop done.")
	}

	if !se.IsOpen {
		err := things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	parts := make([]byte, 100)
	input := make([]byte, 0)

	port := se.uart
	if port != nil {
		for {
			xfer, _ := port.Read(parts)
			if xfer == 0 {
				break
			}
		}
	}

	for se.IsOpen {

		port = se.uart
		if port == nil {
			break
		}

		xfer, _ := port.Read(parts)
		input = append(input, parts[:xfer]...)

		lines := strings.Split(string(input), "\r\n")

		if len(lines) == 0 {
			continue
		}

		input = []byte(lines[len(lines)-1])
		lines = lines[:len(lines)-1]

		for _, line := range lines {
			se.results <- line

			//if se.isProbe {
			//	log.Printf("Line <%s>", line)
			//}
		}
	}
}
