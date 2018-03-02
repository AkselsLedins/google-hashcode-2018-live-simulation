package simulator

import (
	"bufio"
	"fmt"
	"log"
	"os"

	ghashcode "../ghashcode"
)

func ParseOutputFile(filePath string) []*ghashcode.Vehicle {
	fmt.Printf("[PARSING]: [%s]\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Panicf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	/* create a list of vehicles */
	var vehicles []*ghashcode.Vehicle

	for scanner.Scan() {
		line := scanner.Text()
		var numberOfTrips int32
		fmt.Sscanf(line, "%d", &numberOfTrips)
		fmt.Printf("Number of trips [%d]\n", numberOfTrips)
		trips := make([]*ghashcode.Trip, numberOfTrips)
		for i := int32(0); i < numberOfTrips; i++ {
			trip := new(ghashcode.Trip)
			trips[i] = trip
		}

		/* instantiate a vehicle */
		vehicle := new(ghashcode.Vehicle)
		vehicle.SetPosition(0, 0)

		/* add it to the vehicle list */
		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

func ParseInputFile(filePath string) []*ghashcode.Trip {
	fmt.Printf("[PARSING INPUT]: [%s]\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Panicf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	// File format
	//   The first line of the input file contains the following integer numbers separated by single spaces:
	//   ● R – number of rows of the grid (1 ≤ R ≤ 10000)
	//   ● C – number of columns of the grid (1 ≤ C ≤ 10000)
	//   ● F – number of vehicles in the fleet (1 ≤ F ≤ 1000)
	//   ● N – number of rides (1 ≤ N ≤ 10000)
	//   ● B – per-ride bonus for starting the ride on time (1 ≤ B ≤ 10000)
	//   ● T – number of steps in the simulation (1 ≤ T ≤ 10 )
	scanner.Scan()
	// skip first line at the moment
	scanner.Text()
	// fmt.Sscanf("%d %d %d %d %d")

	var trips []*ghashcode.Trip
	for scanner.Scan() {
		line := scanner.Text()
		// N subsequent lines of the input file describe the individual rides, from ride 0 to ride N − 1 . Each line
		//   contains the following integer numbers separated by single spaces:
		//   ● a – the row of the start intersection (0 ≤ a < R)
		//   ● b – the column of the start intersection (0 ≤ b < C)
		//   ● x – the row of the finish intersection (0 ≤ x < R)
		//   ● y – the column of the finish intersection (0 ≤ y < C)
		//   ● s – the earliest start(0 ≤ s < T)
		//   ● f – the latest finish (0 ≤ f ≤ T) , (f ≥ s + |x − a| + |y − b|)
		//   ○ note that f can be equal to T – this makes the latest finish equal to the end of the simulation
		var a, b, x, y, s, f int32
		fmt.Sscanf(line, "%d %d %d %d %d %d", &a, &b, &x, &y, &s, &f)

		trip := new(ghashcode.Trip)
		trip.SetStart(a, b)
		trip.SetEnd(x, y)
		trip.EarliestStart = s
		trip.LatestFinish = f

		trips = append(trips, trip)
	}

	return trips
}
