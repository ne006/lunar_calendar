package moon

import "math"

type Moon struct {
	julianDay int
	Phase     float64
	Age       float64
	Distance  float64
	Longitude float64
	Latitude  float64
}

func MoonFor(julianDay int) *Moon {
	return &Moon{julianDay, 0.0, 0.0, 0.0, 0.0, 0.0}
}

func (moon *Moon) CalcParams() {
	moon.calcPhase()
	moon.calcAge()
	moon.calcCoords()
}

func (moon *Moon) GetHumanPhase() string {
	var humanPhase string

	if moon.Age < 1.84566 {
		humanPhase = "New moon"
	} else if moon.Age < 5.53699 {
		humanPhase = "Waxing crescent"
	} else if moon.Age < 9.22831 {
		humanPhase = "First quarter"
	} else if moon.Age < 12.91963 {
		humanPhase = "Waxing gibbous"
	} else if moon.Age < 16.61096 {
		humanPhase = "Full moon"
	} else if moon.Age < 20.30228 {
		humanPhase = "Waning gibbous"
	} else if moon.Age < 23.99361 {
		humanPhase = "Last quarter"
	} else if moon.Age < 27.68493 {
		humanPhase = "Waning crescent"
	} else {
		humanPhase = "New moon"
	}

	return humanPhase
}

func (moon *Moon) calcPhase() {
	phase := normalize((float64(moon.julianDay) - 2451550.1) / 29.530588853)

	moon.Phase = phase
}

func (moon *Moon) calcAge() {
	age := moon.Phase * 29.53

	moon.Age = age
}

func (moon *Moon) calcCoords() {
	phase_rad := 2 * moon.Phase * math.Pi

	dp := 2 * math.Pi * normalize((float64(moon.julianDay)-2451562.2)/27.55454988)

	distance := 60.4 - 3.3*math.Cos(dp) - 0.6*math.Cos(2*phase_rad-dp) - 0.5*math.Cos(2*phase_rad)

	np := 2 * math.Pi * normalize((float64(moon.julianDay)-2451565.2)/27.212220817)
	latitude := 5.1 * math.Sin(np)

	rp := normalize((float64(moon.julianDay) - 2451555.8) / 27.321582241)
	longitude := 360*rp + 6.3*math.Sin(dp) + 1.3*math.Sin(2*phase_rad-dp) + 0.7*math.Sin(2*phase_rad)

	moon.Distance = distance
	moon.Latitude = latitude
	moon.Longitude = longitude
}

func normalize(num float64) float64 {
	num_fraction := num - math.Floor(num)

	if num_fraction < 0.0 {
		num_fraction = num_fraction + 1.0
	}

	return num_fraction
}
