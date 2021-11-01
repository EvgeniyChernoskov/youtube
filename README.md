#Simple client for YouTube for search last video.
To work you must have YouTubeApiToken

//constructor
yt := youtube.New("YouTubeApiToken")

//Five last video from the channel Maksim Zhashkevych
items, err := yt.Search("", 5)

//10 last video from the channel Maksim Zhashkevych for "sql"
items, err := yt.Search("sql", 10)

//change channelID on АйТиБорода
yt.SetChannelId("UCeObZv89Stb2xLtjLJ0De3Q")


//search in all YouTube, 30 videos for "golang" 
yt.SetAllChannels()
items, err := yt.Search("golang", 30)

//print results
for _, item := range items {
fmt.Println(item.String())
}
