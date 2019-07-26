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

type set struct {
	content map[string]struct{}
}

func (self set) add(str string) {
	self.content[str] = struct{}{}
}

func (self set) has(str string) bool {
	_, ok := self.content[str]
	return ok
}

func (self set) len() int {
	return len(self.content)
}

func (self set) entries() []string {
	keys := make([]string, 0, self.len())
	for key := range self.content {
		keys = append(keys, key)
	}
	return keys
}

func newSet() set {
	s := set{}
	s.content = make(map[string]struct{})
	return s
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
	ipAddresses := newSet()
	for ipAddresses.len() < total {
		ip := generateIPAddress()
		ipAddresses.add(ip)
	}
	return ipAddresses.entries()
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

	writeToFile(generateIPAddresses(50000))
	fmt.Println("Done.")

}
