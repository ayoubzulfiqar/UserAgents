package platformagents

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ayoubzulfiqar/UserAgents/utils"
	"github.com/gocolly/colly"
)

func PlatformAgents() {
	data := getPlatformUserAgents()
	utils.CreateJSONUserAgents("PlatformAgents.json", data)
}
func getPlatformUserAgents() []string {

	links := utils.ReadLinks("PlatformAgents/links.json")

	var userAgents []string = make([]string, 0)

	c := colly.NewCollector(
		// colly.AllowedDomains(URL),
		colly.IgnoreRobotsTxt(),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./PlatformAgents/platform-user-agent-cache"),
	)
	limitErr := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       2 * time.Second, // Add a 2-second delay between requests
		RandomDelay: 1 * time.Second, // Add random delay between 1-2 seconds
	})
	if limitErr != nil {
		fmt.Println("LimitationError:", limitErr)
	}
	c.OnHTML("table.table-useragents tbody tr", func(e *colly.HTMLElement) {
		userAgent := e.ChildText("td.useragent")
		userAgents = append(userAgents, userAgent)
	})
	// only first five
	// remove indexing to get all which is a-lot
	for _, link := range links {
		fmt.Println("VisitingLink:", link)
		err := c.Visit(link)
		if err != nil {
			fmt.Println("ErrorVisiting:", err)
		}

	}
	return userAgents
}

func getPlatformLinks() []string {
	var links []string = make([]string, 0)
	var url string = "https://whatmyuseragent.com/platforms"
	var base string = "https://whatmyuseragent.com"
	c := colly.NewCollector(
		colly.IgnoreRobotsTxt(),
	)
	limitErr := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       2 * time.Second, // Add a 2-second delay between requests
		RandomDelay: 1 * time.Second, // Add random delay between 1-2 seconds
	})
	if limitErr != nil {
		fmt.Println("LimitError:", limitErr)
	}
	c.OnHTML("ul.list-group", func(e *colly.HTMLElement) {
		// Iterate over each <li> with a specific class inside the <ul>
		e.ForEach("li.list-group-item", func(_ int, el *colly.HTMLElement) {
			/* If you want platform and hardware - but we don't need*/
			// platform := e.ChildText("td:nth-child(2)")
			// hardware := e.ChildText("td:nth-child(3)")
			links = append(links, base+el.ChildAttr("a", "href"))

		})
	})
	err := c.Visit(url) // Replace with your target URL
	if err != nil {
		fmt.Println("Error visiting the website:", err)
	}
	// 390
	return links
}

func readPlatformLinks() []string {
	file, err := os.Open("PlatformAgents/links.json") // Replace with your JSON file path
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Step 2: Read the file content
	fileContent, err := os.ReadFile("PlatformAgents/links.json") // Read the entire file
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
func jsonData(fileName string, data []string) {
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
