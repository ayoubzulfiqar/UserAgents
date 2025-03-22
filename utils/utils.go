package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println("FileError:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue // Skip empty lines and comments
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip malformed lines
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return scanner.Err()

}

func ReadEnv(env string) string {
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	envString := os.Getenv(env)
	return envString
}
