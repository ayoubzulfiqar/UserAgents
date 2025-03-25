# **User Agents**



## List Include
 **Implemented**
- [Browser User-Agents](BrowserAgents/BrowserAgents.json)
- [Platform UserAgents](PlatformAgents/PlatformAgents.json)
- [Bots User-Agents](BotsAgents)
- [Device User-Agents](DeviceAgents)
- [Engine User-Agents](EngineAgents)
- [Brand User-Agents](BrandAgents)
- [Application User-Agents](ApplicationAgents)


### How to use

Just open json file and you will get list of all agents
either you can parse the JSON file or you can use as list or array
pick as many as you need or you can use all of them.

**Golang**

```go
var userAgents []string = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.7.5",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.6.9",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.6.3",
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/49.0.2623.75 Safari/537.36 115Browser/7.0.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/24.0.2.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.2.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.6.5",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 115Browser/5.1.6",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/49.0.2623.75 Safari/537.36 115Browser/7.0.0",
	"Mozilla/5.0 (Windows NT 6.1; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.6.5",
}

```

**python**

```py
 userAgents= [
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.7.5",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.6.9",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 115Browser/27.0.6.3",
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/49.0.2623.75 Safari/537.36 115Browser/7.0.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/24.0.2.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.2.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.6.5",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 115Browser/5.1.6",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/49.0.2623.75 Safari/537.36 115Browser/7.0.0",
	"Mozilla/5.0 (Windows NT 6.1; ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36 115Browser/25.0.6.5",
]
```

**To get Latest**

```go
import (
	applicationagents "github.com/ayoubzulfiqar/UserAgents/ApplicationAgents"
	botsagents "github.com/ayoubzulfiqar/UserAgents/BotsAgents"
	brandagents "github.com/ayoubzulfiqar/UserAgents/BrandAgents"
	browseragents "github.com/ayoubzulfiqar/UserAgents/BrowserAgents"
	deviceagents "github.com/ayoubzulfiqar/UserAgents/DeviceAgents"
	engineagents "github.com/ayoubzulfiqar/UserAgents/EngineAgents"
	platformagents "github.com/ayoubzulfiqar/UserAgents/PlatformAgents"
)

func main() {
	// For Latest

	UserAgents()



	// Other Types

	applicationagents.ApplicationAgents()
	botsagents.BotsAgents()
	brandagents.BrandAgents()
	browseragents.BrowserAgents()
	deviceagents.DeviceAgent()
	engineagents.EngineAgents()
	platformagents.PlatformAgents()

}
```

[Latest](UserAgents.json)