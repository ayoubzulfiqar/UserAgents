package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Create the links.json file
func WriteJSONLinks(fileName string, data []string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File Creation Error: %s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print the JSON
	encoderErr := encoder.Encode(data)
	if encoderErr != nil {
		fmt.Printf("EncoderError: %s", encoderErr)
	}
}

// read the links.json file
func ReadLinks(path string) []string {
	file, err := os.Open(path) // Replace with your JSON file path
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Step 2: Read the file content
	fileContent, err := os.ReadFile(path) // Read the entire file
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	var links []string
	err = json.Unmarshal(fileContent, &links)
	if err != nil {
		fmt.Println("Error UnMarshalling JSON:", err)
	}
	// for _, link := range links {
	// 	fmt.Printf("Link: %s\n", link)
	// }
	return links
}

// it create the ____Agents.json
func CreateJSONUserAgents(fileName string, data []string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File Creation Error: %s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print the JSON
	encoderErr := encoder.Encode(data)
	if encoderErr != nil {
		fmt.Printf("EncoderError: %s", encoderErr)
	}
}

func CreateTEXTUserAgents(fileName string, data []string) {
	// Step 1: Format the data as a JSON-like array
	var formattedData strings.Builder
	formattedData.WriteString("[\n") // Start with an opening square bracket and newline

	for i, ua := range data[0:6] {
		formattedData.WriteString("  " + ua) // Add indentation and the user agent string
		if i < len(data)-1 {
			formattedData.WriteString(",\n") // Add a comma and newline after each item (except the last one)
		}
	}

	formattedData.WriteString("\n]") // End with a newline and closing square bracket

	// Step 2: Create a new file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after writing

	// Step 3: Write the formatted data to the file
	_, err = file.WriteString(formattedData.String())
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created successfully:", fileName)
}

func ReadTXT() {}
