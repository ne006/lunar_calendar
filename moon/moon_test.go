package moon

import (
	"testing"
)

func TestCalcParams(t *testing.T) {
	theMoon := MoonFor(2459625)
	moonPhase := 0.44188902550723696550
	moonAge := 13.04898292322870823057
	moonDistance := 62.62038447769674576193
	moonLongitude := 119.71442853800627403871
	moonLatitude := 4.65607436178463718335

	theMoon.CalcParams()

	if theMoon.Phase != moonPhase {
		t.Errorf("moon.phase = %.20f, want %.20f", theMoon.Phase, moonPhase)
	}
	if theMoon.Age != moonAge {
		t.Errorf("moon.age = %.20f, want %.20f", theMoon.Age, moonAge)
	}
	if theMoon.Distance != moonDistance {
		t.Errorf("moon.distance = %.20f, want %.20f", theMoon.Distance, moonDistance)
	}
	if theMoon.Longitude != moonLongitude {
		t.Errorf("moon.longitude = %.20f, want %.20f", theMoon.Longitude, moonLongitude)
	}
	if theMoon.Latitude != moonLatitude {
		t.Errorf("moon.latitude = %.20f, want %.20f", theMoon.Latitude, moonLatitude)
	}
}

func TestGetHumanPhase(t *testing.T) {
	the_moon := MoonFor(2459625)

	the_moon.CalcParams()

	if phase := the_moon.GetHumanPhase(); phase != "Full moon" {
		t.Errorf("Human phase should be \"Full moon\", got \"%s\"", phase)
	}
}

func TestGetZodiacSign(t *testing.T) {
	the_moon := MoonFor(2459625)

	the_moon.CalcParams()

	if zodiacSign := the_moon.GetZodiacSign(); zodiacSign != "Cancer" {
		t.Errorf("Constellation should be in \"Cancer\", got \"%s\"", zodiacSign)
	}
}
