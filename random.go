package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Implemented for Only JSON YET
func Randomize(path string) string {
	content := readFile(path)
	// Initialize PRNG with a seed (e.g., current Unix time)
	prng := NewPRNG(uint32(time.Now().Unix()))

	// Pick a random UserAgent
	selected, err := prng.RandomSelect(content)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	fmt.Println("RandomSelection:", selected)
	return ""
}

// PRNG := Pseudo-RandomNumber Generator

type PRNG struct {
	x uint32
}

// NewPRNG initializes a new PRNG with a seed.
func NewPRNG(seed uint32) *PRNG {
	return &PRNG{x: seed}
}

// Next generates the next pseudo-random uint32
func (r *PRNG) Next() uint32 {
	z := (r.x + 0x9E3779B9) // Weyl sequence step (golden ratio)
	z ^= z >> 16
	z *= 0x21f0aaad
	z ^= z >> 15
	z *= 0x735a2d97
	z ^= z >> 15
	r.x = z // Update state
	return z
}

// RandomSelect picks a random string using the custom PRNG.
func (r *PRNG) RandomSelect(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", fmt.Errorf("Slice is empty")
	}
	index := r.Next() % uint32(len(strings))
	return strings[index], nil
}

func readFile(path string) []string {
	file, err := os.Open(path) // Replace with your JSON file path
	if err != nil {
		fmt.Println("Error Opening File:", err)
	}
	defer file.Close()

	// Step 2: Read the file content
	fileContent, err := os.ReadFile(path) // Read the entire file
	if err != nil {
		fmt.Println("Error Reading File:", err)
	}
	var links []string
	err = json.Unmarshal(fileContent, &links)
	if err != nil {
		fmt.Println("Error UnMarshalling JSON:", err)
	}
	return links
}
