package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func numberRange(min, max int) []int {
	// Make a slice as big as your total number of slots
	integerRange := make([]int, max-min+1)

	// Fill the range with min + index until your slice is full
	for index := range integerRange {
		integerRange[index] = min + index
	}
	return integerRange
}

func octet() int {
	// Load in slice of numbers 1-255
	choices := numberRange(1, 255)
	// Pick a random index from that range
	choiceIndex := rand.Intn(len(choices))
	// Select a choice from the slice of numbers
	choice := choices[choiceIndex]
	return choice
}

func generateIPAddress() string {
	return fmt.Sprintf("%d.%d.%d.%d\n", octet(), octet(), octet(), octet())
}

func generateIPAddresses(total int) []string {
	ipAddresses := []string{}
	for i := 0; i < total; i++ {
		ipAddresses = append(ipAddresses, generateIPAddress())
	}
	return ipAddresses
}

func deduplicate(sourceSlice []string) []string {
	// Make a Map to hold all your entries and a "added" boolean flag
	sourceLength := len(sourceSlice)
	entryMap := make(map[string]bool)
	// Collector slice to store the good values in
	deduplicatedSlice := []string{}

	for len(deduplicatedSlice) < sourceLength {

		for _, entry := range sourceSlice {
			// Discard the Key, and check if the entry has been addedToSlice previously.
			if _, addedToSlice := entryMap[entry]; !addedToSlice {
				// Mark the entry as having been added previously
				entryMap[entry] = true
				// Update the deduplicated slice with the unique entry
				deduplicatedSlice = append(deduplicatedSlice, entry)
			} else {
				sourceSlice = append(sourceSlice, generateIPAddress())
			}
		}
	}

	return deduplicatedSlice
}

func writeToFile(stringSlice []string) {
	fileName := "./go_unique_ips.txt"
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, entry := range stringSlice {
		_, err := writer.WriteString(entry)
		checkError(err)
	}
	writer.Flush()
}

func main() {
	// Initialize Random Seed
	rand.Seed(time.Now().Unix())

	fmt.Println("Generating Unique IP's")
	deduplicatedIps := deduplicate(generateIPAddresses(500000))

	writeToFile(deduplicatedIps)
	fmt.Println("Done.")

}
