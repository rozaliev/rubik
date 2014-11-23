package main

import (
	"math"
	"strconv"
	"time"
)

type Platform struct {
	path   string
	offset int
}

type Grabber struct {
	path string
}

type SensorPositioner struct {
	path string
}

type ColorSensor struct {
	path string
}

func NewPlatform(port string) *Platform {
	path, err := FindDevice("/sys/class/tacho-motor", "port_name", port)
	FatalOnErr(err)

	platform := &Platform{
		path: path,
	}

	FatalOnErr(SetValue(path, "run_mode", "position"))
	FatalOnErr(SetValue(path, "position", "0"))
	FatalOnErr(SetValue(path, "stop_mode", "brake"))
	FatalOnErr(SetValue(path, "regulation_mode", "on"))
	FatalOnErr(SetValue(path, "pulses_per_second_sp", "500"))
	FatalOnErr(SetValue(path, "ramp_down_sp", "100"))
	FatalOnErr(SetValue(path, "ramp_up_sp", "100"))

	return platform
}

const ANGLE_COEF = 2.9944

func (p *Platform) SetAngle(angle int) {
	pos := float64(angle) * ANGLE_COEF
	p.offset = int(pos)
	val := strconv.Itoa(int(pos))
	FatalOnErr(SetValue(p.path, "position_sp", val))
	FatalOnErr(SetValue(p.path, "run", "1"))

	for {
		val, err := GetValue(p.path, "run")
		FatalOnErr(err)

		if val == "0" {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (p *Platform) ClearAngle() {
	RawCPos, err := GetValue(p.path, "position")
	FatalOnErr(err)
	cPos, err := strconv.Atoi(RawCPos)
	FatalOnErr(err)
	nPos := cPos - p.offset
	rawNPos := strconv.Itoa(nPos)
	FatalOnErr(SetValue(p.path, "position", rawNPos))
}

func NewGrabber(port string) *Grabber {
	path, err := FindDevice("/sys/class/tacho-motor", "port_name", port)
	FatalOnErr(err)

	grabber := &Grabber{
		path: path,
	}

	FatalOnErr(SetValue(path, "run_mode", "position"))
	FatalOnErr(SetValue(path, "position", "0"))
	FatalOnErr(SetValue(path, "stop_mode", "brake"))
	FatalOnErr(SetValue(path, "regulation_mode", "on"))
	FatalOnErr(SetValue(path, "pulses_per_second_sp", "500"))

	FatalOnErr(SetValue(path, "ramp_down_sp", "600"))
	FatalOnErr(SetValue(path, "ramp_up_sp", "100"))

	return grabber
}

const GRAB = 250
const MAX_FLIP = 250

func (g *Grabber) Grab() {
	val := strconv.Itoa(GRAB)
	FatalOnErr(SetValue(g.path, "position_sp", val))
	FatalOnErr(SetValue(g.path, "run", "1"))

	for {
		val, err := GetValue(g.path, "run")
		FatalOnErr(err)

		if val == "0" {
			time.Sleep(20 * time.Millisecond)
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (g *Grabber) ToStart() {
	FatalOnErr(SetValue(g.path, "position_sp", "0"))
	FatalOnErr(SetValue(g.path, "run", "1"))

	for {
		val, err := GetValue(g.path, "run")
		FatalOnErr(err)

		if val == "0" {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (g *Grabber) Flip() {
	val := strconv.Itoa(MAX_FLIP)
	FatalOnErr(SetValue(g.path, "position_sp", val))
	FatalOnErr(SetValue(g.path, "run", "1"))

	for {
		val, err := GetValue(g.path, "run")
		FatalOnErr(err)

		if val == "0" {
			time.Sleep(20 * time.Millisecond)
			g.ToStart()
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

const SENSOR_POS_CENTER = -610
const SENSOR_POS_SIDE = -500
const SENSOR_POS_CORNER_SIDE = -475

func NewSensorPositioner(port string) *SensorPositioner {
	path, err := FindDevice("/sys/class/tacho-motor", "port_name", port)
	FatalOnErr(err)

	sensPos := &SensorPositioner{
		path: path,
	}

	FatalOnErr(SetValue(path, "run_mode", "position"))
	FatalOnErr(SetValue(path, "position", "0"))
	FatalOnErr(SetValue(path, "stop_mode", "brake"))
	FatalOnErr(SetValue(path, "regulation_mode", "on"))
	FatalOnErr(SetValue(path, "pulses_per_second_sp", "900"))
	FatalOnErr(SetValue(path, "ramp_down_sp", "50"))
	FatalOnErr(SetValue(path, "ramp_up_sp", "50"))

	return sensPos
}

func (s *SensorPositioner) GoDefault() {
	FatalOnErr(SetValue(s.path, "position_sp", "0"))
	FatalOnErr(SetValue(s.path, "run", "1"))

	for {
		val, err := GetValue(s.path, "run")
		FatalOnErr(err)

		if val == "0" {
			time.Sleep(20 * time.Millisecond)
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (s *SensorPositioner) GoCenter() {
	val := strconv.Itoa(SENSOR_POS_CENTER)
	FatalOnErr(SetValue(s.path, "position_sp", val))
	FatalOnErr(SetValue(s.path, "run", "1"))

	for {
		val, err := GetValue(s.path, "run")
		FatalOnErr(err)

		if val == "0" {
			time.Sleep(20 * time.Millisecond)
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (s *SensorPositioner) GoSide() {
	val := strconv.Itoa(SENSOR_POS_SIDE)
	FatalOnErr(SetValue(s.path, "position_sp", val))
	FatalOnErr(SetValue(s.path, "run", "1"))

	for {
		val, err := GetValue(s.path, "run")
		FatalOnErr(err)

		if val == "0" {
			time.Sleep(20 * time.Millisecond)
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func (s *SensorPositioner) GoCornerSide() {
	val := strconv.Itoa(SENSOR_POS_CORNER_SIDE)
	FatalOnErr(SetValue(s.path, "position_sp", val))
	FatalOnErr(SetValue(s.path, "run", "1"))

	for {
		val, err := GetValue(s.path, "run")
		FatalOnErr(err)

		if val == "0" {
			time.Sleep(20 * time.Millisecond)
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

type Color int

const (
	Red Color = iota
	Blue
	Green
	White
	Orange
	Yellow
)

var (
	BlueValue   = ColorValue{R: 21, G: 55, B: 100}
	GreenValue  = ColorValue{R: 26, G: 100, B: 35}
	OrangeValue = ColorValue{R: 130, G: 35, B: 20}
	YellowValue = ColorValue{R: 180, G: 100, B: 30}
	RedValue    = ColorValue{R: 130, G: 10, B: 10}
	WhiteValue  = ColorValue{R: 255, G: 255, B: 255}
)

var ColorMatches = map[ColorValue]Color{
	BlueValue:   Blue,
	RedValue:    Red,
	WhiteValue:  White,
	GreenValue:  Green,
	OrangeValue: Orange,
	YellowValue: Yellow,
}

var ColorNames = map[Color]string{
	Blue:   "Blue",
	Red:    "Red",
	White:  "White",
	Green:  "Green",
	Orange: "Orange",
	Yellow: "Yellow",
}

type ColorValue struct {
	R, G, B int
}

func (c Color) String() string {
	return ColorNames[c]
}

func NewColorSensor(port string) *ColorSensor {
	path, err := FindDevice("/sys/class/msensor", "port_name", port)
	FatalOnErr(err)

	colorSensor := &ColorSensor{
		path: path,
	}

	FatalOnErr(SetValue(path, "mode", "RGB-RAW"))

	return colorSensor
}

func (s *ColorSensor) GetColor() Color {
	raw := s.GetRawColor()

	smDist := -1.0
	var smColor Color

	for tcolor, color := range ColorMatches {
		sum := math.Pow(float64(raw.R-tcolor.R), 2) + math.Pow(float64(raw.G-tcolor.G), 2) + math.Pow(float64(raw.B-tcolor.B), 2)
		dist := math.Sqrt(sum)
		if smDist == -1 {
			smDist = dist
			smColor = color
		} else {
			if dist < smDist {
				smDist = dist
				smColor = color
			}
		}
	}

	return smColor
}

func (s *ColorSensor) GetRawColor() ColorValue {
	time.Sleep(time.Millisecond * 50)
	rr, err := GetValue(s.path, "value0")
	FatalOnErr(err)
	r, err := strconv.Atoi(rr)
	FatalOnErr(err)

	gr, err := GetValue(s.path, "value1")
	FatalOnErr(err)
	g, err := strconv.Atoi(gr)
	FatalOnErr(err)

	br, err := GetValue(s.path, "value2")
	FatalOnErr(err)
	b, err := strconv.Atoi(br)
	FatalOnErr(err)

	return ColorValue{
		R: r,
		G: g,
		B: b,
	}
}
