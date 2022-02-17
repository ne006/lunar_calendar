package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/ne006/lunarcalendar/juliandate"
	"github.com/ne006/lunarcalendar/moon"
)

func main() {
	args := os.Args[1:]

	if err := validateArgs(args); err == nil {
		showMoon(args)
	} else {
		fmt.Printf("Error: %s\n", err)
	}

}

func validateArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("no date supplied")
	} else if match, _ := regexp.Match(`\d{2}\.\d{2}\.\d{4}`, []byte(args[0])); !match {
		return errors.New("date should look like dd.mm.yyyy")
	} else {
		return nil
	}
}

func showMoon(args []string) {
	date, _ := time.Parse("02.01.2006", args[0])

	jDate := juliandate.GregorianToJulian(date.Year(), int(date.Month()), date.Day())

	theMoon := moon.MoonFor(jDate)

	theMoon.CalcParams()

	fmt.Printf(
		"Moon at %s:\nAge: %.2f days\nPhase: %.2f (%s)\nConstellation: %s\nDistance: %.2f Earth Radii\nLongitude: %.2f degrees\nLatitude: %.2f degrees\n",
		date.Format("2 January 2006"),
		theMoon.Age,
		theMoon.Phase,
		theMoon.GetHumanPhase(),
		theMoon.GetZodiacSign(),
		theMoon.Distance,
		theMoon.Longitude,
		theMoon.Latitude,
	)

}
