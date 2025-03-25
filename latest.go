package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func UserAgents() []string {
	links := []string{
		readEnv("BROWSERAGENT"),
		readEnv("PLATFORMAGENT"),
		readEnv("BRANDAGENT"),
		readEnv("DEVICESAGENT"),
		readEnv("BOTSAGENT"),
		readEnv("APPSAGENT"),
		readEnv("ENGINESAGENT"),
	}
	var latestUserAgents []string = make([]string, 0)

	c := colly.NewCollector()
	tdUseragentSelector := "table.table.table-striped.table-hover.table-bordered.table-useragents > tbody > tr > td.useragent"

	c.OnHTML(tdUseragentSelector, func(e *colly.HTMLElement) {
		// Extract the text content of the td.useragent element
		useragentText := e.Text
		fmt.Println("Useragent:", useragentText)

		// You can also access attributes of the td element if needed
		// Example:
		// useragentClass := e.Attr("class")
		// fmt.Println("Class:", useragentClass)
		latestUserAgents = append(latestUserAgents, useragentText)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	for _, link := range links {
		err := c.Visit(link)
		if err != nil {
			log.Fatal(err)
		}

	}
	saveJSON("UserAgents.json", latestUserAgents)
	return latestUserAgents
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

func readEnv(env string) string {
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	envString := os.Getenv(env)
	envString = strings.Trim(envString, `"`)
	return envString
}

func saveJSON(fileName string, data []string) {
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
