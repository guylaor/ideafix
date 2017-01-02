package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type youtubeFeed struct {
	NextPageToken string                 `json:"nextPageToken"`
	Info          map[string]interface{} `json:"pageInfo"`
	Items         []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			PublishedAt string `json:"publishedAt"`
			Resource    struct {
				VideoID string `json:"videoId"`
			} `json:"resourceId"`
			Thumbnails map[string]struct {
				Url string `json:"url"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

type VideoItem struct {
	Title         string
	Description   string
	PublishDate   string
	YoutubueID    string
	ThumbSmall    string
	ThumbMedium   string
	ThumbHigh     string
	ThumbStandard string
}

type VideoFeed []VideoItem

func doYoutubeAPIRequest(token string, videoFeed *VideoFeed) string {

	req, _ := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/playlistItems", nil)

	q := req.URL.Query()
	q.Add("part", "snippet")
	q.Add("maxResults", "50")
	q.Add("playlistId", fmt.Sprint(config["youtube-playlistid"]))
	q.Add("key", fmt.Sprint(config["youtube-key"]))
	if token != "" {
		q.Add("pageToken", token)
	}
	req.URL.RawQuery = q.Encode()

	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		checkWarnning(err)
	}

	content, err := ioutil.ReadAll(res.Body)
	checkWarnning(err)

	var feed youtubeFeed
	if err := json.Unmarshal(content, &feed); err != nil {
		checkWarnning(err)
	}

	for _, item := range feed.Items {
		v := VideoItem{}
		v.Title = item.Snippet.Title
		v.Description = item.Snippet.Description
		v.PublishDate = item.Snippet.PublishedAt
		v.ThumbSmall = item.Snippet.Thumbnails["default"].Url
		v.ThumbMedium = item.Snippet.Thumbnails["medium"].Url
		v.ThumbHigh = item.Snippet.Thumbnails["high"].Url
		v.ThumbStandard = item.Snippet.Thumbnails["standard"].Url
		v.YoutubueID = item.Snippet.Resource.VideoID
		*videoFeed = append(*videoFeed, v)
	}

	fmt.Printf("\n %s", req.URL.String())
	return feed.NextPageToken

}

func loadYoutube() {

	videoFeed := make(VideoFeed, 0)
	token := doYoutubeAPIRequest("", &videoFeed)
	for {
		if token == "" {
			break
		}
		token = doYoutubeAPIRequest(token, &videoFeed)
	}

	fmt.Println(len(videoFeed))
	fmt.Println(videoFeed[0].Title)

	videoFeedJson, err := json.Marshal(videoFeed)
	checkWarnning(err)

	fmt.Printf("%s", videoFeedJson)
	// fmt.Println(feed.Info["totalResults"])
	// fmt.Println(feed.Items[0].Snippet.Title)
	// fmt.Println(feed.Items[0].Snippet.Resource.VideoID)
	// fmt.Println(feed.Items[0].Snippet.Thumbnails["default"].Url)

	//fmt.Printf("\n %s", req.URL.String())
}
