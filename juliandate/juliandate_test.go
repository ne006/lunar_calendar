package juliandate

import "testing"

func TestGregorianToJulian(t *testing.T) {
	jdExpected := 2459625
	jd := GregorianToJulian(2022, 2, 14)

	if jd != jdExpected {
		t.Errorf("jd = %d, want %d", jd, jdExpected)
	}
}
