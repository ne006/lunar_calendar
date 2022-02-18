package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ne006/lunarcalendar/juliandate"
	"github.com/ne006/lunarcalendar/moon"
)

func main() {
	router := initRouter()
	route(router)
	runServer(router)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func route(router *gin.Engine) {
	router.GET("/moon", moonAt)
}

func runServer(router *gin.Engine) {
	args := os.Args[1:]
	if len(args) == 0 {
		router.Run()
	} else {
		router.Run(args[0])
	}
}

// Handlers
func moonAt(c *gin.Context) {
	at, err := time.Parse("02.01.2006", c.DefaultQuery("at", time.Now().Format("02.01.2006")))

	if err == nil {
		jDate := juliandate.GregorianToJulian(at.Year(), int(at.Month()), at.Day())
		theMoon := moon.MoonFor(jDate)
		theMoon.CalcParams()

		c.JSON(http.StatusOK, gin.H{
			"date": at.Format("2 January 2006"),
			"moon": gin.H{
				"age":        theMoon.Age,
				"phase":      theMoon.Phase,
				"humanPhase": theMoon.GetHumanPhase(),
				"zodiacSign": theMoon.GetZodiacSign(),
				"coords": gin.H{
					"distance":  theMoon.Distance,
					"longitude": theMoon.Longitude,
					"latitdude": theMoon.Latitude,
				},
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
}
