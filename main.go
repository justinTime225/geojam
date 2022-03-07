package main

import "fmt"

/*
Painting the picture

We have 3 bridges to cross and a list of hikers for each bridge
Figure out fastest time that hikers can cross each bridge and total time for all crossing

There are two ways to think through this
1) Find out who the fastest Hiker is and have that hiker pair with all the other hikers.  This optmize the total time because the time going back will almost be based on the fastest Hiker.  In real life, this may not be possible since one person is doing most of the work.
2) Pair any two hikers and send one back each time.  Every hiker except the last one will not have to cross back.  This would be the slower solution but probably more practical in real life.

For this program, we will go with the first approach.  We can pretend the simulated hiker has unlimited stamina.

Bridge 1: 100ft long

Hiker A: 100ft/min
Hiker B: 50ft/min
Hiker C: 20ft/min
Hiker D: 10ft/min

How program should run:
Hiker A and B takes 2 mins from A to B
Hiker A takes 1 min from B to A
Hiker A and C takes 5 mins from A to B
Hiker A takes 1 min from B to A
Hiker A and D takes 10 mins from A to B

Total Time for all hikers to cross bridge 1: 19 mins

Bridge 2: 250ft long
Hiker A: 100ft/min
Hiker B: 50ft/min
Hiker C: 20ft/min
Hiker D: 10ft/min
Hiker E: 2.5ft/min

How program should run:
Hiker A and B takes 5 mins from A to B
Hiker A takes 2.5 min from B to A
Hiker A and C takes 12.5 mins from A to B
Hiker A takes 2.5 min from B to A
Hiker A and D takes 25 mins from A to B
Hiker A takes 2.5 min from B to A
Hiker A and E takes 100 mins from A to B

Total Time for all hikers to cross bridge 2: 150 mins

Bridge 3: 150ft long
Hiker A: 100ft/min
Hiker B: 50ft/min
Hiker C: 20ft/min
Hiker D: 10ft/min
Hiker E: 2.5ft/min
Hiker F: 25ft/min
Hiker G: 15ft/min

How program should run:
Hiker A and B takes 3 mins from A to B
Hiker A takes 1.5 min from B to A
Hiker A and C takes 7.5 mins from A to B
Hiker A takes 1.5 min from B to A
Hiker A and D takes 15 mins from A to B
Hiker A takes 1.5 min from B to A
Hiker A and E takes 60 mins from A to B
Hiker A takes 1.5 min from B to A
Hiker A and F takes 6 mins from A to B
Hiker A takes 1.5 min from B to A
Hiker A and G takes 10 mins from A to B

Total Time for all hikers to cross bridge 3: 109 mins


Sum of all crossing times: 278 mins

Hikers
[
	{
		id: "A"
		speedByFeetPerMin: 100
	},
	{
		id: "B"
		speedByFeetPerMin: 50
	},
	{
		id: "C"
		speedByFeetPerMin: 20
	},
	{
		id: "D"
		speedByFeetPerMin: 10
	},
	{
		id: "E"
		speedByFeetPerMin: 2.5
	},
	{
		id: "F"
		speedByFeetPerMin: 25
	},
	{
		id: "G"
		speedByFeetPerMin: 15
	}
]

Bridges
[
	{
		id: 1
		lengthByFeet: 100
		crossByHikers: [A, B, C, D]
	},
	{
		id: 2
		lengthByFeet: 250
		crossByHikers: [A, B, C, D, E]
	},
	{
		id: 3
		lengthByFeet: 150
		crossByHikers: [A, B, C, D, E, F, G]
	},
]
*/

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

type Bridge struct {
	ID               string   `json:"id"`
	LengthByFeet     uint64   `json:"lengthByFeet"`
	HikerIDsCrossing []string `json:"hikersCrossing"`
	HikersCrossing   []Hiker
	FastestHiker     *Hiker
}

