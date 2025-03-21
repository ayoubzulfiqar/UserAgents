package browseragents

import (
	"fmt"

	// "strings"
	"time"

	"github.com/ayoubzulfiqar/UserAgents/utils"
	"github.com/gocolly/colly"
)

func BrowserAgents() {
	data := getBrowserAgents()
	utils.CreateJSONUserAgents("BrowserAgents/BrowserAgents.json", data)
}

/*


Modify the struct if want to hardware and platform information otherwise it return the latest
user agents


but we only need user-agents so nothing else matter

 type UserAgents struct {
 	UserAgent string `json:"user_agent"`
	platform string `json:"platform"`
	hardware string `json:"platform"`
 }

*/

func getBrowsersLinks() []string {
	var links []string = make([]string, 0)
	var url string = "https://whatmyuseragent.com/browser"
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

func getBrowserAgents() []string {
	/*
	   Here we are scrapping the links and visiting the browser base user-agent page which probably is time consuming so i first scrap the browser links and save into json file so don't have to visit links first
	   we will scrap by directly going to the specific browser url

	   links := getLinks()

	*/
	links := utils.ReadLinks("BrowserAgents/links.json")

	var userAgents []string = make([]string, 0)

	c := colly.NewCollector(
		// colly.AllowedDomains(URL),
		colly.IgnoreRobotsTxt(),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./BrowserAgents/user-agent-cache"),
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
