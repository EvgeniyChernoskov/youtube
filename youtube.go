package youtube

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	YOUTUBE_SEARCH_URL                   = "https://youtube.googleapis.com/youtube/v3/search"
	YOUTUBE_VIDEO_URL                    = "https://www.youtube.com/watch?v="
	YOUTUBE_MAXIM_ZHASHKEVYCH_CHANNEL_ID = "UCHF0TTrKzOASxt4aFByKpnQ"
)

type YouTube struct {
	apiToken  string
	channelId string
}

func New(token string) YouTube {
	return YouTube{apiToken: token, channelId: YOUTUBE_MAXIM_ZHASHKEVYCH_CHANNEL_ID}
}

//search for channelID
func (y *YouTube)SetChannelId(channelID string) {
	y.channelId=channelID
}
//set search for all channels
func (y *YouTube) SetAllChannels(){
	y.channelId=""
}

//search latest video on YouTube, max 50 results
func (y *YouTube) Search(target string, maxResults int) ([]Item, error) {
	req, err := y.findRequest(target, maxResults)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r.Items, nil
}

func (y *YouTube) findRequest(target string, count int) (*http.Request, error) {
	req, err := http.NewRequest("GET", YOUTUBE_SEARCH_URL, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("part", "snippet")
	if y.channelId!="" {
		query.Add("channelId", y.channelId)
	}
	query.Add("maxResults", strconv.Itoa(count))
	query.Add("order", "date")
	query.Add("q", target)
	query.Add("key", y.apiToken)
	req.URL.RawQuery = query.Encode()
	return req, nil
}
