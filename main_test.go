package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCrossingTime(t *testing.T) {
	a := assert.New(t)

	var Bridges = []Bridge{
		{ID: "1", LengthByFeet: 100, HikerIDsCrossing: []string{"A", "B", "C", "D"}},
		{ID: "2", LengthByFeet: 250, HikerIDsCrossing: []string{"A", "B", "C", "D", "E"}},
		{ID: "3", LengthByFeet: 150, HikerIDsCrossing: []string{"A", "B", "C", "D", "E", "F", "G"}},
	}

	var Hikers = []Hiker{
		{ID: "A", SpeedByFeetPerMin: float64(100)},
		{ID: "B", SpeedByFeetPerMin: float64(50)},
		{ID: "C", SpeedByFeetPerMin: float64(20)},
		{ID: "D", SpeedByFeetPerMin: float64(10)},
		{ID: "E", SpeedByFeetPerMin: float64(2.5)},
		{ID: "F", SpeedByFeetPerMin: float64(25)},
		{ID: "G", SpeedByFeetPerMin: float64(15)},
	}

	bridges := AssignHikersToBridge(Bridges, Hikers)

	var totalTime float64

	crossingTimeMap := make(map[string]CrossingTime)
	for _, b := range bridges {
		// process crossing time per bridge
		response := CalculateCrossingTime(b)
		crossingTimeMap[response.BridgeID] = response
		totalTime += response.TotalTimeInMins
	}

	if crossingTimeForBridge1, ok := crossingTimeMap["1"]; ok {
		a.Equal(2, crossingTimeForBridge1.NumOfTripBackForFastestHiker)
		a.Equal(float64(17), crossingTimeForBridge1.FromAToBInMins)
		a.Equal(float64(2), crossingTimeForBridge1.FromBToAInMins)
		a.Equal(float64(19), crossingTimeForBridge1.TotalTimeInMins)
	}

	if crossingTimeForBridge2, ok := crossingTimeMap["2"]; ok {
		a.Equal(3, crossingTimeForBridge2.NumOfTripBackForFastestHiker)
		a.Equal(float64(142.5), crossingTimeForBridge2.FromAToBInMins)
		a.Equal(float64(7.5), crossingTimeForBridge2.FromBToAInMins)
		a.Equal(float64(150), crossingTimeForBridge2.TotalTimeInMins)
	}

	if crossingTimeForBridge3, ok := crossingTimeMap["3"]; ok {
		a.Equal(5, crossingTimeForBridge3.NumOfTripBackForFastestHiker)
		a.Equal(float64(101.5), crossingTimeForBridge3.FromAToBInMins)
		a.Equal(float64(7.5), crossingTimeForBridge3.FromBToAInMins)
		a.Equal(float64(109), crossingTimeForBridge3.TotalTimeInMins)
	}

	a.Equal(float64(278), totalTime)
}
