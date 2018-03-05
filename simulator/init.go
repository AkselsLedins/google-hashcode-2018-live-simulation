package simulator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ghashcode "../ghashcode"
)

// ParseOutputFile it parse the file that your program has created
// It's the file you send to the Google Hashcode Judge System
func (s *Simulation) ParseOutputFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panicf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	// create a slice of vehicles
	var vehicles []*ghashcode.Vehicle

	for scanner.Scan() {
		// retrieve the line
		line := scanner.Text()

		// retrieve the trips' list
		strs := strings.Split(line, " ")
		numberOfTrips, _ := strconv.Atoi(strs[0])
		if numberOfTrips == 0 {
			continue
		}
		trips := make([]int32, numberOfTrips)
		for i := range strs {
			if i > 0 {
				val, _ := strconv.Atoi(strs[i])
				trips[i-1] = int32(val)
			}
		}

		// no point to store a vehicle that has no trips
		if len(trips) == 0 {
			continue
		}

		// instantiate a vehicle and append it to the vehicle slice
		vehicle := ghashcode.NewVehicle(trips)
		vehicles = append(vehicles, vehicle)
	}

	s.Vehicles = vehicles
}

// ParseInputFile it parse the examples
// It's the file you send to the Google Hashcode Judge System
func (s *Simulation) ParseInputFile(filePath string) {
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
	firstLine := scanner.Text()

	fmt.Sscanf(firstLine, "%d %d %d %d %d %d",
		&s.gridRows,
		&s.gridColumns,
		&s.numberOfVehicles,
		&s.numberOfRides,
		&s.perRideBonus,
		&s.numberOfStepsInTheSimulation)

	var trips []*ghashcode.Trip
	for id := 0; scanner.Scan(); id++ {
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

		// initialize a new trip
		trip := ghashcode.NewTrip(id, a, b, x, y, s, f)
		// and add it to the trips slice
		trips = append(trips, trip)
	}

	s.Trips = trips
}
