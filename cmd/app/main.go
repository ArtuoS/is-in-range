package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const EarthRadiusKm = 6371.0

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180.0))*math.Cos(lat2*(math.Pi/180.0))*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadiusKm * c
}

func validateCoordinates(filePath string, initialLat, initialLon, radius float64) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid format:", line)
			continue
		}

		lat, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid latitude:", parts[0])
			continue
		}
		lon, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("Invalid longitude:", parts[1])
			continue
		}

		distance := haversine(initialLat, initialLon, lat, lon)
		if distance <= radius {
			fmt.Printf("Coordinates (%.6f, %.6f) are within the radius.\n", lat, lon)
		} else {
			fmt.Printf("Coordinates (%.6f, %.6f) are outside the radius.\n", lat, lon)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	initialLat := flag.Float64("lat", 0.0, "Initial latitude")
	initialLon := flag.Float64("lon", 0.0, "Initial longitude")
	radius := flag.Float64("radius", 0.0, "Radius in kilometers")
	filePath := flag.String("file", "coordinates.txt", "Path to the coordinates file")

	flag.Parse()

	if *initialLat == 0.0 || *initialLon == 0.0 || *radius == 0.0 {
		fmt.Println("Please provide valid initial latitude, initial longitude, and radius.")
		return
	}

	validateCoordinates(*filePath, *initialLat, *initialLon, *radius)
}
