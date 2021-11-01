package youtube

import (
	"fmt"
	"time"
)

type Response struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id   ItemId   `json:"id"`
	Info ItemInfo `json:"snippet"`
}

type ItemId struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}

type ItemInfo struct {
	PublishTime  time.Time `json:"publishTime"`
	ChannelId    string    `json:"channelId"`
	VideoTitle   string    `json:"title"`
	Description  string    `json:"description"`
	ChannelTitle string    `json:"channelTitle"`
}

func (item Item) GetVideoUrl() string {
	return YOUTUBE_VIDEO_URL + item.Id.VideoId
}

func (item Item) String() string {
	return fmt.Sprintf("Channel: %s [%s]\nVideo: %s\nDescription: %s\nUrl: %s\n",
		item.Info.ChannelTitle,
		item.Info.PublishTime,
		item.Info.VideoTitle,
		item.Info.Description,
		item.GetVideoUrl())
}


