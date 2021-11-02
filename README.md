# YouTube API client.

## Description

Simple API Client with api key for search last videos https://www.youtube.com/

## Install

```
go get github.com/EvgeniyChernoskov/youtube
```

## Example

```go

package main

import (
	"fmt"
	"log"

	"github.com/EvgeniyChernoskov/youtube"
)

func main() {
	youtube := youtube.New("AIzaSyrqCa3_GOXatssdeEMOLwPD3rfdW294345c5pEvrR532105_HJo")
	
    //last video from the channel Maksim Zhashkevych
	items, err := youtube.Search("", 1)
	if err != nil {
		log.Fatal(err)
	}
	printResults(items)

	//Five last video from the channel Maksim Zhashkevych for "sql"
	items, err := youtube.Search("sql", 5)
	if err != nil {
		log.Fatal(err)
	}
	printResults(items)

	//change channelID on АйТиБорода (change channelID)
	youtube.SetChannelId("UCeObZv89Stb2xLtjLJ0De3Q")
	//search...
	
	//search in all YouTube, 30 videos for "golang" 
	youtube.SetAllChannels()
	items, err := youtube.Search("golang", 30)
	if err != nil {
		log.Fatal(err)
	}
	printResults(items)
	
}

func printResults(items []youtube.Item) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

```