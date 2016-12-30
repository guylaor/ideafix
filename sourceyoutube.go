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

type VideoFeed []struct {
	Title       string
	Description string
	PublishDate string
	YoutubueID  string
	Thumbnails  map[string]struct {
		Url string
	}
}

func doYoutubeAPIRequest() []byte {

	req, _ := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/playlistItems", nil)

	q := req.URL.Query()
	q.Add("part", "snippet")
	q.Add("maxResults", "50")
	q.Add("playlistId", fmt.Sprint(config["youtube-playlistid"]))
	q.Add("key", fmt.Sprint(config["youtube-key"]))
	q.Add("pageToken", "CDIQAA")
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

	return content
}

func loadYoutube() {

	content := doYoutubeAPIRequest()

	var feed youtubeFeed
	if err := json.Unmarshal(content, &feed); err != nil {
		checkWarnning(err)
	}

	fmt.Println(feed)
	fmt.Println(feed.Info["totalResults"])
	fmt.Println(feed.Items[0].Snippet.Title)
	fmt.Println(feed.Items[0].Snippet.Resource.VideoID)
	fmt.Println(feed.Items[0].Snippet.Thumbnails["default"].Url)

	//fmt.Printf("\n %s", req.URL.String())
}
