package main

import "log"

func main() {
	platform := NewPlatform("outA")
	sensPos := NewSensorPositioner("outD")
	sensor := NewColorSensor("in1")
	grabber := NewGrabber("outC")

	// side := ReadSide(platform, sensPos, sensor)
	// log.Println(side)

	cube := Cube{}

	for i := 0; i < 3; i++ {
		cube[i] = ReadSide(platform, sensPos, sensor)
		grabber.Flip()
	}

	cube[3] = ReadSide(platform, sensPos, sensor)

	platform.SetAngle(90)
	platform.ClearAngle()

	grabber.Flip()

	cube[4] = ReadSide(platform, sensPos, sensor)

	grabber.Flip()
	grabber.Flip()

	cube[5] = ReadSide(platform, sensPos, sensor)

	log.Println(cube)

}

func ReadSide(p *Platform, sp *SensorPositioner, s *ColorSensor) Side {
	side := Side{}
	sp.GoCenter()

	side[0] = s.GetColor()

	sp.GoSide()

	for i := 1; i < 9; i++ {
		if i%2 == 0 {
			sp.GoCornerSide()
		} else {
			sp.GoSide()
		}
		side[i] = s.GetColor()

		p.SetAngle(45)
		p.ClearAngle()
	}

	sp.GoDefault()

	return side
}