type Hiker struct {
	ID                string  `json:"id"`
	SpeedByFeetPerMin float64 `json:"speedByFeetPerMin"`
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func AssignHikersToBridge(bridges []Bridge, hikers []Hiker) []Bridge {
	// In an effort to not check between two arrays and thus traversing a loop inside a loop, it was necessary to create a hikerIDToHiker map for quick look up
	hikerIDToHikerMap := make(map[string]Hiker)
	for _, h := range hikers {
		hikerIDToHikerMap[h.ID] = h
	}

	// Here we assign hikers' detail onto the bridge so we can access information in one place
	for i := range bridges {
		for _, hikerIdOnBridge := range bridges[i].HikerIDsCrossing {
			if hiker, ok := hikerIDToHikerMap[hikerIdOnBridge]; ok {
				bridges[i].HikersCrossing = append(bridges[i].HikersCrossing, hiker)

				// we want to keep track of the fastest hiker because that person will be crossing back to get all the others.
				// for efficiency, we want the time traversing back to get other hikers be as fast as possible and the same each time
				if bridges[i].FastestHiker == nil || hiker.SpeedByFeetPerMin > bridges[i].FastestHiker.SpeedByFeetPerMin {
					bridges[i].FastestHiker = &hiker
				}
			}
		}
	}

	// now we have detail info on hikers crossing for each bridge
	return bridges
}

type CrossingTime struct {
	BridgeID                     string
	FromAToBInMins               float64
	FromBToAInMins               float64
	NumOfTripBackForFastestHiker int
	TotalTimeInMins              float64
}

func CalculateCrossingTime(bridge Bridge) CrossingTime {
	var crossingTimeResponse CrossingTime
	var minutesFromAToB float64

	// Here we calculate the times it takes for a pair of hikers (one of them being the fastest) to cross, we take the slower hiker's time
	for _, h := range bridge.HikersCrossing {
		if bridge.FastestHiker.ID != h.ID {
			crossingTime := float64(bridge.LengthByFeet) / h.SpeedByFeetPerMin
			minutesFromAToB += crossingTime
		}
	}

	// calculate the time it takes for fastest hiker to cross B to A to get the hiker
	// we can just take his speed and account for amount of times hiker walk back
	// fastest hiker goes crosses B to A len(hikers) - 2 amount of times because we exclude the fastest hiker and the last hiker

	// Calculate the time it takes fro fastest hiker to cross back (B To A) in order to retrieve the next hiker.
	minutesPerTripBackForFastestHiker := float64(bridge.LengthByFeet) / bridge.FastestHiker.SpeedByFeetPerMin
	// Calculate the number of times the fastest hiker crosses back.
	// If there are 5 hikers total, then fastest hiker crosses back 3 times.
	// We can see this as len(hikers) - 2 because fastest hiker does not have to cross back after the last person and fastest hiker makes the trip with the last hiker to conclude bridge crossing
	numberOfTripsFromBToAForFastestHiker := (len(bridge.HikersCrossing) - 2)
	minutesFromBToA := minutesPerTripBackForFastestHiker * float64(numberOfTripsFromBToAForFastestHiker)

	crossingTimeResponse.BridgeID = bridge.ID
	crossingTimeResponse.FromAToBInMins = minutesFromAToB
	crossingTimeResponse.FromBToAInMins = minutesFromBToA
	crossingTimeResponse.NumOfTripBackForFastestHiker = numberOfTripsFromBToAForFastestHiker
	crossingTimeResponse.TotalTimeInMins = minutesFromAToB + minutesFromBToA
	return crossingTimeResponse
}

// Originally I was going to have the input in a json file because I'm not familiar with yaml, but decided to just set the values here to save time.
func main() {
	// One of the first challenge of this problem was knowing hikers will cross what bridges
	// If we know that information ahead of time, then we can simplify the problem rather than repetitively passing in bridges and hikers
	// In an effort to address this, I decided to predetermine hikerIDs on each bridge.
	// Once we have that information, we can correlate hiker's detail on the bridge.
	// In this sense, we simply just need to pass in bridges containing information about what hikers will cross and the hikers' information
	bridges := AssignHikersToBridge(Bridges, Hikers)

	// Once we have bridges with hikers' detail we can calculate the total crossing time per bridge.
	var totalTime float64
	for _, b := range bridges {
		// process crossing time per bridge
		response := CalculateCrossingTime(b)
		totalTime += response.TotalTimeInMins
	}

	fmt.Printf("\n total Time %+v", totalTime)
}
